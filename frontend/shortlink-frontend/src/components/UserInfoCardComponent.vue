<template>
  <div id="components-popover-demo-placement">
    <a-popover placement="bottomLeft">
      <template #content>
        <div v-if="loading" class="loading-container">
          <a-spin /> <!-- 加载动画 -->
        </div>
        <div v-else-if="userInfo" class="user-info-content">
          <div class="user-info-item">
            <UserOutlined class="info-icon" />
            <span class="info-label">用户名:</span>
            <span class="info-value">{{ userInfo.username }}</span>
          </div>
          <div class="user-info-item">
            <IdcardOutlined class="info-icon" />
            <span class="info-label">真实姓名:</span>
            <span class="info-value">{{ userInfo.realName }}</span>
          </div>
          <div class="user-info-item">
            <PhoneOutlined class="info-icon" />
            <span class="info-label">电话:</span>
            <span class="info-value">{{ userInfo.phone }}</span>
          </div>
          <div class="user-info-item">
            <MailOutlined class="info-icon" />
            <span class="info-label">邮箱:</span>
            <span class="info-value">{{ userInfo.email }}</span>
          </div>
        </div>
        <div v-else class="error-message">
          <span>无法加载用户信息</span>
        </div>
      </template>
      <UserOutlined class="user-icon" @click="fetchUserInfo" />
    </a-popover>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue';
import {
  UserOutlined,
  IdcardOutlined,
  PhoneOutlined,
  MailOutlined,
} from '@ant-design/icons-vue';
import { getUserInfo } from '@/api/user/userApi'; // 导入获取用户信息的接口

// 定义用户信息类型
interface UserInfo {
  username: string;
  realName: string;
  phone: string;
  email: string;
}

// 用户信息状态
const userInfo = ref<UserInfo | null>(null);
const loading = ref(false); // 加载状态

// 获取用户信息的方法
const fetchUserInfo = async () => {
  try {
    loading.value = true; // 开始加载
    const username = localStorage.getItem('username');
    if (!username) {
      console.error('未找到用户名');
      return;
    }

    const response = await getUserInfo(username);
    if (response.data.success) {
      userInfo.value = response.data.data; // 存储用户信息
    } else {
      console.error('获取用户信息失败:', response.data.message);
    }
  } catch (error) {
    console.error('请求失败:', error);
  } finally {
    loading.value = false; // 结束加载
  }
};

// 组件挂载时自动获取用户信息
onMounted(() => {
  fetchUserInfo();
});
</script>

<style scoped>
#components-popover-demo-placement .ant-btn {
  width: 70px;
  text-align: center;
  padding: 0;
  margin-right: 8px;
  margin-bottom: 8px;
}

.user-icon {
  font-size: 24px;
  color: #1973f4;
  margin-right: 20px;
  cursor: pointer;
  transition: color 0.3s ease;
}

.user-icon:hover {
  color: #40a9ff; /* 悬停时颜色变化 */
}

.popover-title {
  font-size: 16px;
  font-weight: bold;
  color: #333;
}

.user-info-content {
  padding: 8px 0;
}

.user-info-item {
  display: flex;
  align-items: center;
  margin-bottom: 12px;
}

.user-info-item:last-child {
  margin-bottom: 0;
}

.info-icon {
  font-size: 16px;
  color: #40a9ff;
  margin-right: 8px;
}

.info-label {
  font-weight: 500;
  color: #666;
  margin-right: 8px;
}

.info-value {
  font-weight: 400;
  color: #333;
}

.loading-container {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 16px;
}

.error-message {
  color: #ff4d4f;
  font-weight: 500;
  padding: 16px;
  text-align: center;
}
</style>
