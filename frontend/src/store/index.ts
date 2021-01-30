import { createStore } from "vuex";

export default createStore({
  state: {
    config: {
      outDir: ""
    }
  },
  getters: {
    config(state) {
      return state.config;
    }
  },
  mutations: {},
  actions: {},
  modules: {}
});
