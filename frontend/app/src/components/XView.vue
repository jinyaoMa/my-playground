<script lang="ts" setup>
import { ref, onMounted, onUpdated, onBeforeUnmount } from "vue";

const props = defineProps<{
  src: string;
  theme: string;
  vitepress: boolean;
}>();

const xview = ref();
let curTheme = "";
const setTheme = () => {
  curTheme = props.theme;
  xview.value.contentDocument
    .querySelector(":root")
    ?.classList.add(props.theme);
  if (props.vitepress) {
    xview.value.contentWindow.localStorage.setItem(
      "vitepress-theme-appearance",
      props.theme
    );
  }
};
onMounted(() => {
  xview.value.contentWindow.runtime = (window as any).runtime;
  xview.value.contentWindow.go = (window as any).go;
  xview.value.contentWindow.addEventListener("load", () => {
    if (props.vitepress) {
      const storedVitepressTheme =
        xview.value.contentWindow.localStorage.getItem(
          "vitepress-theme-appearance"
        );
      if (storedVitepressTheme) {
        xview.value.contentDocument
          .querySelector(":root")
          ?.classList.remove(storedVitepressTheme);
      }
    }
    setTheme();
  });
});
onBeforeUnmount(() => {
  if (props.vitepress) {
    const storedVitepressTheme = xview.value.contentWindow.localStorage.getItem(
      "vitepress-theme-appearance"
    );
    if (storedVitepressTheme != props.theme) {
      xview.value.contentWindow.localStorage.setItem(
        "vitepress-theme-appearance",
        props.theme
      );
    }
  }
});
onUpdated(() => {
  if (curTheme != props.theme) {
    xview.value.contentDocument
      .querySelector(":root")
      ?.classList.remove(curTheme);
    setTheme();
  }
});
</script>

<template>
  <iframe class="x-view" ref="xview" :src="props.src"></iframe>
</template>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.x-view {
  flex-grow: 1;
  overflow: auto;
  scroll-behavior: smooth;
  border: none;
  border-top: 2px solid var(--mp-color-bg-2);
}
</style>
