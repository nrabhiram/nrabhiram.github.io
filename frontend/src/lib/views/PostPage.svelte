<script lang="ts">
  import { onDestroy, onMount } from "svelte";
  import { ThemeStore } from "../../stores/theme";
  import type { Artifact, HeadingNode, Reactions, Reply } from "../../types";
  import Comment from "../Comment.svelte";
  import Recast from "../icons/Recast.svelte";
  import Like from "../icons/Like.svelte";
  import Comments from "../icons/Comments.svelte";
  import SingleComment from "../icons/SingleComment.svelte";
  import Accordion from "../ui/Accordion/Accordion.svelte";
  import ChevronRight from "../icons/ChevronRight.svelte";
  import { cn } from "../../utils";
  import AccordionContent from "../ui/Accordion/AccordionContent.svelte";

  export let artifact: Artifact | null = null;
  export let tableOfContents: HeadingNode[] = [];
  let replies: Reply[] = [];
  let reactions: Reactions = {};
  let isDarkMode = false;
  let loading = true;
  let tocOpen = true;
  let enableTocAnim = false;
  let headingId = '';
  let observer: IntersectionObserver | null = null;

  function extractRepliesFromCast(cast: any): Reply[] {
    const allReplies: any[] = [];
    if (cast.replies && cast.replies.casts) {
      cast.replies.casts.forEach((reply: any) => {
        allReplies.push({
          hash: reply.hash,
          text: reply.text,
          username: reply.author.username || reply.author.displayName,
          displayName: reply.author.displayName,
          pfp: reply.author.pfp?.url || null,
          formattedDate: new Date(reply.timestamp).toLocaleString(),
          replies: extractRepliesFromCast(reply)
        });
      });
    }
    return allReplies;
  }

  function extractComments(response: any): {
    allReplies: Reply[];
    reactions: Reactions;
  } {
    const casts = response.result.casts;
    const mainCast = casts.find((cast: any) => !cast.parentHash && cast.castType !== "root-embed");

    if (!mainCast) return { allReplies: [], reactions: {} };

    const reactions = {
      likes: mainCast.reactions.count,
      recasts: mainCast.recasts.count,
      comments: mainCast.replies.count,
    }

    const threadHash = mainCast.threadHash;
    const rootReplies = casts.filter((cast: any) => cast.parentHash === threadHash);
    let allReplies: any[] = [];

    for (let rootReply of rootReplies) {
      const reply = {
        hash: rootReply.hash,
        text: rootReply.text,
        username: rootReply.author.username || rootReply.author.displayName,
        displayName: rootReply.author.displayName,
        pfp: rootReply.author.pfp?.url || null,
        formattedDate: new Date(rootReply.timestamp).toLocaleString(),
        replies: extractRepliesFromCast(rootReply)
      }
      allReplies.push(reply);
    }

    return { allReplies, reactions };
  }

  ThemeStore.getMode.subscribe((value) => {
    isDarkMode = value === 'dark';
  });

  function setupHeadingObserver() {
    if (typeof window === 'undefined') return;
    if (!tableOfContents.length) return;

    // Find the scroll container
    const scrollContainer = document.querySelector('.content-container');

    observer = new IntersectionObserver((entries) => {
      const containerRect = (scrollContainer as Element).getBoundingClientRect();

      // Find all headings that are currently in viewport (even partially)
      const headingsInView = tableOfContents
        .map(heading => {
          const element = document.getElementById(heading.id);
          if (!element) return null;
          const rect = element.getBoundingClientRect();

          // Check if any part of heading is visible
          const isVisible = rect.bottom > containerRect.top && rect.top < containerRect.bottom;

          return {
            id: heading.id,
            element,
            rect,
            isVisible,
            // Distance from top of container (negative if above, positive if below)
            distanceFromTop: rect.top - containerRect.top
          };
        })
        .filter(h => h !== null && h.isVisible);

      // If no headings visible at all, check if we're above all headings
      if (headingsInView.length === 0) {
        const allHeadingsBelow = tableOfContents.every(heading => {
          const element = document.getElementById(heading.id);
          if (!element) return false;
          const rect = element.getBoundingClientRect();
          return rect.top > containerRect.top;
        });

        if (allHeadingsBelow) {
          headingId = "";
        }
        return;
      }

      // Find the heading that's at or just passed the top of the viewport
      // Priority: heading closest to top, but prefer ones that have crossed the top
      const sortedHeadings = headingsInView.sort((a, b) => {
        if (a === null || b === null) return 0;
        // If one is above top and one below, prefer the one above (already scrolled past)
        if (a.distanceFromTop <= 0 && b.distanceFromTop > 0) return -1;
        if (b.distanceFromTop <= 0 && a.distanceFromTop > 0) return 1;

        // Both above or both below, pick closest to top
        return Math.abs(a.distanceFromTop) - Math.abs(b.distanceFromTop);
      });

      const activeHeading = sortedHeadings[0];
      if (activeHeading && tableOfContents.some(h => h.id === activeHeading.id)) {
        headingId = activeHeading.id;
      }
    }, {
      root: scrollContainer as Element,
      rootMargin: '0px 0px -20% 0px',
      threshold: [0, 0.25, 0.5, 0.75, 1.0]
    });

    tableOfContents.forEach(heading => {
      const element = document.getElementById(heading.id);
      if (element && observer) {
        observer.observe(element);
      }
    });
  }

  onMount(() => {
    setupHeadingObserver();
    if (!artifact || !artifact.castHash) return;
    fetch(`https://vaxitas-fc-api-proxy.netlify.app/vaxitas.eth/${artifact.castHash}`)
      .then(res => res.json())
      .then(comments => {
        const commentsData = extractComments(comments);
        replies = commentsData.allReplies;
        reactions = commentsData.reactions;
      })
      .finally(() => {
        loading = false;
      });
  });

  onDestroy(() => {
    if (!observer) return
    observer.disconnect();
  });
</script>

<main>
  <div class="relative w-full flex lg:flex-row flex-col-reverse lg:gap-4">
    <article
      class={cn(
        "w-full",
        tableOfContents.length > 0 && "lg:w-3/5"
      )}
      id="content"
    ></article>
    {#if tableOfContents.length > 0}
      <section class="border-2 border-vaxitas-secondary lg:w-2/5 w-full lg:sticky lg:right-0 lg:top-2 self-start rounded-md">
        <Accordion
          open={tocOpen}
          animDuration={200}
          on:toggleAccordion={() => {
            if (!enableTocAnim) enableTocAnim = true;
            tocOpen = !tocOpen;
          }}
          disableAnim={!enableTocAnim}
        >
          <svelte:fragment
            slot="trigger"
            let:open
            let:toggle
          >
            <button
              class="w-full flex items-center justify-between p-2 bg-vaxitas-secondary"
              on:click={toggle}
            >
              <span class="text-vaxitas-primary font-semibold">Table of Contents</span>
              <ChevronRight
                className={cn(
                  "w-4 h-4 stroke-vaxitas-primary stroke-2 ease-in-out duration-200",
                  open && "rotate-90"
                )}
              />
            </button>
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
              className="overflow-hidden"
            >
              <nav class="p-2">
                <ul class="flex flex-col gap-1 list-none ml-0">
                  {#each tableOfContents as heading (heading.id)}
                    <li>
                      <a
                        href={`#${heading.id}`}
                        class={cn(
                          "no-underline text-vaxitas-secondary text-sm font-medium duration-100",
                          headingId !== heading.id && "lg:opacity-75",
                        )}
                        style={`padding-left: ${(heading.level - 2) * 8}px`}
                      >
                        {heading.text}
                      </a>
                    </li>
                  {/each}
                </ul>
              </nav>
            </AccordionContent>
          </svelte:fragment>
        </Accordion>
      </section>
    {/if}
  </div>
  {#if artifact && artifact.castHash}
    <div class="mt-6 mb-3 flex items-end gap-5 justify-between lg:justify-start border-b-2 border-vaxitas-secondary">
      <div class="flex items-center">
        <Comments className="sm:w-9 sm:h-9 w-6 h-6 stroke-vaxitas-secondary stroke-2" />
        <span class="sm:text-4xl text-2xl font-extrabold text-vaxitas-secondary">Comments</span>
      </div>
      {#if !loading}
        <div class="flex items-end gap-3">
          <div class="flex items-center gap-1">
            <span class="text-vaxitas-secondary font-medium">{reactions.likes || 0}</span>
            <Like className="w-4 h-4 stroke-vaxitas-secondary stroke-2" />
          </div>
          <div class="flex items-center gap-1">
            <span class="text-vaxitas-secondary font-medium">{reactions.recasts || 0}</span>
            <Recast className="w-4 h-4 stroke-vaxitas-secondary stroke-2" />
          </div>
          <div class="flex items-center gap-1">
            <span class="text-vaxitas-secondary font-medium">{reactions.comments || 0}</span>
            <SingleComment className="w-4 h-4 stroke-vaxitas-secondary stroke-2" />
          </div>
        </div>
      {/if}
    </div>
    {#if !loading}
      <div class="px-4">
        {#if replies.length > 0}
          <a
            class="flex items-center gap-1 border-2 rounded-md border-vaxitas-secondary text-vaxitas-secondary px-2 py-1 ml-auto mb-2 no-underline w-max"
            href={`https://farcaster.xyz/vaxitas.eth/${artifact.castHash}`}
            target="_blank"
          >
            <img
              src={isDarkMode ? '/farcaster-logo-dark.png' : '/farcaster-logo-light.png'}
              alt="Farcaster logo"
              class="w-6 h-6 m-0"
            />
            <span class="text-vaxitas-secondary text-sm hidden lg:inline">View on Farcaster</span>
          </a>
          {#each replies as reply (reply.hash)}
            <Comment reply={reply} />
          {/each}
        {/if}
        {#if replies.length === 0}
          <div class="text-center">
            <img
              src={isDarkMode ? '/farcaster-logo-dark.png' : '/farcaster-logo-light.png'}
              alt="Farcaster logo"
              class="w-8 h-8"
            />
            <p class="text-vaxitas-secondary">
              Like, recast, or reply to this post on
              <a
                href={`https://farcaster.xyz/vaxitas.eth/${artifact.castHash}`}
                target="_blank"
                class="text-vaxitas-secondary"
              >
                Farcaster
              </a>
            </p>
          </div>
        {/if}
      </div>
    {/if}
    {#if loading}
      <div class="px-4">
        <img
          src={isDarkMode ? '/farcaster-logo-dark.png' : '/farcaster-logo-light.png'}
          alt="Farcaster logo"
          class="w-8 h-8 animate-pulse"
        />
      </div>
    {/if}
  {/if}
</main>
