import { defineConfig } from "vitepress";

import pkg from "../../package.json";

export default defineConfig({
  lang: "en-US",
  title: "Gouse",
  description: "A modern Golang utility presets",

  logo: "/img/32x32.png",

  lastUpdated: true,

  sitemap: {
    hostname: "https://gouse.vercel.app",
  },

  head: [["link", { rel: "icon", href: "/img/favicon.ico" }]],

  themeConfig: {
    logo: "/img/32x32.png",
    socialLinks: [
      {
        icon: "github",
        link: "https://github.com/thuongtruong109/gouse",
        ariaLabel: "github",
      },
    ],
    nav: [
      {
        text: "💡Introduction",
        link: "/introduction/what-is-gouse",
        activeMatch: "/introduction/",
      },
      {
        text: "🧩Receipts",
        link: "/receipts/array",
        activeMatch: "/receipts/",
      },
      {
        text: `🔖v${pkg.version}`,
        items: [
          {
            text: "Changelog",
            link: "https://github.com/thuongtruong109/gouse/blob/main/CHANGELOG.md",
          },
          {
            text: "Contributing",
            link: "https://github.com/thuongtruong109/gouse/blob/main/.github/CONTRIBUTING.md",
          },
        ],
      },
    ],
    search: {
      provider: "local",
      options: {
        detailedView: true,
      },
    },

    sidebar: {
      "/introduction/": {
        base: "/introduction/",
        items: [
          { text: "💡What is Gouse?", link: "what-is-gouse" },
          { text: "🏃‍➡️Getting Started", link: "getting-started" },
        ],
      },
      "/receipts/": {
        base: "/receipts/",
        items: [
          { text: "⚡ Api", link: "api" },
          { text: "📦 Array", link: "array" },
          { text: "🎫 Cache", link: "cache" },
          { text: "📊 Chart", link: "chart" },
          { text: "🪛 Config", link: "config" },
          { text: "💍 Connection", link: "connection" },
          { text: "🖨️ Console", link: "console" },
          { text: "⌛ Cron", link: "cron" },
          { text: "🔐 Crypto", link: "crypto" },
          { text: "📅 Date", link: "date" },
          { text: "⌛ Time", link: "time" },
          { text: "🫛 Function", link: "function" },
          { text: "🎯 I/O", link: "io" },
          { text: "✍️ Log", link: "log" },
          { text: "➗ Math", link: "math" },
          { text: "📸 Media", link: "media" },
          { text: "🛜 Net", link: "net" },
          { text: "🔢 Number", link: "number" },
          { text: "⭕ OS", link: "os" },
          { text: "💭 Random", link: "random" },
          { text: "🔃 Regex", link: "regex" },
          { text: "🔗 String", link: "string" },
          { text: "🛳️ Struct", link: "struct" },
          { text: "📚 Type", link: "type" },

          // { text: "📧 Email", link: "email" },
          // { text: '🐧Path', link: 'path' },
          // { text: 'Security', link: 'security' },
          // { text: 'Web', link: 'web' },
          // { text: 'Worker', link: 'worker' },
          // { text: 'XML', link: 'xml' },
          // { text: 'Zip', link: 'zip' },
          // { text: 'Collection', link: 'collection' },
          // { text: 'Convert', link: 'convert' },
          // { text: 'Debug', link: 'debug' },
          // { text: 'File', link: 'file' },
          // { text: 'Hash', link: 'hash' },
          // { text: 'Http', link: 'http' },
          // { text: 'Json', link: 'json' },
          // { text: 'Object', link: 'object' },
          // { text: 'Uuid', link: 'uuid' },
          // { text: 'Validate', link: 'validate' },
        ],
      },
    },

    editLink: {
      pattern:
        "https://github.com/thuongtruong109/gouse/edit/main/samples/:path".replace(
          "/receipts",
          ""
        ),
      text: "Improve this page on GitHub",
    },

    footer: {
      message: "Released under the MIT License.",
      copyright: "Copyright © 2024-present Tran Nguyen Thuong Truong",
    },
  },
});
