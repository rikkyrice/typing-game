export interface User {
  userId: string;
  mail: string;
  password: string;
  createdAt: string;
}

export interface TokenInfo {
  token: string
  userId: string
}