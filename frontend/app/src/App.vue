<script setup lang="ts">
import { ref, computed } from "vue";
import { useI18n } from "vue-i18n";
import { useRouter } from "vue-router";
import { routes } from "./router";
import externalApps from "../externalApps";
import {
  EventsOn,
  WindowMinimise,
  WindowToggleMaximise,
  WindowHide,
} from "../../wailsjs/runtime/runtime";
import { ChangeLanguage, ChangeTheme } from "../../wailsjs/go/backend/App";

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
  console.log("onLanguageChanged", lang);
  locale.value = lang;
});

// Click to switch theme
// 点击切换主题
const cTheme = ref("light");
const onclickThemeHandle = (theme: string) => {
  theme !== cTheme.value ? ChangeTheme(theme) : false;
};

EventsOn("OnThemeChanged", (theme: string) => {
  console.log("OnThemeChanged", theme);
  document.querySelector(":root")?.classList.remove(cTheme.value);
  document.querySelector(":root")?.classList.add(theme);
  cTheme.value = theme;
});

const router = useRouter();
const panelApps = ref([
  ...routes.map((r) => {
    return {
      key: r.name?.toString() || "",
      title: computed(() =>
        t(`panel.${r.name?.toString()}`, { default: r.name })
      ),
      link: r.path,
      native: true,
      closeable: r.name != "Home",
    };
  }),
  ...externalApps.map((ea) => {
    return {
      ...ea,
      title: computed(() => ea.title[locale.value]),
      native: false,
      closeable: true,
    };
  }),
]);
const openTabKeys = ref([
  ...routes.map((r) => {
    return r.name?.toString() || "";
  }),
  ...externalApps.map((ea) => {
    return ea.key;
  }),
]);
const openTabs = computed(() => {
  return panelApps.value.filter((pa) => {
    return openTabKeys.value.includes(pa.key);
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
    <router-view
      :apps="panelApps.filter((pa) => pa.key != 'Home')"
    ></router-view>
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
* {
  transition: all 0.2s;
}

html {
  width: 100%;
  height: 100%;
}

body {
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
}

#app {
  position: relative;
  height: 100%;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  color: var(--mp-color-text);
  background-color: var(--mp-color-bg);
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
  align-items: center;
  min-width: 2.5em;
  min-height: 2.5em;
}

.tabbar-icon {
  width: 20px;
  height: 20px;
  display: block;
  pointer-events: none;
  border: 2px solid var(--mp-color-bg-3);
  border-radius: var(--mp-border-radius);
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
