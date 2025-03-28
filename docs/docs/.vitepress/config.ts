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
          // 🔊🔑🧪
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
      copyright: "Copyright © 2024 Tran Nguyen Thuong Truong",
    },
  },
});
