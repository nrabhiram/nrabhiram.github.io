<script lang="ts">
  import { onMount } from "svelte";
  import { ThemeStore } from "../stores/theme";
  import { cn, getLocalStorage, setLocalStorage, VAXITAS_LS_DRAWER_KEY, VAXITAS_LS_THEME_KEY } from "../utils";
  import { Navbar } from "./Navbar";
  import ThemeToggle from "./ThemeToggle.svelte";
  import type { AdjacentLink, Link } from "../types";

  const defaultTheme = typeof window !== 'undefined' && window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
  export let primaryNavData: Link[];
  export let secondaryNavData: Link | null = null;
  export let url: string = "";
  export let heading = "";  
  export let prevLink: AdjacentLink | null = null;
  export let nextLink: AdjacentLink | null = null;
  let deviceType = '';
  // let navOpen = getLocalStorage(VAXITAS_LS_DRAWER_KEY, "false") as "false" | "true" === "true" ? true : false;
  let navOpen = true;
  let isDarkMode = getLocalStorage<'light' | 'dark'>(VAXITAS_LS_THEME_KEY, defaultTheme) === 'dark';
  let disableDrawerAnim = false;

  function openNavbar() {
    navOpen = true;
    // setLocalStorage(VAXITAS_LS_DRAWER_KEY, 'true');
    // disableDrawerAnim = false;
  }

  function closeNavbar() {
    navOpen = false;
    // setLocalStorage(VAXITAS_LS_DRAWER_KEY, 'false');
    // disableDrawerAnim = false;
  }

  function toggleNavbar() {
    navOpen = !navOpen;
    // setLocalStorage(VAXITAS_LS_DRAWER_KEY, String(navOpen));
    // disableDrawerAnim = false;
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

  // function initializeNavDrawerState() {
  //   const prevOpen = getLocalStorage(VAXITAS_LS_DRAWER_KEY, "false") as "true" | "false";

  //   console.log("DEBUG PREV OPEN", prevOpen, navOpen);
  
  //   if (prevOpen === 'true' || prevOpen === 'false') {
  //     console.log("LS EXISTS", prevOpen === 'true');
  //     if (!navOpen) {
  //       navOpen = prevOpen === 'true';
  //     }
  //   } else {
  //     navOpen = false;
  //     console.log("LS DOESNT EXIST");
  //     setLocalStorage(VAXITAS_LS_DRAWER_KEY, 'false');
  //   }

  //   // disableDrawerAnim = false;
  // }

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
    // initializeNavDrawerState();
  });

  console.log("DEBUG LINKS", prevLink, nextLink);
</script>

<svelte:window 
  on:resize={handleResize} 
></svelte:window>

<div class="flex min-h-screen w-screen bg-vaxitas-primary">
  <Navbar
    isDrawerOpen={navOpen}
    deviceType={deviceType}
    on:closeDrawer={closeNavbar}
    on:toggleDrawer={toggleNavbar}
    primaryNavData={primaryNavData} 
    secondaryNavData={secondaryNavData}
    disableDrawerAnim={disableDrawerAnim}
    url={url}
  />
  <div class="w-full h-screen overflow-y-auto custom-scrollbar">
    <div class="w-full max-w-[1200px] px-8 py-4 mx-auto">
      <header>
        <div class="flex items-center justify-between lg:justify-end">
          {#if !deviceType || deviceType === 'mobile'}
            <button 
              class="text-xl font-medium lg:hidden text-vaxitas-secondary" 
              on:click={openNavbar}
              aria-label="open-nav-menu"
              aria-expanded={navOpen}
            >
              V
            </button>
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
        </div>
      </header>
      <slot />
    </div>
  </div>
</div>
