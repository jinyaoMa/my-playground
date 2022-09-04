<script setup lang="ts">
const props = defineProps<{
  name?: string;
}>();
</script>

<template>
  <i :class="`mp-icon-${props.name || 'jinyao-ma'}`"></i>
</template>
