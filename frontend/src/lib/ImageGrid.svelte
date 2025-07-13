
<script lang="ts">
  export let images: string[] = [];
</script>

<div class="columns-3">
  {#each images as src}
    <div class="my-4">
      <img {src} alt="" loading="lazy" />
    </div>
  {/each}
</div>