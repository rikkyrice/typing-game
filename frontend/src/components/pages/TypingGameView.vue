<template>
  <div id="typing-game-view" class="lwtg-secondary-bg" style="height: 100%;">
    <!-- パンくずリスト -->
    <div class="px-16 lwtg-white-bg" style="width: 100%;">
      <lwtg-breadcrumbs
        :class="MdSmXsUtil('py-4', 'py-3', 'py-3')"
        :breadcrumbs="breadcrumbs"
      />
    </div>
    <div class="pa-0" style="width: 100%;">
      <typing-game-content
        :wordList="wordList"
        :typingWords="getTypingWords"
      />
    </div>
    <lwtg-loader :page-loading="true" :loading="false" />
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import WordListApi from '@/api/wordlist';
import WordApi from '@/api/word';
import LwtgBreadcrumbs from '@/components/atoms/LwtgBreadcrumbs.vue';
import LwtgLoader from '@/components/atoms/LwtgLoader.vue';
import LwtgWordCreateButton from '@/components/atoms/LwtgWordCreateButton.vue';
import TypingGameContent from '@/components/organisms/TypingGameContent.vue';
import UtilMixin from '@/mixins/utilMixin';
import { BreadcrumbInfo } from '@/models/types/breadcrumbInfo';
import { WordList } from '@/models/wordlist';
import { TypingWordArray, initializedTypingWordArray } from '@/models/word';

@Component({
  components: {
    LwtgBreadcrumbs,
    LwtgLoader,
    LwtgWordCreateButton,
    TypingGameContent,
  },
})
export default class WordListView extends mixins(UtilMixin) {
  wordList: WordList = {} as WordList;
  typingWordArray: TypingWordArray = {
    matched: 0,
    typingWords: [],
  }
  wordListLoading = false;
  wordsLoading = false;
  get getTypingWords() {
    return this.typingWordArray.matched !== 0
      ? this.typingWordArray.typingWords
      : initializedTypingWordArray.typingWords
  }
  get viewLoading() {
    return (
      this.wordListLoading ||
      this.wordsLoading
    );
  }
  created() {
    this.fetchGetWordList();
    this.fetchGetTypingWords();
  }
  fetchGetWordList() {
    this.wordListLoading = true;
    WordListApi.getWordList(this.$route.params.wordlistId)
      .then((data) => {
        this.wordList = data;
        this.breadcrumbs.push(
          { label: `${this.wordList.title}`, path: `/mypage/${this.wordList.id}` },
          { label: `タイピングゲーム`, path: `/mypage/${this.wordList.id}/typing-game}` },
        );
      })
      .finally(() => (this.wordListLoading = false));
  }
  fetchGetTypingWords() {
    this.wordsLoading = true;
    WordApi.getTypingWords(this.$route.params.wordlistId)
      .then((data) => {
        this.typingWordArray = data;
        for (var i = this.typingWordArray.typingWords.length; i > 1; i--) {
          var k = Math.floor(Math.random() * i);
          [this.typingWordArray.typingWords[k], this.typingWordArray.typingWords[i - 1]] = [this.typingWordArray.typingWords[i - 1], this.typingWordArray.typingWords[k]];
        }
      })
      .finally(() => (this.wordsLoading = false));
  }
  breadcrumbs: BreadcrumbInfo[] = [
    { label: 'トップ', path: '/' },
    { label: 'My Page', path: '/mypage' },
  ];
}
</script>

<style scoped lang="scss">
@import '@/style.scss';
</style>
