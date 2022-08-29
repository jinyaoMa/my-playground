<script setup lang="ts">
const props = defineProps({
  href: String,
})
</script>

<template>
  <span class="mp-open-link">
    <slot></slot>
  </span>
</template>

<style lang="scss">
.mp-open-link {
  cursor: pointer;
  text-decoration: underline;
}
</style>
