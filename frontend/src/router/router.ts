import { PAGES } from '@/router/pages';

export default [
  {
    path: '/',
    name: PAGES.TOP,
    component: () => import('@/components/pages/TopView.vue'),
    meta: { title: 'TOP', saveScroll: false },
    props: true,
  },
  {
    path: '/signup',
    name: PAGES.SIGNUP,
    component: () => import('@/components/pages/SignupView.vue'),
    meta: { title: 'SIGNUP', saveScroll: false },
    props: true,
  },
  {
    path: '/login',
    name: PAGES.LOGIN,
    component: () => import('@/components/pages/LoginView.vue'),
    meta: { title: 'LOGIN', saveScroll: false },
    props: true,
  },
  {
    path: '/pg',
    name: PAGES.PLAYGROUND,
    component: () => import('@/components/pages/PlaygroundView.vue'),
    meta: { title: 'PLAYGROUND', saveScroll: false },
    props: true,
  }
];
