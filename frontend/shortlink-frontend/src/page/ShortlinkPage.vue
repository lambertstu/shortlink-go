<script setup lang="ts">
import { ref } from 'vue';
import { Menu, MenuItem, SubMenu, Table, Input } from 'ant-design-vue';
import AsideMenu from "@/components/AsideMenu.vue";
import ShortlinkMenuComponent from "@/components/ShortlinkMenuComponent.vue";
import ShortlinkContentPage from "@/components/ShortlinkContentPage.vue";
import UserInfoCardComponent from "@/components/UserInfoCardComponent.vue";
const searchText = ref('');

const columns = [
  {
    title: '短信名称',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: '短信网址',
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
    dataIndex: 'ips',
    key: 'ips',
  },
  {
    title: '操作',
    key: 'action',
  },
];

const data = [
  {
    key: '1',
    name: 'best demo',
    url: 'https://bearl.cn/yAxaSdn',
    visits: 0,
    visitors: 0,
    ips: 0,
  },
  {
    key: '2',
    name: '2025-03-11 15:01:13',
    url: 'https://www.google.com/search?q=ME5N8D9k&_',
    visits: 3,
    visitors: 2,
    ips: 2,
  },
];


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
        <div class="menu_content">
          <div>
            项目/ 我的项目
          </div>

          <div>
            +
          </div>
        </div>

        <AsideMenu />

      </aside>
      <main class="main">
        <Input v-model:value="searchText" placeholder="请输入 http:// 或 https:// 开头的链接或应用路径链接" />
        <ShortlinkContentPage />
      </main>
    </div>
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

.menu_content{
  display: flex;
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
</style>
