import { Vue } from 'vue-property-decorator';
import { ToastOptions, ToastObject } from 'vue-toasted';
import { AxiosRequestConfig, AxiosError } from 'axios';
import router from '@/router';
import store from '@/store';
import { TYPES } from '@/store/mutation-types';
import { PAGES } from '@/router/pages';

// set Token to request.headers.Authorization
const reqHandler = (request: AxiosRequestConfig) => {
  request.url = `${process.env.VUE_APP_API_ENTRY_POINT}${request.url}`;
  if (
    process.env.VUE_APP_MODE === 'local' &&
    process.env.VUE_APP_NOAUTH === 'true'
  ) {
    request.headers.noauth = true;
  }
  if (request.url.includes(PAGES.LOGIN) || request.url.includes(PAGES.SIGNUP)) {
    return Promise.resolve(request);
  }
  const token = store.state.auth.token || '';
  if (!token && !request.headers.noauth) {
    return Promise.reject(request);
  }
  request.headers.Authorization = `Bearer ${token}`;
  const userId = store.state.auth.userId || '';
  if (!token && !request.headers.noauth) {
    return Promise.reject(request);
  }
  request.headers['X-User-ID'] = userId;
  console.log(request);
  return Promise.resolve(request);
};

/**
 * APIエラー一覧
 */
enum ErrorStatusCode {
  BAD_REQUEST = 400,
  UNAUTHORIZED = 401,
  NOT_FOUND = 404,
  METHOD_NOT_ALLOWED = 405,
  TIMEOUT = 408,
  CONFLICT = 409,
  INTERNAL_SERVER_ERROR = 500,
  NOT_IMPLEMENTED = 501,
  SERVICE_UNAVAILABLE = 503,
}

const errHandler = async (error: AxiosError) => {
  // Unexpected network error
  if (!error) {
    showError(
      'ネットワーク通信に失敗しました。時間を置いてから再度アクセスしてみてください。'
    );
  }
  const errorCodeMsg = `StatusCode: ${error.response!.status}`;
  switch (error.response!.status) {
    case ErrorStatusCode.BAD_REQUEST:
      showError(
        `送信情報に誤りがあります。再度ご確認ください。${errorCodeMsg}`
      );
      break;
    case ErrorStatusCode.UNAUTHORIZED:
      showError(
        `セッションが切れています。再度ログインを行なって下さい。${errorCodeMsg}`
      );
      const routeInfo = {
        path: router.currentRoute.path,
        name: router.currentRoute.name,
        query: router.currentRoute.query,
      };
      // router.currentRouteをそのまま渡すとstringifyでcircleエラーが起きる
      store.dispatch(TYPES.LAST_PAGE_ROUTE, routeInfo);
      router.push('/login');
      break;
    case ErrorStatusCode.NOT_FOUND:
      showError(`URL先が見つかりません。${errorCodeMsg}`);
      break;
    case ErrorStatusCode.METHOD_NOT_ALLOWED:
      showError(`定義されていないメソッドです。${errorCodeMsg}`);
      break;
    case ErrorStatusCode.TIMEOUT:
      showError(
        `タイムアウトエラーが発生しました。時間を置いてから再度アクセスしてみてください。${errorCodeMsg}`
      );
      break;
    case ErrorStatusCode.CONFLICT:
      showError(
        `送信情報に誤りがあります。再度ご確認ください。${errorCodeMsg}`
      );
      break;
    default:
      showError(
        `ネットワーク通信に失敗しました。時間を置いてから再度アクセスしてみてください。${errorCodeMsg}`
      );
      break;
  }
};

const options: ToastOptions = {
  position: 'bottom-center',
  fullWidth: true,
  action: {
    text: 'キャンセル',
    onClick: function (e: any, toastObject: ToastObject) {
      toastObject.goAway(0);
    },
  },
};

function showError(msg: string) {
  Vue.toasted.clear();
  Vue.toasted.error(msg);
}

export default {
  reqHandler,
  errHandler,
};
