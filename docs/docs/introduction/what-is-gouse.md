# Introduction

## 🧠 What is Gouse?

- Gouse is a modern essentials [`Golang`](https://golang.org/) utility package delivering consistency, modularity, performance, and extras presets.
- Built on top of Go language, and others Go open-source packages.
- Inspired by [`Lodash`](https://lodash.com/) so `Javascript` user-friendly syntax.
- Lightweight package: easy to use, chainable, extendable, available.
- Gouse provides a wide variety of available methods, taking the hassle out of working with any type requests. So you can pick the one of them that fit your project demands.
- Powerful package of built-in functions provide comprehensive, and reliable solutions for any project size and compatible with all OS platforms.
- Open-source, free to use, and contribute.
- Offering over 400 utility functions across more than 25 domains. It serves as a toolkit for Go developers, providing ready-to-use solutions for common programming needs.

## ✨ Motivation

- Go has emerged as a server language, but it still doesn't have complete and consistent packages available to support coding development.
- Developers must write by hand or search manually. That wastes time and even causes many compatibility conflicts.
- Must update each dependent package every time when project update.
- Unexpected errors can easily arise during execution.
- Performance is not optimized.
- The number of lines of code is very long that not easy to read and understand.
- Code logic may not be consistent, making it difficult to maintain and scale.
- Hard to make sure the programs are executed the same across all environments.
- Lack of coherent learning materials and examples.
- And more...

👉 To address that need, Gouse was created as a powerful presets toolkit for Go developers, a collection of built-in functions and best practices that provide comprehensive, powerful, and reliable solutions. Trusted to build services, software platforms, APIs, and web applications.

## 🎯 Why to use

Thanks to Gouse, you can:

- Set up and scale projects rapidly.
- No config - import directly as utility functions.
- Handle complex logic use-cases such as database connection, build APIs, error handling, log management...
- Optimize performance and increase productivity.
- Build easily consistent systems with available functions.
- Avoid writing repetitive code and a unified code style.
- Reduce the number of lines of code and make it easier to read, understand, and maintain.
- Avoid compatibility conflicts and unexpected errors.
- Works smoothly in any environment and platform.
- Easy to learn and use, suitable for beginners and experts.
- Compatible with all frameworks and libraries.
- Referencing examples with comprehensive documentation
- And more...

## 🗝️ Key Philosophy

Gouse is built with several core design principles in mind:

- **User-Friendly & Lightweight**: Gouse features an intuitive JavaScript syntax with no setup required. Import utility functions directly and enjoy a flexible, chainable package available in various builds and formats.
- **Powerful & Versatile**: Access a wide range of methods for arrays, numbers, objects, and strings. Comprehensive documentation and examples make implementation smooth and efficient.
- **Scalable & Efficient**: Ideal for projects of any size, Gouse supports rapid setup, complex logic handling, and performance optimization across all operating systems.
- **Consistent & Maintainable**: Reduce repetitive code and ensure a unified style. Gouse makes your code cleaner, easier to maintain, and minimizes compatibility issues or unexpected errors.

The library follows a modular architecture where related functions are grouped together, making it easy to find and use specific utilities while maintaining a coherent API surface.

## 🗾 Usage Flow

![Flow](../img/flow.png)

## 🚀 Core functionality

Gouse organizes its functionality into logical groupings of related utilities:

|       Category       |                 Description                  |                        Example Functions                         |
| :------------------: | :------------------------------------------: | :--------------------------------------------------------------: |
|   Array Operations   | Functions for working with arrays and slices |           **`MinArr`**, **`MaxArr`**, **`Intersect`**            |
| String Manipulation  |  String processing and formatting utilities  |       **`Capitalize`**, **`CamelCase`**, **`StringSplit`**       |
|     Type System      |   Type checking, validation and conversion   |             **`TypeCheck`**, **`TypeCastToString`**              |
|   Time Operations    | Time retrieval, formatting, and manipulation |       **`Second`**, **`Minute`**, **`Hour`**, **`SleepS`**       |
|    API Utilities     |          Helpers for building APIs           |     **`Validate`**, **`UploadSingle`**, **`NewPagination`**      |
|       OS & I/O       |    System information and file operations    |           **`OsIO`**, **`OsDisk`**, **`IoCheckFile`**            |
|  Function Utilities  | Function manipulation and control utilities  | **`FunctionDelay`**, **`FunctionRetry`**, **`FunctionParallel`** |
| Database Connections |        Database connection utilities         |  **`ConnectRedis`**, **`ConnectPostgres`**, **`ConnectMongo`**   |
|   Chart Generation   |           Visualization utilities            |         **`ChartBar`**, **`ChartLine`**, **`ChartPie`**          |
|  Console Utilities   |     Terminal and command line utilities      |    **`ConsoleCmd`**, **`ConsoleBanner`**, **`ConsoleTable`**     |

> Below is a list of modules that Gouse supports. This project is still in development stage, so some features may be unavailable.

<div style="display: flex; flex-wrap: wrap; gap: 4px;">
  <a href="/receipts/api"><img src="https://img.shields.io/badge/Api-teal" alt="api" /></a>
  <a href="/receipts/array"><img src="https://img.shields.io/badge/Array-teal" alt="array" /></a>
  <a href="/receipts/cache"><img src="https://img.shields.io/badge/Cache-teal" alt="cache" /></a>
  <a href="/receipts/chart"><img src="https://img.shields.io/badge/Chart-teal" alt="chart" /></a>
  <a href="/receipts/config"><img src="https://img.shields.io/badge/Config-teal" alt="config" /></a>
  <a href="/receipts/console"><img src="https://img.shields.io/badge/Console-teal" alt="console" /></a>
  <a href="/receipts/connection"><img src="https://img.shields.io/badge/Connection-teal" alt="connection" /></a>
  <a href="/receipts/cron"><img src="https://img.shields.io/badge/Cron-teal" alt="cron" /></a>
  <a href="/receipts/crypto"><img src="https://img.shields.io/badge/Crypto-teal" alt="crypto" /></a>
  <a href="/receipts/date"><img src="https://img.shields.io/badge/Date-teal" alt="date" /></a>
  <a href="/receipts/function"><img src="https://img.shields.io/badge/Function-teal" alt="function" /></a>
  <a href="/receipts/helper"><img src="https://img.shields.io/badge/Helper-teal" alt="helper" /></a>
  <a href="/receipts/io"><img src="https://img.shields.io/badge/I/O-teal" alt="io" /></a>
  <a href="/receipts/logger"><img src="https://img.shields.io/badge/Logger-teal" alt="logger" /></a>
  <a href="/receipts/math"><img src="https://img.shields.io/badge/Math-teal" alt="math" /></a>
  <a href="/receipts/media"><img src="https://img.shields.io/badge/Media-teal" alt="media" /></a>
  <a href="/receipts/net"><img src="https://img.shields.io/badge/Net-teal" alt="net" /></a>
  <a href="/receipts/number"><img src="https://img.shields.io/badge/Number-teal" alt="number" /></a>
  <a href="/receipts/os"><img src="https://img.shields.io/badge/OS-teal" alt="os" /></a>
  <a href="/receipts/print"><img src="https://img.shields.io/badge/Print-teal" alt="print" /></a>
  <a href="/receipts/random"><img src="https://img.shields.io/badge/Random-teal" alt="random" /></a>
  <a href="/receipts/regex"><img src="https://img.shields.io/badge/Regex-teal" alt="regex" /></a>
  <a href="/receipts/string"><img src="https://img.shields.io/badge/String-teal" alt="string" /></a>
  <a href="/receipts/struct"><img src="https://img.shields.io/badge/Struct-teal" alt="struct" /></a>
  <a href="/receipts/time"><img src="https://img.shields.io/badge/Time-teal" alt="time" /></a>
  <a href="/receipts/type"><img src="https://img.shields.io/badge/Type-teal" alt="type" /></a>
</div>
