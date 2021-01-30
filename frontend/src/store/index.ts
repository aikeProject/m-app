import { createStore } from "vuex";

export default createStore({
  state: {
    config: {
      outDir: "",
      target: ""
    }
  },
  getters: {
    config(state) {
      return state.config;
    }
  },
  mutations: {
    setConfig(state, config) {
      state.config = config;
    }
  },
  actions: {
    getConfig(context) {
      window.backend.Config.GetAppConfig()
        .then(config => {
          console.log("config", config);
          context.commit("setConfig", config);
        })
        .catch(err => console.error(err));
    },
    setConfig(context, config) {
      context.commit("setConfig", config);
    }
  },
  modules: {}
});
