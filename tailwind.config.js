/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.templ", "./views/**/*.templ", "./**/*.html", "./**/*.go"],
  theme: {
    extend: {},
  },
  plugins: [require("@tailwindcss/forms"), require("@tailwindcss/typography")],
};
