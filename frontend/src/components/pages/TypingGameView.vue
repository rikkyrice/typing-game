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
      id: '1',
      title: 'TOEIC',
      explanation: 'TOEICの頻出単語を単語帳にしました！ぜひ使って990点めざしましょう!',
      createdAt: '2021-02-06-00:00:00',
      updatedAt: '2021-02-06-00:00:00',
  }
  wordArray: WordArray = {
    matched: 1,
    words: [
      {
        id: '8e97230e-2c9d-44ed-9a95-e39aa6217dd7',
        word: '橋木陸',
        yomi: 'はしきりく',
        meaning: '2020年度IBMの社員',
        mYomi: '2020ねんどIBMのしゃいん',
        explanation: '2020年に日本IBMに入社した社員の一人。',
        isRemembered: true,
        createdAt: '2020-10-03-21.50.00.000000',
        updatedAt: '2020-10-03-21.50.00.000000',
      },
    ]
  }
  breadcrumbs: BreadcrumbInfo[] = [
    { label: 'トップ', path: '/' },
    { label: 'My Page', path: '/mypage' },
    { label: `${this.wordList.title}`, path: `/mypage/${this.wordList.id}` },
    { label: `タイピングゲーム`, path: `/mypage/${this.wordList.id}/typing-game}` },
  ];
}
</script>

<style scoped lang="scss">
@import '@/style.scss';
</style>
