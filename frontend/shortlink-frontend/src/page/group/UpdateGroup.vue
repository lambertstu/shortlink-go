<template>
  <div>
    <h2>更新群组</h2>
    <form @submit.prevent="handleUpdateGroup">
      <div>
        <label for="gid">群组 ID:</label>
        <input v-model="groupData.gid" id="gid" required />
      </div>
      <div>
        <label for="name">群组名称:</label>
        <input v-model="groupData.name" id="name" required />
      </div>
      <div>
        <label for="username">用户名:</label>
        <input v-model="groupData.username" id="username" required />
      </div>
      <button type="submit">更新群组</button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { updateGroup } from "@/api/group/groupApi";
import { useRouter } from "vue-router";

const groupData = ref({
  gid: "",
  name: "",
  username: "",
});

const router = useRouter();

const handleUpdateGroup = async () => {
  try {
    const response = await updateGroup(groupData.value);
    console.log("群组更新成功:", response);
    router.push("/group/info"); // 重定向到群组信息页面
  } catch (error) {
    console.error("群组更新失败:", error);
  }
};
</script>
