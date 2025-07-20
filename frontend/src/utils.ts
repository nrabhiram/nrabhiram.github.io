import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";
import type { Artifact } from "./types";

const VAXITAS_LS_THEME_KEY = 'theme';
const VAXITAS_LS_DRAWER_KEY = 'drawer_open';

const VAXITAS_BASE_URL = import.meta.env.VITE_BASE_URL || 'http://localhost:8081';

const isBrowser = typeof window !== 'undefined' && typeof localStorage !== 'undefined';

const cn = (...inputs: ClassValue[]) => {
  return twMerge(clsx(inputs));
};

const getLocalStorage = <T>(key: string, defaultValue: T): T => {
  if (!isBrowser) return defaultValue;
  const item = localStorage.getItem(key);
  return item ? JSON.parse(item) : defaultValue;
};

const setLocalStorage = <T>(key: string, value: T): void => {
  if (!isBrowser) return;
  localStorage.setItem(key, JSON.stringify(value));
};

const matchesPath = (currentUrl: string, pattern: string) => {
  if (pattern === '/') {
    return currentUrl === '/';
  }
  
  if (pattern.includes(':slug')) {
    const basePath = pattern.replace('/:slug', '');
    if (basePath === '') {
      // for "/:slug" pattern - matches any single segment
      const segments = currentUrl.replace(/^\//, '').split('/');
      return segments.length === 1 && segments[0] !== '';
    } else {
      // for "/about/:slug" or "/now/:slug" patterns
      return currentUrl.startsWith(basePath + '/') && currentUrl !== basePath;
    }
  }
  
  return currentUrl === pattern;
}

function isPrimaryNavLinkActive(currentUrl: string, linkUrl: string) {
  if (currentUrl === "/" || linkUrl === "/") return currentUrl === linkUrl;
  return currentUrl.startsWith(linkUrl);
}

function isBasePath(url: string) {
  const pathSegments = url.replace(/^\//, '').split('/');
  return pathSegments.length === 1;
}

function getPageTitle(artifact: Artifact | null) {
  const baseTitle = 'Vaxitas';
  const title = `${baseTitle}${artifact && artifact.name && artifact.name !== baseTitle ? ` | ${artifact.name}` : ''}`;
  return title;
}

export { 
  VAXITAS_LS_THEME_KEY,
  VAXITAS_LS_DRAWER_KEY,
  VAXITAS_BASE_URL,
  isBrowser,
  getLocalStorage,
  setLocalStorage,
  cn,
  matchesPath,
  isPrimaryNavLinkActive,
  isBasePath,
  getPageTitle,
}
