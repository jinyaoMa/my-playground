<script>
export default {
  name: 'MpOpenLink',
}
</script>

<script setup>
const props = defineProps({
  href: String,
})

const onClickhandle = () => {
  // You cannot use the a tag directly, you need to call the Go method here to open the link using the default browser.
  // 不能直接使用a标签，需要在这里调用Go方法以使用默认浏览器打开链接。

  window.runtime.BrowserOpenURL(props.href);
};
</script>

<template>
  <span class="mp-openlink" @click="onClickhandle">
    <slot></slot>
  </span>
</template>

<style lang="less">
.mp-openlink {
  cursor: pointer;
  text-decoration: underline;
}
</style>
