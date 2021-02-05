import { PAGES } from '@/router/pages';

export default [
  {
    path: '/',
    name: PAGES.TOP,
    component: () => import('@/components/pages/HelloWorld.vue'),
    meta: { title: 'TOP', saveScroll: false },
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
