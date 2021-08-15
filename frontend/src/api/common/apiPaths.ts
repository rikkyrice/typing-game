// users
const SIGNUP = {
  POST: '/users',
};
const LOGIN = {
  POST: '/users/login',
};

// wordlists
const WORDLISTS = {
  GET: '/wordlists',
};
const WORDLIST = {
  GET: (wordListId: string) => `/wordlists/${wordListId}`,
  POST: '/wordlists',
  PUT: (wordListId: string) => `/wordlists/${wordListId}`,
  DELETE: (wordListId: string) => `/wordlists/${wordListId}`,
};

// words
const WORDS = {
  GET: (wordListId: string) => `/words/wordlist/${wordListId}`,
};
const WORD = {
  POST: '/words',
  GET: (wordId: string) => `/words/${wordId}`,
  PUT: (wordId: string) => `/words/${wordId}`,
  DELETE: (wordId: string) => `/words/${wordId}`,
  PATCH: (wordId: string) => `/words/${wordId}/remembered`,
};
const TYPING = {
  GET: (wordListId: string) => `/words/wordlist/${wordListId}/typing`,
}

// scores
const SCORES = {
  GET: (wordListId: string) => `/scores/${wordListId}`,
};
const SCORE = {
  GET: (wordListId: string) => `/scores/${wordListId}/latest`,
  POST: '/scores',
  DELETE: (wordListId: string) => `/scores/${wordListId}`,
};

export const API_PATHS = {
  SIGNUP,
  LOGIN,
  WORDLISTS,
  WORDLIST,
  WORDS,
  TYPING,
  WORD,
  SCORES,
  SCORE,
};
