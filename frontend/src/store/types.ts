import { TypeWord } from "@/models/types/typeWord";

export interface State {
  userInfo: UserState;
  auth: AuthState;
  scrollTarget: string;
  snackbar: SnackbarInfo;
  typeWord: TypeWordState;
}

export interface UserState {
  userId: string;
  mail: string;
  createdAt: string;
}

export interface AuthState {
  token: string;
  lastPageRoute: RouteInfo;
}

export interface RouteInfo {
  path: string;
  query: any;
}

export interface SnackbarInfo {
  message: string;
  visibility: boolean;
}

export interface TypeWordState {
  typeWord: TypeWord;
  typeMeaning: TypeWord;
}
