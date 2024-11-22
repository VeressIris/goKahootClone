<script>
  import "../app.css";
  import Page from "./+page.svelte";
  let { children } = $props();
</script>

<div class="flex justify-center items-center flex-col my-8 mx-20">
  {@render children()}
</div>
