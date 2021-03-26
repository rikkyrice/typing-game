<template>
  <div class="lwtg-typing-game px-6">
    <div class="d-flex flex-column">
      <div class="text-center">
        <span
          class="bold"
          :class="{
            'main-mono-color': isActivated || wordOnly,
            'mono-30-color': !isActivated && !wordOnly,
          }"
          :style="wordFontStyle"
        >{{ typeWord.word }}</span>
      </div>
      <!-- <div v-if="typeWord.yomi" class="text-center">
        <span
          :class="{
            'main-mono-color': isActivated,
            'mono-30-color': !isActivated,
          }"
          :style="fontSizeUtil(20, 20, 16)"
        >{{ typeWord.yomi }}</span>
      </div> -->
      <v-divider />
      <div class="d-flex justify-center">
        <div class="d-flex flex-wrap">
          <span
            v-for="(cw, i) in clearedWords"
            :key="i"
            class="bold mono-30-color"
            :style="typeFontStyle"
          >
            <span v-if="cw !== ' '"
            >{{ cw }}</span>
            <span v-else class="space"></span>
          </span>
          <span
            v-for="(tw, j) in typeWords"
            :key="j"
          >
            <span
              v-for="(t, k) in tw[0]"
              :key="k"
              :class="{
                'focused-word': j === nextIndex && k === 0 && isActivated,
                'main-mono-color': isActivated,
                'mono-30-color': !isActivated,
              }"
              class="bold"
              :style="typeFontStyle"
            >
              <span v-if="t !== ' '">{{ t }}</span>
              <span v-else class="space"></span>
            </span>
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';
import { TypeWord } from '@/models/types/typeWord';
import store from '@/store';
import { TYPES } from '@/store/mutation-types';

@Component
export default class LwtgTypingGame extends mixins(UtilMixin) {
  @Prop() typeWord!: TypeWord;
  @Prop() isActivated!: boolean;
  @Prop() wordOnly!: boolean;
  typeWords = JSON.parse(JSON.stringify(this.typeWord.typeWord));
  clearedWords: string[] = [];
  typedWord = '';
  nextIndex = 0;
  get wordFontStyle() {
    return this.typeWord.word.length < 30
        ? this.fontSizeUtil(32, 32, 26)
        : this.fontSizeUtil(24, 24, 20);
  }
  get typeFontStyle() {
    return this.typeWord.word.length < 30
        ? this.fontSizeUtil(24, 24, 20)
        : this.fontSizeUtil(16, 16, 10);
  }
  mounted() {
    window.addEventListener('keydown', e => {
      if (this.isActivated && !store.state.gameCleared) {
        this.keyDown(e.key);
      }
    });
  }
  keyDown(code: string) {
    // typeすべき単語の最初の一文字をリストに格納する。
    var target = this.typeWords[this.nextIndex].map((str: string) => str[0]);
    // キーコードを取得し、targetリストに含まれるか判定。
    // 間違っていたら終了。
    this.typedWord = code;
    if(!target.includes(code)) {
      return false;
    }
    // タイプワードがあればクリアワードにキーコードを格納
    this.clearedWords.push(code);
    // 今後判定する必要のないタイプワードは削除するためそのindexを取得
    var removeList = this.findRemoveIndex(target, code);
    // 削除対象のindexを削除
    for (var i = 0; i < removeList.length; i++) {
      this.typeWords[this.nextIndex].splice(removeList[i], 1);
      removeList = removeList.map((v) => (v -= 1));
    }
    // 残ったそれぞれのタイプワードの一文字目を削除
    this.typeWords[this.nextIndex] = this.typeWords[this.nextIndex].map((str: string) => str.slice(1));
    // すでに空文字となったタイプワードがあれば次の文字に進む判定
    if (this.typeWords[this.nextIndex].includes('')) {
      this.nextIndex++;
    }
    // すべての文字をタイプし終えたら完了し、次の単語に進む
    if (this.nextIndex === this.typeWords.length) {
      this.complete();
      return false;
    }
  }
  // 今後判定する必要のないタイプワードは削除するためそのindexを取得
  findRemoveIndex(target: string[], elm: string) {
    var removeList: number[] = [];
    for (var i = 0; i < target.length; i++) {
      if (target[i] !== elm) {
        removeList.push(i);
      }
    }
    return removeList;
  }
  complete() {
    this.$emit('shift');
  }
  @Watch('typeWord')
  setTypeWords(newTypeWord: TypeWord) {
    this.clearedWords = [];
    this.nextIndex = 0;
    this.typeWords = JSON.parse(JSON.stringify(newTypeWord.typeWord));
  }
}
</script>

<style scoped lang="scss">
@import '@/style.scss';
.lwtg-typing-game {
  width: 100%;
  height: 100%;
}
.focused-word {
  border-bottom: solid 2px #666666;
}
.space {
  display: inline-block;
  width: 10px;
}
</style>
