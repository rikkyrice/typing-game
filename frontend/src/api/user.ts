import APIClient from '@/api/common/apiClient';
import { API_PATHS } from '@/api/common/apiPaths';
import { User, TokenInfo } from '@/models/user';
import score from '@/store';

function signup(
  userId: string,
  mail: string,
  password: string,
) {
  const body: SignupPostRequest = {
    userId,
    mail,
    password,
  };
  return APIClient.postAPI<User>(API_PATHS.SIGNUP.POST, body)
    .then((data) => Promise.resolve(data))
    .catch((err) => Promise.reject(err)
  );
}

function login(
  userId: string,
  password: string,
) {
  const body: LoginPostRequest = {
    userId,
    password,
  };
  return APIClient.postAPI<TokenInfo>(API_PATHS.LOGIN.POST, body)
    .then((data) => Promise.resolve(data))
    .catch((err) => Promise.reject(err)
  );
}

interface SignupPostRequest {
  userId: string;
  mail: string;
  password: string;
}

interface LoginPostRequest {
  userId: string;
  password: string;
}

export default {
  signup,
  login,
}
