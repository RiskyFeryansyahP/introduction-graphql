<script lang="ts">
	import Button from "./components/button/button.svelte";
	import { initClient, mutation, operationStore, query } from '@urql/svelte';
	import type { Todos } from "./types/types";

	let todoText: string = "";

	initClient({
		url: 'http://localhost:8080/query',
	});

  const todos = operationStore<Todos>(`
 	query {
		 todos {
			 id
			 text
			 done
		 }
	 } ,
  `,
  );

  // @ts-ignore
  const mutatenAddTodo = mutation({
	  query: `
	 	mutation addTodo($text: String!) {
			createTodo(input: {
				text: $text
			}) {
				id
				text
			}
		}	 
	`,
  })

  // @ts-ignore
  const mutateRemoveTodo = mutation({
	  query: `
	 	mutation deleteTodo($id: ID!) {
			deleteTodo(id: $id){
				id
			}
		}
	`,
  })

  // @ts-ignore
  const mutatenUpdateTodo = mutation({
	  query: `
	  	mutation updateTodo($id: ID!) {
			updateTodo(id: $id) {
				id
				text
				done
			}
		} 
	`,
  })

  const refresh = () => {
	console.log("trigger");
    $todos.context = { requestPolicy: 'network-only' };
  }

  const handleAddTodo = () => {
	mutatenAddTodo({text: todoText}).then(result => {
		console.log(result.data, result.error);
	});
	refresh();
  }

  const handleRemoveTodo = (id: string) => {
	mutateRemoveTodo({id}).then(result => {
		console.log(result.data, result.error);
	});

	refresh();
  }

  const handleUpdateTodo = (id: string) => {
	mutatenUpdateTodo({id}).then(result => {
		console.log(result.data, result.error);
	});
	refresh();
  }

  query(todos);
</script>

<main class="h-full w-full flex items-center justify-center bg-gray-100 font-sans">
	<div class="bg-white rounded shadow p-6 m-4 w-full lg:w-3/4 lg:max-w-lg">
		<div class="mb-4">
			<h1 class="text-grey-darkest">Todo List</h1>
			<div class="flex mt-4">
				<input
				class="shadow appearance-none border rounded w-full py-2 px-3 mr-4 text-gray-600"
				placeholder="Add Todo"
				bind:value={todoText}
				/>
				<Button status="done" on:click={handleAddTodo}>add</Button>
			</div>
		</div>
		<div>
			{#if $todos.fetching}
				<p> Loading... </p>
			{:else if $todos.error}
				<p> Oh no... {$todos.error.message} </p>
			{:else}
				{#each $todos.data.todos as todo}
					<div class="flex mb-4 items-center">
						<p class="w-full text-gray-700" class:line-through="{todo.done}">
							{todo.text}
						</p>
						<Button status={todo.done ? "not done": "done"} on:click={() => handleUpdateTodo(todo.id)}>{todo.done ? "Un-Done": "Done"}</Button>
						<Button status="remove" on:click={() => handleRemoveTodo(todo.id)}>Remove</Button>
					</div>
				{/each}
			{/if}
		</div>
	</div>
</main>

<style global lang="postcss">
	@tailwind base;
	@tailwind components;
	@tailwind utilities;
</style>