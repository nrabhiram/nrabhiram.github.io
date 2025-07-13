<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import DesktopNav from "./DesktopNav/DesktopNav.svelte";
  import MobileNav from "./MobileNav.svelte";
  import type { Link } from "../../types";
  import { getLocalStorage, isBasePath, setLocalStorage, VAXITAS_LS_DRAWER_KEY } from "../../utils";

  export let isDrawerOpen: boolean;
  export let disableDrawerAnim: boolean;
  export let primaryNavData: Link[];
  export let secondaryNavData: Link | null = null;
  export let url: string = "";
  export let isSecondaryNavOpen: boolean = 
    !!secondaryNavData && 
    !!secondaryNavData.items && 
    !!secondaryNavData.items.length &&
    !isBasePath(url);
  export let deviceType: string = '';

  const dispatch = createEventDispatcher();

  function closeDrawer() {
    dispatch('closeDrawer');
  }
  
  function toggleDrawer() {
    dispatch('toggleDrawer');
  }

  function openSecondaryDrawer() {
    isSecondaryNavOpen = true;
    // setLocalStorage(VAXITAS_LS_DRAWER_KEY, "true");
  }
  
  function closeSecondaryDrawer() {
    isSecondaryNavOpen = false;
    // setLocalStorage(VAXITAS_LS_DRAWER_KEY, "false");
  }
</script>

{#if !deviceType || deviceType === 'mobile'}
  <MobileNav 
    isDrawerOpen={isDrawerOpen}
    on:closeDrawer={closeDrawer}
    on:toggleDrawer={toggleDrawer}
  />
{/if}
{#if !deviceType || deviceType === 'desktop'}
  <DesktopNav
    isDrawerOpen={isDrawerOpen}
    secondaryNavOpen={isSecondaryNavOpen}
    on:closeDrawer={closeDrawer}
    on:toggleDrawer={toggleDrawer}
    on:openSecondaryDrawer={openSecondaryDrawer}
    on:closeSecondaryDrawer={closeSecondaryDrawer} 
    primaryNavData={primaryNavData} 
    secondaryNavData={secondaryNavData} 
    disableDrawerAnim={disableDrawerAnim}
    url={url}
  />
{/if}
