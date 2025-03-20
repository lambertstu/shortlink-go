<template>
  <div>
    <h2>删除群组</h2>
    <form @submit.prevent="handleDeleteGroup">
      <div>
        <label for="gid">群组 ID:</label>
        <input v-model="groupData.gid" id="gid" required />
      </div>
      <div>
        <label for="username">用户名:</label>
        <input v-model="groupData.username" id="username" required />
      </div>
      <button type="submit">删除群组</button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { deleteGroup } from "@/api/group/groupApi";
import { useRouter } from "vue-router";

const groupData = ref({
  gid: "",
  username: "",
});

const router = useRouter();

const handleDeleteGroup = async () => {
  try {
    const response = await deleteGroup(groupData.value);
    console.log("群组删除成功:", response);
    router.push("/group/info"); // 重定向到群组信息页面
  } catch (error) {
    console.error("群组删除失败:", error);
  }
};
</script>
