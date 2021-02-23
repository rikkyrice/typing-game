import axios from '@/api/common/axios';
import ApiHandler from '@/api/common/apiHandler';

axios.interceptors.request.use(ApiHandler.reqHandler);

export default {
  /**
   * backendへのgetAPI
   * @param path 各サービスのPath
   * @param body 各サービスのBodyパラメーター
   */
  getAPI: async <T>(path: string, body?: any) => {
    try {
      const res = await axios.get<T>(path, {
        params: body ? encode<any>(body) : undefined,
      });
      return res ? decode<T>(res.data) : ({} as T);
    } catch (error) {
      ApiHandler.errHandler(error);
      throw error;
    }
  },
  /**
   * backendへのpostAPI
   * @param path 各サービスのPath
   * @param body 各サービスのBodyパラメーター
   */
  postAPI: async <T>(path: string, body?: any) => {
    try {
      const res = await axios.post<T>(
        path,
        body ? encode<any>(body) : undefined
      );
      return res ? decode<T>(res.data) : ({} as T);
    } catch (error) {
      ApiHandler.errHandler(error);
      throw error;
    }
  },
  /**
   * backendへのputAPI
   * @param path 各サービスのPath
   * @param body 各サービスのBodyパラメーター
   */
  putAPI: async <T>(path: string, body?: any) => {
    try {
      const res = await axios.put<T>(
        path,
        body ? encode<any>(body) : undefined
      );
      return res ? decode<T>(res.data) : ({} as T);
    } catch (error) {
      ApiHandler.errHandler(error);
      throw error;
    }
  },
  /**
   * backendへのdeleteAPI
   * @param path 各サービスのPath
   * @param body 各サービスのBodyパラメーター
   */
  deleteAPI: async <T>(path: string, body?: any) => {
    try {
      const res = await axios.delete<T>(
        path,
        body ? encode<any>(body) : undefined
      );
      return res ? decode<T>(res.data) : ({} as T);
    } catch (error) {
      ApiHandler.errHandler(error);
      throw error;
    }
  },
};
/* eslint-enable: no-console */

function encode<T>(data: T): T {
  return JSON.parse(JSON.stringify(data, encodeReplacer));
}

function decode<T>(data: T): T {
  return JSON.parse(JSON.stringify(data, decodeReplacer));
}

function encodeReplacer(key: string, value: any) {
  if (typeof value === 'string') {
    return value.replace(/%/g, '%25');
  }
  return value;
}

function decodeReplacer(key: string, value: any) {
  if (typeof value === 'string') {
    return value.replace(/%25/g, '%');
  }
  return value;
}
