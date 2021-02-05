<template>
  <v-overlay
    :value="loading"
    :color="pageLoading ? 'white' : '#212121'"
    :opacity="pageLoading ? 1 : 0.46"
  >
    <v-img
      v-if="pageLoading"
      :src="require('@/assets/common/loading.gif')"
      max-width="360px"
    />
    <v-progress-circular v-else indeterminate color="#A0D0A0" />
  </v-overlay>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';

@Component
export default class LwtgLoader extends mixins(UtilMixin) {
  @Prop({ default: false }) loading!: boolean;
  @Prop() pageLoading!: number;
}
</script>

<style scoped lang="scss">
@import '@/style.scss';
</style>
