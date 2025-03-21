<template>
  <div>
    <h2>创建群组</h2>
    <form @submit.prevent="handleCreateGroup">
      <div>
        <label for="name">群组名称:</label>
        <input v-model="groupData.name" id="name" required />
      </div>
      <div>
        <label for="username">用户名:</label>
        <input v-model="groupData.username" id="username" required />
      </div>
      <button type="submit">创建群组</button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { createGroup } from "@/api/group/groupApi.ts";
import { useRouter } from "vue-router";

const groupData = ref({
  name: "",
  username: "",
});

const router = useRouter();

const handleCreateGroup = async () => {
  try {
    const response = await createGroup(groupData.value);
    console.log("群组创建成功:", response);
    router.push("/group/info"); // 重定向到群组信息页面
  } catch (error) {
    console.error("群组创建失败:", error);
  }
};
</script>
