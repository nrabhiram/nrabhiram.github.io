<script lang="ts">
  import type { Reply } from "../types";
    import SingleComment from "./icons/SingleComment.svelte";

  export let reply: Reply;
</script>

<div class="w-full">
  <a class="flex items-center gap-2 group no-underline" href={`https://farcaster.xyz/${reply.username}`} target="_blank">
    <img class="w-7 h-7 rounded-md m-0" src={reply.pfp} alt={`${reply.displayName}'s profile picture`} />
    <span class="text-sm no-underline group-hover:underline">{reply.username}</span>
  </a>
  <div class="ml-3.5 border-l-2 border-vaxitas-tertiary border-opacity-75 pl-[22px]">
    <p class="text-sm mt-0">
      {reply.text}
      <a 
        href={`https://farcaster.xyz/${reply.username}/${reply.hash}`}
        target="_blank"
      >
        <SingleComment className="w-4 h-4 stroke-vaxitas-secondary stroke-2 inline" />
      </a>
    </p>
    {#each reply.replies as nestedReply (nestedReply.hash)}
      <svelte:self reply={nestedReply} />
    {/each}
  </div>
</div>