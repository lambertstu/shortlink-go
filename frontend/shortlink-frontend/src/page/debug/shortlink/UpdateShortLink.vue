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
    <div class="input-group">
      <label>Favicon：</label>
      <input v-model="favicon" placeholder="请输入 Favicon 链接" />
    </div>
    <div class="input-group">
      <label>点击次数：</label>
      <input v-model.number="clickNum" type="number" />
    </div>
    <div class="input-group">
      <label>总 PV：</label>
      <input v-model.number="totalPv" type="number" />
    </div>
    <div class="input-group">
      <label>总 UV：</label>
      <input v-model.number="totalUv" type="number" />
    </div>
    <div class="input-group">
      <label>总 UIP：</label>
      <input v-model.number="totalUip" type="number" />
    </div>
    <div class="input-group">
      <label>今日 PV：</label>
      <input v-model.number="todayPv" type="number" />
    </div>
    <div class="input-group">
      <label>今日 UV：</label>
      <input v-model.number="todayUv" type="number" />
    </div>
    <div class="input-group">
      <label>今日 UIP：</label>
      <input v-model.number="todayUip" type="number" />
    </div>

    <button @click="updateLink">更新</button>

    <p v-if="message">{{ message }}</p>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { updateShortLink } from "@/api/shortlink/shortlinkApi.ts";

export default defineComponent({
  name: "UpdateShortLink",
  setup() {
    const gid = ref("");
    const shortUri = ref("");
    const originUrl = ref("");
    const describe = ref("");
    const favicon = ref("");
    const clickNum = ref(0);
    const totalPv = ref(0);
    const totalUv = ref(0);
    const totalUip = ref(0);
    const todayPv = ref(0);
    const todayUv = ref(0);
    const todayUip = ref(0);
    const message = ref("");

    const updateLink = async () => {
      try {
        await updateShortLink({
          gid: gid.value,
          origin_url: originUrl.value,
          short_uri: shortUri.value,
          describe: describe.value,
          favicon: favicon.value,
          clickNum: clickNum.value,
          totalPv: totalPv.value,
          totalUv: totalUv.value,
          totalUip: totalUip.value,
          todayPv: todayPv.value,
          todayUv: todayUv.value,
          todayUip: todayUip.value,
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
      favicon,
      clickNum,
      totalPv,
      totalUv,
      totalUip,
      todayPv,
      todayUv,
      todayUip,
      message,
      updateLink,
    };
  },
});
</script>

<style scoped>
.shortlink-container {
  width: 400px;
  margin: 20px auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 5px;
  background-color: #f9f9f9;
}

.input-group {
  margin-bottom: 10px;
}

label {
  display: block;
  font-weight: bold;
}

input {
  width: 100%;
  padding: 5px;
  margin-top: 5px;
  border: 1px solid #ccc;
  border-radius: 3px;
}

button {
  width: 100%;
  padding: 10px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

button:hover {
  background-color: #0056b3;
}

p {
  color: green;
  font-weight: bold;
  text-align: center;
  margin-top: 10px;
}
</style>
