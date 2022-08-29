import { App } from "vue";
import MpOpenLink from "./MpOpenLink.vue";

MpOpenLink.install = (app: App) => {
  app.component(MpOpenLink.__name, MpOpenLink);
  console.log(`Components [${MpOpenLink.__name}] loaded`);
};

export default MpOpenLink;
