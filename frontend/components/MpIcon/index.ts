import { App } from "vue";
import MpIcon from "./MpIcon.vue";

MpIcon.install = (app: App) => {
  app.component(`${MpIcon.__name}`, MpIcon);
};

export default MpIcon;
