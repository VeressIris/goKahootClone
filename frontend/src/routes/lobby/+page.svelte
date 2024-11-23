<script>
  import UserCard from "$lib/userCard.svelte";

  let users = [
    "Alice",
    "Bob",
    "Charlie",
    "David",
    "Eve",
    "Frank",
    "Grace",
    "Hannah",
    "Isaac",
    "Jack",
  ];
</script>

<h1 class="text-3xl font-bold mb-2">Welcome to the lobby!</h1>
<p class="text-xl">Players connected:</p>
<div class="grid grid-cols-4 gap-2 my-3 md:grid-cols-6">
  {#each users as user}
    <UserCard username={user} />
  {/each}
</div>
<button
  class="px-6 py-2 bg-red-400 rounded-lg text-xl font-bold text-white mt-7 hover:bg-red-500 active:bg-red-600"
  >Start</button
>
