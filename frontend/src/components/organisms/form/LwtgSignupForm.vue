<template>
  <lwtg-form>
    <template #form-title>
      <span
        class="bold"
        style="color: #A0D0A0;"
        :style="fontSizeUtil(24, 24, 18)"
      >{{ formItem.title }}</span>
    </template>
    <template #form-comment>
      <div>
        <span :style="fontSizeUtil(10, 10, 8)"
        >{{ formItem.comment }}</span>
      </div>
      <div
        class="page-link pl-1"
        @click="pageTransition()"
      >
        <span
          :style="fontSizeUtil(10, 10, 8)"
          style="color: #A0D0A0;"
        >{{ formItem.link }}</span>
      </div>
    </template>
    <template #form-body>
      <div class="pb-4">
        <v-text-field
          v-model="userId"
          :rules="nameRules"
          label="ユーザーID"
          required
        ></v-text-field>
        <v-text-field
          v-model="email"
          :rules="emailRules"
          label="メールアドレス"
          required
        ></v-text-field>
        <v-text-field
          v-model="password"
          :rules="passwordRules"
          label="パスワード"
          required
        ></v-text-field>
        <v-text-field
          v-model="confirmedPassword"
          :rules="passwordRules"
          label="パスワード確認"
          required
        ></v-text-field>
      </div>
      <v-row justify="end" no-gutters>
        <lwtg-button
          :label="formItem.title"
          :disabled="disabled"
          style="width: 150px;"
          @click="postAnswer"
        />
      </v-row>
    </template>
  </lwtg-form>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';
import LwtgForm from '@/components/molecules/form/LwtgForm.vue'
import LwtgButton from '@/components/atoms/LwtgButton.vue';
import { FormItem } from '@/models/types/formItem';

@Component({
  components: {
    LwtgForm,
    LwtgButton,
  },
})
export default class LwtgSignupForm extends mixins(UtilMixin) {
  formItem: FormItem = { title: '登録', comment: '既に登録済みの方は', link: 'ログイン', path: '/' };

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
