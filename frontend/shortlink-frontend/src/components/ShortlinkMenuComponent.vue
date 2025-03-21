<template>
  <a-menu class= "top_menu" v-model:selectedKeys="current" mode="horizontal" :items="items" />
</template>
<script lang="ts" setup>
import { h, ref } from 'vue';
import { MailOutlined, AppstoreOutlined } from '@ant-design/icons-vue';
import type { MenuProps } from 'ant-design-vue';
const current = ref<string[]>(['longShortlink']);
const items = ref<MenuProps['items']>([
  {
    key: 'longShortlink',
    label: '长期短链接',
    title: 'Navigation One',
  },
  {
    key: 'tempShortlink',
    label: '短期短链接',
    title: 'Navigation Two',
  },
]);
</script>

<style>
.top_menu{
  background-color: #efefef;
  margin-left: 150px;
}
</style>
