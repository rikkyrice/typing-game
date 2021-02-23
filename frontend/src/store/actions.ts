import { ActionTree } from 'vuex';
import { TYPES } from '@/store/mutation-types';
import { State, RouteInfo } from '@/store/types';
import { TypeWord } from '@/models/types/typeWord';
import { TokenInfo } from '@/models/user';

const actions: ActionTree<State, State> = {
  [TYPES.LOGIN]({ commit }, tokenInfo: TokenInfo) {
    commit('setAuthState', tokenInfo);
  },
  [TYPES.LAST_PAGE_ROUTE]({ commit }, lastPageRoute: RouteInfo) {
    commit('setLastPageRoute', lastPageRoute);
  },
  [TYPES.SHIFT_TYPEWORD]({ commit }, typeWord: TypeWord) {
    commit('shiftTypeWord', typeWord);
  },
  [TYPES.SHIFT_TYPEMEANING]({ commit }, typeWord: TypeWord) {
    commit('shiftTypeMeaning', typeWord);
  },
  [TYPES.SWITCH_CLEARED]({ commit }, cleared: boolean) {
    commit('switchCleared', cleared);
  },
}

export default actions;
