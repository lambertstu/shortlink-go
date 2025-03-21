<script setup lang="ts">
import { ref } from "vue";
import { checkLoginStatus } from "@/api/user/userApi.ts";

const form = ref({ token: "", username: "" });
const message = ref("");

const checkStatus = async () => {
  try {
    const response = await checkLoginStatus(form.value);
    message.value = response.data.status ? "已登录" : "未登录";
  } catch (err) {
    message.value = "查询失败：" + err.response.data.message;
  }
};
</script>

<template>
  <div>
    <h2>查询用户登录状态</h2>
    <input v-model="form.token" placeholder="用户 Token" />
    <input v-model="form.username" placeholder="用户名" />
    <button @click="checkStatus">查询</button>
    <p>{{ message }}</p>
  </div>
</template>
