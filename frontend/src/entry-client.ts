import './app.css'
import { hydrate } from 'svelte'
import App from './App.svelte'

const dataElement = document.getElementById('svelteData');
const { url, artifact, primaryNavData, secondaryNavData, tableOfContents } = dataElement ? JSON.parse(dataElement.textContent || '') : {
  url: window.location.pathname,
  primaryNavData: {},
  artifact: null,
  secondaryNavData: {},
  tableOfContents: [],
};

hydrate(App, {
  target: document.getElementById('app')!,
  props: {
    url,
    primaryNavData,
    artifact,
    secondaryNavData,
    tableOfContents,
  }
})
