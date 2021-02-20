<template>
  <v-hover v-model="isCardHover">
    <div
      :id="`word-card-${index}`"
      class="flip-card">
      <lwtg-word-card
        :width="width"
        :border="border"
        :primary="primary"
        class="flip-card-inner"
        :class="{
          'flipped-card': isFlipped,
        }"
        @click="flipCard"
      >
        <template #word>
          <v-row
            class="flip-card-front"
            :class="{
              'flipped-card': isFlipped,
            }"
            style="width: 100%;"
            no-gutters
          >
            <v-spacer v-if="!isFlipped">
            </v-spacer>
            <v-col
              v-if="!isFlipped"
              cols="9"
            >
              <div
                class="d-flex align-center lwtg-white-bg"
                style="height: 100%;"
              >
                <span
                  v-if="isWordRevealed"
                  :style="wordFont"
                >{{ word.word }}</span>
              </div>
            </v-col>
            <v-col
              v-else
              cols="10"
              class="pa-3 pl-6"
            >
              <div class="d-flex flex-column justify-center lwtg-white-bg" style="height: 100%;">
                <div>
                  <span class="bold" :style="meaningFont"
                  >{{ word.meaning }}</span>
                </div>
                <v-divider></v-divider>
                <div>
                  <span :style="explanationFont"
                  >{{ word.explanation }}</span>
                </div>
              </div>
            </v-col>
            <v-col
              cols="2"
              class="pa-3"
            >
              <div
                class="d-flex flex-column lwtg-white-bg pa-0"
                style="height: 100%;"
              >
                <div
                  class="mb-auto text-right"
                  style="height: 36px"
                >
                  <v-menu
                    bottom
                    left
                    internal-activator
                  >
                    <template v-slot:activator="{ on, attrs }">
                      <v-btn
                        icon
                        v-bind="attrs"
                        v-on="on"
                      >
                        <v-icon v-if="isCardHover" x-large>mdi-dots-vertical</v-icon>
                      </v-btn>
                    </template>

                    <v-list>
                      <v-list-item
                        v-for="(menuItem, i) in menuItems"
                        :key="i"
                        :class="{
                          'lwtg-danger-bg': menuItem.danger,
                        }"
                        link
                        @click="$emit(menuItem.action)"
                      >
                        <v-list-item-title
                          :class="{
                            'white-100-color': menuItem.danger,
                          }"
                        >{{ menuItem.title }}</v-list-item-title>
                      </v-list-item>
                    </v-list>
                  </v-menu>
                </div>
                <div
                  class="d-flex justify-end"
                  style="height: 36px;"
                >
                  <lwtg-icon-button
                    v-if="isCardHover"
                    size="large"
                    :src="eyeIconButton"
                    @click="wordVisibility"
                  />
                  <lwtg-icon-button
                    v-if="isCardHover"
                    size="large"
                    :src="checkIconButton"
                    @click="patchRemember"
                  />
                </div>
              </div>
            </v-col>
          </v-row>
        </template>
      </lwtg-word-card>
    </div>
  </v-hover>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';
import LwtgIconButton from '@/components/atoms/LwtgIconButton.vue';
import LwtgWordCard from '@/components/atoms/LwtgWordCard.vue';
import { Word } from '@/models/word';
import { MenuItem } from '@/models/types/menuItem';

@Component({
  components: {
    LwtgIconButton,
    LwtgWordCard,
  },
})
export default class WordListCard extends mixins(UtilMixin) {
  @Prop() width!: number;
  @Prop({ default: '10px' }) border!: number;
  @Prop() primary!: boolean;
  @Prop() word!: Word;
  @Prop() menuItems!: MenuItem[];
  @Prop({ default: -1 }) index!: number;
  @Prop() isFlipped!: boolean;
  tempWidth!: number;
  flipAnimation = false;
  eyeIconSrcList = [
    require('@/assets/word/eye-off.svg'),
    require('@/assets/word/eye.svg'),
  ]
  checkIconSrcList = [
    require('@/assets/word/check.svg'),
    require('@/assets/word/check-green.svg'),
  ]
  isCardHover = false;
  isEyeButtonHover = false;
  isWordRevealed = true;
  get wordCardSrc() {
    return require('@/assets/word/word.svg');
  }
  get eyeIconButton() {
    const i =
      Number(this.isWordRevealed);
    return this.eyeIconSrcList[i];
  }
  get checkIconButton() {
    const i =
      Number(this.word.isRemembered);
    return this.checkIconSrcList[i];
  }
  get wordFont() {
    if (this.tempWidth) {
      this.width = this.tempWidth;
    }
    return this.width >= 1000
      ? 'font-size: 64px;'
      : this.width >= 500
      ? 'font-size: 32px;'
      : 'font-size: 24px;'
  }
  get meaningFont() {
    if (this.tempWidth) {
      this.width = this.tempWidth;
    }
    return this.width >= 1000
      ? 'font-size: 36px;'
      : this.width >= 500
      ? 'font-size: 20px;'
      : 'font-size: 18px;'
  }
  get explanationFont() {
    if (this.tempWidth) {
      this.width = this.tempWidth;
    }
    return this.width >= 1000
      ? 'font-size: 24px;'
      : this.width >= 500
      ? 'font-size: 14px;'
      : 'font-size: 12px;'
  }
  mounted() {
    window.addEventListener('resize', this.handleResize);
    if (!this.width) {
      this.handleResize();
    }
  }
  handleResize() {
    const wordCard = document.getElementById(
    `word-card-${this.index}`
    );
    if (wordCard) {
      const wordCardWidth = wordCard.getBoundingClientRect().width;
      this.width = wordCardWidth;
      this.tempWidth = wordCardWidth;
    }
  }
  flipCard() {
    if (this.isFlipped) {
      this.isFlipped = false;
    } else {
      this.isFlipped = true;
    }
  }
  wordVisibility() {
    if (this.isWordRevealed) {
      this.isWordRevealed = false;
    } else {
      this.isWordRevealed = true;
    }
  }
}
</script>

<style scoped lang="scss">
@import '@/style.scss';
.flip-card {
  background-color: transparent;
  width: 100%;
  height: 100%;
  perspective: 2000px;
}
.flipped-card {
  transform: rotateY(-180deg);
}
.flip-card-inner {
  cursor: pointer;
  position: relative;
  width: 100%;
  height: 100%;
  transition: transform 0.3s;
  transform-style: preserve-3d;
}
.flip-card-front, .flip-card-back {
  position: absolute;
  width: 100%;
  height: 100%;
  -webkit-backface-visibility: hidden; /* Safari */
  backface-visibility: hidden !important;
}
.flip-card-back {
  transform: rotateY(180deg);
}
</style>
