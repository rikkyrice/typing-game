<template>
  <div id="wordlist">
    <!-- パンくずリスト -->
    <div class="px-16 lwtg-white-bg" style="width: 100%;">
      <lwtg-breadcrumbs
        :class="MdSmXsUtil('py-4', 'py-3', 'py-3')"
        :breadcrumbs="breadcrumbs"
      />
    </div>
    <div class="pa-0" style="width: 100%;">
      <word-content
        :wordList="wordList"
        :wordArray="getWordArray"
      />
    </div>
    <lwtg-word-create-button />
    <lwtg-loader :page-loading="true" :loading="viewLoading" />
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
import WordContent from '@/components/organisms/WordContent.vue';
import UtilMixin from '@/mixins/utilMixin';
import { BreadcrumbInfo } from '@/models/types/breadcrumbInfo';
import { WordList } from '@/models/wordlist';
import { WordArray, initializedWordArray } from '@/models/word';

@Component({
  components: {
    LwtgBreadcrumbs,
    LwtgLoader,
    LwtgWordCreateButton,
    WordContent,
  },
})
export default class WordListView extends mixins(UtilMixin) {
  wordList: WordList = {} as WordList;
  wordArray: WordArray = {
    matched: 0,
    words: [],
  }
  wordListLoading = false;
  wordsLoading = false;
  get getWordArray() {
    console.log(this.wordArray);
    return this.wordArray.matched !== 0
      ? this.wordArray
      : initializedWordArray
  }
  get viewLoading() {
    return (
      this.wordListLoading ||
      this.wordsLoading
    );
  }
  created() {
    this.fetchGetWordList();
    this.fetchGetWords();
  }
  fetchGetWordList() {
    this.wordListLoading = true;
    WordListApi.getWordList(this.$route.params.wordlistId)
      .then((data) => {
        this.wordList = data;
        this.breadcrumbs.push(
          { label: `${this.wordList.title}`, path: `/mypage/${this.wordList.id}` },
        );
      })
      .finally(() => (this.wordListLoading = false));
  }
  fetchGetWords() {
    this.wordsLoading = true;
    WordApi.getWords(this.$route.params.wordlistId)
      .then((data) => {
        this.wordArray = data;
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
