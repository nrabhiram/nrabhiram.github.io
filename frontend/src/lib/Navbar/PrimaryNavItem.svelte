<script lang="ts">
  import type { Link } from "../../types";
  import { cn } from "../../utils";

  export let primaryNavItem: Link;
  export let selected: boolean;
</script>

<a 
  class={cn(
    "p-2 w-full no-underline lg:border-2 border-transparent rounded-md hover:text-vaxitas-secondary",
    selected && "text-vaxitas-secondary lg:border-vaxitas-secondary"
  )} 
  href={primaryNavItem.path}
  on:click
>
  {primaryNavItem.name}
</a>
