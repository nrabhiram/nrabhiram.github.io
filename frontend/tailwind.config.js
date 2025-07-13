/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{svelte,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        montserrat: [
          "Montserrat Alternates", 
          "ui-sans-serif", 
          "system-ui",
          "-apple-system", 
          "BlinkMacSystemFont", 
          "Segoe UI", 
          "Roboto", 
          "Helvetica Neue", 
          "Arial", 
          "Noto Sans", 
          "sans-serif", 
          "Apple Color Emoji", 
          "Segoe UI Emoji", 
          "Segoe UI Symbol", 
          "Noto Color Emoji"
        ]
      },
      colors: {
        vaxitas: {
          primary: 'var(--color-primary)',
          secondary: 'var(--color-secondary)',
          tertiary: 'var(--color-tertiary)',
          pale: 'var(--color-pale)',
        }
      }
    },
  },
  plugins: [],
}

