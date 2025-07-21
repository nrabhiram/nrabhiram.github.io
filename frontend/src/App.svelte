<!-- App.svelte -->
<script lang="ts">
  import Layout from "./lib/Layout.svelte";
  import MainPage from "./lib/views/MainPage.svelte";
  import PostPage from "./lib/views/PostPage.svelte";
  import type { Artifact, Link } from "./types";
  import { matchesPath } from "./utils";

  export let url = "";
  export let primaryNavData: Link[] = [];
  export let secondaryNavData: Link | null = null;
  export let artifact: Artifact | null = null;
</script>

<Layout 
  primaryNavData={primaryNavData} 
  secondaryNavData={secondaryNavData}
  url={url}
  heading={artifact?.name}
  prevLink={matchesPath(url, '/now/:slug') && !!artifact?.prev ? artifact.prev : null}
  nextLink={matchesPath(url, '/now/:slug') && !!artifact?.next ? artifact.next : null}
>
  {#if matchesPath(url, '/about/:slug')}
    <MainPage {artifact} />
  {:else if matchesPath(url, '/now/:slug')}
    <PostPage {artifact} />
  {:else if matchesPath(url, '/:slug')}
    <MainPage {artifact} />
  {:else if matchesPath(url, '/')}
    <MainPage {artifact} />
  {/if}
</Layout>
