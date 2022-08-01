import { createApp } from "vue";
import App from "./App.vue";
import i18n from "@/i18n";
import router from "@/router";

const app = createApp(App).use(i18n).use(router);

import { MpOpenLink } from "@my-playground/components";
app.use(MpOpenLink);

app.mount("#app");
