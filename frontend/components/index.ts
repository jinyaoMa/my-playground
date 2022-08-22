import { App } from "vue";
import OpenLink from "./open-link";
import Tabbar from "./tabbar";

const components = [OpenLink, Tabbar];

export const MpOpenLink = OpenLink;
export const MpTabbar = Tabbar;

export default {
  install(app: App) {
    components.forEach((component) => {
      app.use(component);
    });
  },
  ...components,
};
