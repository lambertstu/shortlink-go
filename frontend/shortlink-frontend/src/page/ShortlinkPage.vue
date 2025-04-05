<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { Input, message, Drawer, Select } from 'ant-design-vue';
import AsideMenu from "@/components/AsideMenu.vue";
import ShortlinkMenuComponent from "@/components/ShortlinkMenuComponent.vue";
import ShortlinkContentPage from "@/components/ShortlinkContentPage.vue";
import UserInfoCardComponent from "@/components/UserInfoCardComponent.vue";
import { createShortLink } from "@/api/shortlink/shortlinkApi.ts";
import { getGroupInfo } from "@/api/group/groupApi";

const searchText = ref('');
const selectedGroup = ref<any>(null);
const newShortlinkUrl = ref('');
const newShortlinkDesc = ref('');
const drawerVisible = ref(false);
const createLoading = ref(false);
const shortlinkContentPage = ref<any>(null);
const groupList = ref<any[]>([]);
const selectedGroupForLink = ref<string>('');

// 获取所有分组信息
const fetchGroups = async () => {
  try {
    const username = localStorage.getItem('username');
    if (username) {
      const response = await getGroupInfo(username);
      if (response.data.success && response.data.data.data) {
        groupList.value = response.data.data.data;
      }
    }
  } catch (error) {
    console.error("获取分组信息失败:", error);
    message.error('获取分组信息失败');
  }
};

const handleGroupSelected = (group: any) => {
  selectedGroup.value = group;
  // 当用户选择侧边栏分组时，默认将创建短链的分组也设置为该分组
  selectedGroupForLink.value = group.gid;
};

const showCreateDrawer = () => {
  if (!newShortlinkUrl.value) {
    message.warning('请输入需要转换的链接');
    return;
  }
  
  // 如果当前有选中的分组，默认设置为该分组
  if (selectedGroup.value) {
    selectedGroupForLink.value = selectedGroup.value.gid;
  } else if (groupList.value.length > 0) {
    // 否则默认选择第一个分组
    selectedGroupForLink.value = groupList.value[0].gid;
  } else {
    message.warning('没有可用的分组');
    return;
  }
  
  drawerVisible.value = true;
};

const handleCreateShortlink = async () => {
  if (!selectedGroupForLink.value) {
    message.warning('请选择一个分组');
    return;
  }
  
  if (!newShortlinkUrl.value) {
    message.warning('请输入需要转换的链接');
    return;
  }
  
  createLoading.value = true;
  
  try {
    const response = await createShortLink({
      origin_url: newShortlinkUrl.value,
      gid: selectedGroupForLink.value,
      describe: newShortlinkDesc.value || newShortlinkUrl.value // 如果没有输入描述，默认使用URL
    });
    
    if (response.data && response.data.success) {
      message.success('短链接创建成功');
      // 清空输入框
      newShortlinkUrl.value = '';
      newShortlinkDesc.value = '';
      // 关闭抽屉
      drawerVisible.value = false;
      // 刷新短链接列表
      if (shortlinkContentPage.value) {
        shortlinkContentPage.value.fetchShortLinksForSelectedGroup();
      }
    } else {
      message.error(response?.data?.message || '创建短链接失败');
    }
  } catch (error) {
    console.error("创建短链接失败:", error);
    message.error('创建短链接失败');
  } finally {
    createLoading.value = false;
  }
};

// 在组件挂载时获取分组信息
onMounted(() => {
  fetchGroups();
});
</script>

<template>
  <div class="layout">
    <header class="header">
      <img data-v-0c642287="" class="logo" src="https://cdn.xiaomark.com/portal/main/new-sl-logo.png">
      <shortlink-menu-component class="top-menu"/>
      <Input class="shortlink_search" v-model:value="searchText" placeholder="请输入关键词进行搜索" />
      <UserInfoCardComponent class="userInfo"/>
    </header>
    <div class="content">
      <aside class="sidebar">
        <AsideMenu @group-selected="handleGroupSelected" />
      </aside>
      <main class="main">
        <a-input-search
          class="shortlinkCreatorInput"
          placeholder="请输入 http:// 或 https:// 开头的链接或应用路径链接"
          enter-button="创建短链"
          size="large"
          v-model:value="newShortlinkUrl"
          @search="showCreateDrawer"
        />
        <ShortlinkContentPage :selectedGroup="selectedGroup" ref="shortlinkContentPage" />
      </main>
    </div>
    
    <!-- 创建短链接的抽屉组件 -->
    <a-drawer
      title="创建短链接"
      placement="right"
      :visible="drawerVisible"
      @close="drawerVisible = false"
      :width="400"
    >
      <div class="drawer-content">
        <div class="form-item">
          <div class="form-label">目标链接</div>
          <a-input 
            v-model:value="newShortlinkUrl" 
            placeholder="请输入目标链接" 
          />
        </div>
        
        <div class="form-item">
          <div class="form-label">短链描述</div>
          <a-input 
            v-model:value="newShortlinkDesc" 
            placeholder="请输入短链描述" 
          />
        </div>
        
        <div class="form-item">
          <div class="form-label">所属分组</div>
          <a-select
            v-model:value="selectedGroupForLink"
            style="width: 100%"
            placeholder="请选择分组"
          >
            <a-select-option 
              v-for="group in groupList" 
              :key="group.gid" 
              :value="group.gid"
            >
              {{ group.name }}
            </a-select-option>
          </a-select>
        </div>
        
        <a-button 
          type="primary" 
          block 
          :loading="createLoading" 
          @click="handleCreateShortlink"
          style="margin-top: 20px;"
        >
          确认创建
        </a-button>
      </div>
    </a-drawer>
  </div>
</template>



<style scoped>
.logo{
  width: 80px; /* 设置图标宽度 */
  height: 40px; /* 设置图标高度 */
  float: left; /* 将图标浮动到左边 */
  margin-left: 30px; /* 图标与文本之间的间距 */
}

.userInfo{
  margin-left: auto;
}

.layout {
  display: flex;
  flex-direction: column;
  height: 100vh;
}

.shortlink_search{
  flex: 1; /* 让输入框占据剩余空间 */
  max-width: 450px; /* 控制输入框最大宽度，防止太长 */
  position: relative;
  margin-bottom: -12px;
  margin-left: 30px;
  height: 40px;
}

.shortlinkCreatorInput{
  max-width: 1000px; /* 控制输入框最大宽度，防止太长 */
  margin-left: 40px;
  margin-bottom: 10px;
}


.header {
  display: flex;
  background-color: #efefef;
  color: white;
  padding: 15px;
  margin-top: -5px;
  text-align: center;
  align-items: flex-end; /* 让子元素对齐到底部 */
}

.content {
  display: flex;
  flex: 1;
}

.sidebar {
  width: 256px;
  background-color: #f0f2f5;
}

.main {
  flex: 1;
  padding: 16px;
}

.top-menu{
  margin-bottom: -15px;
}

.drawer-content {
  padding: 0 16px;
}

.form-item {
  margin-bottom: 16px;
}

.form-label {
  margin-bottom: 8px;
  font-weight: 500;
}
</style>
