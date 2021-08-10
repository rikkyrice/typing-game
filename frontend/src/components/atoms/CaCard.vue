<template>
  <v-card
    :height="height"
    width="100%"
    class="hover-animation-card mx-auto"
    :class="{
      'rounded-xl': isHover,
    }"
    @click="handleClick()"
    @click.stop="handleClickStopOption()"
  >
    <v-container fill-height align-content="center">
      <v-row justify="center" no-gutters>
        <div class="d-flex flex-column align-center">
          <span
            class="word-break bold"
            :style="index ? fontSizeUtil(52, 52, 36) : fontSizeUtil(36, 36, 24)"
            >{{ caItem.title }}</span
          >
          <span
            class="word-break bold"
            :style="fontSizeUtil(16, 16, 12)"
            >{{ caItem.subtitle }}</span
          >
        </div>
      </v-row>
    </v-container>
  </v-card>
</template>

<script lang="ts">
import { Component, Emit, Prop, Vue } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';
import { CaItem } from '@/models/types/caItem';

@Component
export default class CaCard extends mixins(UtilMixin) {
  @Prop() caItem!: CaItem;
  @Prop() isHover!: boolean;
  @Prop({ default: '378px' }) height!: string;
  @Prop() index!: boolean;
}
</script>

<style scoped lang="scss">
@import '@/style.scss';
.hover-animation-card {
  transition: all .2s ease-out;
}
.hover-animation-card:hover {
  transform: scale(0.95) !important;
}
</style>
