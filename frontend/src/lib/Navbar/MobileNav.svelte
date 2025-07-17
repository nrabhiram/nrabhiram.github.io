<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import { cn, isPrimaryNavLinkActive } from '../../utils';
  import {  DrawerContainer, Drawer, DrawerOverlay } from '../ui/Drawer';
  import type { Link } from "../../types";
  import SecondaryNavItem from "./SecondaryNavItem.svelte";
    import PrimaryNavItem from "./PrimaryNavItem.svelte";

  export let isDrawerOpen: boolean = true;
  export let primaryNavData: Link[];
  export let secondaryNavData: Link | null = null;
  export let url: string;

  const dispatch = createEventDispatcher();

  function toggleDrawer() {
    dispatch('toggleDrawer');
  }
  
  function closeDrawer() {
    dispatch('closeDrawer');
  }
</script>

<DrawerContainer
  includeOverlay={true}
  open={isDrawerOpen}
  on:closeDrawer={closeDrawer}
  on:toggleDrawer={toggleDrawer}
  direction="left"
  disableAnim={false}
  className="block lg:hidden z-[1]"
  transitionDuration={200}
  >
  <svelte:fragment 
    slot="drawer"
    let:canLoadContent={canLoadContent}
    let:includeOverlay={includeOverlay}
    let:direction={direction}
    let:transitionDuration={transitionDuration}
    let:open={open}
    let:transitionStyle={transitionStyle}
    let:closeDrawer={closeDrawer}
    let:toggleDrawer={toggleDrawer}
    let:disableAnim={disableAnim}
  >
    <Drawer 
      canLoadContent={canLoadContent}
      includeOverlay={includeOverlay}
      direction={direction}
      transitionDuration={transitionDuration}
      open={open}
      transitionStyle={transitionStyle}
      closeDrawer={closeDrawer}
      toggleDrawer={toggleDrawer}
      disableAnim={disableAnim}
      className={cn(
        "relative h-screen w-4/5 bg-vaxitas-primary flex flex-col gap-2 p-2",
      )}
    >
      <svelte:fragment>
        {#each primaryNavData as primaryNavItem}
          {#if secondaryNavData && secondaryNavData.path === primaryNavItem.path}
            <SecondaryNavItem 
              secondaryNavItem={secondaryNavData}
              url={url}
              on:click={(e) => {
                if (url === secondaryNavData.path) {
                  e.preventDefault();
                  return;
                }
              }}
            />
          {:else}
            <PrimaryNavItem 
              primaryNavItem={primaryNavItem}
              selected={isPrimaryNavLinkActive(url, primaryNavItem.path)}
            />
          {/if}
        {/each}
      </svelte:fragment>
    </Drawer>
  </svelte:fragment>
  <svelte:fragment 
    slot="overlay"
    let:open={open}
    let:transitionStyle={transitionStyle}
    let:transitionDuration={transitionDuration}
  >
    <DrawerOverlay 
      open={open}
      transitionStyle={transitionStyle}
      transitionDuration={transitionDuration}   
    />
  </svelte:fragment>
</DrawerContainer>
