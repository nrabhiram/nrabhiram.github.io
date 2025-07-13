<script lang="ts">
  import { cn } from "../../../utils";

  export let className: string = '';
  export let transitionStyle: string = '';
  export let open: boolean = false;
  export let transitionDuration = 0;

  let transitionTimerId: number | undefined;
  let closeOverlay: boolean = false;

  $: {
    if (!open) {
      transitionTimerId = setTimeout(() => {
        closeOverlay = true;
      }, transitionDuration);
    } else {
      closeOverlay = false;
      clearTimeout(transitionTimerId);
    }
  }
</script>

<div 
  class={cn(
    "h-screen transition-opacity bg-black/70 fixed backdrop-blur-sm",
    open ? "opacity-100 w-screen" : "opacity-0",
    closeOverlay ? "w-0" : !open ? "w-screen pointer-events-none" : "",
    className
  )}
  style={transitionStyle}
></div>
