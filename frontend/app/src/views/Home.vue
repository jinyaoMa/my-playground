<script lang="ts" setup>
import { ComputedRef } from "vue";

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
    <mp-tile-item
      v-for="a in props.apps"
      class="tile-item"
      :key="a.key"
      :gap="10"
      @click=""
    >
      <img v-if="a.icon.includes('.')" class="tile-item-img" :src="a.icon" />
      <div v-else class="tile-item-icon">
        <mp-icon :name="a.icon"></mp-icon>
      </div>
      <div class="tile-item-title">{{ a.title }}</div>
    </mp-tile-item>
  </mp-tile>
</template>

<style lang="scss">
.home {
  position: relative;
}

.tile-item {
  &-img {
    display: block;
    object-fit: contain;
    height: 3em;
  }

  &-icon {
    font-size: 3em;
    line-height: 1;
  }

  &-title {
    line-height: 1.3;
  }
}
</style>
