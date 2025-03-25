<template>
  <a-modal
    v-model:visible="visible"
    title="修改用户信息"
    ok-text="提交"
    cancel-text="取消"
    @ok="handleSubmit"
    @cancel="handleCancel"
  >
    <a-form :model="formState" layout="vertical">
      <a-form-item label="用户名" name="username">
        <a-input v-model:value="formState.username" disabled />
      </a-form-item>
      <a-form-item label="密码" name="password">
        <a-input-password v-model:value="formState.password" />
      </a-form-item>
      <a-form-item label="真实姓名" name="realName">
        <a-input v-model:value="formState.realName" />
      </a-form-item>
      <a-form-item label="电话" name="phone">
        <a-input v-model:value="formState.phone" />
      </a-form-item>
      <a-form-item label="邮箱" name="email">
        <a-input v-model:value="formState.email" />
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script lang="ts" setup>
import { ref, reactive } from 'vue';
import { message } from 'ant-design-vue';
import { upsertUser } from '@/api/user/userApi'; // 导入更新用户信息接口

// 定义表单数据类型
interface FormState {
  username: string;
  password: string;
  realName: string;
  phone: string;
  email: string;
}

// 表单状态
const formState = reactive<FormState>({
  username: '',
  password: '',
  realName: '',
  phone: '',
  email: '',
});

// 控制弹窗显示
const visible = ref(false);

// 定义事件
const emit = defineEmits(['update-success']);

// 打开弹窗
const showModal = (user: Partial<FormState>) => {
  Object.assign(formState, user);
  visible.value = true;
};

// 关闭弹窗
const handleCancel = () => {
  visible.value = false;
};

// 提交表单
const handleSubmit = async () => {
  try {
    // 调用更新用户信息接口
    const response = await upsertUser(formState);
    if (response.data.success) {
      message.success('修改成功'); // 弹出成功提示
      visible.value = false; // 关闭弹窗
      emit('update-success'); // 通知父组件数据已更新
    } else {
      message.error('修改失败: ' + response.data.message);
    }
  } catch (error: any) {
    message.error('请求失败: ' + (error?.message || '未知错误'));
  }
};

// 暴露方法，供父组件调用
defineExpose({ showModal });
</script>

<style scoped>
/* 样式可以根据需要调整 */
</style>
