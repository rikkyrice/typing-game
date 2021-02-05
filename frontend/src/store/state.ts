import { Route } from 'vue-router';
import {
  State,
  SnackbarInfo,
} from '@/store/types';

const scrollTarget: string = '';

const snackbar: SnackbarInfo = {
  message: '',
  visibility: false,
};

export const state: State = {
  scrollTarget,
  snackbar,
};
