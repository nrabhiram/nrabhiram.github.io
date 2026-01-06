<script lang="ts">
  import { onMount } from "svelte";
  import { ThemeStore } from "../stores/theme";
  import { getLocalStorage, setLocalStorage, VAXITAS_LS_THEME_KEY } from "../utils";
  import { Navbar } from "./Navbar";
  import ThemeToggle from "./ThemeToggle.svelte";
  import type { AdjacentLink, Link } from "../types";
  import ThreeBars from "./icons/ThreeBars.svelte";

  const defaultTheme = typeof window !== 'undefined' && window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
  export let primaryNavData: Link[];
  export let secondaryNavData: Link | null = null;
  export let url: string = "";
  export let heading = "";
  export let prevLink: AdjacentLink | null = null;
  export let nextLink: AdjacentLink | null = null;
  let deviceType = '';
  let navOpen = true;
  let mobileNavOpen = false;
  let isDarkMode = getLocalStorage<'light' | 'dark'>(VAXITAS_LS_THEME_KEY, defaultTheme) === 'dark';
  let disableDrawerAnim = false;

  function openNavbar() {
    if (!deviceType) return;

    if (deviceType === 'desktop') {
      navOpen = true;
      return;
    }

    if (deviceType === 'mobile') {
      mobileNavOpen = true;
      return;
    }
  }

  function closeNavbar() {
    if (!deviceType) return;

    if (deviceType === 'desktop') {
      navOpen = false;
      return;
    }

    if (deviceType === 'mobile') {
      mobileNavOpen = false;
      return;
    }
  }

  function toggleNavbar() {
    if (!deviceType) return;

    if (deviceType === 'desktop') {
      navOpen = !navOpen;
      return;
    }

    if (deviceType === 'mobile') {
      mobileNavOpen = !mobileNavOpen;
      return;
    }
  }

  function handleResize() {
    ThemeStore.handleResize(window.innerWidth);
  }

  function initializeTheme() {
    const defaultTheme = typeof window !== 'undefined' && window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
    const savedTheme = getLocalStorage(VAXITAS_LS_THEME_KEY, defaultTheme) as 'light' | 'dark';

    if (savedTheme === 'dark' || savedTheme === 'light') {
      ThemeStore.setMode(savedTheme);
    } else {
      const systemPrefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
      const initialTheme = systemPrefersDark ? 'dark' : 'light';
      ThemeStore.setMode(initialTheme);
      localStorage.setItem(VAXITAS_LS_THEME_KEY, initialTheme);
    }
  }

  function scrollToHash() {
    if (typeof window === 'undefined') return;
    const hash = window.location.hash;
    if (!hash) return;
    const element = document.getElementById(hash.substring(1));
    if (!element) return;
    element.scrollIntoView({
      behavior: 'auto',
      block: 'start'
    });
  }

  ThemeStore.getDevice.subscribe((value) => {
    deviceType = value;
  });

  ThemeStore.getMode.subscribe((value) => {
    isDarkMode = value === 'dark';
    setLocalStorage(VAXITAS_LS_THEME_KEY, value);
  });

  onMount(() => {
    handleResize();
    initializeTheme();
    scrollToHash();
  });
</script>

<svelte:window
  on:resize={handleResize}
></svelte:window>

<div class="flex min-h-screen w-screen bg-vaxitas-primary">
  <Navbar
    isDrawerOpen={navOpen}
    isMobileDrawerOpen={mobileNavOpen}
    deviceType={deviceType}
    on:closeDrawer={closeNavbar}
    on:toggleDrawer={toggleNavbar}
    primaryNavData={primaryNavData}
    secondaryNavData={secondaryNavData}
    disableDrawerAnim={disableDrawerAnim}
    url={url}
  />
  <div class="w-full lg:h-screen overflow-y-auto custom-scrollbar">
    <div class="w-full max-w-[1200px] px-4 lg:px-8 py-4 mx-auto">
      <header>
        <div class="flex items-center justify-between lg:justify-end">
          {#if !deviceType || deviceType === 'mobile'}
            <button
              on:click={openNavbar}
              aria-label="open-mobile-nav-drawer"
              aria-expanded={mobileNavOpen}
            >
              <ThreeBars className="h-5 w-5 stroke-vaxitas-secondary stroke-2" />
            </button>
            <a
              href="/"
              class="text-xl font-medium lg:hidden text-vaxitas-secondary no-underline"
            >
              V
            </a>
          {/if}
          <ThemeToggle
            className="w-5 h-5 block cursor-pointer"
            isDarkMode={isDarkMode}
            on:toggle={() => {
              ThemeStore.toggleMode();
              const root = document.documentElement;
              root.classList.toggle('dark');
            }}
          />
        </div>
        <h1>{heading}</h1>
        <div class="my-4">
          {#if prevLink}
            <p>Previous entry: <a href={prevLink?.path}>{prevLink?.name}</a></p>
          {/if}
          {#if nextLink}
            <p>Next entry: <a href={nextLink?.path}>{nextLink?.name}</a></p>
          {/if}
          {#if url.startsWith("/blog") && url !== "/blog"}
            <a href="/blog">← Blog</a>
          {/if}
          {#if url.startsWith("/snippets") && url !== "/snippets"}
            <a href="/snippets">← Snippets</a>
          {/if}
        </div>
      </header>
      <slot />
    </div>
  </div>
</div>
