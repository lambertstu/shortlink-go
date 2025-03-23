<template>
  <div class="register-container">
    <div class="welcome-message">
      <h1>Join Us!</h1>
      <p>Create your account to get started</p>
    </div>

    <a-form
      :model="formState"
      name="register_form"
      class="register-form"
      @finish="onRegister"
      @finishFailed="onRegisterFailed"
    >
      <a-form-item
        label="Username"
        name="username"
        :rules="[{ required: true, message: 'Please input your username!' }]"
      >
        <a-input v-model:value="formState.username" />
      </a-form-item>

      <a-form-item
        label="Password"
        name="password"
        :rules="[{ required: true, message: 'Please input your password!' }]"
      >
        <a-input-password v-model:value="formState.password" />
      </a-form-item>

      <a-form-item
        label="Email"
        name="email"
        :rules="[
          { required: true, type: 'email', message: 'Please enter a valid email!' }
        ]"
      >
        <a-input v-model:value="formState.email" />
      </a-form-item>


      <a-form-item
        label="Phone"
        name="phone"
        :rules="[
          { required: true, message: 'Please enter your phone number!' },
          { pattern: /^1[3-9]\d{9}$/, message: 'Please enter a valid phone number!' }
        ]"
      >
        <a-input v-model:value="formState.phone" />
      </a-form-item>

      <a-form-item
        label="Real Name"
        name="realName"
        :rules="[{ required: true, message: 'Please enter your real name!' }]"
      >
        <a-input v-model:value="formState.realName" />
      </a-form-item>

      <a-form-item>
        <a-button type="primary" html-type="submit" class="register-form-button">
          Register
        </a-button>
      </a-form-item>

      <a-form-item>
        Already have an account?
        <a href="javascript:void(0);" @click="$emit('login')">Login here</a>
      </a-form-item>
    </a-form>
  </div>
</template>

<script lang="ts" setup>
import { reactive } from 'vue';
import { registerUser } from '@/api/user/userApi';

const formState = reactive({
  username: '',
  email: '',
  password: '',
  phone: '',
  realName: '',
});

const onRegister = async (values: any) => {
  try {
    // 调用后端 API 进行注册
    const response = await registerUser(values);

    if (response.data?.success) {
      console.log("User registered successfully");
      // 自动切换到登录界面并填充表单
      const { username, password } = values;
      $emit('login', { username, password });
    } else {
      console.error("Registration failed", response.data?.message);
    }
  } catch (error) {
    console.error("Registration error", error);
  }
};

const onRegisterFailed = (errorInfo: any) => {
  console.log("Registration failed:", errorInfo);
};
</script>

<style scoped>
.register-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
  padding: 20px;
}

.welcome-message {
  text-align: center;
  margin-bottom: 20px;
}

.register-form {
  max-width: 300px;
  width: 100%;
}

.register-form-button {
  width: 100%;
}

.register-container a {
  color: #40a9ff;
  font-weight: bold;
  text-decoration: none;
}

.register-container a:hover {
  color: #1960a2;
}
</style>
