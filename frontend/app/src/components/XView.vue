<script lang="ts" setup>
import { ref, onMounted, onUpdated } from "vue";
const props = defineProps({
  src: String,
  theme: String,
});

let origTheme = "";
const xview = ref();
onMounted(() => {
  xview.value.contentWindow.runtime = (window as any).runtime;
  xview.value.contentWindow.go = (window as any).go;
  xview.value.contentWindow.addEventListener("load", () => {
    xview.value.contentDocument
      .querySelector(":root")
      ?.classList.remove(origTheme);
    origTheme = props.theme || "";
    xview.value.contentDocument
      .querySelector(":root")
      ?.classList.add(props.theme);
    xview.value.contentWindow.localStorage.setItem(
      "vitepress-theme-appearance",
      props.theme
    );
  });
});
onUpdated(() => {
  xview.value.contentDocument
    .querySelector(":root")
    ?.classList.remove(origTheme);
  origTheme = props.theme || origTheme;
  xview.value.contentDocument
    .querySelector(":root")
    ?.classList.add(props.theme);
  xview.value.contentWindow.localStorage.setItem(
    "vitepress-theme-appearance",
    props.theme
  );
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
