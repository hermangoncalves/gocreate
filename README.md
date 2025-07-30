# gocreate

Uma CLI simples escrita em Go para gerar rapidamente a estrutura inicial de um novo projeto Go.

## ✨ Funcionalidades

- Executa `go mod init` com o repositório correto
- Gera um arquivo `main.go` básico com `Hello, <nome-do-projeto>`
- Salva o repositório Git base na primeira execução (ex: `github.com/seunome/`)
- Reutiliza o repositório automaticamente em execuções futuras

---

## 🚀 Instalação

### ✅ Opção 1: Instalar com `go install` (recomendado)

Se o projeto estiver hospedado no GitHub:

```bash
go install github.com/hermangoncalves/gocreate@latest
```

> Isso instalará o binário em: `~/go/bin/gocreate`

Certifique-se de que `~/go/bin` está no seu `PATH`:

```bash
echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc
source ~/.bashrc
```

Se usa `zsh`:

```bash
echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.zshrc
source ~/.zshrc
```

---

### 🛠️ Opção 2: Compilar manualmente

1. Clone o repositório:

```bash
git clone https://github.com/hermangoncalves/gocreate.git
cd gocreate
```

2. Compile o binário:

```bash
go build -o gocreate
```

3. (Opcional) Instale globalmente:

```bash
sudo mv gocreate /usr/local/bin/
```

---

## 📦 Uso

```bash
gocreate nome-do-projeto
```

### Primeira vez:

Você será solicitado a informar seu repositório base, ex:

```bash
Informe o seu repositório Git (ex: github.com/seunome): seunome
```

A estrutura gerada será:

```
nome-do-projeto/
├── go.mod
└── main.go
```

### Próximas execuções:

O repositório será reutilizado automaticamente.

---

## ⚙️ Configuração

- O repositório base informado é salvo em um arquivo `~/.gocreateconfig
`.
- Para alterar, basta editar ou remover esse arquivo.
