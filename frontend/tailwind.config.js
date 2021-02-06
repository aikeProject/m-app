module.exports = {
  purge: { content: ["./public/**/*.html", "./src/**/*.{vue,js,ts,jsx,tsx}"] },
  darkMode: false, // or 'media' or 'class'
  theme: {
    // extend: {},
    colors: {
      blue: {
        default: "#27d1ff"
      },
      gray: {
        100: "#cbccd2",
        200: "#b3b3b3",
        300: "#808080",
        400: "#666666",
        700: "#3a3a42",
        800: "#212128",
        900: "#18181f"
      },
      green: {
        default: "#27ffa7"
      },
      pink: {
        default: "#ff27ff"
      },
      purple: {
        400: "#d690ff",
        default: "#ba45ff"
      },
      red: {
        default: "#f84d53"
      },
      yellow: {
        default: "#ffe027"
      }
    }
  },
  variants: {
    extend: {}
  },
  plugins: []
};
