import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";
import Home from "@/views/Home.vue";
import svgHome from "@/assets/images/home.svg";
import About from "@/views/About.vue";
import svgSettings from "@/assets/images/settings.svg";

export const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "Home",
    component: Home,
    meta: {
      icon: svgHome,
    },
  },
  {
    path: "/about",
    name: "About",
    component: About,
    meta: {
      icon: svgSettings,
    },
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
