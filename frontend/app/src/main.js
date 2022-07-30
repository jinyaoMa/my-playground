import { createApp } from "vue";
import { createI18n } from "vue-i18n";
import App from "./App.vue";
import UIKit from "@frontend/ui-kit";
import router from "@/router";
import messages from "@intlify/vite-plugin-vue-i18n/messages";

createApp(App)
  .use(UIKit)
  .use(router)
  .use(
    createI18n({
      legacy: false,
      locale: "en",
      messages,
    })
  )
  .mount("#app");
