# gocreate

A simple CLI written in Go to quickly generate the initial structure of a new Go project.

## ✨ Features

- Runs `go mod init` with the correct repository path
- Generates a basic `main.go` file with `Hello, <project-name>`
- Prompts for your Git repository base on the first run (e.g., `github.com/yourname`)
- Automatically reuses the saved repository in future executions

---

## 🚀 Installation

### ✅ Option 1: Install via `go install` (recommended)

If the project is hosted on GitHub:

```bash
go install github.com/hermangoncalves/gocreate@latest
````

> This will install the binary at: `~/go/bin/gocreate`

Make sure `~/go/bin` is in your `PATH`:

```bash
echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc
source ~/.bashrc
```

For `zsh` users:

```bash
echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.zshrc
source ~/.zshrc
```

---

### 🛠️ Option 2: Manual compilation

1. Clone the repository:

```bash
git clone https://github.com/hermangoncalves/gocreate.git
cd gocreate
```

2. Build the binary:

```bash
go build -o gocreate
```

3. (Optional) Install globally:

```bash
sudo mv gocreate /usr/local/bin/
```

---

## 📦 Usage

```bash
gocreate my-project
```

### First time:

You’ll be prompted to enter your base Git repository:

```bash
Enter your default Git repository (e.g., github.com/yourname): yourname
```

The generated structure will look like:

```
my-project/
├── go.mod
└── main.go
```

### Next runs:

The CLI will automatically reuse your saved repository prefix.

---

## ⚙️ Configuration

* The base Git repository is saved in a config file: `~/.gocreateconfig`
* To change it, simply edit or delete that file:

```bash
nano ~/.gocreateconfig
```

or

```bash
rm ~/.gocreateconfig
```

The next time you run `gocreate`, it will prompt you again for your Git repository base.