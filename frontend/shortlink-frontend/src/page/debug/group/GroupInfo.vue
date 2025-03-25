<template>
  <div>
    <h2>群组信息</h2>
    <div v-if="groupList.length > 0">
      <div v-for="group in groupList" :key="group.gid" class="group-item">
        <p><strong>群组 ID:</strong> {{ group.gid }}</p>
        <p><strong>群组名称:</strong> {{ group.name }}</p>
        <p><strong>用户名:</strong> {{ group.username }}</p>
      </div>
    </div>
    <div v-else>
      <p>暂无群组信息</p>
    </div>
  </div>
</template>

<script setup lang="ts">
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

<style scoped>
.group-item {
  border: 1px solid #eee;
  padding: 15px;
  margin-bottom: 15px;
  border-radius: 4px;
}

.group-item p {
  margin: 5px 0;
}
</style>