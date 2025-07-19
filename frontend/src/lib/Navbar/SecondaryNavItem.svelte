<script lang="ts">
  import type { Link } from "../../types";
  import { cn, isBrowser } from "../../utils";
  import ChevronRight from "../icons/ChevronRight.svelte";
  import Accordion from "../ui/Accordion/Accordion.svelte";
  import AccordionContent from "../ui/Accordion/AccordionContent.svelte";

  export let secondaryNavItem: Link;
  export let url: string;
  export let level = 1;
  let accordionOpen = shouldBeOpenInitially();
  let interacted = false;
  let initialized = false;

  function shouldBeOpenInitially() {
    const shouldBeOpen = url.startsWith(secondaryNavItem.path + '/') || 
      hasNestedActiveItem(secondaryNavItem.items);
    return shouldBeOpen;
  }
  
  // recursively check if any nested item should be active
  function hasNestedActiveItem(items?: Link[]): boolean {
    if (!items) return false;
    return items.some(item => 
      url.startsWith(item.path + '/') || 
      url === item.path ||
      hasNestedActiveItem(item.items)
    );
  }

  function computeAccordionsMaxHeights(url: string) {
    if (!isBrowser) return;
    
    const currentAccordion = document.querySelector(`[data-url="${url}"]`);
    if (!currentAccordion) return;
    let contentHeight = currentAccordion.scrollHeight;
    let currentParent = currentAccordion.closest('.accordion-content');

    while (!!currentParent) {       
      // found a parent accordion, update its height
      if (accordionOpen) {
        contentHeight += currentParent.scrollHeight;
        (currentParent as HTMLElement).style.maxHeight = contentHeight + 'px';
      } else {
        const openChildren = currentParent.querySelectorAll('.accordion-content[data-open="true"]');
        if (openChildren.length === 0) {
          (currentParent as HTMLElement).style.maxHeight = '0px';
        }
      }
      
      // move up to the next parent accordion
      const nextParent = currentParent.parentElement;
      currentParent = nextParent ? nextParent.closest('.accordion-content') : null;
    }
  }
</script>

{#if secondaryNavItem.items?.length}
  <Accordion 
    open={accordionOpen}
    className={cn(
      "w-full",
      level === 1 ? "p-2" : "pl-4 pr-0 py-1",
      level > 1 && "border-l-2 border-vaxitas-pale",
      level > 1 && url === secondaryNavItem.path && "border-vaxitas-secondary",
      level > 1 && url !== secondaryNavItem.path && "hover:border-vaxitas-tertiary"
    )}
    animDuration={200}
    on:toggleAccordion={() => {
      accordionOpen = !accordionOpen;
      if (!interacted) interacted = true;
      computeAccordionsMaxHeights(secondaryNavItem.path);
    }}
  >
    <svelte:fragment 
      slot="trigger"
      let:open
      let:toggle
    >
      <div class="w-full shrink-0 cursor-pointer flex items-center justify-between group">
        <a 
          class={cn(
            "no-underline text-vaxitas-pale grow block",
            url === secondaryNavItem.path && "text-vaxitas-secondary",
            url !== secondaryNavItem.path && "group-hover:text-vaxitas-tertiary",
            level === 2 && "text-sm",
            level > 2 && "text-xs"
          )}
          href={secondaryNavItem.path}
          on:click
        >
          {secondaryNavItem.name}
        </a>
        <button on:click={toggle}>
          <ChevronRight 
            className={cn(
              "w-4 h-4 stroke-vaxitas-pale stroke-2 ease-in-out duration-200",
              url === secondaryNavItem.path && "stroke-vaxitas-secondary",
              url !== secondaryNavItem.path && "group-hover:stroke-vaxitas-tertiary",
              open && "rotate-90"
            )}
          />
        </button>
      </div>
    </svelte:fragment>
    <svelte:fragment
      slot="content"
      let:open
      let:animDuration
      let:disableAnim
    >
      <AccordionContent 
        open={open} 
        animDuration={animDuration} 
        disableAnim={disableAnim}
        contentIdAttr={secondaryNavItem.path}
        className={cn(
          "overflow-x-clip",
          accordionOpen ? (interacted ? "overflow-y-hidden" : "") : "overflow-y-hidden"
        )}
      >
        {#each secondaryNavItem.items as subNavItem}
          <svelte:self 
            secondaryNavItem={subNavItem}
            url={url}
            level={level + 1}
          />
        {/each}
      </AccordionContent>
    </svelte:fragment>
  </Accordion>
{:else}
  <a 
    class={cn(
      "w-full no-underline text-vaxitas-pale shrink-0 cursor-pointer block",
      url === secondaryNavItem.path && "text-vaxitas-secondary border-vaxitas-secondary",
      url !== secondaryNavItem.path && "hover:text-vaxitas-tertiary hover:border-vaxitas-tertiary",
      level === 1 ? "p-2" : "pl-4 pr-2 py-1",
      level === 2 && "text-sm",
      level > 2 && "text-xs",
      level > 1 && "border-l-2 border-vaxitas-pale",
      level > 1 && url === secondaryNavItem.path && "border-vaxitas-secondary",
      level > 1 && url !== secondaryNavItem.path && "hover:border-vaxitas-tertiary"
    )} 
    href={secondaryNavItem.path}
    on:click
  >
    {secondaryNavItem.name}
  </a>
{/if}
