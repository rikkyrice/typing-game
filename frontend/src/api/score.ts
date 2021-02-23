import APIClient from '@/api/common/apiClient';
import { API_PATHS } from '@/api/common/apiPaths';
import { ScoreArray, Score } from '@/models/score';

function getScores(wordListId: string) {
  return APIClient.getAPI<ScoreArray>(API_PATHS.SCORES.GET(wordListId))
    .then((data) => Promise.resolve(data))
    .catch((err) => Promise.reject(err));
}

function getScoreLatest(scoreId: string) {
  return APIClient.getAPI<Score>(API_PATHS.SCORE.GET(scoreId))
    .then((data) => Promise.resolve(data))
    .catch((err) => Promise.reject(err))
}

function postScore(
  playCount: number,
  clearTypeCount: number,
  missTypeCount: number,
) {
  const body: ScoreRequestBody = {
    playCount,
    clearTypeCount,
    missTypeCount,
  };
  return APIClient.postAPI<Score>(API_PATHS.SCORE.POST, body)
    .then((data) => Promise.resolve(data))
    .catch((err) => Promise.reject(err))
}

function deleteScore(wordListId: string) {
  return APIClient.deleteAPI(API_PATHS.SCORE.DELETE(wordListId)).catch((err) =>
    Promise.reject(err)
  );
}

interface ScoreRequestBody {
  playCount: number;
  clearTypeCount: number;
  missTypeCount: number;
}

export default {
  getScores,
  getScoreLatest,
  postScore,
  deleteScore,
}