import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";
import Home from "@/views/Home.vue";
import Settings from "@/views/Settings.vue";

export const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "Home",
    component: Home,
    meta: {
      icon: "home",
    },
  },
  {
    path: "/settings",
    name: "Settings",
    component: Settings,
    meta: {
      icon: "settings",
    },
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
