import { App } from "vue";
import MpTile from "./MpTile.vue";
import MpTileItem from "./MpTileItem.vue";

MpTile.install = (app: App) => {
  app.component(`${MpTile.__name}`, MpTile);
  app.component(`${MpTileItem.__name}`, MpTileItem);
};

export default MpTile;
