module.exports = {
  purge: { content: ["./public/**/*.html", "./src/**/*.{vue,js,ts,jsx,tsx}"] },
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {},
    colors: {
      gray: {
        200: "#b3b3b3",
        300: "#808080",
        400: "#666666",
        700: "#3a3a42",
        800: "#24242a",
        900: "#18181f"
      },
      green: {
        default: "#07fdbc"
      },
      purple: {
        default: "#ce51ed"
      },
      red: {
        default: "#f84d53"
      }
    }
  },
  variants: {
    extend: {}
  },
  plugins: []
};
