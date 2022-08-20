import { createApp } from "vue";
import App from "./App.vue";
import i18n from "@/i18n";
import router from "@/router";

import OpenLink from "@jinyaoma/my-playground-components/open-link";

const app = createApp(App).use(i18n).use(router);
app.use(OpenLink);

app.mount("#app");
