import Vue from 'vue';
import Vuetify from 'vuetify';
import 'vuetify/dist/vuetify.min.css';
import App from '@/App.vue';
import router from '@/router';
import store from '@/store';
import 'animate.css';

Vue.config.productionTip = false;
Vue.use(Vuetify);

new Vue({
  vuetify: new Vuetify({
    icons: {
      iconfont: 'mdi',
    },
  }),
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
