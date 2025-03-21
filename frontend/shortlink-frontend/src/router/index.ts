import {createRouter, createWebHistory} from "vue-router";
import ShortLink from "@/page/debug/shortlink/ShortLink.vue";
import CreateShortLink from "@/page/debug/shortlink/CreateShortLink.vue";
import UpdateShortLink from "@/page/debug/shortlink/UpdateShortLink.vue";
import DeleteShortLink from "@/page/debug/shortlink/DeleteShortLink.vue";
import UserRegister from "@/page/debug/user/UserRegister.vue";
import UserLogin from "@/page/debug/user/UserLogin.vue";
import UserLogout from "@/page/debug/user/UserLogout.vue";
import UserExist from "@/page/debug/user/UserExist.vue";
import UserLoginStatus from "@/page/debug/user/UserLoginStatus.vue";
import UserInfo from "@/page/debug/user/UserInfo.vue";
import UserUpsert from "@/page/debug/user/UserUpsert.vue";
import UserDelete from "@/page/debug/user/UserDelete.vue";
import CreateGroup from "@/page/debug/group/CreateGroup.vue";
import UpdateGroup from "@/page/debug/group/UpdateGroup.vue";
import DeleteGroup from "@/page/debug/group/DeleteGroup.vue";
import GroupInfo from "@/page/debug/group/GroupInfo.vue";

import LoginPage from "@/page/LoginPage.vue";
const routes = [
  {path: "/", redirect: "/login"},
  {path: "/login", component: LoginPage},
  {path: "/debug/shortlink/page", component: ShortLink},
  {path: "/debug/shortlink/create", component: CreateShortLink},
  {path: "/debug/shortlink/update", component: UpdateShortLink},
  {path: "/debug/shortlink/delete", component: DeleteShortLink},
  {path: "/debug/user/register", component: UserRegister},
  {path: "/debug/user/login", component: UserLogin},
  {path: "/debug/user/logout", component: UserLogout},
  {path: "/debug/user/exist", component: UserExist},
  {path: "/debug/user/loginStatus", component: UserLoginStatus},
  {path: "/debug/user/info", component: UserInfo},
  {path: "/debug/user/upsert", component: UserUpsert},
  {path: "/debug/user/delete", component: UserDelete},
  {path: "/debug/group/create", component: CreateGroup},
  {path: "/debug/group/update", component: UpdateGroup},
  {path: "/debug/group/delete", component: DeleteGroup},
  {path: "/debug/group/info", component: GroupInfo},


];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
