declare module "*.vue" {
  import { defineComponent } from "vue";
  const component: ReturnType<typeof defineComponent>;
  export default component;
}

interface Todos {
  SaveList(string: String): Promise<string>;
  LoadList(): Promise<any>;
  SaveAs(string: String): Promise<any>;
}

declare interface Window {
  backend: {
    basic(): Promise<string>;
    Todos: Todos;
  };
}
