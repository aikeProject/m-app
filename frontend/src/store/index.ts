import { createStore } from "vuex";

export default createStore({
  state: {
    config: {
      outDir: "",
      target: "",
      prefix: "",
      suffix: ""
    },
    stats: {
      byteCount: 0,
      imageCount: 0
    }
  },
  getters: {
    config(state) {
      return state.config;
    },
    stats(state) {
      return state.stats;
    }
  },
  mutations: {
    setConfig(state, config) {
      state.config = config;
    },
    setStats(state, s) {
      state.stats = s;
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
      window.backend.Config.SetConfig(JSON.stringify(config))
        .then(() => {
          context.dispatch("getConfig").then();
        })
        .catch(console.error);
    },
    getStats(context) {
      window.backend.Stat.GetStats()
        .then(s => {
          context.commit("setStats", s);
        })
        .catch(err => {
          console.error(err);
        });
    },
    setStats(context, s) {
      context.commit("setStats", s);
    }
  },
  modules: {}
});
