const plugin = require('tailwindcss/plugin')
const defaultTheme = require('tailwindcss/defaultTheme')
const colors = require('tailwindcss/colors')

/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ['*.html'],
    theme: {
        extend: {

            colors: {
                teal: colors.teal,
                cyan: colors.cyan,
            },

            fontFamily: {

                // Default font
                sans: defaultTheme.fontFamily.sans,

                // Include: <link rel="stylesheet" href="https://rsms.me/inter/inter.css">
                'inter': ["'Inter var'", ...defaultTheme.fontFamily.sans],

                // Include: <link rel="preconnect" href="https://fonts.googleapis.com">
                // Include: <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
                // Include: <link href="https://fonts.googleapis.com/css2?family=DM+Sans:ital,wght@0,400;0,500;0,700;1,400;1,500;1,700&display=swap" rel="stylesheet"> 
                'dm-sans': ['"DM Sans"', ...defaultTheme.fontFamily.sans],

                // Include: <link rel="preconnect" href="https://fonts.googleapis.com" >
                // Include: <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
                // Include: <link href="https://fonts.googleapis.com/css2?family=Work+Sans:ital,wght@0,100;0,200;0,300;0,400;0,500;0,600;0,700;0,800;0,900;1,100;1,200;1,300;1,400;1,500;1,600;1,700;1,800;1,900&display=swap" rel="stylesheet"> 
                'work-sans': ["'Work Sans'", ...defaultTheme.fontFamily.sans],

                // Include: <link rel="preconnect" href="https://fonts.googleapis.com">
                // Include: <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
                // Include: <link href="https://fonts.googleapis.com/css2?family=Poppins:ital,wght@0,100;0,200;0,300;0,400;0,500;0,600;0,700;0,800;0,900;1,100;1,200;1,300;1,400;1,500;1,600;1,700;1,800;1,900&display=swap" rel="stylesheet"> 
                'poppins': ["'Poppins'", ...defaultTheme.fontFamily.sans],

                // Include: <link rel="preconnect" href="https://fonts.googleapis.com">
                // Include: <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
                // Include: <link href="https://fonts.googleapis.com/css2?family=Space+Grotesk:wght@300;400;500;600;700&display=swap" rel="stylesheet"> 
                'space-grotesk': ["'Space Grotesk'", ...defaultTheme.fontFamily.sans],

                // Include: <link rel="preconnect" href="https://fonts.googleapis.com">
                // Include: <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
                // Include: <link href="https://fonts.googleapis.com/css2?family=Open+Sans:ital,wght@0,300;0,400;0,500;0,600;0,700;0,800;1,300;1,400;1,500;1,600;1,700;1,800&display=swap" rel="stylesheet"> 
                'open-sans': ["'Open Sans'", ...defaultTheme.fontFamily.sans],

                // Include: <link rel="preconnect" href="https://fonts.googleapis.com">
                // Include: <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
                // Include: <link href="https://fonts.googleapis.com/css2?family=Montserrat:ital,wght@0,100;0,200;0,300;0,400;0,500;0,600;0,700;0,800;0,900;1,100;1,200;1,300;1,400;1,500;1,600;1,700;1,800;1,900&display=swap" rel="stylesheet"> 
                'montserrat': ["'Montserrat'", ...defaultTheme.fontFamily.sans],

                // Include: <link rel="preconnect" href="https://fonts.googleapis.com">
                // Include: <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
                // Include: <link href="https://fonts.googleapis.com/css2?family=Nunito+Sans:ital,wght@0,200;0,300;0,400;0,600;0,700;0,800;0,900;1,200;1,300;1,400;1,600;1,700;1,800;1,900&display=swap" rel="stylesheet">
                'nunito-sans': ["'Nunito Sans'", ...defaultTheme.fontFamily.sans],
            },

            colors: {
                // Sky
                // primary: {
                //     50: '#F0F9FF',
                //     100: '#E0F2FE',
                //     200: '#BAE6FD',
                //     300: '#7DD3FC',
                //     400: '#38BDF8',
                //     500: '#0EA5E9',
                //     600: '#0284C7',
                //     700: '#0369A1',
                //     800: '#075985',
                //     900: '#0C4A6E',
                // },

                primary: {
                    25: '#FCFCFD',  // WCAG(Contrast) AA 4.84:  gray-500
                    50: '#F9FAFB',  // WCAG(Contrast) AA 4.63:  gray-500
                    100: '#F2F4F7', // WCAG(Contrast) AA 4.49:  gray-500
                    200: '#EAECF0', // WCAG(Contrast) 4.19:     gray-500
                    300: '#D0D5DD', // WCAG(Contrast) 1.48:     white
                    400: '#98A2B3', // WCAG(Contrast) 2.58:     white
                    500: '#667085', // WCAG(Contrast) AA 4.95:  white
                    600: '#475467', // WCAG(Contrast) AAA:      white
                    700: '#344054', // WCAG(Contrast) AAA:      white
                    800: '#1D2939', // WCAG(Contrast) AAA:      white
                    900: '#101828', // WCAG(Contrast) AAA:      white
                },

                // Brown
                // primary: {
                //     50: '#fdf8f6',
                //     100: '#f2e8e5',
                //     200: '#eaddd7',
                //     300: '#e0cec7',
                //     400: '#d2bab0',
                //     500: '#bfa094',
                //     600: '#a18072',
                //     700: '#977669',
                //     800: '#846358',
                //     900: '#43302b',
                // },
            }
        },
    },
    plugins: [
        require('@tailwindcss/forms'),
        require('@tailwindcss/typography'),
        require('@tailwindcss/aspect-ratio'),

        plugin(function ({ addUtilities }) {
            addUtilities({
                '.scrollbar-w-0::-webkit-scrollbar': { 'width': '0px' },
                '.scrollbar-w-1::-webkit-scrollbar': { 'width': '4px' },
                '.scrollbar-w-2::-webkit-scrollbar': { 'width': '8px' },
                '.scrollbar-w-3::-webkit-scrollbar': { 'width': '12px' },
                '.scrollbar-w-4::-webkit-scrollbar': { 'width': '16px' },
                '.scrollbar-w-5::-webkit-scrollbar': { 'width': '20px' },
            })
        })
    ],

}