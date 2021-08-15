import internal from "stream";

export interface WordListArray {
  matched: number;
  wordlists: WordList[];
}

export interface WordList {
  id: string;
  title: string;
  explanation: string;
  createdAt: string;
  updatedAt: string;
}

export interface WordListSummaryArray {
  matched: number;
  wordlists: WordListSummary[];
}

export interface WordListSummary {
  id: string;
  title: string;
  explanation: string;
  wordCount: number;
  playCount: number;
  playedAt: string;
  createdAt: string;
  updatedAt: string;
}

