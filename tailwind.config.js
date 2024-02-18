module.exports = {
  content: [
    "public/*.{html,js,templ,go}",
    "public/**/*.{html,js,templ,go}",
    "public/**/**/*.{html,js,templ,go}"
  ],
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ]
}
