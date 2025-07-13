import { derived, writable, type Writable } from "svelte/store";
import { getLocalStorage, VAXITAS_LS_THEME_KEY } from "../utils";

interface Theme {
  mode: "dark" | "light";
  device: "desktop" | "mobile";
}

const theme: Writable<Theme> = writable({
  mode: function(){
    const defaultTheme = typeof window !== 'undefined' && window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
    const savedTheme = getLocalStorage(VAXITAS_LS_THEME_KEY, defaultTheme) as 'light' | 'dark';
    return savedTheme;
  }(),
  device: "desktop",
});

const getMode = derived(theme, ($theme) => $theme.mode);
const getDevice = derived(theme, ($theme) => $theme.device);

const ThemeStore = {
	subscribe: theme.subscribe,
  set: theme.set,
  update: theme.update,
  setMode: (mode: 'dark' | 'light') => {
    theme.update((val) => ({
      ...val,
      mode
    }));
  },
  toggleMode: () => {
    theme.update((val) => ({
      ...val,
      mode: val.mode === "light" ? "dark" : "light",
    }));
  },
  handleResize: (width: number) => {
    theme.update((val) => ({
      ...val,
      device: width >= 1024 ? "desktop" : "mobile",
    }));
  },
  getMode: getMode,
  getDevice: getDevice,
};

export { ThemeStore }
