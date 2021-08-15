export class TypeWord {
  name: string;
  yomi: string;
  types: string[][];

  constructor(name: string, yomi: string, types: string[][]) {
    this.name = name;
    this.yomi = yomi;
    this.types = types;
  }
}
