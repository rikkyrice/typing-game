<template>
  <div id="app">
    <v-app>
      <!-- System Bar -->
      <!-- <v-system-bar app /> -->
      <!-- Snackbar -->
      <lwtg-snackbar v-model="snackbarVisibility" :text="snackbarInfo.message" />
      <!-- HeaderNav Bar -->
      <header-navbar
        :nav-item-list="navItemList"
        @open-side-navbar="openSideNavbar"
      />
      <!-- Side Bar -->
      <!-- <v-navigation-drawer app /> -->
      <side-navbar ref="side-navbar" :nav-item-list="navItemList" />
      <v-main id="main">
        <router-view />
      </v-main>
    </v-app>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator';
import Toasted from 'vue-toasted';
import LwtgSnackbar from '@/components/atoms/LwtgSnackbar.vue';
import HeaderNavbar from '@/components/molecules/HeaderNavbar.vue';
import SideNavbar from '@/components/molecules/SideNavbar.vue';
import { NavItemInfo } from '@/models/types/navItemInfo';
import store from '@/store';
import { SnackbarInfo } from '@/store/types';
import { TYPES } from '@/store/mutation-types';

Vue.use(Toasted, {
  position: 'top-center',
  duration: 3000,
});

@Component({
  components: { LwtgSnackbar, HeaderNavbar, SideNavbar },
})
export default class App extends Vue {
  userId = store.state.auth.userId;
  navItemListBeforeLogin: NavItemInfo[] = [
    { type: 'img', label: require('@/assets/lwtg-logo-3.svg'), path: '/' },
    { type: 'span', label: 'ユーザー登録', path: '/signup' },
    { type: 'span', label: 'ログイン', path: '/login' },
  ];
  navItemListAfterLogin: NavItemInfo[] = [
    { type: 'img', label: require('@/assets/lwtg-logo-3.svg'), path: '/' },
    { type: 'span', label: 'マイページ', path: '/mypage' },
    { type: 'span', label: 'PlayGround', path: '/pg' },
  ];

  get navItemList() {
    return this.userId
      ? this.navItemListAfterLogin
      : this.navItemListBeforeLogin;
  }

  get snackbarInfo() {
    return store.state.snackbar;
  }
  get snackbarVisibility() {
    return this.snackbarInfo.visibility;
  }
  set snackbarVisibility(visibility: boolean) {
    store.dispatch(TYPES.SNACKBAR, '');
  }

  openSideNavbar() {
    const sideNavbar: any = this.$refs['side-navbar'];
    sideNavbar.openSideNavbar();
  }
}
</script>

<style lang="scss">
@import '@/style.scss';
// global styles
#app {
  font-family: '游ゴシック', 'HiraginoSans-W3', 'Meiryo UI', sans-serif !important;
  min-height: 100vh;
}
#main {
  max-height: 100vh;
}
.toasted-container .toasted {
  background-color: #e32d2d !important;
  box-shadow: 0 2px 4px 0 rgba(0, 0, 0, 0.5);
  opacity: 0.75 !important;
}
</style>
