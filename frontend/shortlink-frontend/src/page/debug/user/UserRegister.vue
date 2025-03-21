<script setup lang="ts">
import { ref } from "vue";
import { registerUser } from "@/api/user/userApi.ts";

const form = ref({ email: "", password: "", phone: "", realName: "", username: "" });
const message = ref("");

const submit = async () => {
  try {
    await registerUser(form.value);
    message.value = "注册成功！";
  } catch (err) {
    message.value = "注册失败：" + err.response.data.message;
  }
};
</script>

<template>
  <div>
    <h2>用户注册</h2>
    <form @submit.prevent="submit">
      <input v-model="form.email" placeholder="邮箱" />
      <input v-model="form.password" type="password" placeholder="密码" />
      <input v-model="form.phone" placeholder="手机号" />
      <input v-model="form.realName" placeholder="真实姓名" />
      <input v-model="form.username" placeholder="用户名" />
      <button type="submit">注册</button>
    </form>
    <p>{{ message }}</p>
  </div>
</template>
