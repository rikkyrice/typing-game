<template>
  <div id="app">
    <v-app>
      <v-main id="main">
        <router-view />
      </v-main>
    </v-app>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator';
import Toasted from 'vue-toasted';

Vue.use(Toasted, {
  position: 'top-center',
  duration: 3000,
});

@Component
export default class App extends Vue {
}
</script>

<style lang="scss">
@import '@/style.scss';
// global styles
#app {
  font-family: 'HiraginoSans-W3', 'Meiryo UI', sans-serif !important;
  min-height: 100vh;
}
#main {
  max-height: 100vh;
}
.toasted-container .toasted {
  background-color: #e32d2d !important;
  box-shadow: 0 2px 4px 0 rgba(0, 0, 0, 0.5);
  opacity: 0.75 !important;
}
</style>
