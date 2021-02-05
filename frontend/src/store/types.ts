export interface State {
  scrollTarget: string;
  snackbar: SnackbarInfo;
}

export interface SnackbarInfo {
  message: string;
  visibility: boolean;
}