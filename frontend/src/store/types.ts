export interface State {
  userInfo: UserState;
  auth: AuthState;
  scrollTarget: string;
  snackbar: SnackbarInfo;
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