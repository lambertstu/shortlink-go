<template>
  <div class="shortlink-container">
    <h2>创建短链接</h2>
    <div class="input-group">
      <label>原始链接：</label>
      <input v-model="originUrl" placeholder="请输入原始链接" />
    </div>
    <div class="input-group">
      <label>GID：</label>
      <input v-model="gid" placeholder="请输入 GID" />
    </div>
    <div class="input-group">
      <label>描述：</label>
      <input v-model="describe" placeholder="请输入描述信息" />
    </div>
    <button @click="createLink">创建</button>

    <p v-if="message">{{ message }}</p>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { createShortLink } from "@/api/shortlink/shortlinkApi";

export default defineComponent({
  name: "CreateShortLink",
  setup() {
    const originUrl = ref("");
    const gid = ref("");
    const describe = ref("");
    const message = ref("");

    const createLink = async () => {
      try {
        await createShortLink({
          origin_url: originUrl.value,
          gid: gid.value,
          describe: describe.value,
        });
        message.value = "短链接创建成功！";
      } catch (error) {
        message.value = "创建失败：" + error;
      }
    };

    return {
      originUrl,
      gid,
      describe,
      message,
      createLink,
    };
  },
});
</script>
