import { createApp } from "vue";
import App from "./App.vue";
import i18n from "./i18n";
import router from "./router";

import Components from "@jinyaoma/my-playground-components";

const app = createApp(App).use(i18n).use(router);
app.use(Components);

app.mount("#app");
