<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import { cn, isPrimaryNavLinkActive } from '../../utils';
  import { ChevronDoubleLeft, ChevronDoubleRight, Cross } from '../icons';
  import { DrawerContainer, Drawer } from '../ui/Drawer';
  import type { Link } from "../../types";
  import PrimaryNavItem from "./PrimaryNavItem.svelte";
  import SecondaryNavItem from "./SecondaryNavItem.svelte";
  
  export let isDrawerOpen: boolean;
  export let disableDrawerAnim: boolean;
  export let secondaryNavOpen: boolean = false;
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
  
  function openSecondaryDrawer() {
    dispatch('openSecondaryDrawer');
  }
  
  function closeSecondaryDrawer() {
    dispatch('closeSecondaryDrawer');
  }
</script>

<div class="lg:flex hidden">
  <DrawerContainer
    includeOverlay={false}
    open={isDrawerOpen}
    on:closeDrawer={closeDrawer}
    on:toggleDrawer={toggleDrawer}
    direction="left"
    disableAnim={disableDrawerAnim}
    navLabel="primary-navigation"
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
          "relative lg:block hidden h-screen w-0",
          isDrawerOpen && "lg:w-56 2xl:w-72 border-r-2 border-vaxitas-tertiary"
        )}
        drawerId="desktop-nav"
      >
        <svelte:fragment>
          <div class="w-full h-full p-2 flex flex-col gap-2">
            {#each primaryNavData as primaryNavItem (primaryNavItem.path)}
              <PrimaryNavItem
                primaryNavItem={primaryNavItem}
                selected={isPrimaryNavLinkActive(url, primaryNavItem.path)}
                on:click={(e) => {
                  if (url !== primaryNavItem.path) return;
                  e.preventDefault();
                  openSecondaryDrawer();
                }}
              />
            {/each}
          </div>
        </svelte:fragment>
        <svelte:fragment slot="toggle" let:open={open} let:toggleDrawer={toggleDrawer}>
          <button 
            class={cn(
              "absolute top-6 translate-x-full translate-y-full border-2 border-l-0 border-vaxitas-tertiary p-1 rounded-r-md transition-all z-[1]",
              isDrawerOpen && secondaryNavOpen ? "lg:-right-80 xl:-right-96" : "right-0"
            )}
            aria-label="toggle-primary-nav-drawer"
            on:click={() => toggleDrawer()}
          >
            {#if open}
              <Cross className="w-4 h-4 stroke-vaxitas-tertiary stroke-2" />
            {:else}
              <ChevronDoubleRight className="w-4 h-4 stroke-vaxitas-tertiary stroke-2" /> 
            {/if}
          </button>
        </svelte:fragment>
      </Drawer>
    </svelte:fragment>
  </DrawerContainer>
  <nav 
    class={cn(
      "h-screen transition-all flex flex-col relative p-2 z-0",
      isDrawerOpen && secondaryNavOpen ? "lg:w-80 xl:w-96 border-r-2 border-vaxitas-tertiary" : "w-0"
    )}
    aria-label="secondary-navigation"
  >
    {#if isDrawerOpen && !secondaryNavOpen && secondaryNavData?.items && secondaryNavData.items.length}
      <button 
        class="absolute top-14 translate-x-full translate-y-full border-2 border-l-0 border-vaxitas-tertiary transition-colors p-1 rounded-r-md right-4"
        on:click={openSecondaryDrawer}
        aria-label="open-secondary-nav-drawer"
      >
        <ChevronDoubleRight className="w-4 h-4 stroke-vaxitas-tertiary stroke-2" /> 
      </button>
    {/if}
    {#if isDrawerOpen && secondaryNavOpen && secondaryNavData?.items && secondaryNavData.items.length}
      <button 
        class="absolute top-14 translate-x-full translate-y-full border-2 border-l-0 border-vaxitas-tertiary transition-colors p-1 rounded-r-md -right-0.5"
        on:click={closeSecondaryDrawer}
        aria-label="close-secondary-nav-drawer"
      >
      <ChevronDoubleLeft className="w-4 h-4 stroke-vaxitas-tertiary stroke-2" /> 
      </button>
    {/if}
    {#if isDrawerOpen && secondaryNavOpen && secondaryNavData && secondaryNavData.items}
      {#each secondaryNavData.items as secondaryNavItem (secondaryNavItem.path)}
        <SecondaryNavItem 
          secondaryNavItem={secondaryNavItem}
          url={url}
          on:click={(e) => {
            if (url === secondaryNavItem.path) {
              e.preventDefault();
              return;
            }
          }}
        />
      {/each}
    {/if}
  </nav>
</div>
