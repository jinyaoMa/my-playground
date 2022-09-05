<script setup lang="ts">
const props = defineProps<{
  name?: string;
  fw?: boolean;
}>();
</script>

<template>
  <i
    :class="`mp-icon-${props.name || 'jinyao-ma'}${fw ? ' fix-width' : ''}`"
  ></i>
</template>
