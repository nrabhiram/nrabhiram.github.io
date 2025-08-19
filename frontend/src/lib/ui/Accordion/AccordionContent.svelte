<script lang="ts">
  import { onDestroy, onMount } from "svelte";
  import { cn, isBrowser } from "../../../utils";

  export let className: string = '';
  export let animDuration: number;
  export let disableAnim: boolean;
  export let open: boolean = false;
  export let contentIdAttr = "";
  let contentEle: HTMLElement;
  let styles = "";

  function calculateStyles(disableAnim: boolean, open: boolean, animDuration: number, contentEle: HTMLElement) {
    const styles = `
      ${!disableAnim ? 
        `max-height: ${open && contentEle ? contentEle.scrollHeight : 0}px;` : 
        `${!open ? 'max-height: 0px;' : `${contentEle ? `max-height: ${contentEle.scrollHeight}px;` : ''}`}`
      }
      transition-duration: ${!disableAnim && animDuration ? animDuration : 0}ms;
    `;
    return styles;
  }

  function handleResize() {
    styles = calculateStyles(disableAnim, open, animDuration, contentEle);
  }

  $: {
    styles = calculateStyles(disableAnim, open, animDuration, contentEle);
  }

  onMount(() => {
    window.addEventListener("resize", handleResize);
  });

  onDestroy(() => {
    if (!isBrowser) return;
    window.removeEventListener("resize", handleResize);
  })
</script>

<div
  bind:this={contentEle}
  class={cn("ease-in-out w-full accordion-content", className)}
  style={styles}
  data-url={contentIdAttr}
  data-open={open}
>
  <slot></slot>
</div>
