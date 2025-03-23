<script setup lang="ts">
import LoginComponent from "@/components/LoginComponent.vue";
</script>

<template>
  <div class="container">
    <div class="centered-box">
      <div class="left-section">
        <img src="@/assets/login.png" alt="Login Image" class="login-image">
      </div>
      <div class="right-section">
        <LoginComponent />
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
  background: #3494E6;  /* fallback for old browsers */
  background: -webkit-linear-gradient(to right, #EC6EAD, #3494E6);  /* Chrome 10-25, Safari 5.1-6 */
  background: linear-gradient(to right, #EC6EAD, #3494E6); /* W3C, IE 10+/ Edge, Firefox 16+, Chrome 26+, Opera 12+, Safari 7+ */
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
  width: 70%; /* 占据 left-section 的 80% */
  height: auto;
  object-fit: contain; /* 保持图片比例 */
}
</style>
