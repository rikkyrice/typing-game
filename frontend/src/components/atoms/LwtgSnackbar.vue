<template>
  <v-snackbar
    v-model="visibility"
    :timeout="3000"
    :top="yPos === 'top'"
    :bottom="yPos === 'bottom'"
    color="#A0D0A0"
    class="white--text"
    style="opacity: 0.9;"
    >{{ text
    }}<template #action="{ attrs }">
      <v-icon
        class="mx-1"
        color="white"
        v-bind="attrs"
        @click="visibility = false"
        >mdi-close</v-icon
      >
    </template>
  </v-snackbar>
</template>

<script lang="ts">
import { Component, Vue, Prop, Model, Emit } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import UtilMixin from '@/mixins/utilMixin';

@Component
export default class LwtgSnackbar extends mixins(UtilMixin) {
  @Prop() value!: boolean;
  @Prop() text!: string;
  @Prop({ default: 'top' }) yPos!: string;

  @Emit()
  input(value: boolean) {
    // emit event
  }

  get visibility() {
    return this.value;
  }

  set visibility(value: boolean) {
    this.input(value);
  }
}
</script>

<style scoped lang="scss">
@import '@/style.scss';
</style>
