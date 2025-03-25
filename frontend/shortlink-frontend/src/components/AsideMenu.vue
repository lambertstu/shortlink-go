<template>
  <div>
    <div class="menu_content">
      <div class="title">
        短链分组
      </div>

      <a-button class="createGroup" type="link" >
        <PlusSquareOutlined />
      </a-button>
    </div>
    <a-menu
      class = "groupMenu"
      style="width: 256px"
      :default-selected-keys="['1']"
      mode="inline"
    >
      <a-menu-item key="1">
        默认分组
      </a-menu-item>
      <a-menu-item key="2">
        分组1
      </a-menu-item>
      <a-menu-item key="3">
        分组2
      </a-menu-item>
    </a-menu>
  </div>
</template>
<script lang="ts" setup>
import {PlusSquareOutlined} from "@ant-design/icons-vue";
import { ref, onMounted } from "vue";
import { getGroupInfo } from "@/api/group/groupApi";

interface GroupData {
  gid: string;
  name: string;
  username: string;
}

const groupList = ref<GroupData[]>([]);

onMounted(async () => {
  try {
    const username = localStorage.getItem('username');
    if (username) {
      const response = await getGroupInfo(username);
      if (response.data.success) {
        groupList.value = response.data.data.data;
      }
    }
  } catch (error) {
    console.error("获取群组信息失败:", error);
  }
});
</script>

<style>
.menu_content {
  display: flex;
  align-items: center; /* 让子元素在垂直方向上居中对齐 */
  margin-top: 15px;
}

.title {
  margin-left: 20px;
  font-size: 18px; /* 根据需要调整字体大小 */
  line-height: 1; /* 让文本与按钮保持一致高度 */
  display: flex;
  align-items: center; /* 确保文字垂直居中 */
}

.createGroup {
  margin-left: auto;
  font-size: 24px; /* 让按钮里的内容更大 */
  padding: 10px 20px; /* 调整按钮的整体大小 */
  display: flex;
  align-items: center; /* 垂直居中 */
  justify-content: center; /* 水平居中 */
}

.groupMenu{
  margin-top: 10px;
}
</style>
