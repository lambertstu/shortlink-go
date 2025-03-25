<script setup lang="ts">
import { ref } from "vue";
import LoginComponent from "@/components/LoginComponent.vue";
import RegisterComponent from "@/components/RegisterComponent.vue";

const isRegistering = ref(false);
const formData = ref<{ username: string; password: string } | null>(null);

const switchToRegister = () => {
  isRegistering.value = true;
};

const switchToLogin = (data: { username: string; password: string }) => {
  isRegistering.value = false;
  formData.value = data;
};
</script>

<template>
  <div class="container">
    <div class="centered-box">
      <div class="left-section">
        <img src="@/assets/login.png" alt="Login Image" class="login-image">
      </div>
      <div class="right-section">
        <LoginComponent v-if="!isRegistering"
                        :username="formData?.username ?? null"
                        :password="formData?.password ?? null"
                        @register="switchToRegister" />
        <RegisterComponent v-else @login="switchToLogin" />
      </div>
    </div>
  </div>
</template>

<style scoped>
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  width: 100vw;
  background: linear-gradient(to right, #ec6ead, #3494e6);
}

.centered-box {
  display: flex;
  width: 65vw;
  height: 80vh;
  background-color: white;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  border-radius: 8px;
  padding: 20px;
}

.left-section {
  flex: 5;
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
}

.right-section {
  flex: 4;
  display: flex;
  justify-content: center;
  align-items: center;
}

.login-image {
  width: 70%;
  height: auto;
  object-fit: contain;
}
</style>
