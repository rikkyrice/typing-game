<template>
  <div
    id="typing-word-card"
    class="d-flex justify-center align-center flip-card"
    style="width: 100%; height: 100%;"
  >
    <lwtg-word-card
      :width="width"
      :border="border"
      :primary="primary"
      class="lwtg-word-card flip-card-inner"
      :class="{
        'flipped-card': !isActivated,
      }"
    >
      <template #word>
        <div
          style="width: 100%; height: 100%;"
          class="d-flex justify-center align-center"
          :class="{
            'flipped-card': !isActivated,
          }"
        >
          <div style="width: 100%;" class="text-center">
            <lwtg-typing-game
              v-if="!clear && isActivated"
              :key="key"
              :typeWord="typeWord"
              :isActivated="getIsActivated"
              @shift="shift"
            />
            <lwtg-typing-game
              v-else-if="!clear && !isActivated"
              :key="key"
              :typeWord="typeMeaning"
              :isActivated="getIsActivated"
              :wordOnly="wordOnly"
              @shift="shift"
            />
            <div
              v-else
              style="width: 100%"
              class="text-center"
            >
              <span
                class="bold main-mono-color"
                :style="fontSizeUtil(48, 48, 32)"
              >Clear!</span><br>
              <span
                class="bold main-mono-color"
                :style="fontSizeUtil(24, 24, 18)"
              >- Press Space to Restart -</span>
            </div>
          </div>
        </div>
      </template>
    </lwtg-word-card>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';
import LwtgWordCard from '@/components/atoms/LwtgWordCard.vue';
import LwtgTypingGame from '@/components/atoms/LwtgTypingGame.vue';
import { TypingWord, Word } from '@/models/word';
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
  @Prop() words!: TypingWord[];
  @Prop() isActivated!: boolean;
  @Prop({ default: true }) isWords!: boolean;
  @Prop() wordOnly!: boolean;
  clear: boolean = false;
  index = 0;
  key = 0;
  get typeWord() {
    return store.state.typeWord.typeWord;
  }
  get typeMeaning() {
    return store.state.typeWord.typeMeaning;
  }
  get getIsActivated() {
    return this.isActivated;
  }
  dispatchTypeWord() {
    store.dispatch(TYPES.SHIFT_TYPEWORD, this.words[this.index].word);
    store.dispatch(TYPES.SHIFT_TYPEMEANING, this.words[this.index].meaning);
  }
  mounted() {
    console.log(this.width);
    window.addEventListener('resize', this.handleResize);
    if (this.width === 0) {
      this.handleResize();
    }
    window.addEventListener('keydown', e => {
      if (store.state.gameCleared) {
        this.keyDown(e.key);
      }
    });
  }
  @Watch('words')
  updateWords(newWords: Word[]) {
    Promise.resolve()
      .then(() => {
        this.dispatchTypeWord();
      });
  }
  handleResize() {
    const typingWordCard = document.getElementById(
      'typing-word-card'
    );
    if (typingWordCard) {
      const typingWordCardWidth = typingWordCard.getBoundingClientRect().width;
      if (typingWordCardWidth > 1000) {
        this.$emit('updateWidth', 1000);
      } else {
        this.$emit('updateWidth', typingWordCardWidth - (typingWordCardWidth / 9));
      }
    }
  }
  keyDown(code: string) {
    if (code !== ' ') {
      return false;
    }
    this.$emit('reset');
  }
  shift() {
    this.$emit('shift');
  }
  shiftIndex() {
    this.index += 1;
    if (this.index === this.words.length) {
      this.clear = true;
      store.dispatch(TYPES.SWITCH_CLEARED, true);
    } else {
      this.dispatchTypeWord();
    }
  }
  shuffle(randomNum: number) {
    for (var i = this.words.length; i > 1; i--) {
      var k = Math.floor(randomNum * i);
      [this.words[k], this.words[i - 1]] = [this.words[i - 1], this.words[k]];
    }
    this.reset();
  }
  reset() {
    store.dispatch(TYPES.SWITCH_CLEARED, false);
    this.clear = false;
    this.index = 0;
    this.dispatchTypeWord();
    this.key = this.key ? 0 : 1;
  }
}
</script>

<style scoped lang="scss">
@import '@/style.scss';
.meaning-card {
  transform: rotateY(180deg);
}
.flip-card {
  background-color: transparent;
  width: 100%;
  height: 100%;
  perspective: 2000px;
}
.flipped-card {
  transform: rotateY(-180deg);
}
.flip-card-inner {
  cursor: pointer;
  position: relative;
  width: 100%;
  height: 100%;
  transition: transform 0.3s;
  transform-style: preserve-3d;
}
.flip-card-front {
  position: absolute;
  width: 100%;
  height: 100%;
  -webkit-backface-visibility: hidden; /* Safari */
  backface-visibility: hidden !important;
}
</style>
