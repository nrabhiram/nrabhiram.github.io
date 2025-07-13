<script lang="ts">
  import type { Artifact } from "../../types";

  export let artifact: Artifact | null = null;

  console.log("DEBUG 1234", artifact?.next, artifact?.prev);
</script>

<main>
  <article>{@html artifact?.content}</article>
</main>
