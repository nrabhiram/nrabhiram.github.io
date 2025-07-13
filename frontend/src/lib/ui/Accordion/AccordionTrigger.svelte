<script lang="ts">
  import { cn } from "../../../utils";

  export let className: string = '';
  export let open: boolean = false;
  export let toggle: Function = () => {};
</script>

<button 
  class={cn("flex items-center justify-between group", className)}
  on:click={() => toggle()}
>
  <slot></slot>
</button>