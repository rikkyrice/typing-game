import { MutationTree } from 'vuex';
import { State, RouteInfo } from '@/store/types';
import { TypeWord } from '@/models/types/typeWord';
import { TokenInfo } from '@/models/user';

const mutations: MutationTree<State> = {
  setAuthState(state: State, tokenInfo: TokenInfo) {
    state.auth.token = tokenInfo.token;
    state.auth.userId = tokenInfo.userId;
  },
  setLastPageRoute(state: State, lastPageRoute: RouteInfo) {
    state.auth.lastPageRoute = lastPageRoute;
  },
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
