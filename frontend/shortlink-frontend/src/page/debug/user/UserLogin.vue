<script setup lang="ts">
import { ref } from "vue";
import { loginUser } from "@/api/user/userApi.ts";

const form = ref({ username: "", password: "" });
const message = ref("");

const submit = async () => {
  try {
    await loginUser(form.value);
    message.value = "登录成功！";
  } catch (err) {
    message.value = "登录失败：" + err.response.data.message;
  }
};
</script>

<template>
  <div>
    <h2>用户登录</h2>
    <form @submit.prevent="submit">
      <input v-model="form.username" placeholder="用户名" />
      <input v-model="form.password" type="password" placeholder="密码" />
      <button type="submit">登录</button>
    </form>
    <p>{{ message }}</p>
  </div>
</template>
