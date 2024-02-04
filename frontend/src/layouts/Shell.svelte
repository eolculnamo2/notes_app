<script lang="ts">
  import {
    Header,
    HeaderNav,
    HeaderNavItem,
    // HeaderNavMenu,
    SkipToContent,
    Content,
    Grid,
    Row,
    Column,
  } from "carbon-components-svelte";
  import { goToPage } from "../globalState";
  let isSideNavOpen = false;
</script>

<Header company="Rob" platformName="Notes App" bind:isSideNavOpen>
  <svelte:fragment slot="skip-to-content">
    <SkipToContent />
  </svelte:fragment>
  <HeaderNav>
    <HeaderNavItem
      on:click={() => goToPage({ page: "home" })}
      text="Create File"
    />
    <HeaderNavItem
      on:click={() => goToPage({ page: "list-files" })}
      text="Files List"
    />
    <!-- <HeaderNavMenu text="Menu"> -->
    <!--   <HeaderNavItem href="/" text="Link 1" /> -->
    <!--   <HeaderNavItem href="/" text="Link 2" /> -->
    <!--   <HeaderNavItem href="/" text="Link 3" /> -->
    <!-- </HeaderNavMenu> -->
  </HeaderNav>
</Header>

<Content>
  <Grid>
    <Row>
      <Column>
        <slot />
      </Column>
    </Row>
  </Grid>
</Content>
