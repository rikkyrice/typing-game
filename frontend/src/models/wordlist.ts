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
