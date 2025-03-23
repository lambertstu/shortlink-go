<script lang="ts" setup>
import { reactive, computed, toRefs } from 'vue';
import { UserOutlined, LockOutlined } from '@ant-design/icons-vue';
import { useRouter } from 'vue-router';
import { loginUser } from '@/api/user/userApi';

interface FormState {
  username: string;
  password: string;
  remember: boolean;
}

// 接收外部传入的用户名和密码
const props = defineProps<{
  username: string | null;
  password: string | null;
}>();

const formState = reactive<FormState>({
  username: props.username || '',
  password: props.password || '',
  remember: true,
});

const router = useRouter();

const onFinish = async (values: any) => {
  try {
    const response = await loginUser({
      username: values.username,
      password: values.password,
    });
    console.log('登录成功:', response.data);

    localStorage.setItem('username', values.username);
    localStorage.setItem('token', response.data.data.token);

    router.push('/shortlink');
  } catch (error) {
    console.error('登录失败:', error.response?.data?.message || error.message);
  }
};

const onFinishFailed = (errorInfo: any) => {
  console.log('表单验证失败:', errorInfo);
};

const disabled = computed(() => {
  return !(formState.username && formState.password);
});
</script>

<template>
  <div class="login-container">
    <div class="welcome-message">
      <h1>Hello Again!</h1>
      <p>Welcome back, you've been missed!</p>
    </div>

    <a-form
      :model="formState"
      name="normal_login"
      class="login-form"
      @finish="onFinish"
      @finishFailed="onFinishFailed"
    >
      <a-form-item
        label="Username"
        name="username"
        :rules="[{ required: true, message: 'Please input your username!' }]"
      >
        <a-input v-model:value="formState.username">
          <template #prefix>
            <UserOutlined class="site-form-item-icon" />
          </template>
        </a-input>
      </a-form-item>

      <a-form-item
        label="Password"
        name="password"
        :rules="[{ required: true, message: 'Please input your password!' }]"
      >
        <a-input-password v-model:value="formState.password">
          <template #prefix>
            <LockOutlined class="site-form-item-icon" />
          </template>
        </a-input-password>
      </a-form-item>

      <a-form-item>
        <a-form-item name="remember" no-style>
          <a-checkbox v-model:checked="formState.remember">Remember me</a-checkbox>
        </a-form-item>
        <a class="login-form-forgot" href="">Forgot password</a>
      </a-form-item>

      <a-form-item>
        <a-button :disabled="disabled" type="primary" html-type="submit" class="login-form-button">
          Log in
        </a-button>
        Or
        <a href="javascript:void(0);" @click="$emit('register')">register now!</a>
      </a-form-item>
    </a-form>
  </div>
</template>

<style scoped>
.login-container {
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

.welcome-message h1 {
  font-size: 24px;
  font-weight: bold;
  color: #333;
}

.welcome-message p {
  font-size: 14px;
  color: #666;
}

.login-form {
  max-width: 300px;
  width: 100%;
}

.login-form-forgot {
  float: right;
}

.login-form-button {
  width: 100%;
}

/* 修改链接颜色 */
.login-form a {
  color: #40a9ff; /* 浅蓝色 */
  text-decoration: none; /* 可选，去除下划线 */
  font-weight: bold; /* 可选，加粗 */
  background: transparent !important;
}

/* 悬停时颜色变化 */
.login-form a:hover {
  color: #1960a2; /* 深一点的蓝色 */
  background: transparent !important;
}
</style>
