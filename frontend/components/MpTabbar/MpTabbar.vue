<script setup lang="ts"></script>

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
  align-items: center;
  font-size: var(--mp-font-size-s);
  line-height: var(--mp-line-height-none);
  background-color: var(--mp-color-bg-1);
  color: var(--mp-color-text);
  user-select: none;
  padding: 0.75em;

  &_tabs {
    flex-grow: 1;
    display: flex;
    flex-direction: row;
    overflow-x: clip;
    overflow-y: visible;
  }
}
</style>
