export class TypeWord {
  word: string;
  yomi: string;
  typeWord: string[][];

  constructor(word: string, yomi: string, typeWord: string[][]) {
    this.word = word;
    this.yomi = yomi;
    this.typeWord = typeWord;
  }
}
