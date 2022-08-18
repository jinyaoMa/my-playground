// import { defineConfigWithTheme } from 'vitepress'
// import { ThemeConfig } from 'your-theme'
import { defineConfig } from "vitepress";

export default defineConfig({
  lang: "en-US",
  title: "My Text Playground",
  titleTemplate: false,
  description:
    "Everything about and in Jinyao's playground includes introduction, blog posts, technical docs, etc.",
  appearance: true,
  lastUpdated: true,
  markdown: {
    theme: "material-palenight",
    lineNumbers: true,
  },

  locales: {},

  themeConfig: {
    // Type is `DefaultTheme.Config`
  },

  head: [],

  scrollOffset: 0,

  ignoreDeadLinks: false,

  // set it to subdirectory in production inserting into /backend/.assets
  base: process.env.NODE_ENV === "production" ? "/docs/" : "/",
  outDir: "../../backend/.assets/docs",
  vite: {
    server: {
      port: 10002,
    },
  },
});
