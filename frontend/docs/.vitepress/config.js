/**
 * @type {import('vitepress').UserConfig}
 */
const config = {
  base: "/docs/",
  outDir: "../../backend/.assets/docs",
  vite: {
    server: {
      port: 10002,
    },
  },
};

export default config;
