import http from 'k6/http';
import { sleep, check, group } from 'k6';
import { randomString } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';
import { Trend, Rate, Counter } from 'k6/metrics';

// 创建自定义指标
const dbOpsLatency = new Trend('db_operation_latency');
const dbOpsFailRate = new Rate('db_operation_failures');
const dbErrors = new Counter('db_errors');
const cacheHits = new Counter('cache_hits');
const cacheMisses = new Counter('cache_misses');

// 注意：K6主要是HTTP负载测试工具，不能直接监控MongoDB
// 建议在运行K6时，同步使用以下方法监控MongoDB:
// 1. 使用 'docker stats mongoContainer' 查看容器资源使用情况
// 2. 在另一个终端运行 'docker exec -it mongoContainer mongo --port 27017 --eval "db.serverStatus()"'
// 3. 使用 MongoDB Compass 连接到 localhost:27019 进行实时监控

export const options = {
  stages: [
    { duration: '15s', target: 500 },    // 快速增加到500用户
    { duration: '30s', target: 500 },    // 保持500用户30秒
    { duration: '20s', target: 1500 },   // 增加到1500用户
    { duration: '1m', target: 1500 },    // 保持1500用户1分钟
    { duration: '20s', target: 3000 },   // 峰值测试：增加到3000用户
    { duration: '1m', target: 3000 },    // 保持3000用户1分钟
    { duration: '20s', target: 0 },      // 缓慢降低到0用户
  ],
  thresholds: {
    http_req_duration: ['p(95)<1000'], // 放宽响应时间要求
    http_req_failed: ['rate<0.10'],    // 允许更高失败率
    'http_req_duration{requestType:正确请求}': ['p(95)<500'],  // 放宽正常请求时间要求
    'db_operation_latency': ['p(95)<500', 'p(99)<1000'],      // 放宽数据库延迟要求
    'db_operation_failures': ['rate<0.05'],                    // 允许更高数据库失败率
  },
};

// 创建一个热门短链接列表，用于测试高频访问场景
const popularShortLinks = [
  '2tdaue',  // 已知存在的短链接
  '2cx3dT',  // 已知存在的短链接
];

// 创建长连接测试用短链接列表（至少需保证存在于数据库中）
const longConnectionTestLinks = Array(50).fill(0).map(() => randomString(6, 'abcdefghijklmnopqrstuvwxyz0123456789'));

// 创建一组额外的测试用例，针对数据库性能
const dbStressTestCases = [
  { type: '批量重复请求', count: 200 },   // 增加到200个请求
  { type: '随机混合请求', count: 100 },   // 增加到100个请求
  { type: '超大批量请求', count: 300 },   // 增加到300个请求
  { type: '长事务请求', count: 80 },      // 增加到80个请求
  { type: '复杂查询', count: 40 }         // 增加到40个请求
];

// 定义错误请求类型
const errorRequestTypes = [
  { type: '随机错误请求', weight: 5 },
  { type: '格式正确但不存在', weight: 5 },
  { type: '高频重复请求', weight: 4 },
  { type: '格式错误请求', weight: 3 },
  { type: '带额外参数请求', weight: 3 }
];

// 按权重随机选择错误请求类型
function selectRandomErrorType() {
  const totalWeight = errorRequestTypes.reduce((sum, type) => sum + type.weight, 0);
  let random = Math.random() * totalWeight;
  
  for (const errorType of errorRequestTypes) {
    random -= errorType.weight;
    if (random <= 0) {
      return errorType.type;
    }
  }
  return errorRequestTypes[0].type; // 默认返回第一种
}

// 记录所有请求过的短链接及其响应时间，用于模拟缓存效果分析
const requestedLinks = new Map();

// 创建连续批次请求用于模拟持续高负载
const continuousBatchSettings = {
  enabled: true,        // 是否启用连续批次请求
  batchSize: 100,       // 每批次请求数量
  batchCount: 5,        // 连续批次数
  batchInterval: 0.01,  // 批次间隔(秒)
  targetUrls: [         // 目标URL列表
    { url: `http://zeroLink:8001/${popularShortLinks[0]}`, weight: 20 }, // 热门URL权重高
    { url: `http://zeroLink:8001/${popularShortLinks[1]}`, weight: 10 },
    { url: 'http://zeroLink:8001/NOTEXIST', weight: 5 }, // 不存在的URL
  ]
};

export default function() {
  // 10%概率执行连续批次请求
  if (continuousBatchSettings.enabled && Math.random() < 0.10) {
    group('连续批次压力测试', function() {
      console.log(`开始连续批次压力测试: 批次数=${continuousBatchSettings.batchCount}, 每批请求数=${continuousBatchSettings.batchSize}`);
      
      // 根据权重选择目标URL
      function selectTargetUrl() {
        const totalWeight = continuousBatchSettings.targetUrls.reduce((sum, item) => sum + item.weight, 0);
        let random = Math.random() * totalWeight;
        
        for (const target of continuousBatchSettings.targetUrls) {
          random -= target.weight;
          if (random <= 0) {
            return target.url;
          }
        }
        return continuousBatchSettings.targetUrls[0].url;
      }
      
      // 执行连续批次请求
      for (let batch = 0; batch < continuousBatchSettings.batchCount; batch++) {
        for (let i = 0; i < continuousBatchSettings.batchSize; i++) {
          const targetUrl = selectTargetUrl();
          sendRequest(targetUrl, `连续批次_${batch}_${i}`);
        }
        
        // 批次间短暂休眠
        if (batch < continuousBatchSettings.batchCount - 1) {
          sleep(continuousBatchSettings.batchInterval);
        }
      }
    });
  }
  // 减少普通压力测试概率到50%，为连续批次测试腾出空间
  else if (Math.random() < 0.50) {
    group('数据库压力测试', function() {
      // 选择压力测试类型
      const testCase = dbStressTestCases[Math.floor(Math.random() * dbStressTestCases.length)];
      console.log(`开始数据库压力测试: ${testCase.type}, 请求数: ${testCase.count}`);
      
      // 批量发送请求，删除内部休眠
      for (let i = 0; i < testCase.count; i++) {
        if (testCase.type === '批量重复请求') {
          // 重复请求同一个热门链接
          sendRequest(`http://zeroLink:8001/${popularShortLinks[0]}`, `${testCase.type}_${i}`);
        } else if (testCase.type === '随机混合请求') {
          // 混合正确和错误请求
          if (i % 2 === 0) {
            const linkIndex = Math.floor(Math.random() * popularShortLinks.length);
            sendRequest(`http://zeroLink:8001/${popularShortLinks[linkIndex]}`, `${testCase.type}_正确_${i}`);
          } else {
            const randomId = randomString(6, 'abcdefghijklmnopqrstuvwxyz0123456789');
            sendRequest(`http://zeroLink:8001/${randomId}`, `${testCase.type}_错误_${i}`);
          }
        } else if (testCase.type === '超大批量请求') {
          // 大批量随机请求
          const randomId = randomString(6, 'abcdefghijklmnopqrstuvwxyz0123456789');
          sendRequest(`http://zeroLink:8001/${randomId}`, `${testCase.type}_${i}`);
        } else if (testCase.type === '长事务请求') {
          // 模拟长事务 - 发送带特殊参数的请求，服务端可能需要特殊处理
          const linkId = longConnectionTestLinks[i % longConnectionTestLinks.length];
          sendRequest(`http://zeroLink:8001/${linkId}?simulate=longTx&duration=${Math.floor(Math.random() * 500)}`, `${testCase.type}_${i}`);
        } else if (testCase.type === '复杂查询') {
          // 模拟复杂查询 - 发送带复杂查询参数的请求
          const complexQueries = [
            `?simulate=complexQuery&field=all&ts=${Date.now()}`,
            `?simulate=analytics&groupBy=hour&from=${Date.now() - 86400000}&to=${Date.now()}`,
            `?simulate=search&term=${randomString(10)}&filters=domain,pv,uv&sort=desc`
          ];
          const queryIndex = i % complexQueries.length;
          const linkId = popularShortLinks[0]; // 使用已知存在的链接
          sendRequest(`http://zeroLink:8001/${linkId}${complexQueries[queryIndex]}`, `${testCase.type}_${i}`);
        }
      }
    });
  } else {
    // 40%的概率发送正确请求，60%的概率发送错误请求
    if (Math.random() < 0.4) {
      group('正常请求', function() {
        // 正确请求 - 已知存在的短链接
        const linkIndex = Math.floor(Math.random() * popularShortLinks.length);
        sendRequest(`http://zeroLink:8001/${popularShortLinks[linkIndex]}`, '正确请求');
      });
    } else {
      group('错误请求', function() {
        // 选择一种错误请求类型
        const errorType = selectRandomErrorType();
        
        switch (errorType) {
          case '随机错误请求':
            // 完全随机的短链接ID
            const randomId = randomString(6, 'abcdefghijklmnopqrstuvwxyz0123456789');
            sendRequest(`http://zeroLink:8001/${randomId}`, errorType);
            break;
            
          case '格式正确但不存在':
            // 格式正确但不存在的短链接（使用大写字母开头，降低与现有链接冲突的可能）
            const nonExistingId = 'Z' + randomString(5, 'abcdefghijklmnopqrstuvwxyz0123456789');
            sendRequest(`http://zeroLink:8001/${nonExistingId}`, errorType);
            break;
            
          case '高频重复请求':
            // 重复访问同一个不太可能存在的短链接，测试数据库缓存
            sendRequest(`http://zeroLink:8001/NOTEXIST`, errorType);
            break;
            
          case '格式错误请求':
            // 格式不正确的请求（长度不是6位，或包含特殊字符）
            const invalidFormats = [
              randomString(3, 'abcdefghijklmnopqrstuvwxyz0123456789'), // 太短
              randomString(10, 'abcdefghijklmnopqrstuvwxyz0123456789'), // 太长
              randomString(4, 'abcdefghijklmnopqrstuvwxyz0123456789') + '!@', // 含特殊字符
            ];
            const invalidFormat = invalidFormats[Math.floor(Math.random() * invalidFormats.length)];
            sendRequest(`http://zeroLink:8001/${invalidFormat}`, errorType);
            break;
            
          case '带额外参数请求':
            // 带额外查询参数的请求，测试数据库查询处理
            const baseId = randomString(6, 'abcdefghijklmnopqrstuvwxyz0123456789');
            const queries = [
              `?utm_source=test`,
              `?t=${Date.now()}`,
              `?test=1&debug=true`,
            ];
            const query = queries[Math.floor(Math.random() * queries.length)];
            sendRequest(`http://zeroLink:8001/${baseId}${query}`, errorType);
            break;
        }
      });
    }
  }
  
  // 压缩休眠时间，增加请求频率
  if (Math.random() < 0.05) {
    // 只有5%概率长休眠，大幅减少
    sleep(Math.random() * 0.2);
  } else if (Math.random() < 0.3) {
    // 只有30%概率短休眠
    sleep(0.05);
  } else {
    // 65%概率极短休眠，模拟高压力场景
    sleep(0.01);
  }
}

function sendRequest(url, requestType) {
  const startTime = new Date();
  
  // 提取短链接ID，用于缓存分析
  const urlPath = url.split('?')[0]; // 去除查询参数
  const linkId = urlPath.split('/').pop(); // 获取路径最后部分
  
  // 判断是否已经请求过这个链接（模拟缓存分析）
  const isCacheHit = requestedLinks.has(linkId);
  if (isCacheHit) {
    cacheHits.add(1);
  } else {
    cacheMisses.add(1);
    // 记录新请求的链接
    requestedLinks.set(linkId, true);
  }
  
  // 对于模拟长事务的请求，添加自定义头部
  let headers = {};
  if (url.includes('simulate=longTx')) {
    headers['X-Custom-Simulation'] = 'longTransaction';
  } else if (url.includes('simulate=complexQuery')) {
    headers['X-Custom-Simulation'] = 'complexQuery';
  } else if (url.includes('simulate=analytics')) {
    headers['X-Custom-Simulation'] = 'analytics';
  }
  
  let res = http.get(url, {
    tags: { requestType: requestType, cacheHit: isCacheHit ? 'true' : 'false' },
    headers: headers
  });
  
  const endTime = new Date();
  const totalTime = endTime - startTime;
  
  // 假设响应时间超过100ms的请求主要是数据库查询导致
  if (res.timings.duration > 100) {
    dbOpsLatency.add(res.timings.duration);
  }
  
  // 假设500错误和超时错误可能与数据库有关
  if (res.status >= 500 || res.timings.duration > 500) {
    dbOpsFailRate.add(1);
    dbErrors.add(1);
  } else {
    dbOpsFailRate.add(0);
  }
  
  check(res, { 
    "状态码检查": (res) => requestType.includes('正确') ? res.status === 200 : true,
    "响应时间 < 200ms": (res) => res.timings.duration < 200,
    "总请求时间 < 300ms": () => totalTime < 300
  }, { requestType: requestType, cacheHit: isCacheHit ? 'true' : 'false' });
  
  // 输出请求详情，区分正确和错误请求
  console.log(`${requestType} | URL: ${url} | 状态码: ${res.status} | 请求总时间: ${totalTime}ms | 服务器处理时间: ${res.timings.duration}ms | 缓存命中: ${isCacheHit}`);
  
  // 添加数据库性能相关错误监控
  if (res.status === 500 || res.timings.duration > 300) {
    console.warn(`可能的数据库性能问题 | ${url} | 状态: ${res.status} | 处理时间: ${res.timings.duration}ms`);
  }
}
