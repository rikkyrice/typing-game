import Vue from 'vue';
import Vuex, { StoreOptions } from 'vuex';
import createPersistedState from 'vuex-persistedstate';
import { State } from '@/store/types';
import { state } from '@/store/state';

Vue.use(Vuex);

const storeOptions: StoreOptions<State> = {
  state,
  plugins: [createPersistedState({ storage: window.localStorage })],
};

export default new Vuex.Store(storeOptions);
