<template>
  <div
    id="wordContent"
    class="lwtg-secondary-bg"
    style="width: 100%;"
  >
    <div class="px-12 pt-12">
      <div
        class="d-flex flex-column mx-auto px-6 pt-3 pb-8 lwtg-white-bg"
        style="width: 600px;"
      >
        <v-row style="width: 100%;">
          <v-spacer />
          <v-col cols="6" style="text-align: center;">
            <span class="bold" :style="fontSizeUtil(24, 24, 20)"
            >{{ wordList.title }}</span>
          </v-col>
          <v-col cols="3">
            <lwtg-icon-button
              size="small"
              :icon="'lead-pencil'"
            />
          </v-col>
        </v-row>
        <div style="text-align: center;width: 100%;">
          <span :style="fontSizeUtil(14, 14, 12)"
          >{{ wordList.explanation }}</span>
        </div>
      </div>
    </div>
    <div class="px-12 pb-12 lwtg-secondary-bg" style="width: 100%;">
      <div class="pa-12 lwtg-secondary-bg">
        <div class="d-flex pa-2 lwtg-secondary-bg">
          <span class="bold main-mono-color" :style="fontSizeUtil(24, 24, 20)"
          >単語</span>
          <span
            class="bold ml-auto pt-3 main-mono-color"
            style="cursor: pointer;"
            :style="fontSizeUtil(14, 14, 12)"
            @click="changeList"
          >一覧を表示</span>
        </div>
        <div class="lwtg-primary-bg" style="height: 2px;" />
      </div>
      <v-row v-if="isShowEach" class="lwtg-secondary-bg">
        <v-col cols="1" class="d-flex justify-center align-center">
          <v-icon
            v-if="index > 0"
            color="#A0D0A0"
            size="72px"
            @click="moveBack"
          >mdi-chevron-left</v-icon>
        </v-col>
        <v-col cols="10" class="d-flex justify-center align-center">
          <word-card
            :word="getWordCard"
            :border="'10'"
            :primary="true"
            :isFlipped="isFlipped"
            :menuItems="menuItemsEach"
          />
        </v-col>
        <v-col cols="1" class="d-flex justify-center align-center">
          <v-icon
            v-if="index < wordArray.matched - 1"
            color="#A0D0A0"
            size="72px"
            @click="moveForward"
          >mdi-chevron-right</v-icon>
        </v-col>
      </v-row>
      <v-row v-else class="lwtg-secondary-bg">
        <v-spacer />
        <v-col cols="10" class="d-flex flex-wrap">
          <div
            v-for="(word, i) in wordArray.words"
            :key="word.id"
            class="ma-1"
          >
            <word-card
              :word="word"
              :width="'450'"
              :border="'5'"
              :primary="true"
              :index="i"
              :menuItems="menuItemsAll"
              @detail='detail(i)'
            />
          </div>
        </v-col>
        <v-spacer />
      </v-row>
      <div style="height: 100px; width: 100%;" />
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';
import LwtgIconButton from '@/components/atoms/LwtgIconButton.vue';
import WordCard from '@/components/organisms/card/WordCard.vue';
import { WordList } from '@/models/wordlist';
import { WordArray } from '@/models/word';
import { MenuItem } from '@/models/types/menuItem';

@Component({
  components: {
    LwtgIconButton,
    WordCard,
  },
})
export default class WordContent extends mixins(UtilMixin) {
  @Prop() wordList!: WordList;
  @Prop() wordArray!: WordArray;
  menuItemsEach: MenuItem[] = [
    { title: '編集', danger: false, action: 'edit' },
    { title: '削除', danger: true, action: 'delete' },
  ]
  menuItemsAll: MenuItem[] = [
    { title: '詳細', danger: false, action: 'detail' },
    { title: '編集', danger: false, action: 'edit' },
    { title: '削除', danger: true, action: 'delete' },
  ]
  isShowEach = true;
  isFlipped = false;
  index = 0;
  get getWordCard() {
    return this.wordArray.words[this.index];
  }
  changeList() {
    if (this.isShowEach) {
      this.isShowEach = false;
    } else {
      this.isShowEach = true;
    }
  }
  detail(i: number) {
    this.index = i;
    this.isShowEach = true;
  }
  moveForward() {
    this.index += 1;
    this.isFlipped = false;
  }
  moveBack() {
    this.index -= 1;
    this.isFlipped = false;
  }
}
</script>

<style scoped lang="scss">
@import '@/style.scss';
</style>
