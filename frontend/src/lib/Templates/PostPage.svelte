<script lang="ts">
  import type { Artifact } from "../../types";

  export let artifact: Artifact | null = null;
</script>

<main>
  <article>{@html artifact?.content}</article>
</main>
