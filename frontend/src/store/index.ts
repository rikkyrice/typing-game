import Vue from 'vue';
import Vuex, { StoreOptions } from 'vuex';
import createPersistedState from 'vuex-persistedstate';
import { State } from '@/store/types';
import mutations from '@/store/mutations';
import actions from '@/store/actions';
import { state } from '@/store/state';

Vue.use(Vuex);

const storeOptions: StoreOptions<State> = {
  state,
  mutations,
  actions,
  plugins: [createPersistedState({ storage: window.localStorage })],
};

export default new Vuex.Store(storeOptions);
