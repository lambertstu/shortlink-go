<template>
  <div class="shortlink-container">
    <h2>更新短链接</h2>
    <div class="input-group">
      <label>GID：</label>
      <input v-model="gid" placeholder="请输入 GID" />
    </div>
    <div class="input-group">
      <label>短链接标识：</label>
      <input v-model="shortUri" placeholder="请输入 ShortUri" />
    </div>
    <div class="input-group">
      <label>原始链接：</label>
      <input v-model="originUrl" placeholder="请输入原始链接" />
    </div>
    <div class="input-group">
      <label>描述：</label>
      <input v-model="describe" placeholder="请输入描述信息" />
    </div>
    <button @click="updateLink">更新</button>

    <p v-if="message">{{ message }}</p>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { updateShortLink } from "@/api/shortlink/shortlinkApi";

export default defineComponent({
  name: "UpdateShortLink",
  setup() {
    const gid = ref("");
    const shortUri = ref("");
    const originUrl = ref("");
    const describe = ref("");
    const message = ref("");

    const updateLink = async () => {
      try {
        await updateShortLink({
          gid: gid.value,
          short_uri: shortUri.value,
          origin_url: originUrl.value,
          describe: describe.value,
          favicon: "",
          clickNum: 0,
          totalPv: 0,
          totalUv: 0,
          totalUip: 0,
          todayPv: 0,
          todayUv: 0,
          todayUip: 0,
        });
        message.value = "短链接更新成功！";
      } catch (error) {
        message.value = "更新失败：" + error;
      }
    };

    return {
      gid,
      shortUri,
      originUrl,
      describe,
      message,
      updateLink,
    };
  },
});
</script>
