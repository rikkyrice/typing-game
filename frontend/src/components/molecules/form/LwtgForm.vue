<template>
  <v-card
    class="rounded-md mx-auto"
    :class="MdSmXsUtil('px-8 py-4', 'pa-4', 'pa-3')"
    :width="MdSmXsUtil('400px', '400px', '')"
  >
    <v-container>
      <div>
        <slot name="form-title" />
      </div>
      <div class="lwtg-primary-bg" style="height: 2px;" />
      <div class="d-flex align-center">
        <slot name="form-comment" />
      </div>
      <v-form v-model="valid">
        <slot name="form-body" />
      </v-form>
    </v-container>
    <lwtg-loader :loading="loading" />
  </v-card>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';
import LwtgLoader from '@/components/atoms/LwtgLoader.vue';
import { FormItem } from '@/models/types/formItem';

@Component({
  components: {
    LwtgLoader,
  },
})
export default class LwtgForm extends mixins(UtilMixin) {
  @Prop() formItem!: FormItem;
  @Prop() loading!: boolean;

  pageTransition() {
    this.$router.push(this.formItem.path);
  }
}
</script>

<style scoped lang="scss">
@import '@/style.scss';
.page-link {
  cursor: pointer;
  &:hover {
    color: #A0D0A0;
    text-decoration: underline;
  }
}
</style>
