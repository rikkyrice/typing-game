export interface WordArray {
  matched: number;
  words: Word[];
}

export interface Word {
  wordId: string;
  word: string;
  yomi: string;
  meaning: string;
  explanation: string;
  isRemembered: boolean;
  createdAt: string;
  updatedAt: string;
}