# 📰 NewsGO

**NewsGO** é uma aplicação web simples e moderna desenvolvida em Go que consome a [NewsAPI](https://newsapi.org/) para exibir as 10 notícias mais relevantes sobre um assunto pesquisado. Com um front-end estilizado usando TailwindCSS.

## 🚀 Funcionalidades

- 🔍 Campo de busca por temas de interesse
- 📄 Listagem de até 10 notícias mais relevantes
- 🖼️ Layout responsivo e moderno com TailwindCSS

## 🧱 Estrutura do Projeto (MVC)

```
newsgo/
├── controllers/
│   └── news_controller.go
├── models/
│   └── news.go
├── views/
│   └── index.gohtml
│   └── results.gohtml
├── main.go
├── go.mod
└── README.md
```

## 🛠️ Tecnologias

- [Go](https://golang.org/)
- [NewsAPI](https://newsapi.org/)
- [TailwindCSS](https://tailwindcss.com/)

## ⚙️ Como executar

1. Clone este repositório:
   ```bash
   git clone https://github.com/seu-usuario/newsgo.git
   cd newsgo
   ```

2. Instale as dependências necessárias (Go instalado):
   ```bash
   go mod init newsgo
   go get
   ```

3. Informe sua chave de API e Execute o servidor:
   ```bash
   export NEWSAPI_KEY=sua_chave
   go run main.go
   ```

4. Acesse no navegador:
   ```
   http://localhost:8080
   ```