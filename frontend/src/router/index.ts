import Vue from 'vue';
import Router from 'vue-router';
import routes from '@/router/router';
import { PAGES } from '@/router/pages';
import store from '@/store';

Vue.use(Router);

const router = new Router({
  mode: 'history',
  base: process.env.VUE_APP_PUBLIC_PATH,
  routes,
  scrollBehavior() {
    return { x: 0, y: 0 };
  },
});

router.beforeEach((to, from, next) => {
  if (to.fullPath === from.fullPath && to.meta.title === from.meta.title) {
    return;
  }
  const isAfterLogin = !from.name && to.name === PAGES.TOP && !!to.hash;
  next();
});

export default router;
