<template>
  <div
    class="lwtg-word-card"
    :style="{
      width: width + 'px',
      height: getHeight,
    }"
    @click="handleClick()"
    @click.stop="handleClickStopOption()"
  >
    <div
      class="word-card lwtg-white-bg"
      :style="wordCardStyle"
    >
      <slot name="word" />
    </div>
    <div
      class="word-card-ring"
      :style="wordCardRingStyle"
    />
    <div
      class="word-card-ring-hole lwtg-primary-bg"
      :style="wordCardRingHoleStyle"
    />
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';

@Component
export default class LwtgWordCard extends mixins(UtilMixin) {
  @Prop() width!: number;
  @Prop({ default: '10px' }) border!: number;
  @Prop() primary!: boolean;
  baseColor = '#666666';
  primaryColor = '#A0D0A0';

  get getHeight() {
    const height = this.width / 3;
    return height + 'px';
  }
  get wordCardStyle() {
    const color = this.primary
      ? this.primaryColor
      : this.baseColor;
    const width = this.width - (this.width / 9);
    return `
      width: ${width}px;
      border: solid ${this.border}px ${color} !important;
      top: 0px;
    `;
  }
  get wordCardRingStyle() {
    const color = this.primary
      ? this.primaryColor
      : this.baseColor
    const height = this.width / 6;
    const top = this.width / 12;
    const border = this.border - 2;
    return `
      width: ${height}px;
      height: ${height}px;
      border: solid ${border}px ${color} !important;
      top: ${top}px;
    `;
  }
  get wordCardRingHoleStyle() {
    const height = this.border * 2;
    const top = (this.width / 6) - (this.border * 1.5);
    return `
      width: ${height}px;
      height: ${height}px;
      top: ${top}px;
      left: ${top}px;
    `;
  }
  handleClick(ev: any) {
    this.$emit('click', ev);
  }

  handleClickStopOption(ev: any) {
    this.$emit('clickStop', ev);
  }
}
</script>

<style scoped lang="scss">
@import '@/style.scss';
.lwtg-word-card {
  position: relative;
  transform-style: preserve-3d;
}
.word-card {
  position: absolute;
  right: 0px;
  height: 100%;
  transform-style: preserve-3d;
}
.word-card-ring {
  position: absolute;
  left: 0px;
  background-color: transparent;
  border-radius: 50% !important;
  transform: rotate3d(1,0,0,-1deg);
}
.word-card-ring-hole {
  position: absolute;
  border-radius: 50% !important;
}
</style>
