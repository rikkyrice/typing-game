<template>
  <div id="formTemplate" style="height: 100%;">
    <lwtg-content :bg="`secondary`" style="height: 100%;" class="pa-16">
      <slot name="form-card" />
    </lwtg-content>
    <lwtg-loader :page-loading="true" :loading="false" />
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';
import LwtgContent from '@/components/atoms/LwtgContent.vue';

@Component({
  components: {
    LwtgContent,
  }
})
export default class formTemplate extends mixins(UtilMixin) {

}
</script>

<style scoped lang="scss">
@import '@/style.scss';
</style>
