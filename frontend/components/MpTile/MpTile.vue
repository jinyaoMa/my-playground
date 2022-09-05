<script setup lang="ts">
// MpTile
const props = defineProps<{
  gap: number;
}>();
</script>

<template>
  <div
    class="mp-tile"
    :style="{
      gap: `${props.gap}px`,
      padding: `${props.gap}px`,
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
}
</style>
