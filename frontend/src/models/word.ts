export interface WordArray {
  matched: number;
  words: Word[];
}

export interface Word {
  id: string;
  word: string;
  yomi: string;
  meaning: string;
  mYomi: string;
  explanation: string;
  isRemembered: boolean;
  createdAt: string;
  updatedAt: string;
}

export const initializedWordArray: WordArray = {
    matched: 0,
    words: [
      {
        id: '',
        word: '',
        yomi: '',
        meaning: '',
        mYomi: '',
        explanation: '',
        isRemembered: false,
        createdAt: '',
        updatedAt: '',
      },
    ],
}
