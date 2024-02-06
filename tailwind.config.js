module.exports = {
    mode: 'jit',
    content: ["public/view/**/*.{html,js,templ}"],
    plugins: [
        require('@tailwindcss/forms'),
        require('@tailwindcss/typography'),
    ]
}