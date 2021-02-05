<template>
  <v-app-bar app fixed hide-on-scroll elevation="1" color="white" style="max-height: 64px;">
    <template v-if="$vuetify.breakpoint.xsOnly">
      <v-row style="position: relative;">
        <v-icon
          v-if="!isLoginPage"
          style="position: absolute; top: 17px; left: 16px;"
          @click="$emit('open-side-navbar')"
          >mdi-menu</v-icon
        >
        <img
          :src="require('@/assets/lwtg-logo-3.svg')"
          height="56"
          class="mx-auto"
          @click="$router.push('/')"
        />
      </v-row>
    </template>
    <template v-else>
      <div
        v-for="(navItem, index) in navItemList"
        :key="navItem.path"
        v-ripple="navItem.type !== 'btn' && !isLoginPage"
        class="px-2 mx-2"
        :class="{
          'mr-auto': index === 0,
          'header-nav-item':
            navItem.type !== 'btn' &&
            !isCurrentPage(navItem.path) &&
            !isLoginPage,
          'current-nav-item':
            navItem.type !== 'btn' && isCurrentPage(navItem.path) && index > 0,
        }"
        @click="pageTransition(navItem.path)"
      >
        <div
          class="d-flex align-center justify-center"
        >
          <span
            v-if="navItem.type === 'span' && !isLoginPage"
            class="mx-auto"
            >{{ navItem.label }}</span
          >
          <img
            v-if="navItem.type === 'img'"
            :src="navItem.label"
            v-ripple="false"
            width="160"
            height="56"
            class="mx-auto"
          />
        </div>
      </div>
    </template>
  </v-app-bar>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';
import { NavItemInfo } from '@/models/types/navItemInfo';
import { PAGES } from '@/router/pages';

@Component
export default class HeaderNavbar extends mixins(UtilMixin) {
  @Prop() navItemList!: NavItemInfo[];
  drawerVisibility = false;

  get isLoginPage() {
    return this.$route.name === PAGES.LOGIN;
  }

  isCurrentPage(path: string) {
    return this.$route.path === path;
  }

  pageTransition(path: string) {
    if (!this.isCurrentPage(path) && !this.isLoginPage) {
      this.$router.push(path);
    }
  }
}
</script>

<style scoped lang="scss">
@import '@/style.scss';
.current-nav-item {
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  position: relative;
  cursor: normal;
  &:after {
    content: '';
    width: 100%;
    height: 4px;
    display: block;
    position: absolute;
    bottom: -4px;
    background: linear-gradient(90deg, #E2F0D9, #A0D0A0);
  }
  &:hover {
    color: #A0D0A0;
    -webkit-transition-duration: 0.3s;
    transition-duration: 0.3s;
  }
}
.header-nav-item {
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  position: relative;
  cursor: pointer;
  &:hover {
    color: #A0D0A0;
    -webkit-transition-duration: 0.3s;
    transition-duration: 0.3s;
  }
}
</style>
