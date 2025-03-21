<script setup lang="ts">
import LoginComponent from "@/components/userComponent/LoginComponent.vue";
</script>

<template>
  <div class="container">
    <div class="centered-box">
      <div class="left-section">
        <!-- 预留放置UI图片 -->
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
  width: 80vw;
  height: 80vh;
  background-color: white;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  border-radius: 8px;
  padding: 20px;
}

.left-section {
  flex: 5;
  /* 预留放置UI图片 */
}

.right-section {
  flex: 4;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
