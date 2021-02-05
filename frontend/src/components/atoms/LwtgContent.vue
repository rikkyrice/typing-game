<template>
  <div :class="`lwtg-${bg}-bg`" :style="{ minWidth }">
    <div
      class="mx-auto"
      :style="maxWidthUtil('960px', '720px', 'calc(100vw - 32px)')"
    >
      <slot />
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Emit, Prop, Vue } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';

@Component
export default class LwtgContent extends mixins(UtilMixin) {
  @Prop({ default: 'white' }) bg!: string;
  @Prop({ default: '100vw' }) minWidth!: string;
}
</script>

<style scoped lang="scss">
@import '@/style.scss';
</style>
