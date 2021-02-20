<template>
  <div id="mypage">
    <!-- パンくずリスト -->
    <div class="px-16 lwtg-white-bg" style="width: 100%;">
      <lwtg-breadcrumbs
        :class="MdSmXsUtil('py-4', 'py-3', 'py-3')"
        :breadcrumbs="breadcrumbs"
      />
    </div>
    <div class="px-16 mb-6 lwtg-white-bg" style="width: 100%;">
      <span class="bold main-mono-color" :style="fontSizeUtil(28, 28, 24)"
      >マイページ</span>
    </div>
    <div class="pa-16 lwtg-secondary-bg" style="width: 100%;">
      <word-list-content :wordListArray="wordListArray" />
    </div>
    <lwtg-loader :page-loading="true" :loading="false" />
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';
import LwtgBreadcrumbs from '@/components/atoms/LwtgBreadcrumbs.vue';
import LwtgLoader from '@/components/atoms/LwtgLoader.vue';
import WordListContent from '@/components/organisms/WordListContent.vue';
import { BreadcrumbInfo } from '@/models/types/breadcrumbInfo';
import { WordListArray } from '@/models/wordlist';

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
    matched: 3,
    wordLists: [
      { wordListId: '1',
        wordListTitle: 'TOEIC',
        explanation: 'TOEICの頻出単語を単語帳にしました！ぜひ使って990点めざしましょう!',
        createdAt: '2021-02-06-00:00:00',
        updatedAt: '2021-02-06-00:00:00',
      },
      { wordListId: '2',
        wordListTitle: '基本情報',
        explanation: 'TOEICの頻出単語を単語帳にしました！ぜひ使って990点めざしましょう!',
        createdAt: '2021-02-06-00:00:00',
        updatedAt: '2021-02-06-00:00:00',
      },
      { wordListId: '3',
        wordListTitle: 'OpenShift',
        explanation: 'TOEICの頻出単語を単語帳にしました！ぜひ使って990点めざしましょう!',
        createdAt: '2021-02-06-00:00:00',
        updatedAt: '2021-02-06-00:00:00',
      },
    ]
  }
}
</script>

<style scoped lang="scss">
@import '@/style.scss';
</style>
