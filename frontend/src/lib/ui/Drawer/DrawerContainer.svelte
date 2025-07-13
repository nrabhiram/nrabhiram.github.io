<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import { cn } from "../../../utils";
  import type { Direction } from "./types";

  export let className: string = '';
  export let includeOverlay: boolean = false;
  export let transitionDuration: number = 200;
  export let direction: Direction = 'left';
  export let open: boolean = false;
  export let disableAnim: boolean = false;
  export let navLabel: string = "";
  let canLoadContent: boolean = false;
  let containerEle: HTMLElement;
  let contentLoadTimerId: number | undefined;
  let outroTimerId: number | undefined;

  $: transitionStyle = disableAnim ? '' : `transition-duration: ${transitionDuration}ms`;

  $: {
    if (open) {
      contentLoadTimerId = setTimeout(() => {
        canLoadContent = true;
      }, transitionDuration);
    } else {
      clearTimeout(contentLoadTimerId);
      canLoadContent = false;
    }
  }

  $: {
    if (includeOverlay) {
      if (open) {
        if (!disableAnim) {
          clearTimeout(outroTimerId);
        }
        if (containerEle) {
          containerEle.classList.remove("w-0");
        }
      } else {
        if (!disableAnim) {
          outroTimerId = setTimeout(() => {
            if (!containerEle) return;
            containerEle.classList.add("w-0");
          }, transitionDuration);
        } else {
          if (containerEle) {
            containerEle.classList.add("w-0");
          }
        }
      }
    }
  }

  const dispatch = createEventDispatcher();

  function closeDrawer() {
    dispatch('closeDrawer');
  }

  function toggleDrawer() {
    dispatch('toggleDrawer');
  }
</script>

<nav
  class={cn(
    !includeOverlay && "relative w-max",
    includeOverlay && "absolute top-0 left-0 z-2",
    className
  )}
  aria-label={navLabel}
  bind:this={containerEle}
>
  {#if includeOverlay}
    <slot 
      name="overlay"
      open={open}
      transitionStyle={transitionStyle}
      disableAnim={disableAnim}
      transitionDuration={transitionDuration}
    ></slot>
  {/if}
  <slot
    name="drawer"
    canLoadContent={canLoadContent}
    includeOverlay={includeOverlay}
    direction={direction}
    transitionDuration={transitionDuration}
    open={open}
    transitionStyle={''}
    closeDrawer={closeDrawer}
    toggleDrawer={toggleDrawer}
    disableAnim={disableAnim}
  ></slot>
</nav>
