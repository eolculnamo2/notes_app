<script lang="ts">
  import {
    TextInput,
    TextArea,
    ButtonSet,
    Button,
    ComposedModal,
    ModalHeader,
    ModalBody,
    ModalFooter,
    Grid,
    Row,
    Column,
    PasswordInput,
  } from "carbon-components-svelte";
  import { SaveFile } from "../../wailsjs/go/main/App.js";

  let resultText = "";
  let fileName = "";
  let noteText = "";
  let password = "";
  let confirmPassword = "";
  let showConfirmCancelModal = false;

  function clearForm() {
    fileName = "";
    noteText = "";
    password = "";
    resultText = "";
  }

  async function saveFile() {
    if (!fileName.length || !noteText.length) {
      resultText = "File name and note text are required";
      return;
    }
    if (password !== confirmPassword) {
      resultText = "Password and confirm password do not match";
      return;
    }
    resultText = await SaveFile(fileName, noteText, password.trim());
  }
</script>

<h2>Save a new file</h2>
{#if resultText.length > 0}
  <p>{resultText}</p>
{/if}
<Grid padding>
  <Row>
    <Column>
      <TextInput
        labelText="File Name (with ext)"
        placeholder="foo.txt"
        bind:value={fileName}
      />
    </Column>
  </Row>
  <Row>
    <Column>
      <TextArea
        labelText="Content"
        placeholder="Once upon a time..."
        bind:value={noteText}
      />
    </Column>
  </Row>

  <Row>
    <Column>
      <PasswordInput
        labelText="Password"
        placeholder="Leave blank for default encryption"
        bind:value={password}
      />
    </Column>
  </Row>
  <Row>
    <Column>
      <PasswordInput
        labelText="Confirm Password"
        placeholder="Leave blank for default encryption"
        bind:value={confirmPassword}
      />
    </Column>
  </Row>
  <Row>
    <Column>
      <ButtonSet>
        <Button
          kind="secondary"
          on:click={() => (showConfirmCancelModal = true)}>Cancel</Button
        >
        <Button on:click={saveFile}>Submit</Button>
      </ButtonSet>
    </Column>
  </Row>
</Grid>

<ComposedModal
  open={showConfirmCancelModal}
  on:click:button--primary={clearForm}
  on:close={() => (showConfirmCancelModal = false)}
>
  <ModalHeader label="Changes" title="Confirm changes" />
  <ModalBody hasForm>
    <p>
      Are you sure you wish to erase unsaved changes? This action cannot be
      undone
    </p>
  </ModalBody>
  <ModalFooter primaryButtonText="Proceed" secondaryButtonText="Cancel" />
</ComposedModal>
