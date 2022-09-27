/*
 * @ts-check
 * Note: type annotations allow type checking and IDEs autocompletion
 */

const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');

const GITHUB_URL = 'https://github.com/go-task/task';
const TWITTER_URL = 'https://twitter.com/taskfiledev';
const DISCORD_URL = 'https://discord.gg/6TY36E39UK';

/** @type {import('@docusaurus/types').Config} */
const config = {
  baseUrl: '/',
  deploymentBranch: 'gh-pages',
  favicon: 'img/favicon.ico',
  i18n: {
    defaultLocale: 'en',
    locales: ['en']
  },
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'throw',
  organizationName: 'go-task',
  presets: [
    [
      'classic',
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        blog: false,
        docs: {
          routeBasePath: '/',
          sidebarPath: require.resolve('./sidebars.js')
        },
        gtag: {
          anonymizeIP: true,
          trackingID: 'G-4RT25NXQ7N'
        },
        sitemap: {
          changefreq: 'weekly',
          ignorePatterns: ['/tags/**'],
          priority: 0.5
        },
        theme: {
          customCss: [
            require.resolve('./src/css/custom.css'),
            require.resolve('./src/css/carbon.css')
          ]
        }
      })
    ]
  ],
  projectName: 'task',
  scripts: [
    {
      async: true,
      src: '/js/carbon.js'
    }
  ],
  tagline: 'A task runner / simpler Make alternative written in Go ',
  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      // NOTE(@andreynering): Don't worry, these keys are meant to be public =)
      algolia: {
        apiKey: '34b64ae4fc8d9da43d9a13d9710aaddc',
        appId: '7IZIJ13AI7',
        indexName: 'taskfile'
      },
      footer: {
        links: [
          {
            items: [
              {
                label: 'Installation',
                to: '/installation/'
              },
              {
                label: 'Usage',
                to: '/usage/'
              },
              {
                label: 'Donate',
                to: '/donate/'
              }
            ],
            title: 'Pages'
          },
          {
            items: [
              {
                href: GITHUB_URL,
                label: 'GitHub'
              },
              {
                href: TWITTER_URL,
                label: 'Twitter'
              },
              {
                href: DISCORD_URL,
                label: 'Discord'
              },
              {
                href: 'https://opencollective.com/task',
                label: 'OpenCollective'
              }
            ],
            title: 'Community'
          }
        ],
        style: 'dark'
      },
      metadata: [
        {
          content: 'https://taskfile.dev/img/og-image.png',
          name: 'og:image'
        }
      ],
      navbar: {
        items: [
          {
            docId: 'installation',
            label: 'Installation',
            position: 'left',
            type: 'doc'
          },
          {
            docId: 'usage',
            label: 'Usage',
            position: 'left',
            type: 'doc'
          },
          {
            docId: 'api_reference',
            label: 'API',
            position: 'left',
            type: 'doc'
          },
          {
            docId: 'donate',
            label: 'Donate',
            position: 'left',
            type: 'doc'
          },
          {
            href: GITHUB_URL,
            label: 'GitHub',
            position: 'right'
          },
          {
            href: TWITTER_URL,
            label: 'Twitter',
            position: 'right'
          },
          {
            href: DISCORD_URL,
            label: 'Discord',
            position: 'right'
          }
        ],
        logo: {
          alt: 'Task Logo',
          src: 'img/logo.svg'
        },
        title: 'Task'
      },
      prism: {
        darkTheme: darkCodeTheme,
        theme: lightCodeTheme
      }
    }),
  title: 'Task',
  url: 'https://taskfile.dev'
};

module.exports = config;
