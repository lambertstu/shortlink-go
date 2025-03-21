<script setup lang="ts">
import { ref } from "vue";
import { checkUserExist } from "@/api/user/userApi.ts";

const username = ref("");
const exists = ref<boolean | null>(null);
const message = ref("");

const checkExistence = async () => {
  try {
    const response = await checkUserExist(username.value);
    exists.value = response.data.exists;
    message.value = exists.value ? "用户已存在" : "用户不存在";
  } catch (err) {
    message.value = "查询失败：" + err.response.data.message;
  }
};
</script>

<template>
  <div>
    <h2>查询用户是否存在</h2>
    <input v-model="username" placeholder="用户名" />
    <button @click="checkExistence">查询</button>
    <p>{{ message }}</p>
  </div>
</template>
