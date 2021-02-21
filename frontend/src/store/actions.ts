import { ActionTree } from 'vuex';
import { TYPES } from '@/store/mutation-types';
import { State, RouteInfo } from '@/store/types';
import { TypeWord } from '@/models/types/typeWord';

const actions: ActionTree<State, State> = {
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
