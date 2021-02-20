<template>
  <lwtg-icon-button
    icon="plus"
    :fab="true"
    :outlined="true"
    color="#A0D0A0"
    :size="MdSmXsUtil('large', '', 'small')"
    elevation="1"
    class="create-word"
    @click="handleClick()"
    @click.stop="handleClickStopOption()"
  />
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import LwtgIconButton from '@/components/atoms/LwtgIconButton.vue';
import UtilMixin from '@/mixins/utilMixin';

@Component({
  components: {
    LwtgIconButton
  }
})
export default class LwtgWordCreateButton extends mixins(UtilMixin) {
  handleClick(ev: any) {
    this.$emit('click', ev);
  }

  handleClickStopOption(ev: any) {
    this.$emit('clickStop', ev);
  }
}
</script>

<style scoped lang="scss">
.create-word {
  position: fixed;
  right: 24px;
  bottom: 24px;
  z-index: 100;
  animation-duration: 0.4s;
}
</style>
