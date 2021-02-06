<template>
  <form-template>
    <template #form-card>
      <lwtg-login-form />
    </template>
  </form-template>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';
import LwtgLoginForm from '@/components/organisms/form/LwtgLoginForm.vue';
import FormTemplate from '@/components/templates/FormTemplate.vue';

@Component({
  components: {
    LwtgLoginForm,
    FormTemplate,
  }
})
export default class LoginView extends mixins(UtilMixin) {

}
</script>

<style scoped lang="scss">
@import '@/style.scss';
</style>
