<script lang="ts" setup>
import { ComputedRef } from "vue";
import { useI18n } from "vue-i18n";
const { t } = useI18n({ useScope: "global" });

const props = defineProps<{
  apps: {
    key: string;
    title: ComputedRef<string>;
    link: string;
    native: boolean;
    closeable: boolean;
  }[];
}>();
</script>

<template>
  <div class="home">
    <div v-for="a in props.apps" :key="a.key">
      {{ a.title }}
    </div>
  </div>
</template>

<style lang="scss">
.home {
  position: relative;
  padding: 40px;
}
</style>
