<script setup lang="ts">
import { ref } from "vue";
import { logoutUser } from "@/api/user/userApi.ts";

const form = ref({ token: "", username: "" });
const message = ref("");

const submit = async () => {
  try {
    await logoutUser(form.value);
    message.value = "登出成功！";
  } catch (err) {
    message.value = "登出失败：" + err.response.data.message;
  }
};
</script>

<template>
  <div>
    <h2>用户登出</h2>
    <form @submit.prevent="submit">
      <input v-model="form.token" placeholder="用户 Token" />
      <input v-model="form.username" placeholder="用户名" />
      <button type="submit">登出</button>
    </form>
    <p>{{ message }}</p>
  </div>
</template>
