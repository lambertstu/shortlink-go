<template>
  <div>
    <h2>群组信息</h2>
    <div v-if="groupInfo">
      <p><strong>群组 ID:</strong> {{ groupInfo.gid }}</p>
      <p><strong>群组名称:</strong> {{ groupInfo.name }}</p>
      <p><strong>用户名:</strong> {{ groupInfo.username }}</p>
    </div>
    <div v-else>
      <p>加载群组信息中...</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { getGroupInfo } from "@/api/group/groupApi.ts";
import { useRoute } from "vue-router";

const groupInfo = ref(null);
const route = useRoute();

onMounted(async () => {
  try {
    const gid = route.query.gid as string;
    const response = await getGroupInfo(gid);
    groupInfo.value = response.data;
  } catch (error) {
    console.error("获取群组信息失败:", error);
  }
});
</script>
