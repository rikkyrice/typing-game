<template>
  <v-card
    outlined
    class="rounded-md d-flex flex-column my-2"
    :class="MdSmXsUtil('pl-4', 'pa-4', 'pa-3')"
    :width="MdSmXsUtil('100%', '1000px', '600px')"
    height="150px"
    @click="pageTransition"
  >
    <v-container
      class="py-2"
      style="height: 100%;"
    >
      <v-row style="height: 100%;">
        <v-col cols="11" class="py-0">
          <v-row style="height: 100%">
            <v-col cols="12" sm="6" md="9">
              <div style="width: 100%;">
                <span
                  class="bold"
                  :style="fontSizeUtil(24, 24, 20)"
                >{{ wordlist.wordListTitle }}</span>
              </div>
              <div style="width: 100%;">
                <span
                  class="mono-60-color"
                  :style="fontSizeUtil(16, 16, 12)"
                >{{ wordlist.explanation }}</span>
              </div>
            </v-col>
            <v-col cols="6" sm="3" md="3" class="px-0" style="height: 100%;">
              <v-row
                style="height: 100%;"
                class="flex-column"
                no-gutters
              >
                <v-col style="text-align: right" :style="fontSizeUtil(16, 16, 12)">
                  <span class="bold">単語数: {{ wordArray.matched }}</span>
                </v-col>
                <v-spacer />
                <v-col class="d-flex" :style="fontSizeUtil(12, 12, 12)">
                  <div class="ml-auto">
                    <div>
                      <span>プレイ回数: {{ score.playCount }}</span>
                    </div>
                    <div>
                      <span>プレイ日時: {{ score.playedAt }}</span>
                    </div>
                  </div>
                </v-col>
              </v-row>
            </v-col>
          </v-row>
        </v-col>
        <v-col cols="1" class="py-0" style="text-align: right;">
          <v-menu
            bottom
            left
          >
            <template v-slot:activator="{ on, attrs }">
              <v-btn
                icon
                v-bind="attrs"
                v-on="on"
              >
                <v-icon>mdi-dots-vertical</v-icon>
              </v-btn>
            </template>

            <v-list>
              <v-list-item
                v-for="(menuItem, i) in menuItems"
                :key="i"
                :class="{
                  'lwtg-danger-bg': menuItem.danger,
                }"
                link
              >
                <v-list-item-title
                  :class="{
                    'white-100-color': menuItem.danger,
                  }"
                >{{ menuItem.title }}</v-list-item-title>
              </v-list-item>
            </v-list>
          </v-menu>
        </v-col>
      </v-row>
    </v-container>
  </v-card>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';
import LwtgIconButton from '@/components/atoms/LwtgIconButton.vue';
import { WordList } from '@/models/wordlist';
import { WordArray } from '@/models/word';
import { Score } from '@/models/score';
import store from '@/store';
import { DeviceType, deviceType } from '@/models/types/deviceType';
import { MenuItem } from '@/models/types/menuItem';

@Component({
  components: {
    LwtgIconButton,
  }
})
export default class WordListCard extends mixins(UtilMixin) {
  @Prop() wordlist!: WordList;
  // @Prop() wordArray!: WordArray;
  wordArray: WordArray = {
    matched: 0,
    words: [],
  }
  // @Prop() score!: Score;
  score: Score = {
    scoreId: 0,
    playCount: 1,
    clearTypeCount: 100,
    missTypeCount: 1,
    playedAt: '2021-02-06-00:00:00',
  }
  menuItems: MenuItem[] = [
    { title: '編集', danger: false, action: 'edit' },
    { title: '削除', danger: true, action: 'delete' },
  ]
  isTrashcanButtonHover = false;
  get isDesktop() {
    return deviceType === DeviceType.DESKTOP;
  }
  get trashcanButtonIcon() {
    return this.isDesktop && this.isTrashcanButtonHover
      ? require('@/assets/common/trash-can.svg')
      : require('@/assets/common/trash-can-outline.svg');
  }
  pageTransition() {
    this.$router.push(`/mypage/${this.wordlist.id}`, () => {});
  }
}
</script>

<style scoped lang="scss">
@import '@/style.scss';
</style>
