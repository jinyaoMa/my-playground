<script setup lang="ts">
import { ref, computed } from "vue";
import { useI18n } from "vue-i18n";
import { useRouter } from "vue-router";
import {
  EventsOn,
  WindowMinimise,
  WindowToggleMaximise,
  WindowHide,
} from "../../wailsjs/runtime/runtime";
import { ChangeLanguage } from "../../wailsjs/go/backend/App";
import MpTabbar from "../../components/MpTabbar/MpTabbar.vue";
import MpTabbarItem from "../../components/MpTabbar/MpTabbarItem.vue";

const { t, availableLocales, locale } = useI18n();

// List of supported languages
// 支持的语言列表
const languages = availableLocales;

// Click to switch language
// 点击切换语言
const onclickLanguageHandle = (lang: string) => {
  lang !== locale.value ? ChangeLanguage(lang) : false;
};

EventsOn("onLanguageChanged", (lang: string) => {
  console.log(lang);
  locale.value = lang;
});

const testTabs = ref([
  {
    key: "home",
    title: "Home",
    link: "/",
    native: true,
    closeable: false,
  },
  {
    key: "about",
    title: "About",
    link: "/about",
    native: true,
    closeable: true,
  },
  {
    key: "docs",
    title: "Docs",
    link: "/docs/index.html",
    native: false,
    closeable: true,
  },
]);

const router = useRouter();
const openTabKeys = ref(["home", "about", "docs"]);
const openTabs = computed(() => {
  return testTabs.value.filter((tab) => {
    return openTabKeys.value.includes(tab.key);
  });
});
const activeTabIndex = ref(0);
const xviewSrc = ref("");
const onClickTab = (index: number) => {
  if (activeTabIndex.value != index) {
    activeTabIndex.value = index;
  }
  if (activeTabIndex.value >= openTabKeys.value.length) {
    activeTabIndex.value = openTabKeys.value.length - 1;
  } else if (activeTabIndex.value < 0) {
    activeTabIndex.value = 0;
  }
  const tab = openTabs.value[activeTabIndex.value];
  if (tab.native) {
    router.replace(tab.link);
  } else {
    xviewSrc.value = tab.link;
  }
};
const onCloseTab = (prop: string) => {
  const target = openTabKeys.value.indexOf(prop);
  openTabKeys.value.splice(target, 1);
};

const isWindowMaximized = ref(false);
window.addEventListener("resize", (ev: UIEvent) => {
  if (ev.target instanceof Window) {
    isWindowMaximized.value = ev.target.innerWidth >= screen.availWidth;
  }
});
const onclickMinimize = () => {
  WindowMinimise();
};
const onclickToggleMaximize = () => {
  WindowToggleMaximise();
};
const onclickQuit = () => {
  WindowHide();
};
</script>

<template>
  <!-- Header -->
  <mp-tabbar data-wails-drag @dblclick="onclickToggleMaximize">
    <template #prepend>
      <div class="tabbar-prepend">
        <img class="tabbar-icon" src="./assets/icon.png" />
      </div>
    </template>
    <mp-tabbar-item
      v-for="(tab, i) in openTabs"
      :active="i == activeTabIndex"
      :closeable="tab.closeable"
      data-wails-no-drag
      @click="onClickTab(i)"
      @close-tab="onCloseTab(tab.key)"
    >
      {{ tab.title }}
    </mp-tabbar-item>
    <template #append>
      <div class="tabbar-append" data-wails-no-drag>
        <div class="tabbar-btn" @click="onclickMinimize">
          <i class="mp-icon-window-minimize"></i>
        </div>
        <div class="tabbar-btn" @click="onclickToggleMaximize">
          <i v-if="isWindowMaximized" class="mp-icon-window-restore"></i>
          <i v-else class="mp-icon-window-maximize"></i>
        </div>
        <div class="tabbar-btn danger" @click="onclickQuit">
          <i class="mp-icon-window-close"></i>
        </div>
      </div>
    </template>
  </mp-tabbar>
  <!-- Page -->
  <div v-if="openTabs[activeTabIndex].native" class="view">
    <router-view></router-view>
  </div>
  <iframe v-else class="xview" :src="xviewSrc"></iframe>
  <!--
    <div class="header" data-wails-drag>
      <div class="nav" data-wails-no-drag>
        <router-link to="/"
          ><i class="mp-icon-jinyao-ma"></i>{{ t("nav.home") }}</router-link
        >
        <router-link to="/about">{{ t("nav.about") }}</router-link>
      </div>
      <div class="menu" data-wails-no-drag>
        <div class="language">
          <div
            v-for="item in languages"
            :key="item"
            :class="{ active: item === locale }"
            @click="onclickLanguageHandle(item)"
            class="lang-item"
          >
            {{ t("languages." + item) }}
          </div>
        </div>
        <div class="bar">
          <div class="bar-btn" @click="onclickDocs">DOCS</div>
        </div>
      </div>
    </div>
  -->
</template>

<style lang="scss">
html {
  width: 100%;
  height: 100%;
}

body {
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
  background-color: #ffffff;
}

#app {
  position: relative;
  height: 100%;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

.view,
.xview {
  flex-grow: 1;
  overflow: auto;
  scroll-behavior: smooth;
  border: none;
  border-top: 2px solid var(--mp-color-bg-2);
}

.tabbar-prepend {
  display: flex;
  flex-direction: row;
  justify-content: center;
  min-width: 2.5em;
}

.tabbar-icon {
  width: 2em;
  height: 2em;
  display: block;
  pointer-events: none;
}

.tabbar-append {
  display: flex;
  flex-direction: row;
  min-width: 7.5em;
}

.tabbar-btn {
  width: 2.5em;
  height: 2.5em;
  color: var(--mp-color-text-1);
  display: flex;
  justify-content: center;
  align-items: center;
  border-radius: var(--mp-border-radius);

  &:hover {
    color: var(--mp-color-text);
    background-color: var(--mp-color-bg-2);
  }

  &.danger i {
    font-size: 1.3em;
  }

  &.danger:hover {
    color: var(--mp-color-danger);
  }

  i {
    font-size: 1em;
  }
}

//===============================================================

.header {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  align-items: center;
  justify-content: space-between;
  height: 50px;
  padding: 0 10px;
  background-color: rgba(171, 126, 220, 0.9);

  .nav {
    a {
      display: inline-block;
      min-width: 50px;
      height: 30px;
      line-height: 30px;
      padding: 0 5px;
      margin-right: 8px;
      background-color: #ab7edc;
      border-radius: 2px;
      text-align: center;
      text-decoration: none;
      color: #000000;
      font-size: 14px;
      white-space: nowrap;

      &:hover,
      &.router-link-exact-active {
        background-color: #d7a8d8;
        color: #ffffff;
      }
    }
  }

  .menu {
    display: flex;
    flex-direction: row;
    flex-wrap: nowrap;
    align-items: center;
    justify-content: space-between;

    .language {
      margin-right: 20px;
      border-radius: 2px;
      background-color: #c3c3c3;
      overflow: hidden;

      .lang-item {
        display: inline-block;
        min-width: 50px;
        height: 30px;
        line-height: 30px;
        padding: 0 5px;
        background-color: transparent;
        text-align: center;
        text-decoration: none;
        color: #000000;
        font-size: 14px;

        &:hover {
          background-color: #ff050542;
          cursor: pointer;
        }

        &.active {
          background-color: #ff050542;
          color: #ffffff;
          cursor: not-allowed;
        }
      }
    }

    .bar {
      display: flex;
      flex-direction: row;
      flex-wrap: nowrap;
      align-items: center;
      justify-content: flex-end;
      min-width: 150px;

      .bar-btn {
        display: inline-block;
        min-width: 80px;
        height: 30px;
        line-height: 30px;
        padding: 0 5px;
        margin-left: 8px;
        background-color: #ab7edc;
        border-radius: 2px;
        text-align: center;
        text-decoration: none;
        color: #000000;
        font-size: 14px;

        &:hover {
          background-color: #d7a8d8;
          color: #ffffff;
          cursor: pointer;
        }
      }
    }
  }
}
</style>
