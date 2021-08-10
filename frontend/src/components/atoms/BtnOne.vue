<template>
  <div
    class="btn btn-one d-flex align-center justify-center"
    :style="style"
    @click="handleClick()"
    @click.stop="handleClickStopOption()"
  >
    <div class="d-flex align-center">
      <v-icon v-if="prependIcon" left :style="{ fontSize: iconSize }">{{
        `mdi-${prependIcon}`
      }}</v-icon>
      <img
        v-else-if="prependSrc"
        class="mr-2"
        :src="prependSrc"
        :height="iconSize"
      />
      <span>{{ label }}</span>
      <v-icon v-if="appendIcon" right :style="{ fontSize: iconSize }">{{
        `mdi-${appendIcon}`
      }}</v-icon>
      <img
        v-else-if="appendSrc"
        class="ml-2"
        :src="appendSrc"
        :height="iconSize"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';
import { DeviceType, deviceType } from '@/models/types/deviceType';

@Component
export default class BtnOne extends mixins(UtilMixin) {
  @Prop() label!: string;
  @Prop() prependIcon!: string;
  @Prop() appendIcon!: string;
  @Prop() prependSrc!: string;
  @Prop() appendSrc!: string;
  @Prop() size!: string;
  @Prop({ default: '36px' }) height!: string;
  get iconSize() {
    return this.size === 'x-small'
      ? '14px'
      : this.size === 'small'
      ? '16px'
      : this.size === 'large'
      ? '20px'
      : this.size === 'x-large'
      ? '22px'
      : '20px';
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
.btn {
  cursor: pointer;
  position: relative;
  color: #666666;
  width: 256px;
  height: 64px;
  line-height: 64px;
  background: #ffffff;
  transition: all 0.3s;
  span {
    transition: all 0.3s;
    transform: scale(1, 1);
  }
}
.btn::before, .btn::after {
  content: '';
  position: absolute;
  transition: all 0.3s;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 1;
}
.btn-one::before {
  left: 4px;
  z-index: 1;
  opacity: 0;
  background: rgba(102, 102, 102, 0.3);
  transform: scale(0.1, 1);
}
.btn-one:hover::before {
  opacity: 1;
  transform: scale(1, 1);
}
.btn-one::after {
  transition: all 0.3s;
  border: 1px solid #666666;
}
.btn-one:hover::after {
  transform: scale(1, .1);
  opacity: 0;
}
</style>
