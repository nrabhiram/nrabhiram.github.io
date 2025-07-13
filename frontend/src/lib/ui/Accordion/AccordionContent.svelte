<script lang="ts">
  import { cn } from "../../../utils";

  export let className: string = '';
  export let animDuration: number;
  export let disableAnim: boolean;
  export let open: boolean = false;
  export let contentIdAttr = "";
  let contentEle: HTMLElement;

  $: contentStyles = `
    max-height: ${open && contentEle ? contentEle.scrollHeight : 0}px;
    transition-duration: ${!disableAnim && animDuration ? animDuration : 0}ms;
  `;
</script>

<div
  bind:this={contentEle}
  class={cn("ease-in-out w-full accordion-content", className)}
  style={contentStyles}
  data-url={contentIdAttr}
  data-open={open}
>
  <slot></slot>
</div>
