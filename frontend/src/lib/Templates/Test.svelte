<script lang="ts">
  import type { Artifact } from "../../types";

  export let artifact: Artifact | null = null;
</script>

<article>
  <h1>{artifact?.metadata.title || artifact?.name}</h1>
  {@html artifact?.content}
</article>
