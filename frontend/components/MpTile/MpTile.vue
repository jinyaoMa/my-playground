<script setup lang="ts">
// MpTile
const props = defineProps<{
  gap?: number;
}>();
</script>

<template>
  <div
    class="mp-tile"
    :style="{
      gap: `${props.gap || 0}px`,
      padding: `${props.gap || 0}px`,
    }"
  >
    <slot></slot>
  </div>
</template>

<style lang="scss">
.mp-tile {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  font-size: var(--mp-font-size);
  line-height: var(--mp-line-height-none);
  user-select: none;
}
</style>
