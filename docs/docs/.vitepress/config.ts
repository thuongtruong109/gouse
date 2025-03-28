import { defineConfig } from "vitepress";

import pkg from "../../package.json";
import { editLink } from "./util";

export default defineConfig({
  lang: "en-US",
  title: "Gouse",
  description: "A modern Golang utility presets",
  lastUpdated: true,
  head: [["link", { rel: "icon", href: "/img/favicon.ico" }]],
  ignoreDeadLinks: true,
  sitemap: {
    hostname: "https://gouse.vercel.app",
  },
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
          { text: "📚About", link: "about" },
        ],
      },
      "/receipts/": {
        base: "/receipts/",
        items: [
          { text: "⚡ Api", link: "api" },
          { text: "📦 Array", link: "array" },
          { text: "🎫 Cache", link: "cache" },
          { text: "📊 Chart", link: "chart" },
          { text: "🎨 Color", link: "color" },
          { text: "🪛 Config", link: "config" },
          { text: "💍 Connection", link: "connection" },
          { text: "🖨️ Console", link: "console" },
          { text: "🗂️ Context", link: "context" },
          { text: "⌛ Cron", link: "cron" },
          { text: "🔐 Crypto", link: "crypto" },
          { text: "📅 Date", link: "date" },
          { text: "🔍 Error", link: "error" },
          { text: "🫛 Function", link: "function" },
          { text: "🛜 Http", link: "http" },
          { text: "🎯 I/O", link: "io" },
          { text: "✍️ Logger", link: "logger" },
          { text: "➗ Math", link: "math" },
          { text: "📸 Media", link: "media" },
          { text: "🔢 Number", link: "number" },
          { text: "⭕ OS", link: "os" },
          { text: "📝 Print", link: "print" },
          { text: "💭 Random", link: "random" },
          { text: "🔃 Regex", link: "regex" },
          { text: "💻 Server", link: "server" },
          { text: "🔗 String", link: "string" },
          { text: "🛳️ Struct", link: "struct" },
          { text: "⌛ Time", link: "time" },
          { text: "📚 Type", link: "type" },

          // { text: "📦 Middleware", link: "middleware" },
          // { text: "🗂️ Mime", link: "mime" },
          // { text: "🔊 Notify", link: "notify" },
          // { text: "🔑 Oauth", link: "oauth" },
          // { text: "📦 Option", link: "option" },
          // { text: "📦 Parser", link: "parser" },
          // { text: "🗂️ Path", link: "path" },
          // { text: "📦 Plugin", link: "plugin" },
          // { text: "🗂️ Directory", link: "directory" },
          // { text: "💾 Disk", link: "disk" },
          // { text: "🎨 Emoji", link: "emoji" },
          // { text: "📦 Env", link: "env" },
          // { text: "🗂️ Event", link: "event" },
          // { text: "⚙️ Flag", link: "flag" },
          // { text: "🗃️ Copy", link: "copy" },
          // { text: "🗂️ Database", link: "database" },
          // { text: "📁 File", link: "file" },
          // { text: "📊 Filter", link: "filter" },
          // { text: "🔍 Find", link: "find" },
          // { text: "🎨 Format", link: "format" },
          // { text: "🧩 Gouse", link: "gouse" },
          // { text: "🔗 Url", link: "url" },
          // { text: "🗂️ Zip", link: "zip" },
          // { text: "🧪 Test", link: "test" },
          // { text: "🔌 Plugin", link: "plugin" },
          // { text: "📧 Email", link: "email" },
          // { text: '🐧Path', link: 'path' },
          // { text: 'Security', link: 'security' },
          // { text: 'Web', link: 'web' },
          // { text: 'Worker', link: 'worker' },
          // { text: 'XML', link: 'xml' },
          // { text: 'Collection', link: 'collection' },
          // { text: 'Convert', link: 'convert' },
          // { text: 'Debug', link: 'debug' },
          // { text: 'Hash', link: 'hash' },
          // { text: 'Object', link: 'object' },
        ],
      },
    },

    editLink: {
      pattern: editLink(
        "https://github.com/thuongtruong109/gouse/edit/main/samples/:path"
      ),
      text: "Improve this page on GitHub",
    },

    footer: {
      message: "Released under the MIT License.",
      copyright: "Copyright © 2024-present Tran Nguyen Thuong Truong",
    },
  },
});
