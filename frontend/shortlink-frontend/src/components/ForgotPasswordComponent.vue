<template>
  <div class="forgot-password-container">
    <div class="welcome-message">
      <h1>Reset Password</h1>
      <p>Please enter your new password</p>
    </div>

    <a-form :model="formState" layout="vertical" @finish="handleSubmit">
      <a-form-item label="Username" name="username">
        <a-input v-model:value="formState.username" disabled />
      </a-form-item>
      <a-form-item label="New Password" name="password">
        <a-input-password v-model:value="formState.password" />
      </a-form-item>
      <a-form-item label="Confirm Password" name="confirmPassword">
        <a-input-password v-model:value="formState.confirmPassword" />
      </a-form-item>

      <a-form-item>
        <a-button type="primary" html-type="submit" class="submit-button">
          Reset Password
        </a-button>
        <div class="back-to-login">
          <a href="javascript:void(0);" @click="$emit('back-to-login')">Back to Login</a>
        </div>
      </a-form-item>
    </a-form>
  </div>
</template>

<script lang="ts" setup>
import { reactive } from 'vue';
import { message } from 'ant-design-vue';
import { upsertUser } from '@/api/user/userApi';

interface FormState {
  username: string;
  password: string;
  confirmPassword: string;
}

const props = defineProps<{
  username: string;
}>();

const emit = defineEmits(['back-to-login']);

const formState = reactive<FormState>({
  username: props.username,
  password: '',
  confirmPassword: '',
});

const handleSubmit = async () => {
  if (formState.password !== formState.confirmPassword) {
    message.error('两次输入的密码不一致');
    return;
  }

  try {
    const response = await upsertUser({
      username: formState.username,
      password: formState.password,
      realName: '',
      phone: '',
      email: ''
    });
    
    if (response.data.success) {
      message.success('密码重置成功');
      emit('back-to-login');
    } else {
      message.error('密码重置失败: ' + response.data.message);
    }
  } catch (error: any) {
    message.error('请求失败: ' + (error?.message || '未知错误'));
  }
};
</script>

<style scoped>
.forgot-password-container {
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

.submit-button {
  width: 100%;
  margin-bottom: 16px;
}

.back-to-login {
  text-align: center;
}

.back-to-login a {
  color: #40a9ff;
  text-decoration: none;
  font-weight: bold;
}

.back-to-login a:hover {
  color: #1960a2;
}
</style> 