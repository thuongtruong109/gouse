---
layout: home

title: Gouse
titleTemplate: Utility presets package

hero:
  name: Gouse
  text: A modern Golang utility presets
  tagline: Resuable functions rapidly
  actions:
    - theme: brand
      text: What is Gouse?
      link: /introduction/what-is-gouse
    - theme: alt
      text: Quickstart
      link: /introduction/getting-started
    - theme: alt
      text: GitHub
      link: https://github.com/thuongtruong109/gouse
  image:
    src: /img/512x512.png
    alt: Gouse

features:
  - icon: 📝
    title: User-Friendly and Lightweight
    details: Gouse features an intuitive JavaScript syntax with no setup required. Import utility functions directly and enjoy a flexible, chainable package available in various builds and formats.
  - icon: 🛠️
    title: Powerful and Versatile
    details: Access a wide range of methods for arrays, numbers, objects, and strings. Comprehensive documentation and examples make implementation smooth and efficient.
  - icon: 🧩
    title: Scalable and Efficient
    details: Ideal for projects of any size, Gouse supports rapid setup, complex logic handling, and performance optimization across all operating systems.
  - icon: 🚀
    title: Consistent and Maintainable
    details: Reduce repetitive code and ensure a unified style. Gouse makes your code cleaner, easier to maintain, and minimizes compatibility issues or unexpected errors.
---

<style>
@import './style.css';

:root {
  --vp-home-hero-name-color: transparent !important;
  --vp-home-hero-name-background: -webkit-linear-gradient(120deg, #bd34fe 30%, #41d1ff) !important;
  --vp-home-hero-image-background-image: linear-gradient(-45deg, #bd34fe 50%, #47caff 50%) !important;
  --vp-home-hero-image-filter: blur(44px) !important;
}

@media (min-width: 640px) {
  :root {
    --vp-home-hero-image-filter: blur(56px);
  }
}

@media (min-width: 960px) {
  :root {
    --vp-home-hero-image-filter: blur(68px);
  }
}

</style>
