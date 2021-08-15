import APIClient from '@/api/common/apiClient';
import { API_PATHS } from '@/api/common/apiPaths';
import { WordArray, TypingWordArray, Word } from '@/models/word';

function getWords(wordListId: string) {
  return APIClient.getAPI<WordArray>(API_PATHS.WORDS.GET(wordListId))
    .then((data) => Promise.resolve(data))
    .catch((err) => Promise.reject(err));
}

function getWord(wordId: string) {
  return APIClient.getAPI<Word>(API_PATHS.WORD.GET(wordId))
    .then((data) => Promise.resolve(data))
    .catch((err) => Promise.reject(err))
}

function postWord(
  wordListId: string,
  word: string,
  meaning: string,
  explanation: string,
  isRemembered: boolean,

) {
  const body: WordRequestBody = {
    wordListId,
    word,
    meaning,
    explanation,
    isRemembered,
  };
  return APIClient.postAPI<Word>(API_PATHS.WORDLIST.POST, body)
    .then((data) => Promise.resolve(data))
    .catch((err) => Promise.reject(err))
}

function putWordList(
  wordId: string,
  wordListId: string,
  word: string,
  meaning: string,
  explanation: string,
  isRemembered: boolean,
) {
  const body: WordRequestBody = {
    wordListId,
    word,
    meaning,
    explanation,
    isRemembered,
  };
  return APIClient.putAPI<Word>(API_PATHS.WORD.PUT(wordId), body)
    .then((data) => Promise.resolve(data))
    .catch((err) => Promise.reject(err))
}

function deleteWordList(wordId: string) {
  return APIClient.deleteAPI(API_PATHS.WORDLIST.DELETE(wordId)).catch((err) =>
    Promise.reject(err)
  );
}

function getTypingWords(wordListId: string) {
  return APIClient.getAPI<TypingWordArray>(API_PATHS.TYPING.GET(wordListId))
    .then((data) => Promise.resolve(data))
    .catch((err) => Promise.reject(err));
}

interface WordRequestBody {
  wordListId: string;
  word: string;
  meaning: string;
  explanation: string;
  isRemembered: boolean;
}

export default {
  getWords,
  getWord,
  postWord,
  putWordList,
  deleteWordList,
  getTypingWords
}
