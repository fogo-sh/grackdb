module.exports = {
  purge: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {},
    colors: {
      accent: "#fff66c",
      background: "#23272A",
      gray: {
        darkest: "#0c0d0e",
        dark: "#23272a",
        DEFAULT: "#464e54",
        light: "#656769",
        lightest: "#bdbebf",
      },
      white: "#fafafa",
    },
  },
  variants: {
    extend: {},
  },
  plugins: [],
};
