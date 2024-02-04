<script lang="ts">
  import Home from "./pages/Home.svelte";
  import { currentPage } from "./globalState";
  import ListFiles from "./pages/ListFiles.svelte";
  import FileView from "./pages/FileView.svelte";
</script>

<main>
  {#if $currentPage.page === "home"}
    <Home />
  {:else if $currentPage.page === "list-files"}
    <ListFiles />
  {:else if $currentPage.page === "view-file"}
    <FileView fileName={$currentPage.fileName} />
  {/if}
</main>
