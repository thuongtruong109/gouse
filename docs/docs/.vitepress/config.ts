import { defineConfig } from 'vitepress';
import { createRequire } from 'module'

const require = createRequire(import.meta.url)
const pkg = require('../../package.json')

export default defineConfig({
  lang: 'en-US',
  title: 'Gouse',
  description: 'A modern Golang utility presets',

  sitemap: {
    hostname: 'https://gouse.vercel.app',
  },

  head: [
    [
      'favicon',
      { rel: 'icon', type: 'image/*', href: '/img/favicon.ico' }
    ],
  ],

  themeConfig: {
    nav: [
      {
        text: 'Introduction',
        link: '/introduction/what-is-gouse',
        activeMatch: '/introduction/'
      },
      {
        text: 'Receipts',
        link: '/receipts/array',
        activeMatch: '/receipts/'
      },
      {
        text: `v${pkg.version}`,
        items: [
          {
            text: 'Changelog',
            link: 'https://github.com/thuongtruong109/gouse/blob/main/CHANGELOG.md'
          },
          {
            text: 'Contributing',
            link: 'https://github.com/thuongtruong109/gouse/blob/main/.github/CONTRIBUTING.md'
          }
        ]
      }
    ],

    sidebar: {
      '/introduction/': { 
        base: '/introduction/', 
        items: [
            { text: 'What is Gouse?', link: 'what-is-gouse' },
            { text: 'Getting Started', link: 'getting-started' },
          ]
        
      },
      '/receipts/': { 
        base: '/receipts/', 
        items: [
          { text: 'Api', link: 'api' },
          { text: 'Array', link: 'array' },
          { text: 'Cache', link: 'cache' },
          { text: 'Chart', link: 'chart' },
          { text: 'Config', link: 'config' },
          { text: 'Connection', link: 'connection' },
          { text: 'Console', link: 'console' },
          { text: 'Cron', link: 'cron' },
          { text: 'Crypto', link: 'crypto' },
          { text: 'Date', link: 'date' },
          { text: 'Function', link: 'function' },
          { text: 'I/O', link: 'io' },
          { text: 'Log', link: 'Log' },
          { text: 'Math', link: 'math' },
          { text: 'Media', link: 'media' },
          { text: 'Net', link: 'net' },
          { text: 'Number', link: 'number' },
          { text: 'OS', link: 'os' },
          { text: 'Random', link: 'random' },
          { text: 'Regex', link: 'regex' },
          { text: 'String', link: 'string' },
          { text: 'Struct', link: 'struct' },
          { text: 'Type', link: 'type' },

          // { text: 'Path', link: 'path' },
          // { text: 'Security', link: 'security' },
          // { text: 'Web', link: 'web' },
          // { text: 'Worker', link: 'worker' },
          // { text: 'XML', link: 'xml' },
          // { text: 'Zip', link: 'zip' },
          // { text: 'Zlib', link: 'zlib' },
          // { text: 'Collection', link: 'collection' },
          // { text: 'Convert', link: 'convert' },
          // { text: 'Debug', link: 'debug' },
          // { text: 'File', link: 'file' },
          // { text: 'Hash', link: 'hash' },
          // { text: 'Http', link: 'http' },
          // { text: 'Json', link: 'json' },
          // { text: 'Object', link: 'object' },
          // { text: 'Time', link: 'time' },
          // { text: 'Uuid', link: 'uuid' },
          // { text: 'Validate', link: 'validate' },
          // { text: 'Web', link: 'web' },
          // { text: 'Xml', link: 'xml' }
        ]
      },
    },

    editLink: {
      pattern: 'https://github.com/thuongtruong109/gouse/edit/main/docs/docs/:path',
      text: 'Edit this page on GitHub'
    },

    footer: {
      message: 'Released under the MIT License.',
      copyright: 'Copyright © 2024-present Tran Nguyen Thuong Truong'
    }
  },
});