import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "Games Across Time",
  description: "A History and Wiki for Abstract Strategy Games and other Board Games",
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: 'Home', link: '/' },
    ],

    sidebar: [
      {
        text: 'Examples',
        items: [
          { text: 'Browse Games', link: '/g' },
          { text: 'Game Mechanisms', link: '/m' },
          { text: 'Game Designers', link: '/d' },
          { text: 'Artists & Authors', link: '/a' },
          { text: 'Game Publishers', link: '/p' }
        ]
      }
    ],

    socialLinks: [
      {
        icon: 'github',
        link: 'https://github.com/SymbolNotFound/gamesacrosstime'
      }
    ]
  }
})
