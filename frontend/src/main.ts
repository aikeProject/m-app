import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import * as Wails from "@wailsapp/runtime";
import installElementPlus from "plugins/element";
import "assets/css/base.css";
import "assets/css/app.css";

Wails.Init(() => {
  const app = createApp(App);
  installElementPlus(app);
  app.use(router);
  app.mount("#app");
});
