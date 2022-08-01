import { MpOpenLink } from "@my-playground/components/open-link";

const components = [MpOpenLink];

export const INSTALLED_KEY = Symbol("INSTALLED_KEY");
export default {
  install(app) {
    if (app[INSTALLED_KEY]) return;
    app[INSTALLED_KEY] = true;
    components.forEach((c) => {
      app.use(c);
    });
  },
};
