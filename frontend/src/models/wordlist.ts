export interface WordListArray {
  matched: number;
  wordLists: WordList[];
}

export interface WordList {
  id: string;
  title: string;
  explanation: string;
  createdAt: string;
  updatedAt: string;
}
