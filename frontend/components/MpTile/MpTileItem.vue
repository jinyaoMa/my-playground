<script setup lang="ts">
// MpTileItem
const props = defineProps<{
  gap?: number;
}>();
</script>

<template>
  <div
    class="mp-tile-item"
    :style="{
      gap: `${props.gap || 0}px`,
      padding: `${props.gap || 0}px`,
    }"
  >
    <slot></slot>
  </div>
</template>

<style lang="scss">
.mp-tile-item {
  height: fit-content;
  width: 50%;
  max-width: 8em;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
  align-items: center;
  border-radius: var(--mp-border-radius);
  cursor: pointer;

  &:hover {
    background-color: var(--mp-color-bg-1);
  }
}
</style>
