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
  switchCleared(state: State, cleared: boolean) {
    state.gameCleared = cleared;
  }
};

export default mutations;
