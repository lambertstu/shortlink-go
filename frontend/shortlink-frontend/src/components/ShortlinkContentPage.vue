<template>
  <div>
    <!-- Show loading state or message when no group selected -->
    <div v-if="loading" class="loading-state">
      <a-spin tip="Loading..." />
    </div>
    <div v-else-if="!selectedGroup" class="empty-state">
      <p>请选择一个分组查看短链接</p>
    </div>
    <div v-else-if="noShortlinks" class="empty-state">
      <p>当前分组还未创建短链哟</p>
    </div>
    <!-- Table with data -->
    <a-table v-else :columns="columns" :data-source="data">
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
            <a :href="getFullUrl(record.url1)" target="_blank">{{ record.url1 }}</a>
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
  </div>
  <a-modal
    title="创建分组"
  >
    <a-input placeholder="请输入分组名称" />
  </a-modal>
</template>

<script lang="ts" setup>
import { DownOutlined, LinkOutlined } from '@ant-design/icons-vue';
import { fetchShortLinks } from "@/api/shortlink/shortlinkApi.ts";
import { onMounted, ref, watch } from "vue";
import { message } from 'ant-design-vue';

// Properly define the props with the correct prop name (selectedGroup)
const props = defineProps({
  selectedGroup: {
    type: Object,
    default: null
  }
});

// 添加处理URL的方法
const getFullUrl = (url: string) => {
  if (!url) return '';
  // 检查URL是否已经包含http或https前缀
  if (url.startsWith('http://') || url.startsWith('https://')) {
    return url;
  }
  // 添加http前缀
  return `http://${url}`;
};

const data = ref([]);
const loading = ref(false);
const noShortlinks = ref(false);

// Watch for changes to selectedGroup prop
watch(() => props.selectedGroup, async (newGroup) => {
  if (newGroup) {
    await fetchShortLinksForSelectedGroup();
  }
}, { immediate: true });

// Fetch shortlinks for the selected group
const fetchShortLinksForSelectedGroup = async () => {
  if (!props.selectedGroup) return;
  
  loading.value = true;
  noShortlinks.value = false;
  
  try {
    const response = await fetchShortLinks({
      gid: props.selectedGroup.gid,
      page: 1,
      size: 10,
      orderTag: 1
    });
    
    if (response.data && response.data.success) {
      // Check if list is null or empty
      if (!response.data.data.list || response.data.data.list.length === 0) {
        noShortlinks.value = true;
        data.value = [];
        return;
      }
      
      data.value = response.data.data.list.map((link: any) => ({
        key: link.short_uri,
        info: link.describe,
        date: link.update_at,
        url1: link.full_short_url,
        url2: link.origin_url,
        visitsToday: link.today_pv,
        visitsTotal: link.total_pv,
        visitorsToday: link.today_uv,
        visitorsTotal: link.total_uv,
        ipToday: link.today_uip,
        ipTotal: link.total_uip,
      }));
    } else {
      message.error(response?.data?.message || '获取短链接数据失败');
      data.value = [];
    }
  } catch (error) {
    console.error("获取短链接数据失败:", error);
    message.error('获取短链接数据失败');
    data.value = [];
  } finally {
    loading.value = false;
  }
};

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

// Initialize when component mounts
onMounted(() => {
  if (props.selectedGroup) {
    fetchShortLinksForSelectedGroup();
  }
});
</script>

<style scoped>
.loading-state, .empty-state {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 300px;
  color: #999;
}
</style>
