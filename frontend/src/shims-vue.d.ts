declare module "*.vue" {
  import { defineComponent } from "vue";
  const component: ReturnType<typeof defineComponent>;
  export default component;
}

namespace backend {
  function basic(): Promise<string>;
  function saveList(string): Promise<string>;
}
