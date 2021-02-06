<template>
  <form-template>
    <template #form-card>
      <lwtg-signup-form />
    </template>
  </form-template>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';
import LwtgSignupForm from '@/components/organisms/form/LwtgSignupForm.vue';
import FormTemplate from '@/components/templates/FormTemplate.vue';

@Component({
  components: {
    LwtgSignupForm,
    FormTemplate,
  }
})
export default class SignupView extends mixins(UtilMixin) {

}
</script>

<style scoped lang="scss">
@import '@/style.scss';
</style>
