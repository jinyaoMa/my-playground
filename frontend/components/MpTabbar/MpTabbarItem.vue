<script setup lang="ts">
import MpIcon from "../MpIcon";

// MpTabbarItem
const props = defineProps<{
  active?: boolean;
  closeable?: boolean;
}>();

const emit = defineEmits<{
  (event: "close-tab"): void;
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
      @click="emit('close-tab')"
    >
      <MpIcon name="window-close"></MpIcon>
    </div>
  </div>
</template>

<style lang="scss">
.mp-tabbar-item {
  box-sizing: border-box;
  line-height: 2.8;
  padding: 0 0.75em;
  margin-left: 0.75em;
  gap: 0.5em;
  max-width: 50%;
  min-width: calc(2.5em);
  width: fit-content;
  overflow: hidden;
  text-overflow: ellipsis;
  position: relative;
  border-radius: var(--mp-border-radius);
  display: flex;
  flex-direction: row;
  align-items: center;

  &.closeable {
    padding-right: 0;
  }

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
  }

  &_close {
    opacity: 0;
    pointer-events: none;
    width: 0;
    height: 100%;
    margin-left: 0.25em;
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
    border-left: 2px dashed var(--mp-color-bg-3);

    &:hover {
      color: var(--mp-color-danger);
    }
  }

  &.active &_close {
    width: 2.5em;
    opacity: 1;
    pointer-events: all;
  }
}
</style>
