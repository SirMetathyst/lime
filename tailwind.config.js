/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ['*.html'],
    plugins: [
        require('@tailwindcss/forms'),
        require('@tailwindcss/typography')
    ],
}