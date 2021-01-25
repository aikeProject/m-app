<template>
  <div>
    <h2 v-if="errorMessage.length > 0">{{ errorMessage }}</h2>
    <section class="todoapp" v-cloak>
      <header class="header">
        <h1>todos</h1>
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
              v-model="todo.title"
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
      todos: [{ id: 0, title: "My test todo item", completed: true }] as Todo[]
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
        // 打印日志
        Wails.Log.Info("Todo List: " + JSON.stringify(todos));
        window.backend.saveList(JSON.stringify(todos));
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
    }
  },
  mounted() {
    window.backend
      .loadList()
      .then(list => {
        try {
          this.todos = JSON.parse(list);
        } catch (e) {
          this.errorMessage = "Unable to load todo list";
          setTimeout(() => {
            this.errorMessage = "";
          }, 3000);
        }
      })
      .catch(error => {
        this.errorMessage = error;
        setTimeout(() => {
          this.errorMessage = "";
        }, 3000);
      });
  }
});
</script>

<style scoped lang="stylus">
h2
  text-align: center;
  color: white;
  background-color: red;
  min-width: 230px;
  max-width: 550px;
  padding: 1rem;
  border-radius: 0.5rem;
</style>
