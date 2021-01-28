declare module "*.vue" {
  import { defineComponent } from "vue";
  const component: ReturnType<typeof defineComponent>;
  export default component;
}

interface Todos {
  SaveList(string: String): Promise<string>;
  LoadList(): Promise<string>;
  SaveAs(string: String): Promise<any>;
  LoadNewList(): Promise<string>;
}

interface FileManager {
  HandleFile(string: string): Promise<string>;
  Convert(): Promise<any>;
  SetOutDir(): Promise<any>;
}

declare interface Window {
  backend: {
    basic(): Promise<string>;
    Todos: Todos;
    FileManager: FileManager;
  };
}
