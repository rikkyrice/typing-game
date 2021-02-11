<template>
  <v-row class="flip-card-parent" justify="center" no-gutters>
    <v-hover v-model="isCardHover">
      <v-img
        :src="wordCardSrc"
        class="flip-card"
        :class="{
          'flipped-card': isFlipped,
        }"
      >
        <div
          class="flip-card-inner"
          :class="{
            'flipped-card': isFlipped,
          }"
          @click="flipCard"
        >
          <v-row class="flip-card-front" width="100%" height="100%" no-gutters>
            <v-spacer>
            </v-spacer>
            <v-col
              cols="10"
              class="pa-6 pl-12"
            >
              <div class="lwtg-white-bg" style="width: 100%;">
                <v-row>
                  <v-col
                    cols="12"
                    class="py-0"
                    style="text-align: right; height: 36px"
                  >
                    <v-menu
                      bottom
                      left
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
                        >
                          <v-list-item-title
                            :class="{
                              'white-100-color': menuItem.danger,
                            }"
                          >{{ menuItem.title }}</v-list-item-title>
                        </v-list-item>
                      </v-list>
                    </v-menu>
                  </v-col>
                </v-row>
                <v-row>
                  <v-col cols="12" class="d-flex align-center" style="height: 280px;">
                    <div>
                      <span v-if="isWordRevealed" :style="fontSizeUtil(64, 64, 48)"
                      >{{ word.word }}</span>
                    </div>
                  </v-col>
                </v-row>
                <v-row>
                  <v-col cols="12" class="d-flex align-center" style="height: 36px;">
                    <div class="ml-auto">
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
                  </v-col>
                </v-row>
              </div>
            </v-col>
          </v-row>
          <v-row class="flip-card-back" width="100%" height="100%" no-gutters>
            <v-col
              cols="10"
              class="pa-6 pr-12"
            >
              <div class="lwtg-white-bg" style="width: 100%;">
                <v-row>
                  <v-col
                    cols="12"
                    class="py-0"
                    style="text-align: right; height: 36px"
                  >
                    <v-menu
                      bottom
                      left
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
                        >
                          <v-list-item-title
                            :class="{
                              'white-100-color': menuItem.danger,
                            }"
                          >{{ menuItem.title }}</v-list-item-title>
                        </v-list-item>
                      </v-list>
                    </v-menu>
                  </v-col>
                </v-row>
                <v-row>
                  <v-col cols="12" class="d-flex align-center" style="height: 280px;">
                    <div>
                      <span v-if="isWordRevealed" :style="fontSizeUtil(64, 64, 48)"
                      >{{ word.word }}</span>
                    </div>
                  </v-col>
                </v-row>
                <v-row>
                  <v-col cols="12" class="d-flex align-center" style="height: 36px;">
                    <div class="ml-auto">
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
                  </v-col>
                </v-row>
              </div>
            </v-col>
            <v-spacer>
            </v-spacer>
          </v-row>
        </div>
      </v-img>
    </v-hover>
  </v-row>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';
import LwtgIconButton from '@/components/atoms/LwtgIconButton.vue';
import { Word } from '@/models/word';
import { MenuItem } from '@/models/types/menuItem';

@Component({
  components: {
    LwtgIconButton,
  },
})
export default class WordListCard extends mixins(UtilMixin) {
  // @Prop() word!: Word;
  word: Word = {
    wordId: '1',
    word: 'organization',
    meaning: '企業、組織',
    explanation: 'TOEICの頻出単語を単語帳にしました！ぜひ使って990点めざしましょう!',
    isRemembered: false,
    createdAt: '2021-02-06-00:00:00',
    updatedAt: '2021-02-06-00:00:00',
  }
  menuItems: MenuItem[] = [
    { title: '詳細', danger: false },
    { title: '編集', danger: false },
    { title: '削除', danger: true },
  ]
  flipAnimation = false;
  isFlipped = false;
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
  flipCard() {
    if (this.isFlipped) {
      this.isFlipped = false;
    } else {
      this.isFlipped = true;
    }
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
.flip-card-parent {
  background-color: transparent;
  perspective: 1000px;
}
.flip-card {
  max-width: 1200px;
  min-height: 100%;
  transition: transform 0.8s;
  transform-style: preserve-3d;
}
.flipped-card {
  transform: rotateY(-180deg);
}
.flip-card-inner {
  cursor: pointer;
  position: relative;
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
