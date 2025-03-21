<template>
  <div id="components-popover-demo-placement">
    <a-popover placement="bottomLeft">
      <template #content>
        <p>Content</p>
        <p>Content</p>
      </template>
      <template #title>
        <span>Title</span>
      </template>
      <UserOutlined class="user-icon" />
    </a-popover>
  </div>
</template>
<script lang="ts" setup>

import {UserOutlined} from "@ant-design/icons-vue";
</script>
<style scoped>
#components-popover-demo-placement .ant-btn {
  width: 70px;
  text-align: center;
  padding: 0;
  margin-right: 8px;
  margin-bottom: 8px;
}

.user-icon {
  font-size: 24px; /* 调整图标大小 */
  color: #1973f4; /* 轮廓颜色（蓝色） */
  margin-right: 20px  ;
}
</style>
