<script setup lang="ts">
import { ref } from "vue";
import { deleteUser } from "@/api/user/userApi.ts";

const username = ref("");
const message = ref("");

const submit = async () => {
  try {
    await deleteUser(username.value);
    message.value = "删除成功！";
  } catch (err) {
    message.value = "删除失败：" + err.response.data.message;
  }
};
</script>

<template>
  <div>
    <h2>删除用户</h2>
    <input v-model="username" placeholder="用户名" />
    <button @click="submit">删除</button>
    <p>{{ message }}</p>
  </div>
</template>
