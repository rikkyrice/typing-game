<template>
  <div id="wordlist">
    <!-- パンくずリスト -->
    <div class="px-16 lwtg-white-bg" style="width: 100%;">
      <lwtg-breadcrumbs
        :class="MdSmXsUtil('py-4', 'py-3', 'py-3')"
        :breadcrumbs="breadcrumbs"
      />
    </div>
    <div class="px-16 mb-6 lwtg-white-bg" style="width: 100%;">
      <span class="bold main-mono-color" :style="fontSizeUtil(28, 28, 24)"
      >{{ wordList.wordListTitle }}</span>
    </div>
    <div class="pa-0" style="width: 100%;">
      <word-content
        :wordList="wordList"
        :wordArray="wordArray"
      />
    </div>
    <lwtg-word-create-button />
    <lwtg-loader :page-loading="true" :loading="false" />
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import LwtgBreadcrumbs from '@/components/atoms/LwtgBreadcrumbs.vue';
import LwtgLoader from '@/components/atoms/LwtgLoader.vue';
import LwtgWordCreateButton from '@/components/atoms/LwtgWordCreateButton.vue';
import WordContent from '@/components/organisms/WordContent.vue';
import UtilMixin from '@/mixins/utilMixin';
import { BreadcrumbInfo } from '@/models/types/breadcrumbInfo';
import { WordList } from '@/models/wordlist';
import { WordArray } from '@/models/word';

@Component({
  components: {
    LwtgBreadcrumbs,
    LwtgLoader,
    LwtgWordCreateButton,
    WordContent,
  },
})
export default class WordListView extends mixins(UtilMixin) {
  wordList: WordList = {
      id: '1',
      title: 'TOEIC',
      explanation: 'TOEICの頻出単語を単語帳にしました！ぜひ使って990点めざしましょう!',
      createdAt: '2021-02-06-00:00:00',
      updatedAt: '2021-02-06-00:00:00',
  }
  wordArray: WordArray = {
    matched: 4,
    words: [
      {
        id: '8e97230e-2c9d-44ed-9a95-e39aa6217dd7',
        word: 'organization',
        yomi: '',
        meaning: '組織、企業',
        mYomi: 'かきくけこ',
        explanation: 'おーがにぜーしょん',
        isRemembered: true,
        createdAt: '2020-10-03-21.50.00.000000',
        updatedAt: '2020-10-03-21.50.00.000000',
      },
      {
        id: '90443017-7096-4e67-9ada-863f581bd1a7',
        word: 'communication',
        yomi: '',
        meaning: 'コミュニケーション',
        mYomi: 'かきくけこ',
        explanation: '大事',
        isRemembered: false,
        createdAt: '2020-10-03-21.51.00.000000',
        updatedAt: '2020-10-03-21.51.00.000000',
      },
      {
        id: 'b5f888bd-d8c6-4d74-a61c-ef838d8b52d9',
        word: 'government',
        yomi: '',
        meaning: '陰謀ありあり',
        mYomi: 'かきくけこ',
        explanation: '政府',
        isRemembered: false,
        createdAt: '2020-10-03-21.52.00.000000',
        updatedAt: '2020-10-03-21.52.00.000000',
      },
      {
        id: '48d4becc-f73c-45f7-8cb5-a9dee4c4f214',
        word: 'authentication',
        yomi: '',
        meaning: 'ここには何でも',
        mYomi: 'かきくけこ',
        explanation: '認証',
        isRemembered: true,
        createdAt: '2020-10-03-21.53.00.000000',
        updatedAt: '2020-10-03-21.53.00.000000',
      }
    ]
  }
  breadcrumbs: BreadcrumbInfo[] = [
    { label: 'トップ', path: '/' },
    { label: 'My Page', path: '/mypage' },
    { label: `${this.wordList.title}`, path: `/mypage/${this.wordList.id}` },
  ];
}
</script>

<style scoped lang="scss">
@import '@/style.scss';
</style>
