import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import path from "path";

import pkg from "../../wails.json";
import { createHtmlPlugin } from "vite-plugin-html";
import { vueI18n } from "@intlify/vite-plugin-vue-i18n";

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    port: 10001,
  },
  plugins: [
    vue(),
    createHtmlPlugin({
      minify: false,
      // entry: "src/main.js",
      template: "index.html",
      inject: {
        data: {
          title: `${pkg.name}`,
        },
      },
    }),
    vueI18n({
      include: path.resolve(__dirname, "src/i18n/locales/**"),
      compositionOnly: true,
    }),
  ],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "src"),
      "@runtime": path.resolve(__dirname, "wailsjs/runtime/runtime.js"),
      "@backend": path.resolve(__dirname, "wailsjs/go/backend/App.js"),
      "vue-i18n": "vue-i18n/dist/vue-i18n.runtime.esm-bundler.js",
    },
  },
  build: {
    outDir: "../../backend/.assets",
    rollupOptions: {
      output: {
        entryFileNames: `assets/[name].js`,
        chunkFileNames: `assets/[name].js`,
        assetFileNames: `assets/[name].[ext]`,
      },
    },
  },
});
