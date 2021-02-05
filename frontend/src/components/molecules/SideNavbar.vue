<template>
  <v-navigation-drawer v-model="visibility" app temporary class="py-10">
    <div
      v-for="navItem in navItemList"
      :key="navItem.path"
      :class="{
        'current-page': navItem.type === 'span' && isCurrentPage(navItem.path),
      }"
    >
      <div
        v-if="navItem.type === 'span'"
        v-ripple
        class="px-3 py-2"
        @click="pageTransition(navItem.path)"
      >
        <span class="mx-auto text-button">{{ navItem.label }}</span>
      </div>
      <img
        v-if="navItem.type === 'img'"
        class="ma-3"
        :src="navItem.label"
        height="56"
        @click="pageTransition(navItem.path)"
      />
    </div>
  </v-navigation-drawer>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';
import { NavItemInfo } from '@/models/types/navItemInfo';

@Component
export default class SideNavbar extends mixins(UtilMixin) {
  @Prop() navItemList!: NavItemInfo[];
  visibility = false;

  openSideNavbar() {
    this.visibility = true;
  }

  isCurrentPage(path: string) {
    return this.$route.path === path;
  }

  pageTransition(path: string) {
    if (!this.isCurrentPage(path)) {
      this.$router.push(path);
    }
  }
}
</script>

<style scoped lang="scss">
@import '@/style.scss';
.current-page {
  color: white;
  background: #006699;
}
</style>
