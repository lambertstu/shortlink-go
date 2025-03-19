import { createRouter, createWebHistory } from "vue-router";
import ShortLink from "@/page/ShortLink.vue";
import CreateShortLink from "@/page/CreateShortLink.vue";
import UpdateShortLink from "@/page/UpdateShortLink.vue";
import DeleteShortLink from "@/page/DeleteShortLink.vue";

const routes = [
  { path: "/", redirect: "/shortlink" },
  { path: "/shortlink/page", component: ShortLink },
  { path: "/shortlink/create", component: CreateShortLink },
  { path: "/shortlink/update", component: UpdateShortLink },
  { path: "/shortlink/delete", component: DeleteShortLink },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
