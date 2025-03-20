<script setup lang="ts">
import { ref } from "vue";
import { upsertUser } from "@/api/user/userApi";

const form = ref({ username: "", password: "", realName: "", phone: "", email: "" });
const message = ref("");

const submit = async () => {
  try {
    await upsertUser(form.value);
    message.value = "操作成功！";
  } catch (err) {
    message.value = "操作失败：" + err.response.data.message;
  }
};
</script>

<template>
  <div>
    <h2>用户新增或更新</h2>
    <form @submit.prevent="submit">
      <input v-model="form.username" placeholder="用户名" />
      <input v-model="form.password" type="password" placeholder="密码" />
      <input v-model="form.realName" placeholder="真实姓名" />
      <input v-model="form.phone" placeholder="手机号" />
      <input v-model="form.email" placeholder="邮箱" />
      <button type="submit">提交</button>
    </form>
    <p>{{ message }}</p>
  </div>
</template>
