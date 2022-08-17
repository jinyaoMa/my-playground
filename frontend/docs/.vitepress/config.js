/**
 * @type {import('vitepress').UserConfig}
 */
const config = {
  // set it to subdirectory in production inserting into /backend/.assets
  base: process.env.NODE_ENV === "production" ? "/docs/" : "/",
  outDir: "../../backend/.assets/docs",
  vite: {
    server: {
      port: 10002,
    },
  },
};

export default config;
