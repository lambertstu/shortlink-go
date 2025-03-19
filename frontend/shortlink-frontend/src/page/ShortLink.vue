<template>
  <div class="shortlink-container">
    <h2>短链接查询</h2>
    <div class="input-group">
      <label>GID：</label>
      <input v-model="gid" placeholder="请输入 GID" />
    </div>
    <div class="input-group">
      <label>Page：</label>
      <input v-model.number="page" type="number" min="1" placeholder="页码" />
    </div>
    <button @click="loadShortLinks">查询</button>

    <div v-if="shortLinks.length > 0">
      <h3>查询结果：</h3>
      <pre>{{ shortLinks }}</pre>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { fetchShortLinks } from "@/api/shortlink/shortlinkApi";

export default defineComponent({
  name: "ShortLink",
  setup() {
    const gid = ref("TrvVab"); // 默认 GID
    const page = ref(1); // 默认页码
    const shortLinks = ref<any[]>([]); // 存储短链接数据

    const loadShortLinks = async () => {
      try {
        const response = await fetchShortLinks({
          gid: gid.value,
          page: page.value,
          size: 2, // 每页 2 条
          orderTag: 1,
        });

        if (response.data.code === "0" && response.data.data) {
          shortLinks.value = response.data.data.list; // 只获取 list 数组
        } else {
          console.error("查询失败，后端返回错误:", response.data.message);
        }
      } catch (error) {
        console.error("查询失败：", error);
      }
    };

    return {
      gid,
      page,
      shortLinks,
      loadShortLinks,
    };
  },
});
</script>

<style scoped>
.shortlink-container {
  max-width: 500px;
  margin: 20px auto;
  padding: 20px;
  border: 1px solid #ddd;
  border-radius: 5px;
}
.input-group {
  margin-bottom: 10px;
}
input {
  padding: 5px;
  width: 100%;
  border: 1px solid #ccc;
  border-radius: 3px;
}
button {
  padding: 8px 12px;
  background: #007bff;
  color: #fff;
  border: none;
  border-radius: 3px;
  cursor: pointer;
}
button:hover {
  background: #0056b3;
}
</style>
