<template>
  <lwtg-form>
    <template #form-title>
      <span
        class="bold"
        style="color: #A0D0A0;"
        :style="fontSizeUtil(24, 24, 20)"
      >{{ formItem.title }}</span>
    </template>
    <template #form-comment>
      <div>
        <span :style="fontSizeUtil(12, 12, 10)"
        >{{ formItem.comment }}</span>
      </div>
      <div
        class="page-link pl-1 bold"
        @click="pageTransition()"
      >
        <span
          :style="fontSizeUtil(12, 12, 10)"
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
          v-model="password"
          :rules="passwordRules"
          :type="'password'"
          label="パスワード"
          required
        ></v-text-field>
      </div>
      <v-row justify="end" no-gutters>
        <lwtg-button
          :label="formItem.title"
          :disabled="disabled"
          style="width: 150px;"
          @click="login"
        />
      </v-row>
    </template>
  </lwtg-form>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';
import UserApi from '@/api/user';
import LwtgForm from '@/components/molecules/form/LwtgForm.vue'
import LwtgButton from '@/components/atoms/LwtgButton.vue';
import { FormItem } from '@/models/types/formItem';
import { TokenInfo } from '@/models/user';
import store from '@/store';
import { TYPES } from '@/store/mutation-types';

@Component({
  components: {
    LwtgForm,
    LwtgButton,
  },
})
export default class LwtgLoginForm extends mixins(UtilMixin) {
  userId = '';
  password = '';
  tokenInfo: TokenInfo = {} as TokenInfo;
  formItem: FormItem = { title: 'ログイン', comment: 'まだ登録していない方は', link: '登録', path: '/signup' };
  nameRules: any = [
    (v: string) => !!v || 'ユーザーIDは必須です',
    (v: string) => this.alphanumericRule(v),
    (v: string) => this.lengthRule(v, 20),
  ];
  passwordRules: any = [
    (v: string) => !!v || 'パスワードは必須です',
    (v: string) => this.alphanumericRule(v),
  ];
  alphanumericRule(value: string) {
    return value.match(/^[0-9a-z]+$/) || '英数字で入力してください';
  }
  lengthRule(value: string, length: number) {
    return value.length <= length || length + '字以内で入力してください';
  }
  login() {
    UserApi.login(this.userId, this.password)
      .then((data) => {
        this.tokenInfo = data;
        store.dispatch(TYPES.LOGIN, this.tokenInfo);
    })
    .finally(() => {
      this.$router.push('/mypage');
    });
  }
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
