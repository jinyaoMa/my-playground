import { App } from "vue";
import OpenLink from "./open-link";

const components = [OpenLink];

export const MpOpenLink = OpenLink;

export default {
  install(app: App) {
    components.forEach((component) => {
      app.use(component);
    });
  },
  ...components,
};
