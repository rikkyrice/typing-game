<template>
  <v-container
    id="wordListContent"
    class="lwtg-white-bg"
  >
    <div class="pl-12 py-6">
      <v-row class="px-12">
        <v-col cols="11" class="pa-0">
          <div class="d-flex align-center">
            <div class="mr-auto">
              <span
                :style="fontSizeUtil(24, 24, 20)"
                class="bold main-mono-color"
              >単語帳</span>
              <v-btn
                :icon="true"
                :x-small="true"
                class="ml-4"
              >
                <v-icon size="20px" color="A0D0A0"
                >mdi-sort</v-icon>
              </v-btn>
            </div>
            <lwtg-chip
              text="+ 作成"
              color="#666666"
              class="mr-1 bold"
              :clickable="true"
              :outlined="true"
            />
          </div>
        </v-col>
        <v-spacer />
      </v-row>
      <v-row v-for="wordList in wordListArray.wordlists" :key="wordList.id" class="px-12">
        <v-col cols="11" class="pa-0">
          <word-list-card
            :wordlist="wordList"
          />
        </v-col>
        <v-col cols="1">
          <div class="d-flex justify-center align-center">
            <div class="d-flex flex-column play-content">
              <span class="play-text bold">PLAY!</span>
              <v-btn
                :icon="true"
                :large="true"
                class="play-button"
                @click="playTypingGame(wordList.id)"
              >
                <v-icon size="36px" color="#A0D0A0"
                >mdi-arrow-right-circle</v-icon>
              </v-btn>
            </div>
          </div>
        </v-col>
      </v-row>
      <v-row class="px-12">
        <v-col cols="12" class="pa-0">
          <div class="text-center">
            <v-pagination
              v-model="page"
              :length="wordListArray.matched / 3"
            ></v-pagination>
          </div>
        </v-col>
      </v-row>
    </div>
  </v-container>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';
import LwtgIconButton from '@/components/atoms/LwtgIconButton.vue';
import LwtgChip from '@/components/atoms/LwtgChip.vue';
import WordListCard from '@/components/organisms/card/WordListCard.vue';
import { WordListSummaryArray } from '@/models/wordlist';

@Component({
  components: {
    LwtgIconButton,
    LwtgChip,
    WordListCard,
  },
})
export default class WordListContent extends mixins(UtilMixin) {
  @Prop() wordListArray!: WordListSummaryArray;
  playTypingGame(id: string) {
    this.$router.push(`mypage/${id}/typing-game`, () => {});
  }
}
</script>

<style scoped lang="scss">
@import '@/style.scss';
.play-content {
  width: 100%;
  height: 100%;
}
.play-text {
  cursor: pointer;
  opacity: 0;
  transition: opacity 0.2s;
}
.play-content:hover .play-text {
  opacity: 1.0;
}
.play-button {
  transition: 0.2s;
}
.play-content:hover .play-button {
  border: solid 2px #666666 !important;
}
</style>
