<template>
  <div class="shortlink-container">
    <h2>删除短链接</h2>
    <div class="input-group">
      <label>短链接标识：</label>
      <input v-model="shortUri" placeholder="请输入 ShortUri" />
    </div>
    <div class="input-group">
      <label>原始链接：</label>
      <input v-model="originUrl" placeholder="请输入原始链接" />
    </div>
    <button @click="deleteLink">删除</button>

    <p v-if="message">{{ message }}</p>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { deleteShortLink } from "@/api/shortlink/shortlinkApi";

export default defineComponent({
  name: "DeleteShortLink",
  setup() {
    const shortUri = ref("");
    const originUrl = ref("");
    const message = ref("");

    const deleteLink = async () => {
      if (!shortUri.value || !originUrl.value) {
        message.value = "短链接和原始链接不能为空！";
        return;
      }

      try {
        const response = await deleteShortLink({
          shortUri: shortUri.value,
          origin_url: originUrl.value,
        });
        message.value = "短链接删除成功：" + JSON.stringify(response.data);
      } catch (error: any) {
        console.error("删除失败:", error);
        message.value =
          "删除失败：" + (error.response?.data?.message || "未知错误");
      }
    };

    return {
      shortUri,
      originUrl,
      message,
      deleteLink,
    };
  },
});
</script>
