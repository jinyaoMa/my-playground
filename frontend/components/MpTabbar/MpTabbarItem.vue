<script setup lang="ts">
const props = defineProps<{
  prop: string;
  active?: boolean;
  closeable?: boolean;
}>();

const emit = defineEmits<{
  (event: "close-tab", prop: string): void;
}>();
</script>

<template>
  <div
    class="mp-tabbar-item"
    :class="{
      active: props.active,
      closeable: props.closeable,
    }"
  >
    <slot></slot>
    <div
      v-if="props.closeable"
      class="mp-tabbar-item_close"
      @click="emit('close-tab', props.prop)"
    >
      <i class="mp-icon-window-close"></i>
    </div>
  </div>
</template>

<style lang="scss">
.mp-tabbar-item {
  box-sizing: border-box;
  line-height: 2.5;
  padding: 0 0.75em;
  margin-left: 0.75em;
  max-width: 50%;
  min-width: calc(2.5em);
  width: fit-content;
  overflow: hidden;
  text-overflow: ellipsis;
  position: relative;
  border-radius: var(--mp-border-radius);

  &:last-child {
    margin-right: 0.75em;
  }

  &:hover {
    background-color: var(--mp-color-bg-2);
  }

  &.active {
    background-color: var(--mp-color-bg);
    box-shadow: var(--mp-shadow);
    transform: translateY(-0.5px);

    &.closeable {
      padding-right: 2.5em;
    }
  }

  &_close {
    opacity: 0;
    pointer-events: none;
    position: absolute;
    top: 0.5em;
    right: 0.5em;
    width: 1.5em;
    line-height: 1.5em;
    text-align: center;
    cursor: pointer;
    border-radius: var(--mp-border-radius-s);

    &:hover {
      background-color: var(--mp-color-bg-2);
    }
  }

  &.active &_close {
    opacity: 1;
    pointer-events: all;
  }
}
</style>
