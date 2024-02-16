module.exports = {
  mode: 'jit',
  content: ["public/**/*.{html,js,templ,go}"],
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ]
}
