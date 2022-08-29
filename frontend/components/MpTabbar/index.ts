import { App } from "vue";
import MpTabbar from "./MpTabbar.vue";
import MpTabbarItem from "./MpTabbarItem.vue";

MpTabbar.install = (app: App) => {
  app.component(MpTabbar.__name, MpTabbar);
};

MpTabbar.MpTabbarItem = MpTabbarItem;

export default MpTabbar;
