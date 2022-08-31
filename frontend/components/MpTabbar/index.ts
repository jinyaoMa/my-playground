import { App } from "vue";
import MpTabbar from "./MpTabbar.vue";
import MpTabbarItem from "./MpTabbarItem.vue";

MpTabbar.install = (app: App) => {
  app.component(`${MpTabbar.__name}`, MpTabbar);
  app.component(`${MpTabbarItem.__name}`, MpTabbarItem);
};

export default MpTabbar;
