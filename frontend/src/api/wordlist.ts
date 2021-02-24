import APIClient from '@/api/common/apiClient';
import { API_PATHS } from '@/api/common/apiPaths';
import { WordListArray, WordList } from '@/models/wordlist';

function getWordLists() {
  return APIClient.getAPI<WordListArray>(API_PATHS.WORDLISTS.GET)
    .then((data) => Promise.resolve(data))
    .catch((err) => Promise.reject(err));
}

function getWordList(wordListId: string) {
  return APIClient.getAPI<WordList>(API_PATHS.WORDLIST.GET(wordListId))
    .then((data) => Promise.resolve(data))
    .catch((err) => Promise.reject(err))
}

function postWordList(
  wordListTitle: string,
  explanation: string,
) {
  const body: WordListRequestBody = {
    wordListTitle,
    explanation,
  };
  return APIClient.postAPI<WordList>(API_PATHS.WORDLIST.POST, body)
    .then((data) => Promise.resolve(data))
    .catch((err) => Promise.reject(err))
}

function putWordList(
  wordListId: string,
  wordListTitle: string,
  explanation: string,
) {
  const body: WordListRequestBody = {
    wordListTitle,
    explanation,
  };
  return APIClient.putAPI<WordList>(API_PATHS.WORDLIST.PUT(wordListId), body)
    .then((data) => Promise.resolve(data))
    .catch((err) => Promise.reject(err))
}

function deleteWordList(wordListId: string) {
  return APIClient.deleteAPI(API_PATHS.WORDLIST.DELETE(wordListId)).catch((err) =>
    Promise.reject(err)
  );
}

interface WordListRequestBody {
  wordListTitle: string;
  explanation: string;
}

export default {
  getWordLists,
  getWordList,
  postWordList,
  putWordList,
  deleteWordList,
}
