import { Route } from 'vue-router';
import {
  State,
  UserState,
  AuthState,
  RouteInfo,
  SnackbarInfo,
} from '@/store/types';

const userInfo: UserState = {
  userId: '',
  mail: '',
  createdAt: '',
}

const auth: AuthState = {
  token: '',
  lastPageRoute: {} as RouteInfo,
}

const scrollTarget: string = '';

const snackbar: SnackbarInfo = {
  message: '',
  visibility: false,
};

export const state: State = {
  userInfo,
  auth,
  scrollTarget,
  snackbar,
};
