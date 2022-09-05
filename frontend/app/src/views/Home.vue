<script lang="ts" setup>
import { ComputedRef } from "vue";
import { useI18n } from "vue-i18n";
const { t } = useI18n({ useScope: "global" });

const props = defineProps<{
  apps: {
    key: string;
    title: ComputedRef<string>;
    link: string;
    icon: string;
    native: boolean;
    closeable: boolean;
  }[];
}>();
</script>

<template>
  <mp-tile class="home" :gap="40">
    <mp-tile-item v-for="a in props.apps" :key="a.key">
      <img :src="a.icon" />
      {{ a.title }}
    </mp-tile-item>
  </mp-tile>
</template>

<style lang="scss">
.home {
  position: relative;
}
</style>
