declare module "*.vue" {
  import { defineComponent } from "vue";
  const component: ReturnType<typeof defineComponent>;
  export default component;
}

interface Todos {
  SaveList(string): Promise<string>;
  LoadList(): Promise<any>;
}

declare interface Window {
  backend: {
    basic(): Promise<string>;
    Todos: Todos;
  };
}
