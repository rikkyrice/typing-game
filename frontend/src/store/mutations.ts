import { MutationTree } from 'vuex';
import { State } from '@/store/types';
import { TypeWord } from '@/models/types/typeWord';

const mutations: MutationTree<State> = {
  shiftTypeWord(state: State, typeWord: TypeWord) {
    state.typeWord.typeWord = typeWord;
  },
  shiftTypeMeaning(state: State, typeWord: TypeWord) {
    state.typeWord.typeMeaning = typeWord;
  },
  resetGame(state: State) {
    state.typeWord.typeWord = new TypeWord('', '', [[]]);
    state.typeWord.typeMeaning = new TypeWord('', '', [[]]);
  }
};

export default mutations;
