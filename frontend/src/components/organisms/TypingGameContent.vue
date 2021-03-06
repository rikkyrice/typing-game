<template>
  <div
    id="typing-game-content"
    class="lwtg-secondary-bg px-12 py-6"
    style="width: 100%;"
  >
    <v-row>
      <v-col cols="7">
        <div class="d-flex flex-column lwtg-white-bg pa-6">
          <div>
            <span
              class="bold main-mono-color"
              :style="fontSizeUtil(24, 24, 20)"
            >{{ wordList.wordListTitle }}</span>
          </div>
          <div>
            <span
              class="main-mono-color"
              :style="fontSizeUtil(14, 14, 12)"
            >{{ wordList.explanation }}</span>
          </div>
        </div>
      </v-col>
      <v-col cols="5">
        <div class="lwtg-white-bg pa-6" style="height: 100%;"></div>
      </v-col>
    </v-row>
    <div class="lwtg-white-bg pa-6">
      <v-row>
        <v-spacer />
        <v-col cols="10">
          <typing-word-card
            ref="childWords"
            :border="'10'"
            :primary="true"
            :words="wordArray.words"
            :isActivated="shiftController"
            @shift="shift"
            @reset="reset"
          />
        </v-col>
        <v-col cols="1">
          <div @click="reset">
            <span
              class="bold main-mono-color"
              style="cursor: pointer;"
              :style="fontSizeUtil(14, 14, 12)"
            >リスタート</span>
          </div>
        </v-col>
      </v-row>
      <v-row>
        <v-spacer />
        <v-col cols="10">
          <typing-word-card
            ref="childMeanings"
            :border="'10'"
            :primary="true"
            :words="wordArray.words"
            :isWords="false"
            :isActivated="!shiftController"
            :wordOnly="wordOnly"
            @shift="shift"
            @reset="reset"
          />
        </v-col>
        <v-spacer />
      </v-row>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';
import TypingWordCard from '@/components/organisms/card/TypingWordCard.vue';
import { WordList } from '@/models/wordlist';
import { WordArray } from '@/models/word';
import store from '@/store';
import { TYPES } from '@/store/mutation-types';

@Component({
  components: {
    TypingWordCard,
  },
})
export default class TypingGameContent extends mixins(UtilMixin) {
  @Prop() wordList!: WordList;
  @Prop() wordArray!: WordArray;
  shiftController = true;
  wordOnly = false;
  created() {
    store.dispatch(TYPES.SWITCH_CLEARED, false);
  }
  refs():any {
    return this.$refs;
  }
  shift() {
    if(!this.wordOnly) {
      this.shiftController = !this.shiftController;
    }
    if (this.shiftController || this.wordOnly) {
      this.refs().childWords.shiftIndex();
      this.refs().childMeanings.shiftIndex();
    }
  }
  reset() {
    this.shiftController = true;
    this.refs().childWords.reset();
    this.refs().childMeanings.reset();
  }
}
</script>

<style scoped lang="scss">
@import '@/style.scss';
</style>
