<script lang="ts">
  import Todo from "./lib/Todo.svelte";
  import type { TodoItem } from "./lib/types";

  let todos: TodoItem[] = $state([]);

  async function fetchTodos() {
    try {
      const response = await fetch("http://localhost:8080/");
      if (response.status !== 200) {
        console.error("Error fetching data. Response status not 200");
        return;
      }

      todos = await response.json();
    } catch (e) {
      console.error("Could not connect to server. Ensure it is running.", e);
    }
  }

  // Initially fetch todos on page load
  $effect(() => {
    fetchTodos();
  });
</script>

<main class="app">
  <header class="app-header">
    <h1>TODO</h1>
  </header>

  <div class="todo-list">
    {#each todos as todo}
      <Todo title={todo.title} description={todo.description} />
    {/each}
  </div>

  <h2 class="todo-list-form-header">Add a Todo</h2>
  <form class="todo-list-form">
    <input placeholder="Title" name="title" />
    <input placeholder="Description" name="description" />
    <button>Add Todo</button>
  </form>
</main>

<style>
  .app {
    color: white;
    background-color: #282c34;

    text-align: center;
    font-size: 24px;

    min-height: 100vh;
    padding: 20px;
  }

  .app-header {
    font-size: calc(10px + 4vmin);
    margin-top: 50px;
  }

  .todo-list {
    margin: 50px 100px 0px 100px;
  }

  .todo-list-form-header {
    margin-top: 100px;
  }

  .todo-list-form {
    margin-top: 10px;
  }
</style>
