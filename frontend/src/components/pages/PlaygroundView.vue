<template>
  <div id="playground">
    <!-- パンくずリスト -->
    <lwtg-content>
      <lwtg-breadcrumbs
        :class="MdSmXsUtil('py-4', 'py-3', 'py-3')"
        :breadcrumbs="breadcrumbs"
      />
    </lwtg-content>
    <!-- lwtgButton -->
    <lwtg-content>
      <v-row>
        <v-col
          cols="12"
          sm="6"
        >
          <v-hover v-model="isPlayButtonHover" open-delay="300">
            <lwtg-button
              class="ml-3"
              label="PLAY!"
              size="large"
              :append-src="playButtonSvg"
              :primary="true"
            />
          </v-hover>
        </v-col>
        <v-col
          cols="12"
          sm="6"
        >
          <lwtg-button
            class="ml-3"
            label="PLAY!"
            size="large"
            :append-src="require('@/assets/common/LwtgRightArrowWhite.svg')"
          />
        </v-col>
      </v-row>
    </lwtg-content>
    <lwtg-content>
      <lwtg-chip
        text="+ 作成"
        color="#666666"
        class="mr-1 bold"
        :clickable="true"
        :outlined="true"
      />
      <v-hover v-model="isTrashcanButtonHover" open-delay="100" close-delay="50">
        <lwtg-icon-button
          size="small"
          :src="trashcanButtonIcon"
        />
      </v-hover>
    </lwtg-content>
    <div class="lwtg-white-bg">
      <ca-content
        :caItems="caItems"
      />
    </div>
    <lwtg-content class="lwtg-secondary-bg">
      <lwtg-signup-form />
    </lwtg-content>
    <lwtg-loader :page-loading="true" :loading="false" />
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import CaContent from '@/components/organisms/CaContent.vue';
import LwtgBreadcrumbs from '@/components/atoms/LwtgBreadcrumbs.vue';
import LwtgButton from '@/components/atoms/LwtgButton.vue';
import LwtgChip from '@/components/atoms/LwtgChip.vue';
import LwtgContent from '@/components/atoms/LwtgContent.vue';
import LwtgIconButton from '@/components/atoms/LwtgIconButton.vue';
import LwtgLoader from '@/components/atoms/LwtgLoader.vue';
import LwtgSignupForm from '@/components/organisms/form/LwtgSignupForm.vue';
import UtilMixin from '@/mixins/utilMixin';
import { BreadcrumbInfo } from '@/models/types/breadcrumbInfo';
import { DeviceType, deviceType } from '@/models/types/deviceType';
import { CaItem } from '@/models/types/caItem';

@Component({
  components: {
    CaContent,
    LwtgBreadcrumbs,
    LwtgButton,
    LwtgChip,
    LwtgContent,
    LwtgIconButton,
    LwtgLoader,
    LwtgSignupForm,
  },
})
export default class PlaygroundView extends mixins(UtilMixin) {
  breadcrumbs: BreadcrumbInfo[] = [
    { label: 'トップ', path: '/' },
    { label: 'PlayGround', path: '/pg' },
  ];
  isPlayButtonHover = false;
  isTrashcanButtonHover = false;
  caItems: CaItem[] = [
    { title: 'CyberAgent Way', subtitle: '人と企業が成長する仕組み', path: '/', img: '#' },
    { title: 'CSR', subtitle: '社会的責任', path: '/', img: '#' },
    { title: 'CSR', subtitle: '社会的責任', path: '/', img: '#' },
    { title: 'CSR', subtitle: '社会的責任', path: '/', img: '#' },
  ];
  get isDesktop() {
    return deviceType === DeviceType.DESKTOP;
  }
  get playButtonSvg() {
    return this.isDesktop && this.isPlayButtonHover
      ? require('@/assets/common/LwtgRightArrowGreen.svg')
      : require('@/assets/common/LwtgRightArrowWhite.svg');
  }
  get trashcanButtonIcon() {
    return this.isDesktop && this.isTrashcanButtonHover
      ? require('@/assets/common/trash-can.svg')
      : require('@/assets/common/trash-can-outline.svg');
  }
}
</script>

<style scoped lang="scss">
@import '@/style.scss';
</style>
