const withMT = require("@material-tailwind/react/utils/withMT");

/** @type {import('tailwindcss').Config} */
export default withMT({
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        grey: {
          1: "#1E1F22",
          2: "#2B2D31",
          3: "#313338",
        },
      },
    },
  },
  plugins: [],
});
