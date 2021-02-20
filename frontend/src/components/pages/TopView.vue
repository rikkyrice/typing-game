<template>
  <div
    id="top"
    class="d-flex align-center justify-center"
    :class="$vuetify.breakpoint.xsOnly ? 'lwtg-top-mobile-bg' : 'lwtg-top-bg'"
    style="height: 100%;"
  >
    <div class="lwtg-overlay"></div>
    <div
      class="text-center d-flex flex-column"
      :style="{
        width: MdSmXsUtil('760px', '650px', '100%'),
      }"
    >
      <span
        id="type-writer"
        class="bold mb-10"
        style="color: #666666; z-index: 5;"
        :style="fontSizeUtil(56, 48, 40)"
      >
        Learning
        <br v-if="$vuetify.breakpoint.xsOnly" />w/ Typing Game.
      </span>
      <v-row class="mb-3" justify="center" no-gutters>
        <v-col cols="6" class="d-flex justify-center">
          <lwtg-button
            :label="'ユーザー登録'"
            class="lwtg-top-button"
            :style="{
              fontSize: MdSmXsUtil('18px', '18px', '12px'),
              width: MdSmXsUtil('200px', '200px', '100px'),
            }"
            :height="'50px'"
            @click="signupTransition"
          />
        </v-col>
        <v-col cols="6" class="d-flex justify-center">
          <lwtg-button
            :label="'ログイン'"
            class="lwtg-top-button"
            :style="{
              fontSize: MdSmXsUtil('18px', '18px', '12px'),
              width: MdSmXsUtil('200px', '200px', '100px'),
            }"
            :height="'50px'"
            @click="loginTransition"
          />
        </v-col>
      </v-row>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';
import LwtgButton from '@/components/atoms/LwtgButton.vue'

@Component({
  components: {
    LwtgButton,
  },
})
export default class TopView extends mixins(UtilMixin) {
  mounted() {
    const typewriter = document.getElementById('type-writer');
    if (typewriter) {
      const string = typewriter.innerText.split('');
      typewriter.textContent = '';
      string.forEach((char, index) => {
        setTimeout(() => {
          typewriter.textContent += char;
        }, 110 * index);
      });
    }
  }

  signupTransition() {
    this.$router.push('/signup');
  }
  loginTransition() {
    this.$router.push('/login');
  }
}
</script>

<style scoped lang="scss">
@import '@/style.scss';
.lwtg-top-bg {
  background-image: url('~@/assets/lwtg-home.png');
  background-size: cover;
}
.lwtg-top-mobile-bg {
  background-image: url('~@/assets/lwtg-home.png');
  background-size: cover;
}
.lwtg-overlay {
  height: 100%;
  width: 100%;
  background: linear-gradient(rgba(255,255,255,.3) 70%, 80%, #A0D0A0 93%);
  position: fixed;
}
.lwtg-top-button {
  background-color: rgba(255,255,255,.3) !important;
}
#type-writer::after {
  content: "|";
  animation-name: blink;
  animation-duration: 1.5s;
  animation-iteration-count: infinite;
}
@keyframes blink {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}
</style>
