<script lang="ts">
  import Shell from "../layouts/Shell.svelte";
  import { ReadSaved } from "../../wailsjs/go/main/App";
  import { Button, PasswordInput } from "carbon-components-svelte";

  export let fileName: string;
  let ready = false;
  let password = "";

  let fileContents: Promise<string> | null = null;
  function readFile() {
    fileContents = ReadSaved(fileName, password);
  }
</script>

<Shell>
  {#if ready && fileContents}
    <h1>File: <strong>{fileName}</strong></h1>
    {#await fileContents}
      <p>Loading...</p>
    {:then content}
      {content}
    {:catch err}
      <p>{err}</p>
    {/await}
  {:else}
    <PasswordInput
      labelText="Password"
      placeholder="Leave blank if no password"
      bind:value={password}
    />
    <br />
    <Button
      on:click={() => {
        ready = true;
        readFile();
      }}
    >
      View File
    </Button>
  {/if}
</Shell>
