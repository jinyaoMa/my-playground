import OpenLink from "./open-link.vue";

OpenLink.install = (app) => {
  app.component(OpenLink.name, OpenLink);
};

export default OpenLink;
