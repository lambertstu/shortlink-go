<script setup lang="ts">
import { ref } from "vue";
import { getUserInfo } from "@/api/user/userApi.ts";

const username = ref("");
const userInfo = ref(null);
const message = ref("");

const fetchUserInfo = async () => {
  try {
    const response = await getUserInfo(username.value);
    userInfo.value = response.data;
    message.value = "获取成功";
  } catch (err) {
    message.value = "查询失败：" + err.response.data.message;
  }
};
</script>

<template>
  <div>
    <h2>获取用户信息</h2>
    <input v-model="username" placeholder="用户名" />
    <button @click="fetchUserInfo">查询</button>
    <p>{{ message }}</p>
    <pre v-if="userInfo">{{ userInfo }}</pre>
  </div>
</template>
