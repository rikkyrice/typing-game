<template>
  <div
    id="typing-word-card"
    class="d-flex justify-center align-center"
    style="width: 100%; height: 100%;"
  >
    <lwtg-word-card
      :width="width"
      :border="border"
      :primary="primary"
      class="lwtg-word-card"
    >
      <template #word>
        <div
          style="width: 100%; height: 100%;"
          class="d-flex justify-center align-center"
        >
          <div
            v-for="(typeWord, i) in typeWords"
            :key="i"
          >
            <lwtg-typing-game
              ref="offspring"
              v-if="i === index"
              :typeWord="typeWord"
              :isActivated="isActivated"
              @shift="shiftIndex"
            />
          </div>
          <span v-if="clear">Clear!</span>
        </div>
      </template>
    </lwtg-word-card>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';
import LwtgWordCard from '@/components/atoms/LwtgWordCard.vue';
import LwtgTypingGame from '@/components/atoms/LwtgTypingGame.vue';
import { Word } from '@/models/word';
import { TypeWord } from '@/models/types/typeWord';
import { mp } from '@/models/types/kana';
import store from '@/store';
import { TYPES } from '@/store/mutation-types';

@Component({
  components: {
    LwtgWordCard,
    LwtgTypingGame,
  },
})
export default class TypingWordCard extends mixins(UtilMixin) {
  @Prop() width!: number;
  @Prop({ default: '10px' }) border!: number;
  @Prop() primary!: boolean;
  @Prop() words!: Word[];
  @Prop({ default: 0 }) index!: number;
  @Prop() isActivated!: boolean;
  typeWords: TypeWord[] = [];
  clear: boolean = false;
  refs():any {
    return this.$refs;
  }
  created() {
    this.createTypeWords();
  }
  mounted() {
    window.addEventListener('resize', this.handleResize);
    if (!this.width) {
      this.handleResize();
    }
  }
  handleResize() {
    const typingWordCard = document.getElementById(
      'typing-word-card'
    );
    if (typingWordCard) {
      const typingWordCardWidth = typingWordCard.getBoundingClientRect().width;
      if (typingWordCardWidth > 1000) {
        this.width = 1000;
      } else {
        this.width = typingWordCardWidth;
      }
    }
  }
  shiftIndex() {
    this.index += 1;
    if (this.index === this.typeWords.length) {
      this.clear = true;
    }
  }
  reset() {
    this.clear = false;
    this.index = 0;
    this.refs().offspring.reset();
  }
  // Wordsでfor文を作成
  // Yomiがあるか、ないか判定 -> アルファベットのみの場合変換の必要がない
  // Yomiがない場合、Stringの二次元配列に変換してtypeWordsにpush
  // Yomiがある場合、まずparseして細かいひらがなの配列に変換
  // その後typing用のアルファベットに変換し、typeWordsにpush
  createTypeWords() {
    for (var i = 0; i < this.words.length; i++) {
      var word = this.words[i].word;
      var yomi = '';
      var tw: string[][] = [];
      if (!this.words[i].yomi) {
        var letters = this.words[i].word.split('');
        for (var j = 0; j < letters.length; j++) {
          tw[j] = [];
          tw[j].push(letters[j]);
        }
        console.log(tw);
      } else {
        var parsedSentence = this.parseKanaSentence(this.words[i].yomi);
        var typingSentence = this.convertToTypingSentences(parsedSentence);
        tw = typingSentence;
        yomi = this.words[i].yomi;
      }
      var typeWord = new TypeWord(word, yomi, tw)
      this.typeWords.push(typeWord);
    }
  }
  parseKanaSentence(str: string) {
    var res: string[] = [];
    var i = 0;
    var uni, bi: string;
    while (i < str.length) {
      uni = str[i].toString();
      if (i + 1 < str.length) {
        bi = str[i].toString() + str[i + 1].toString();
      } else {
        bi = '';
      }
      if (mp[bi]) {
        i += 2;
        res.push(bi);
      } else {
        i++;
        res.push(uni);
      }
    }
    return res;
  }
  convertToTypingSentences(str: string[]) {
    var res: string[][] = [];
    var s, ns;
    for (var i = 0; i < str.length; i++) {
      s = str[i];
      if (i + 1 < str.length) {
        ns = str[i + 1];
      } else {
        ns = '';
      }
      var tmpList: string[] = [];
      if (s === 'ん') {
        var isValidSingleN: boolean;
        var nList = mp[s];
        if (str.length - 1 === i) {
          isValidSingleN = false;
        } else if (i + 1 < str.length && (
            ns === 'あ' || ns === 'い' || ns === 'う' || ns === 'え' || ns === 'お' ||
            ns === 'な' || ns === 'に' || ns === 'ぬ' || ns === 'ね' || ns === 'の' ||
            ns === 'や' || ns === 'ゆ' || ns === 'よ'
        )) {
          isValidSingleN = false;
        } else {
          isValidSingleN = true;
        }
        for (var t of nList) {
          if (!isValidSingleN && t === 'n') {
            continue;
          }
          tmpList.push(t);
        }
      } else if (s === 'っ') {
        var ltuList = mp[s];
        var nextList = mp[ns];
        var hs: string[] = [];
        for (var v of nextList) {
          hs.push(v[0]);
        }
        var ltuTypeList = hs.concat(ltuList);
        tmpList = ltuTypeList;
      } else if (s.length === 2 && s[0] !== 'ん') {
        tmpList = tmpList.concat(mp[s])
        var fstList = mp[s[0]];
        var sndList = mp[s[1]];
        var resList: string[] = [];
        for (var fstStr of fstList) {
          for (var sndStr of sndList) {
            var u = fstStr + sndStr;
            resList.push(u);
          }
        }
        tmpList = tmpList.concat(resList);
      } else {
        tmpList = mp[s];
      }
      res.push(tmpList);
    }
    return res;
  }
}
</script>

<style scoped lang="scss">
@import '@/style.scss';
</style>
