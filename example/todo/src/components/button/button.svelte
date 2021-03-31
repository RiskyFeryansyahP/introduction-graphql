<script lang="ts">
    type Status = "done" | "not done" | "remove";
    export let status: Status;
</script>

<button on:click class="flex-no-shrink p-2 ml-4 mr-2 border-2 rounded hover:text-white" class:text-green-300="{status === "done"}" class:border-green-400="{status === "done"}" class:hover:bg-green-500="{status === "done"}" class:text-gray-300="{status === "not done"}" class:hover:border-gray-400="{status === "not done"}" class:hover:bg-gray-500="{status === "not done"}" class:hover:bg-red-500="{status === "remove"}" class:border-red-400="{status === "remove"}" class:text-red-300="{status === "remove"}">
    <slot></slot>
</button>

<style></style>