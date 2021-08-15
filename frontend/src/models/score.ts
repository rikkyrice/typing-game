export interface ScoreArray {
  matched: number;
  scores: Score[];
}

export interface Score {
  scoreId: number;
  clearTypeCount: number;
  missTypeCount: number;
  playedAt: string;
}