import { App } from "vue";
import OpenLink from "./open-link.vue";

OpenLink.install = (app: App) => {
  app.component(OpenLink.name, OpenLink);
};

export default OpenLink;
