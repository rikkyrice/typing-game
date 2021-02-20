<template>
  <div id="typing-game-view">
    <!-- パンくずリスト -->
    <div class="px-16 lwtg-white-bg" style="width: 100%;">
      <lwtg-breadcrumbs
        :class="MdSmXsUtil('py-4', 'py-3', 'py-3')"
        :breadcrumbs="breadcrumbs"
      />
    </div>
    <div class="px-16 mb-6 lwtg-white-bg" style="width: 100%;">
      <span class="bold main-mono-color" :style="fontSizeUtil(28, 28, 24)"
      >タイピングゲーム</span>
    </div>
    <div class="pa-0" style="width: 100%;">
      <typing-game-content
        :wordList="wordList"
        :wordArray="wordArray"
      />
    </div>
    <lwtg-loader :page-loading="true" :loading="false" />
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import LwtgBreadcrumbs from '@/components/atoms/LwtgBreadcrumbs.vue';
import LwtgLoader from '@/components/atoms/LwtgLoader.vue';
import LwtgWordCreateButton from '@/components/atoms/LwtgWordCreateButton.vue';
import TypingGameContent from '@/components/organisms/TypingGameContent.vue';
import UtilMixin from '@/mixins/utilMixin';
import { BreadcrumbInfo } from '@/models/types/breadcrumbInfo';
import { WordList } from '@/models/wordlist';
import { WordArray } from '@/models/word';

@Component({
  components: {
    LwtgBreadcrumbs,
    LwtgLoader,
    LwtgWordCreateButton,
    TypingGameContent,
  },
})
export default class WordListView extends mixins(UtilMixin) {
  wordList: WordList = {
      wordListId: '1',
      wordListTitle: 'TOEIC',
      explanation: 'TOEICの頻出単語を単語帳にしました！ぜひ使って990点めざしましょう!',
      createdAt: '2021-02-06-00:00:00',
      updatedAt: '2021-02-06-00:00:00',
  }
  wordArray: WordArray = {
    matched: 5,
    words: [
      {
        wordId: '8e97230e-2c9d-44ed-9a95-e39aa6217dd7',
        word: '愛飢えお',
        yomi: 'あいうえおきゅきゅ',
        meaning: 'かきくけこ',
        mYomi: 'かきくけこ',
        explanation: '50音最初の5文字',
        isRemembered: true,
        createdAt: '2020-10-03-21.50.00.000000',
        updatedAt: '2020-10-03-21.50.00.000000',
      },
      {
        wordId: '8e97230e-2c9d-44ed-9a95-e39aa6217dd7',
        word: 'oc login',
        yomi: '',
        meaning: 'かきくけこ',
        mYomi: 'かきくけこ',
        explanation: 'おーがにぜーしょん',
        isRemembered: true,
        createdAt: '2020-10-03-21.50.00.000000',
        updatedAt: '2020-10-03-21.50.00.000000',
      },
      {
        wordId: '90443017-7096-4e67-9ada-863f581bd1a7',
        word: 'communication',
        yomi: '',
        meaning: 'かきくけこ',
        mYomi: 'かきくけこ',
        explanation: '大事',
        isRemembered: false,
        createdAt: '2020-10-03-21.51.00.000000',
        updatedAt: '2020-10-03-21.51.00.000000',
      },
      {
        wordId: 'b5f888bd-d8c6-4d74-a61c-ef838d8b52d9',
        word: 'government',
        yomi: '',
        meaning: 'かきくけこ',
        mYomi: 'かきくけこ',
        explanation: '政府',
        isRemembered: false,
        createdAt: '2020-10-03-21.52.00.000000',
        updatedAt: '2020-10-03-21.52.00.000000',
      },
      {
        wordId: '48d4becc-f73c-45f7-8cb5-a9dee4c4f214',
        word: 'authentication',
        yomi: '',
        meaning: 'かきくけこ',
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
    { label: `${this.wordList.wordListTitle}`, path: `/mypage/${this.wordList.wordListId}` },
    { label: `タイピングゲーム`, path: `/mypage/${this.wordList.wordListId}/typing-game}` },
  ];
}
</script>

<style scoped lang="scss">
@import '@/style.scss';
</style>
