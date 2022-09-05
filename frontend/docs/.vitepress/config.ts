// import { defineConfigWithTheme } from 'vitepress'
// import { ThemeConfig } from 'your-theme'
import { defineConfig } from "vitepress";

export default defineConfig({
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

  locales: {
    "/en/": {
      lang: "en-US",
      title: "My Text Playground",
      description:
        "Everything about and in Jinyao's playground includes introduction, blog posts, technical docs, etc.",
    },
    "/zh/": {
      lang: "zh-CN",
      title: "我的文本游乐场",
      description:
        "有关于我的游乐场和包含在我的游乐场中的所有东西，如简介、博客、技术文档等",
    },
  },

  themeConfig: {
    // Type is `DefaultTheme.Config`
  },

  head: [],

  scrollOffset: 0,

  ignoreDeadLinks: false,

  // set it to subdirectory in production inserting into /backend/.frontend
  base: process.env.NODE_ENV === "production" ? "/docs/" : "/",
  outDir: "../../backend/.frontend/docs",
  vite: {
    server: {
      port: 10002,
    },
  },
});
