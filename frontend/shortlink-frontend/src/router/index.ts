import {createRouter, createWebHistory} from "vue-router";
import ShortLink from "@/page/shortlink/ShortLink.vue";
import CreateShortLink from "@/page/shortlink/CreateShortLink.vue";
import UpdateShortLink from "@/page/shortlink/UpdateShortLink.vue";
import DeleteShortLink from "@/page/shortlink/DeleteShortLink.vue";
import UserRegister from "@/page/user/UserRegister.vue";
import UserLogin from "@/page/user/UserLogin.vue";
import UserLogout from "@/page/user/UserLogout.vue";
import UserExist from "@/page/user/UserExist.vue";
import UserLoginStatus from "@/page/user/UserLoginStatus.vue";
import UserInfo from "@/page/user/UserInfo.vue";
import UserUpsert from "@/page/user/UserUpsert.vue";
import UserDelete from "@/page/user/UserDelete.vue";
import CreateGroup from "@/page/group/CreateGroup.vue";
import UpdateGroup from "@/page/group/UpdateGroup.vue";
import DeleteGroup from "@/page/group/DeleteGroup.vue";
import GroupInfo from "@/page/group/GroupInfo.vue";


const routes = [
  {path: "/", redirect: "/shortlink"},
  {path: "/shortlink/page", component: ShortLink},
  {path: "/shortlink/create", component: CreateShortLink},
  {path: "/shortlink/update", component: UpdateShortLink},
  {path: "/shortlink/delete", component: DeleteShortLink},
  {path: "/user/register", component: UserRegister},
  {path: "/user/login", component: UserLogin},
  {path: "/user/logout", component: UserLogout},
  {path: "/user/exist", component: UserExist},
  {path: "/user/loginStatus", component: UserLoginStatus},
  {path: "/user/info", component: UserInfo},
  {path: "/user/upsert", component: UserUpsert},
  {path: "/user/delete", component: UserDelete},
  {path: "/group/create", component: CreateGroup},
  {path: "/group/update", component: UpdateGroup},
  {path: "/group/delete", component: DeleteGroup},
  {path: "/group/info", component: GroupInfo},
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
