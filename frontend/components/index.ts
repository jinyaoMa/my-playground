import { App } from "vue";
import "@jinyaoma/my-playground-icons/dist/mp-icon.scss";
import "./theme/style.scss";
import MpOpenLink from "./MpOpenLink";
import MpTabbar from "./MpTabbar";

const components = [MpOpenLink, MpTabbar];

export default {
  install(app: App) {
    components.forEach((component) => {
      app.use(component);
    });
  },
  ...components,
};
