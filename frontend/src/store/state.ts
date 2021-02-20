import { Route } from 'vue-router';
import {
  State,
  UserState,
  AuthState,
  RouteInfo,
  SnackbarInfo,
  TypeWordState,
} from '@/store/types';
import { TypeWord } from '@/models/types/typeWord';

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

const typeWord: TypeWordState = {
  typeWord: {} as TypeWord,
  typeMeaning: {} as TypeWord,
};

export const state: State = {
  userInfo,
  auth,
  scrollTarget,
  snackbar,
  typeWord,
};
