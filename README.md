# Projeto Multi-Stack: Go + Rust + C/C++

![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go&logoColor=white)
![Rust](https://img.shields.io/badge/Rust-1.80+-000000?logo=rust&logoColor=white)
![C/C++](https://img.shields.io/badge/C%2FC++-17+-00599C?logo=c%2B%2B&logoColor=white)
![gRPC](https://img.shields.io/badge/gRPC-Protocol-4285F4?logo=google&logoColor=white)
![Bubbletea](https://img.shields.io/badge/Bubbletea-TUI-FF69B4?logo=github&logoColor=white)
![Fyne](https://img.shields.io/badge/Fyne-GUI-563D7C?logo=go&logoColor=white)

---

## 📌 Visão Geral

Este projeto tem como objetivo criar uma plataforma **modular e extensível**, com um **backend principal em Go** e suporte a **agentes/plugins em Rust e C/C++**.  
A comunicação entre os módulos é feita via **gRPC**, permitindo fácil integração e expansão.  

---

## 🏗️ Estrutura do Projeto

### **Go (backend principal)**
- 🔌 Loader de plugins
- 💻 Gerenciador CLI → [Bubbletea](https://github.com/charmbracelet/bubbletea) + [Lipgloss](https://github.com/charmbracelet/lipgloss)
- 🖥️ Interface gráfica (GUI) → [Fyne](https://github.com/fyne-io/fyne)
- 🌐 Gerenciador Web + Monitorador Web

### **Rust e C/C++ (agentes e plugins)**
- Agentes de monitoramento
- Plugins externos (exemplo inicial em Go → replicado em Rust/C++)

---

## 📡 Comunicação

- **gRPC** para comunicação entre **backend** e **agentes/plugins**  
- Contratos definidos em arquivos `.proto` (pasta `/proto`)  

---

## 🛠️ Funcionalidades Planejadas

- [ ] Definir requisitos do projeto  
- [ ] Estrutura inicial do backend em Go  
- [ ] Suporte a CLI com Bubbletea + Lipgloss  
- [ ] Interface GUI (Fyne)  
- [ ] Painel Web + Monitorador Web  
- [ ] Contratos gRPC para agentes/plugins  
- [ ] Plugin de exemplo em Go (replicação em Rust/C++)  
- [ ] Testes e documentação  

---

## 📂 Estrutura de Pastas

sentinel/
  ├─ api/
  │   └─ agent.proto
  ├─ cmd/
  │   ├─ cli.go
  │   └─ gui.go
  ├─ internal/
  │   ├─ config/
  │   │   └─ config.go
  │   ├─ logging/
  │   │   └─ logger.go
  │   └─ plugin/
  │       └─ loader.go
  ├─ pkg/
  │   ├─ ui/
  │   │   └─ manager.go
  │   └─ web/
  │       ├─ monitor.go
  │       └─ server.go
  ├─ plugins/
  │   └─ greeter/
  │       └─ main.go
  ├─ webui/
  │   └─ index.html
  ├─ main.go
  ├─ go.mod
  └─ README.md


WIP