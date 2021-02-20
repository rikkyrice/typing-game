<template>
  <div class="mb-1" style="width: 100%;">
    <v-expansion-panels v-model="expandPanel" multiple style="border-bottom: 2px solid #A0D0A0;">
      <v-expansion-panel>
        <v-expansion-panel-header hide-actions>
          <div class="d-flex justify-center align-center mx-12" style="position: relative;">
            <span>キーワード、フィルター検索、ソートを行う</span>
            <v-hover v-model="isPlayButtonHover" open-delay="300">
              <lwtg-button
                label="PLAY!"
                height="24px"
                style="position: absolute; right: 0;"
                size="small"
                :append-src="playButtonSvg"
                :primary="true"
                @click="playTypingGame"
              />
            </v-hover>
          </div>
        </v-expansion-panel-header>
        <v-expansion-panel-content>Lorem ipsum.</v-expansion-panel-content>
      </v-expansion-panel>
    </v-expansion-panels>
    <div class="d-flex justify-center pb-2">
      <v-img
        v-ripple="{class: 'white--text'}"
        alt="コントロールパネルを開く"
        :src="expansionButton"
        :class="{
          'hover-animation-card': expandPanel.length === 0,
          'clicked-card': expandPanel.length !== 0,
        }"
        max-width="56px"
        width="56px"
        style="z-index: 5; cursor: pointer;"
        @click="expand"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';
import LwtgButton from '@/components/atoms/LwtgButton.vue';
import LwtgIconButton from '@/components/atoms/LwtgIconButton.vue';
import { DeviceType, deviceType } from '@/models/types/deviceType';

@Component({
  components: {
    LwtgButton,
    LwtgIconButton,
  }
})
export default class FilterExpansionPanel extends mixins(UtilMixin) {
  expandPanel: number[] = [];
  expansionButtonRotate = 0;
  expansionButtonTranslate = 0;
  isPlayButtonHover = false;
  get isDesktop() {
    return deviceType === DeviceType.DESKTOP;
  }
  get expansionButton() {
    return require('@/assets/common/expansion-button.svg');
  }
  get playButtonSvg() {
    return this.isDesktop && this.isPlayButtonHover
      ? require('@/assets/common/LwtgRightArrowGreen.svg')
      : require('@/assets/common/LwtgRightArrowWhite.svg');
  }
  expand() {
    if (this.expandPanel.length !== 0) {
      this.expandPanel = [];
      this.expansionButtonRotate = 0;
      this.expansionButtonTranslate = 0;
    } else {
      this.expandPanel = [0];
      this.expansionButtonRotate = 180;
      this.expansionButtonTranslate = 20;
    }
  }
  playTypingGame() {
    const path = this.$route.path
    this.$router.push(`${path}/typing-game`, () => {});
  }
};
</script>

<style scoped lang="scss">
@import '@/style.scss';
.v-expansion-panel::before {
   box-shadow: none !important;
}
.hover-animation-card {
  transition: all .2s ease-out;
}
.hover-animation-card:hover {
  transform: scale(1.1) !important;
}
.clicked-card {
  transform: rotate(180deg) translateY(20px);
  transition: all .2s ease-out;
}
.clicked-card:hover {
  transform: rotate(180deg) translateY(20px) scale(1.1) !important;
}
</style>
