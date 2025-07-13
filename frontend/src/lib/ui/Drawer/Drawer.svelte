<script lang="ts">
  import { cn } from "../../../utils";
  import type { Direction } from "./types";

  export let className: string = '';
  export let transitionStyle: string = '';
  export let canLoadContent: boolean;
  export let includeOverlay: boolean;
  export let direction: Direction;
  export let transitionDuration: number;
  export let open: boolean;
  export let disableAnim: boolean;
  export let closeDrawer: Function;
  export let toggleDrawer: Function;
  export let drawerId: string = "";
  
  let drawerEle: HTMLDivElement;
  let outroTimerId: number | undefined;

  function stopPropagation(e: MouseEvent) {
    e.stopPropagation();
  }

  $: {
    if (includeOverlay) {
      if (open) {
        if (!disableAnim) {
          clearTimeout(outroTimerId);
        }
        if (drawerEle) {
          if (direction === 'left' || direction === 'right') {
            drawerEle.classList.remove("h-0");
            drawerEle.classList.add("h-screen");
          } else {
            drawerEle.classList.remove("w-0");
            drawerEle.classList.add("w-screen");
          }
        }
      } else {
        if (!disableAnim) {
          outroTimerId = setTimeout(() => {
            if (!drawerEle) return;
            if (direction === 'left' || direction === 'right') {
              drawerEle.classList.add("h-0");
              drawerEle.classList.remove("h-screen");
            } else {
              drawerEle.classList.add("w-0");
              drawerEle.classList.remove("w-screen");
            }
          }, transitionDuration);
        } else {
          if (drawerEle) {
            if (direction === 'left' || direction === 'right') {
              drawerEle.classList.add("h-0");
              drawerEle.classList.remove("h-screen");
            } else {
              drawerEle.classList.add("w-0");
              drawerEle.classList.remove("w-screen");
            }
          }
        }
      }
    }
  }
</script>

<!-- svelte-ignore a11y-click-events-have-key-events a11y-no-noninteractive-element-interactions a11y-no-noninteractive-tabindex a11y-no-static-element-interactions -->
<div 
  class={cn(
    !includeOverlay && className,
    includeOverlay && "relative h-screen w-screen",
    !includeOverlay && "h-full",
    !includeOverlay && !open && "w-0 overflow-hidden",
    !includeOverlay && !disableAnim && "transition-all ",
    !includeOverlay && className
  )}
  style={!includeOverlay ? transitionStyle : ''}
  on:click={() => {
    if (!includeOverlay) return;
    closeDrawer();
  }}
  id={drawerId}
  bind:this={drawerEle}
>
  {#if !includeOverlay}
    <slot canLoadContent={canLoadContent}></slot>
  {/if}
  {#if includeOverlay}
    <div 
      class={cn(
        "transition-all bg-white absolute",
        !open && direction === 'left' && "-translate-x-full",
        !open && direction === 'right' && "translate-x-full",
        !open && direction === 'top' && "-translate-y-full",
        !open && direction === 'bottom' && "translate-y-full",
        direction === 'left' && 'top-0 left-0 h-full',
        direction === 'right' && 'top-0 right-0 h-full',
        direction === 'top' && "top-0 left-0 w-full",
        direction === 'bottom' && "bottom-0 left-0 w-full",
        className,
      )}
      style={transitionStyle}
      on:click={stopPropagation}
    >
      <!-- Toggle slot that exposes the open state -->
      <slot 
        name="toggle" 
        open={open}
        closeDrawer={closeDrawer}
      ></slot>
      <slot canLoadContent={canLoadContent}></slot>
    </div>
  {/if}
</div>
{#if !includeOverlay}
  <slot 
    name="toggle" 
    open={open}
    closeDrawer={closeDrawer}
    toggleDrawer={toggleDrawer}
  ></slot>
{/if}
