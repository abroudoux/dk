# dk

🦧 ❤️ 🐳 A minimalist TUI to interract faster with Docker

![Monkey loves Whale by marde (https://drawception.com/player/922067/marde/)](./ressources/monkey-loves-whale.png)

## 🚀 Installation

### Homebrew

Add my tap

```bash
brew tap abroudoux/tap
```

Then you'll be able to download it using

```bash
brew install abroudoux/tap/dk
```

Congratulations!

```bash
dk --version
#       _ _
#    __| | | __
#   / _` | |/ /
#  | (_| |   <
#   \__,_|_|\_\

# dk version 0.2.2
```

### Manual

You can create the binary with

```bash
go build -o dk ./cmd/main.go
```

Then paste it in your `bin` directory (e.g., on MacOS it's `/usr/bin/local`) \
Don't forget to grant execution perssions

```bash
chmod +x dk
```

You can now use `dk`!

```bash
dk --version
#       _ _
#    __| | | __
#   / _` | |/ /
#  | (_| |   <
#   \__,_|_|\_\

# dk version 0.2.2
```

## 💻 Usage

### Taskfile

Install [Taskfile](https://taskfile.dev/installation/) and use it to run the program

```bash
task run
```

### Manual

Execute the binary by using

```bash
go build -o ./bin/dk ./cmd/main.go && ./bin/dk
```

## 🧑‍🤝‍🧑 Contributing

To contribute, fork the repository and open a pull request detailling your changes. \
Create a branch with a [conventionnal name](https://tilburgsciencehub.com/building-blocks/collaborate-and-share-your-work/use-github/naming-git-branches/).

- fix: `bugfix/bug-fixed`
- features: `feature/amazing-feature`
- test: `test/famous-test`
- refactor `refactor/great-change`

## 📌 Roadmap

- [ ] `-it` mode
- [ ] Volumes management
- [ ] Networks management
- [x] Create images from source
- [x] `-env` flag when running a container

## 📑 License

This project is under [MIT License](LICENSE).
