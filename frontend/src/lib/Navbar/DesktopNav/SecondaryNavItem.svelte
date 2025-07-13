<script lang="ts">
  import type { Link } from "../../../types";
  import { cn, isBrowser, isDoublyNested, isNested } from "../../../utils";
  import ChevronRight from "../../icons/ChevronRight.svelte";
  import Accordion from "../../ui/Accordion/Accordion.svelte";
  import AccordionContent from "../../ui/Accordion/AccordionContent.svelte";

  export let secondaryNavItem: Link;
  export let url: string;
  export let level = 1;
  let accordionOpen = shouldBeOpenInitially();
  let interacted = false;

  function shouldBeOpenInitially() {
    return url.startsWith(secondaryNavItem.path + '/') || 
           secondaryNavItem.items?.some(item => url.startsWith(item.path + '/')) ||
           hasNestedActiveItem(secondaryNavItem.items);
  }
  
  // Recursively check if any nested item should be active
  function hasNestedActiveItem(items?: Link[]): boolean {
    if (!items) return false;
    return items.some(item => 
      url.startsWith(item.path + '/') || 
      url === item.path ||
      hasNestedActiveItem(item.items)
    );
  }
</script>

{#if secondaryNavItem.items?.length}
  <Accordion 
    open={accordionOpen}
    className={cn(
      "w-full",
      level === 1 ? "p-2" : "pl-4 pr-2 py-1",
      level > 1 && "border-l-2 border-vaxitas-pale",
      level > 1 && url === secondaryNavItem.path && "border-vaxitas-secondary",
      level > 1 && url !== secondaryNavItem.path && "hover:border-vaxitas-tertiary"
    )}
    animDuration={200}
    on:toggleAccordion={() => {
      accordionOpen = !accordionOpen;

      if (!interacted) interacted = true;

      const currentAccordion = document.querySelector(`[data-url="${secondaryNavItem.path}"]`);
      if (!currentAccordion) return;

      currentAccordion.classList.add("overflow-y-hidden");
      
      const allAccordionContents = document.querySelectorAll('.accordion-content');

      // Find which ones contain the current accordion
      const parentAccordions = Array.from(allAccordionContents).filter(accordion => 
        accordion.contains(currentAccordion) && accordion !== currentAccordion
      );

      console.log(accordionOpen, parentAccordions);

      parentAccordions.forEach(parent => {
        if (accordionOpen) {
          parent.classList.remove('overflow-y-hidden');
        } else {
          const openSiblings = parent.querySelectorAll('.accordion-content[data-open="true"]');
          if (openSiblings.length === 0) {
            parent.classList.add('overflow-y-hidden');
          }
        }
      });
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
            "no-underline text-vaxitas-pale grow",
            url === secondaryNavItem.path && "text-vaxitas-secondary",
            url !== secondaryNavItem.path && "group-hover:text-vaxitas-tertiary",
            level === 2 && "text-sm",
            level > 2 && "text-xs"
          )}
          href={secondaryNavItem.path}
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
        className={accordionOpen ? (interacted ? "overflow-y-hidden" : "") : "overflow-y-hidden"}
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
      "w-full no-underline text-vaxitas-pale shrink-0 cursor-pointer",
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
