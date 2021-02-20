export interface WordListArray {
  matched: number;
  wordLists: WordList[];
}

export interface WordList {
  wordListId: string;
  wordListTitle: string;
  explanation: string;
  createdAt: string;
  updatedAt: string;
}

export interface WordListDetail {
  wordListId: string;
  wordListTitle: string;
  explanation: string;
  createdAt: string;
  updatedAt: string;
  words: WordArray;
}

export interface WordArray {
  matched: number;
  words: Word[];
}

export interface Word {
  wordId: string;
  word: string;
  meaning: string;
  explanation: string;
  isRemembered: boolean;
  createdAt: string;
  updatedAt: string;
}