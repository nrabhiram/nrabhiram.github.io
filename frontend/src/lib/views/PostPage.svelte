<script lang="ts">
  import { onMount } from "svelte";
  import { ThemeStore } from "../../stores/theme";
  import type { Artifact, Reactions, Reply } from "../../types";
  import Comment from "../Comment.svelte";
  import Recast from "../icons/Recast.svelte";
  import Like from "../icons/Like.svelte";
  import Comments from "../icons/Comments.svelte";
  import SingleComment from "../icons/SingleComment.svelte";

  export let artifact: Artifact | null = null;
  let replies: Reply[] = [];
  let reactions: Reactions = {};
  let isDarkMode = false;
  let loading = true;

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

  onMount(() => {
    if (!artifact || !artifact.castHash) return;
    fetch(`https://client.farcaster.xyz/v2/user-thread-casts?castHashPrefix=${artifact.castHash}&username=vaxitas.eth`)
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
</script>

<main>
  <article id="content"></article>
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
