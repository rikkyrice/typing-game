import { ActionTree } from 'vuex';
import { TYPES } from '@/store/mutation-types';
import { State, RouteInfo } from '@/store/types';

const actions: ActionTree<State, State> = {
  [TYPES.RESET_GAME]({ commit }) {
    commit('resetGame');
  }
}

export default actions;
