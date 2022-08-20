import OpenLink from "./open-link";

const components = [OpenLink];

export default {
  install(app) {
    components.forEach((c) => {
      app.use(c);
    });
  },
  ...components,
};
