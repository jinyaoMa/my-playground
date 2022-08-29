<script setup lang="ts">
const props = defineProps({
  active: {
    type: Boolean,
  }
})

</script>

<template>
  <div class="mp-tabbar-item" :class="{
    active: props.active
  }">
    <slot></slot>
  </div>
</template>

<style lang="scss">
.mp-tabbar-item {
  box-sizing: border-box;
  padding: 0.75em 1em;
  margin: 0.5em 0 0.5em 0.5em;
  max-width: 100%;
  min-width: calc(3em);
  width: fit-content;
  cursor: pointer;
  overflow: hidden;
  text-overflow: ellipsis;

  &:hover {
    background-color: var(--mp-color-bg-2);
  }

  &.active {
    background-color: var(--mp-color-bg);
    box-shadow: var(--mp-shadow);

  }
}
</style>
