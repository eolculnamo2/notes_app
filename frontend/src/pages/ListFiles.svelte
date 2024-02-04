<script lang="ts">
  import Shell from "../layouts/Shell.svelte";
  import { ListFiles } from "../../wailsjs/go/main/App.js";
  import { goToPage } from "../globalState";
  const files = ListFiles();
</script>

<Shell>
  <h1>Saved Files</h1>
  <br />
  {#await files}
    <p>Loading...</p>
  {:then files}
    {#each files as file}
      <h3 on:click={() => goToPage({ page: "view-file", fileName: file })}>
        {file}
      </h3>
    {/each}
  {:catch err}
    <p>{err}</p>
  {/await}
</Shell>
