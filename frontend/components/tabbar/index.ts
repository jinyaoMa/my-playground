import { App } from "vue";
import Tabbar from "./tabbar.vue";
import TabbarItem from "./tabbar-item.vue";

Tabbar.install = (app: App) => {
  app.component(Tabbar.name, Tabbar);
};

Tabbar.TabbarItem = TabbarItem;

export default Tabbar;
