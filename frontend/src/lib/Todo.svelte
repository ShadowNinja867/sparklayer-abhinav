<script lang="ts">
  import type { TodoItem } from "./types";

  const { title, description }: TodoItem = $props();
</script>

<div class="todo">
  <div class="todo-details">
    <p class="todo-title">{title}</p>
    <p class="todo-description">{description}</p>
  </div>
</div>

<style>
  .todo {
    display: flex;
    justify-content: space-between;
    align-items: center;

    padding: 0px 20px;

    background-color: #303540;

    border-style: solid;
    border-color: #464e61;
    border-width: 1px;
    border-radius: 15px;

    margin: 10px 0px;
  }

  .todo-details {
    display: flex;
    flex-direction: column;
    align-items: start;

    padding: 15px 0px;
  }

  .todo-title {
    font-size: 28px;
    font-weight: bold;

    margin: 0px;
    margin-bottom: 5px;
  }

  .todo-description {
    display: flex;
    margin: 0px;

    font-size: 20px;
    color: #ddd;

    text-align: left;
  }
</style>
