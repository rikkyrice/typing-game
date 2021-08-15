import { TypeWord } from './types/typeWord'

export interface WordArray {
  matched: number;
  words: Word[];
}

export interface TypingWordArray {
  matched: number;
  typingWords: TypingWord[];
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

export interface TypingWord {
  word: TypeWord;
  meaning: TypeWord;
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

export const initializedTypingWordArray: TypingWordArray = {
    matched: 0,
    typingWords: [
      {
        word: {
          name: '',
          yomi: '',
          types: [['']],
        },
        meaning: {
          name: '',
          yomi: '',
          types: [['']],
        },
        explanation: '',
        isRemembered: false,
        createdAt: '',
        updatedAt: '',
      },
    ],
}
