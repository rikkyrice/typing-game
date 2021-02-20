export interface ScoreArray {
  matched: number;
  scores: Score[];
}

export interface Score {
  scoreId: number;
  playCount: number;
  clearTypeCount: number;
  missTypeCount: number;
  playedAt: string;
}