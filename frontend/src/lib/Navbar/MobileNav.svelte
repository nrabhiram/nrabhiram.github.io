<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import { cn } from '../../utils';
  import {  DrawerContainer, Drawer, DrawerOverlay } from '../ui/Drawer';

  export let isDrawerOpen: boolean = true;

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
        "relative h-screen w-4/5 bg-vaxitas-primary",
      )}
    >
      <svelte:fragment>
        <!-- <div class="w-full h-full p-2">
          <button class="p-2 rounded-md bg-teal-500 text-white w-full" on:click={() => secondaryNavOpen = !secondaryNavOpen}>Toggle</button>
        </div> -->
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
