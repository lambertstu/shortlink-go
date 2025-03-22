<template>
  <a-table :columns="columns" :data-source="data">
    <template #bodyCell="{ column, record }">
      <template v-if="column.key === 'info'">
        <div style="display: flex; align-items: center;">
          <link-outlined style="margin-right: 8px;" />
          <div>
            <div>{{ record.info }}</div>
            <div style="color: #666; font-size: 12px;">{{ record.date }}</div>
          </div>
        </div>
      </template>
      <template v-else-if="column.key === 'url'">
        <div>
          <a :href="record.url1" target="_blank">{{ record.url1 }}</a>
          <br />
          <a :href="record.url2" target="_blank" style="color: #666; font-size: 12px;">{{ record.url2 }}</a>
        </div>
      </template>
      <template v-else-if="column.key === 'visits'">
        <div>
          <div>今日 {{ record.visitsToday }}</div>
          <div style="color: #666;">累计 {{ record.visitsTotal }}</div>
        </div>
      </template>
      <template v-else-if="column.key === 'visitors'">
        <div>
          <div> {{ record.visitorsToday }}</div>
          <div style="color: #666; font-size: 12px;"> {{ record.visitorsTotal }}</div>
        </div>
      </template>
      <template v-else-if="column.key === 'ip'">
        <div>
          <div> {{ record.ipToday }}</div>
          <div style="color: #666; font-size: 12px;"> {{ record.ipTotal }}</div>
        </div>
      </template>
      <template v-else-if="column.key === 'action'">
        <span>
          <a>数据</a>
          <a-divider type="vertical" />
          <a>编辑</a>
          <a-divider type="vertical" />
          <a class="ant-dropdown-link">
            更多
            <down-outlined />
          </a>
        </span>
      </template>
    </template>
  </a-table>
</template>

<script lang="ts" setup>
import { DownOutlined, LinkOutlined } from '@ant-design/icons-vue';

const columns = [
  {
    title: '短链信息',
    dataIndex: 'info',
    key: 'info',
  },
  {
    title: '短链网址',
    dataIndex: 'url',
    key: 'url',
  },
  {
    title: '访问次数',
    dataIndex: 'visits',
    key: 'visits',
  },
  {
    title: '访问人数',
    dataIndex: 'visitors',
    key: 'visitors',
  },
  {
    title: 'IP数',
    dataIndex: 'ip',
    key: 'ip',
  },
  {
    title: '操作',
    key: 'action',
  },
];

const data = [
  {
    key: '1',
    info: 'test demo',
    date: '2025-03-11 11:50:13',
    url1: 'https://sourl.cn/yAxsSn',
    url2: 'https://www.google.com/search?q...',
    visitsToday: 0,
    visitsTotal: 6,
    visitorsToday: 0,
    visitorsTotal: 5,
    ipToday: 0,
    ipTotal: 5,
  },
];
</script>
