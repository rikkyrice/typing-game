<template>
  <v-btn
    :icon="!fab"
    :fab="fab"
    :dark="dark"
    :elevation="fab ? elevation : 0"
    :loading="loading"
    color="white"
    :x-small="size === 'x-small'"
    :small="size === 'small'"
    :large="size === 'large'"
    :x-large="size === 'x-large'"
    :style="style"
    @click="handleClick()"
    @click.stop="handleClickStopOption()"
  >
    <img
      v-if="src"
      :src="src"
      :width="svgSize"
      :style="{ fill: color || '#666666' }"
    />
    <v-icon v-else :size="iconSize" :color="color || '#666666'"
      >mdi-{{ icon }}</v-icon
    >
  </v-btn>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';

@Component
export default class LwtgIconButton extends mixins(UtilMixin) {
  @Prop() icon!: string;
  @Prop() src!: string;
  @Prop() fab!: boolean;
  @Prop() outlined!: boolean;
  @Prop() dark!: boolean;
  @Prop() color!: string;
  @Prop({ default: 2 }) elevation!: number | string;
  @Prop() loading!: boolean;
  @Prop() size!: string;
  @Prop() disabled!: boolean;

  get style() {
    return this.outlined
      ? `color: ${this.color} !important; border: solid 2px ${this.color} !important;`
      : '';
  }

  get iconSize() {
    return this.size === 'x-small'
      ? '20px'
      : this.size === 'small'
      ? '24px'
      : this.size === 'large'
      ? '32px'
      : this.size === 'x-large'
      ? '36px'
      : '28px';
  }

  get svgSize() {
    return this.size === 'x-small'
      ? '20px'
      : this.size === 'small'
      ? '24px'
      : this.size === 'large'
      ? '32px'
      : this.size === 'x-large'
      ? '36px'
      : '28px';
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
</style>
