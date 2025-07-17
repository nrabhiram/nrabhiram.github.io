<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import DesktopNav from "./DesktopNav.svelte";
  import MobileNav from "./MobileNav.svelte";
  import type { Link } from "../../types";
  import { isBasePath } from "../../utils";

  export let isDrawerOpen: boolean;
  export let isMobileDrawerOpen: boolean;
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
  }
  
  function closeSecondaryDrawer() {
    isSecondaryNavOpen = false;
  }
</script>

{#if !deviceType || deviceType === 'mobile'}
  <MobileNav 
    isDrawerOpen={isMobileDrawerOpen}
    primaryNavData={primaryNavData} 
    secondaryNavData={secondaryNavData} 
    url={url}
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
