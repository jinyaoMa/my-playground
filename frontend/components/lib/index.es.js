import { defineComponent, openBlock, createElementBlock, renderSlot } from "vue";
var OpenLink_vue_vue_type_style_index_0_lang = "";
const _hoisted_1 = { class: "openlink" };
const __default__ = defineComponent({
  name: "MpOpenLink"
});
const _sfc_main = /* @__PURE__ */ defineComponent({
  ...__default__,
  props: {
    href: String
  },
  setup(__props) {
    return (_ctx, _cache) => {
      return openBlock(), createElementBlock("span", _hoisted_1, [
        renderSlot(_ctx.$slots, "default")
      ]);
    };
  }
});
_sfc_main.install = (app) => {
  app.component(_sfc_main.name, _sfc_main);
};
const components = [_sfc_main];
var index = {
  install(app) {
    components.forEach((component) => {
      app.use(component);
    });
  },
  ...components
};
export { index as default };
