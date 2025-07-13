<script lang="ts">
  import { Tween } from 'svelte/motion';
  import { elasticOut } from 'svelte/easing';
  import { cn } from '../utils';
  import { createEventDispatcher } from "svelte";
  
  export let isDarkMode = false;
  export let className: string = "";

  const dispatch = createEventDispatcher();
  
  const properties = {
    sun: {
      r: 9,
      transform: 40,
      cx: 12,
      cy: 4,
      opacity: 0
    },
    moon: {
      r: 5,
      transform: 90,
      cx: 30,
      cy: 0,
      opacity: 1
    }
  };

  const config = { 
    duration: 1000,  // Increased duration for more noticeable bounce
    easing: elasticOut
  };

  // For opacity, we'll use a shorter duration without bounce
  const fastConfig = {
    duration: 400,
    easing: elasticOut
  };

  // Create Tween instances for each property
  const r = new Tween(properties.sun.r, config);
  const transform = new Tween(properties.sun.transform, config);
  const cx = new Tween(properties.sun.cx, config);
  const cy = new Tween(properties.sun.cy, config);
  const opacity = new Tween(properties.sun.opacity, fastConfig);

  $: {
    const currentProps = isDarkMode ? properties.moon : properties.sun;
    r.set(currentProps.r);
    transform.set(currentProps.transform);
    cx.set(currentProps.cx);
    cy.set(currentProps.cy);
    opacity.set(currentProps.opacity);
  }
  
  // Get the current color based on theme
  $: currentThemeColor = isDarkMode ? "#E8CF7B" : "#336A83";
</script>

<button on:click={() => dispatch("toggle")} aria-label="theme-toggle" class={cn(className)}>
  <svg
    xmlns="http://www.w3.org/2000/svg"
    viewBox="0 0 24 24"
    fill="none"
    stroke={currentThemeColor}
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
    style="transform: rotate({transform.current}deg);"
    class={className}
  >
    <mask id="mask">
      <rect x="0" y="0" width="100%" height="100%" fill="white" />
      <circle
        cx={cx.current}
        cy={cy.current}
        r="9"
        fill="black"
      />
    </mask>
    <circle
      fill={currentThemeColor}
      cx="12"
      cy="12"
      r={r.current}
      mask="url(#mask)"
    />
    <g style="opacity: {opacity.current};" fill={currentThemeColor}>
      <line x1="12" y1="1" x2="12" y2="3" />
      <line x1="12" y1="21" x2="12" y2="23" />
      <line x1="4.22" y1="4.22" x2="5.64" y2="5.64" />
      <line x1="18.36" y1="18.36" x2="19.78" y2="19.78" />
      <line x1="1" y1="12" x2="3" y2="12" />
      <line x1="21" y1="12" x2="23" y2="12" />
      <line x1="4.22" y1="19.78" x2="5.64" y2="18.36" />
      <line x1="18.36" y1="5.64" x2="19.78" y2="4.22" />
    </g>
  </svg>
</button>
