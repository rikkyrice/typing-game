<template>
  <v-chip
    :close="clearable"
    :color="outlined ? 'white' : color"
    :dark="!outlined"
    :label="tile"
    :ripple="false"
    :draggable="false"
    style="height: 24px;"
    :style="style"
    @click="handleClick()"
    @click.stop="handleClickStopOption()"
    >{{ text }}</v-chip
  >
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';

@Component
export default class LwtgChip extends mixins(UtilMixin) {
  @Prop() text!: string;
  @Prop() clearable!: boolean;
  @Prop() clickable!: boolean;
  @Prop() tile!: boolean;
  @Prop() outlined!: boolean;
  @Prop() color!: string;

  get style() {
    const cursor = `cursor: ${this.clickable ? 'pointer;' : 'default;'}`;
    const border = `border: solid 1px ${this.color} !important;`;
    const textColor = this.outlined ? `color: ${this.color} !important;` : '';
    const fontSize = `font-size: ${this.smallChip ? 10 : 12}px`;
    return cursor + border + textColor + fontSize;
  }

  get smallChip() {
    return (
      this.textWidth > 32 && (window.innerWidth < 375 || this.textWidth > 40)
    );
  }

  get textWidth() {
    const halfChars = `${this.text}`.match(/[ -~]/g) || [];
    return this.text.length * 2 - halfChars.length;
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
.dark-outline {
  border: solid 1px;
}
</style>
