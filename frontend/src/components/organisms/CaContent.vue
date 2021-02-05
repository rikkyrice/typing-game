<template>
  <v-hover v-model="ishover">
    <div
      id="caContent"
      style="width: 100%"
    >
      <div>
        <ca-card
          :caItem="caItems[0]"
          :index="true"
          :isHover="ishover ? true : false"
        />
      </div>
      <v-row no-gutters>
        <v-col
          v-for="(caItem, index) in caItems.slice(1,caItems.length)"
          :key="index"
          cols="12"
          sm="4"
        >
          <ca-card
            :caItem="caItem"
            :height="`218px`"
            :index="false"
            :isHover="ishover ? true : false"
          />
        </v-col>
      </v-row>
    </div>
  </v-hover>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
import { mixins } from 'vue-class-component';
import CaCard from '@/components/atoms/CaCard.vue';
import UtilMixin from '@/mixins/utilMixin';
import { CaItem } from '@/models/types/caItem';

@Component({
  components: {
    CaCard,
  },
})
export default class CaContent extends mixins(UtilMixin) {
  @Prop() caItems!: CaItem[];
  ishover = false;
}
</script>

<style scoped lang="scss">
@import '@/style.scss';
</style>
