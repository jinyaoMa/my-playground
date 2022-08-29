<script setup lang="ts">
const props = defineProps({
  test: String
})
</script>

<template>
  <div class="mp-tabbar">
    <div class="mp-tabbar_prepend">
      <slot name="prepend"></slot>
    </div>
    <div class="mp-tabbar_tabs">
      <slot></slot>
    </div>
    <div class="mp-tabbar_append">
      <slot name="append"></slot>
    </div>
  </div>
</template>

<style lang="scss">
.mp-tabbar {
  display: flex;
  flex-direction: row;
  font-size: var(--mp-font-size-s);
  line-height: var(--mp-line-height-none);
  background-color: var(--mp-color-bg-1);

  &_tabs {
    flex-grow: 1;
    display: flex;
    flex-direction: row;
  }
}
</style>
