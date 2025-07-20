<!-- App.svelte -->
<script lang="ts">
  import Layout from "./lib/Layout.svelte";
  import MainPage from "./lib/views/MainPage.svelte";
  import PostPage from "./lib/views/PostPage.svelte";
  import type { Artifact, Link } from "./types";
  import { getPageTitle, matchesPath, VAXITAS_BASE_URL } from "./utils";

  export let url = "";
  export let primaryNavData: Link[] = [];
  export let secondaryNavData: Link | null = null;
  export let artifact: Artifact | null = null;
  
  const title = getPageTitle(artifact);
</script>

<svelte:head>
  <title>{title}</title>
  <meta name="title" content={title}>
  {#if artifact?.summary}
    <meta name="description" content={artifact.summary}>
  {/if}
  <meta name="author" content="Abhiram Reddy">
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <link rel="icon" type="image/png" href="/favicon-96x96.png" sizes="96x96">
  <link rel="icon" type="image/svg+xml" href="/favicon.svg" />
  <link rel="shortcut icon" href="/favicon.ico" />
  <link rel="apple-touch-icon" sizes="180x180" href="/apple-touch-icon.png">
  <link rel="alternate" type="application/atom+xml" title="Posts & Logs" href={`${VAXITAS_BASE_URL}/feed.atom`}>
  <meta property="og:title" content={title}>
  {#if artifact?.summary}
    <meta property="og:description" content={artifact.summary}>
  {/if}
  <meta property="og:locale" content="en_US">
  <meta property="og:image" content={`${VAXITAS_BASE_URL}/card.png`}>
  <meta property="og:image:secure_url" content={`${VAXITAS_BASE_URL}/card.png`}>
  <meta property="og:image:type" content="image/png">
  <meta property="og:image:width" content={'1920'}>
  <meta property="og:image:height" content={'1280'}>
  <meta property="og:site_name" content="Vaxitas">
  <meta property="og:url" content={`${VAXITAS_BASE_URL}${artifact?.path}`}>
  <!-- condition for article or website -->
  <meta property="og:type" content="website">
  {#if artifact?.date}
    <meta property="og:published_time" content={artifact.date} />
  {/if}
  {#if artifact?.dateEdited}
    <meta property="og:updated_time" content={artifact.dateEdited}>
  {/if}
  <meta name="twitter:card" content="summary_large_image">
  <meta name="twitter:creator" content="@nrabhiram">
  <meta name="twitter:title" content={title}>
  {#if artifact?.summary}
    <meta name="twitter:description" content={artifact.summary}>
  {/if}
  <meta name="twitter:image" content={`${VAXITAS_BASE_URL}/card.png`}>
  <meta name="twitter:image:secure_url" content={`${VAXITAS_BASE_URL}/card.png`}>
</svelte:head>

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
