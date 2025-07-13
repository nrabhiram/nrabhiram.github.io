<script lang="ts">
  import { createEventDispatcher } from "svelte";
  
  export let className: string = '';
  export let open: boolean = false;
  export let animDuration: number = 500;
  export let disableAnim: boolean = false;

  const dispatch = createEventDispatcher();

  function toggle() {
    dispatch('toggleAccordion');
  }
</script>

<div 
  class={className}
>
  <slot 
    name="trigger"
    toggle={toggle}
    open={open}
  ></slot>
  <slot
    name="content"
    open={open}
    animDuration={animDuration}
    disableAnim={disableAnim}
  ></slot>
</div>