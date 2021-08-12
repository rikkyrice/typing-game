<template>
  <div id="mypage">
    <!-- パンくずリスト -->
    <div class="px-16 lwtg-white-bg" style="width: 100%;">
      <lwtg-breadcrumbs
        :class="MdSmXsUtil('py-4', 'py-3', 'py-3')"
        :breadcrumbs="breadcrumbs"
      />
    </div>
    <div class="pa-16 lwtg-secondary-bg" style="width: 100%;">
      <word-list-content :wordListArray="getWordLists" />
    </div>
    <lwtg-loader :page-loading="true" :loading="false" />
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';
import WordListApi from '@/api/wordlist';
import LwtgBreadcrumbs from '@/components/atoms/LwtgBreadcrumbs.vue';
import LwtgLoader from '@/components/atoms/LwtgLoader.vue';
import WordListContent from '@/components/organisms/WordListContent.vue';
import { BreadcrumbInfo } from '@/models/types/breadcrumbInfo';
import { WordListArray } from '@/models/wordlist';
import store from '@/store';
import { TYPES } from '@/store/mutation-types';

@Component({
  components: {
    LwtgBreadcrumbs,
    LwtgLoader,
    WordListContent,
  },
})
export default class MyPageView extends mixins(UtilMixin) {
  breadcrumbs: BreadcrumbInfo[] = [
    { label: 'トップ', path: '/' },
    { label: 'My Page', path: '/mypage' },
  ];
  wordListArray: WordListArray = {
    matched: 0,
    wordlists: [],
  }
  userId = store.state.auth.userId;
  wordListLoading = false;

  get viewLoading() {
    return (
      this.wordListLoading
    );
  }
  get getWordLists() {
    return this.wordListArray;
  }
  created() {
    this.fetchGetWordList();
  }
  fetchGetWordList() {
    this.wordListLoading = true;
    WordListApi.getWordLists()
      .then((data) => {
        this.wordListArray = data;
      })
      .finally(() => (this.wordListLoading = false));
  }
}
</script>

<style scoped lang="scss">
@import '@/style.scss';
</style>
