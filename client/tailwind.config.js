import colors from "tailwindcss/colors";
/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {},
    colors: {
      ...colors,
      primary: "#B56D41",
      secondary: "#F1D9CA",
      white: "#FFFDFA",
    },
    fontFamily: {
      handwriting: ["Gamja Flower", "cursive"],
      sans: ['Ubuntu', "sans-serif"]
    },
  },
  plugins: [],
};
