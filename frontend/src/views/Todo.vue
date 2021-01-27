<template>
  <div>
    <h2 v-if="errorMessage.length > 0">{{ errorMessage }}</h2>
    <section class="todoapp" v-cloak>
      <header class="header">
        <h1>todos</h1>
        <div class="buttons">
          <ul class="filters">
            <li>
              <a @click="saveAs">Save As</a>
            </li>
            <li>
              <a @click="loadNewList">Load</a>
            </li>
          </ul>
        </div>
        <input
          class="new-todo"
          autofocus
          autocomplete="off"
          placeholder="What needs to be done?"
          v-model="newTodo"
          @keyup.enter="addTodo"
        />
      </header>
      <section class="main">
        <ul class="todo-list" v-show="todos.length">
          <li
            class="todo"
            v-for="todo in todos"
            :key="todo.id"
            :class="{ completed: todo.completed, editing: todo === editedTodo }"
          >
            <div class="view">
              <input class="toggle" type="checkbox" v-model="todo.completed" />
              <label @dblclick="editTodo(todo)">{{ todo.title }}</label>
              <button class="destroy" @click="removeTodo(todo)"></button>
            </div>
            <input
              class="edit"
              type="text"
              v-model.lazy="todo.title"
              v-todo-focus="todo === editedTodo"
              @keyup.enter="doneEdit(todo)"
              @blur="doneEdit(todo)"
              @keyup.esc="cancelEdit(todo)"
            />
          </li>
        </ul>
      </section>
    </section>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Wails from "@wailsapp/runtime";

interface Todo {
  id: number;
  title: string;
  completed: boolean;
}

export default defineComponent({
  name: "Todo",
  data() {
    return {
      errorMessage: "",
      newTodo: "",
      beforeEditCache: "",
      editedTodo: null as Todo | null,
      todos: [{ id: 0, title: "My test todo item", completed: true }] as Todo[],
      loading: false
    };
  },
  directives: {
    "todo-focus": function(el, binding) {
      if (binding.value) {
        el.focus();
      }
    }
  },
  watch: {
    todos: {
      handler(todos) {
        if (this.loading) {
          this.loading = false;
          return;
        }
        // 打印日志
        Wails.Log.Info("Todo List: " + JSON.stringify(todos));
        window.backend.Todos.SaveList(JSON.stringify(todos, null, 2));
      },
      deep: true
    }
  },
  methods: {
    addTodo() {
      const value = this.newTodo && this.newTodo.trim();
      if (!value) {
        return;
      }
      this.todos.push({
        id: this.todos.length,
        title: value,
        completed: false
      });
      this.newTodo = "";
    },
    removeTodo(todo: Todo) {
      const index = this.todos.indexOf(todo);
      this.todos.splice(index, 1);

      for (let i = 0; i < this.todos.length; i++) {
        this.todos[i].id = i;
      }
    },
    editTodo: function(todo: Todo) {
      this.beforeEditCache = todo.title;
      this.editedTodo = todo;
    },
    doneEdit: function(todo: Todo) {
      if (!this.editedTodo) {
        return;
      }
      this.editedTodo = null;
      todo.title = todo.title.trim();
      if (!todo.title) {
        this.removeTodo(todo);
      }
    },
    cancelEdit: function(todo: Todo) {
      this.editedTodo = null;
      todo.title = this.beforeEditCache;
    },
    saveAs() {
      window.backend.Todos.SaveAs(JSON.stringify(this.todos, null, 2));
    },
    loadNewList() {
      window.backend.Todos.LoadNewList();
    },
    setErrorMessage(message: string) {
      this.errorMessage = message;
      setTimeout(() => {
        this.errorMessage = "";
      }, 3000);
    },
    // 加载列表
    loadList: function() {
      window.backend.Todos.LoadList()
        .then(list => {
          try {
            this.todos = JSON.parse(list);
            this.loading = true;
          } catch (e) {
            this.setErrorMessage("Unable to load todo list");
          }
        })
        .catch(error => {
          this.setErrorMessage(error.message);
        });
    }
  },
  mounted() {
    Wails.Events.On("fileModified", () => {
      // this.setErrorMessage("File Modified");
      this.loadList();
    });

    // 监听来自Go程序的消息
    Wails.Events.On("error", message => {
      this.setErrorMessage(message);
    });

    this.loadList();
  }
});
</script>

<style scoped lang="stylus">
h2
  text-align center
  color white
  background-color red
  min-width 230px
  max-width 550px
  padding 1rem
  border-radius 0.5rem

.buttons {
  padding: 10px 20px;
  box-shadow: inset 0 -2px 1px rgba(0, 0, 0, 0.1);
  text-align: center;
  border-color: rgba(175, 47, 47, 0.2);
}

.buttons ul li a {
  margin: 10px;
}

.buttons li {
  border-color: rgba(175, 47, 47, 0.1);
}
</style>
