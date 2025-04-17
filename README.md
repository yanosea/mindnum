<div align="right">

![golangci-lint](https://github.com/yanosea/mindnum/actions/workflows/golangci-lint.yml/badge.svg)
![release](https://github.com/yanosea/mindnum/actions/workflows/release.yml/badge.svg)

</div>

<div align="center">

# 🧠 mindnum

![Language:Go](https://img.shields.io/static/v1?label=Language&message=Go&color=blue&style=flat-square)
![License:MIT](https://img.shields.io/static/v1?label=License&message=MIT&color=blue&style=flat-square)
[![Latest Release](https://img.shields.io/github/v/release/yanosea/mindnum?style=flat-square)](https://github.com/yanosea/mindnum/releases/latest)
<br/>
[Coverage Report](https://yanosea.github.io/mindnum/coverage.html)
</div>

## ℹ️ About

`mindnum` is a CLI tool to get the mind number.

Mind number is the number obtained by breaking down all dates of birth in the Western calendar and adding them to a single digit.
To calculate the mind number, you take the full date of birth (year, month, and day) and break it down into individual digits. Then, you sum these digits together. If the resulting sum is a multi-digit number, you continue to add the digits together until you get a single-digit number.

For example, if the date of birth is July 19, 1990 (1990-07-19):

1. Break down the date into individual digits: 1, 9, 9, 0, 0, 7, 1, 9
2. Sum these digits: 1 + 9 + 9 + 0 + 0 + 7 + 1 + 9 = 36
3. Since 36 is a multi-digit number, break it down again: 3 + 6 = 9
4. The mind number is 9.

This single-digit number is believed to have a specific meaning or personality description associated with it.

## 💻 Usage

```
Usage:
  mindnum [command]

Available Commands:
  get         🧠 Get a mind number from an argument or the birthday flag and show the personality description
  help        🤝 Help about any command
  completion  🔧 Generate the autocompletion script for the specified shell
  version     🔖 Show the version of mindnum

Flags:
  -h, --help     🤝 help for mindnum
  -v, --version  🔖 version for mindnum
```

## 🔧 Installation

### 🐭 Using go

```
go install github.com/yanosea/mindnum/v2/app/presentation/cli/mindnum@latest 
```

### 🍺 Using homebrew

```
brew tap yanosea/tap
brew install yanosea/tap/mindnum
```

### 📦 Download from release

Go to the [Releases](https://github.com/yanosea/mindnum/releases) and download the latest binary for your platform.  

## ✨ Update

### 🐭 Using go

Reinstall `mindnum`!

```
go install github.com/yanosea/mindnum/v2/app/presentation/cli/mindnum@latest 
```

### 🍺 Using homebrew

```
brew update
brew upgrade mindnum
```

### 📦 Download from release

Download the latest binary from the [Releases](https://github.com/yanosea/mindnum/releases) page and replace the old binary in your `$PATH`.

## 📃 License

[🔓MIT](./LICENSE)

## 🖊️ Author

[🏹yanosea](https://github.com/yanosea)

## 🤝 Contributing

This is my first golang application and also first GitHub public repository,

so please feel free to point me in the right direction🙏
