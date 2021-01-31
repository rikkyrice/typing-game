import { PAGES } from '@/router/pages';

export default [
  {
    path: '/',
    name: PAGES.TOP,
    component: () => import('@/components/HelloWorld.vue'),
    meta: { title: 'TOP', saveScroll: false },
    props: true,
  },
];
