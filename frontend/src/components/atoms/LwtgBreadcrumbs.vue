<template>
  <div
    id="breads-wrapper"
    class="d-flex align-center text--secondary"
    style="font-size: 12px;"
  >
    <div
      v-for="(bread, index) in breadcrumbs"
      :key="bread.path"
      class="d-flex align-center"
      :style="fontSizeUtil(14, 12, 12)"
    >
      <v-icon v-if="index" class="mx-2" color="#A0D0A0"
        >mdi-chevron-right</v-icon
      >
      <div
        :id="`bread-${index}`"
        class="text-overflow-ellipsis"
        :class="{ 'breadcrumb-link': index < breadcrumbs.length -1 }"
        :style="{
          width:
            index < breadcrumbs.length -1 ? 'auto' : `${lastBreadWidth}px`,
        }"
        @click="pageTransition(bread)"
      >
        <span>{{ bread.label }}</span>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';
import { BreadcrumbInfo } from '@/models/types/breadcrumbInfo';

@Component
export default class LwtgBreadcrumbs extends mixins(UtilMixin) {
  @Prop() breadcrumbs!: BreadcrumbInfo[];
  lastBreadWidth = 0;

  get iconTemplate() {
    return `<span class="mr-1" color="#A0D0A0">mdi-chevron-right</span>`;
  }

  mounted() {
    const breadsWrapper = document.getElementById('breads-wrapper');
    const lastBread = document.getElementById(
      `bread-${this.breadcrumbs.length - 1}`
    );
    if (breadsWrapper && lastBread) {
      const wrapperWidth = breadsWrapper.clientWidth;
      const wrapperLeft = breadsWrapper.getBoundingClientRect().left;
      const lastBreadcrumbLeft = lastBread.getBoundingClientRect().left;
      this.lastBreadWidth = wrapperWidth + wrapperLeft - lastBreadcrumbLeft - 4;
    }
  }

  pageTransition(bread: BreadcrumbInfo) {
    if (this.breadcrumbs.indexOf(bread) < this.breadcrumbs.length - 1) {
      this.$router.push(bread.path);
    }
  }
}
</script>

<style scoped lang="scss">
@import '@/style.scss';
.breadcrumb-link {
  cursor: pointer;
  &:hover {
    color: #A0D0A0;
    text-decoration: underline;
  }
}
</style>
