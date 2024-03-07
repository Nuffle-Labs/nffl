// @ts-check
// `@type` JSDoc annotations allow editor autocompletion and type checking
// (when paired with `@ts-check`).
// There are various equivalent ways to declare your Docusaurus config.
// See: https://docusaurus.io/docs/api/docusaurus-config

import {themes as prismThemes} from 'prism-react-renderer';

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'NEAR SFFL',
  tagline: 'NEAR Super Fast Finality Layer',
  favicon: 'img/favicon.ico',

  url: 'https://near-sffl.nethermind.io',
  baseUrl: '/',

  organizationName: 'NethermindEth',
  projectName: 'near-sffl',

  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',

  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

  presets: [
    [
      'classic',
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          sidebarPath: './sidebars.js',
          routeBasePath: '/',
        },
        blog: false,
        theme: {
          customCss: './src/css/custom.css',
        },
      }),
    ],
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      image: 'img/near-sffl-social-card.jpg',
      navbar: {
        title: 'NEAR SFFL',
        logo: {
          alt: 'NEAR Logo',
          src: 'img/near-icon.svg',
        },
        items: [
          {
            type: 'docSidebar',
            sidebarId: 'sidebar',
            position: 'left',
            label: 'Docs',
          },
          {
            href: 'https://github.com/NethermindEth/near-sffl',
            label: 'GitHub',
            position: 'right',
          },
        ],
      },
      footer: {
        style: 'dark',
        links: [
          {
            title: 'Docs',
            items: [
              {
                label: 'Introduction',
                to: '/',
              },
              {
                label: 'Protocol Design',
                to: '/category/protocol-design',
              },
              {
                label: 'Milestones',
                to: '/milestones',
              },
            ],
          },
          {
            title: 'Community',
            items: [
              {
                label: 'Discord',
                href: 'http://near.chat/',
              },
              {
                label: 'Discourse',
                href: 'https://gov.near.org/',
              },
              {
                label: 'Reddit',
                href: 'https://www.reddit.com/r/nearprotocol/',
              },
              {
                label: 'Telegram',
                href: 'https://t.me/cryptonear',
              },
              {
                label: 'WeChat',
                href: 'https://pages.near.org/wechat',
              },
              {
                label: 'X',
                href: 'https://twitter.com/nearprotocol',
              },
              {
                label: 'YouTube',
                href: 'https://www.youtube.com/channel/UCuKdIYVN8iE3fv8alyk1aMw',
              },
            ],
          },
          {
            title: 'More',
            items: [
              {
                label: 'GitHub',
                href: 'https://github.com/NethermindEth/near-sffl',
              },
            ],
          },
        ],
        copyright: `Copyright © ${new Date().getFullYear()} Near.org. Built with Docusaurus.`,
      },
      prism: {
        theme: prismThemes.github,
        darkTheme: prismThemes.dracula,
      },
    }),
};

export default config;
