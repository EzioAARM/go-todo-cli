<h1 align="center">🧩 ToDo CLI</h1>

<p align="center">
  <!-- Tech Stack Badges -->
  <a href="https://go.dev/"><img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white"/></a>
  <a href="https://www.postgresql.org/"><img src="https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white"/></a>
  <a href="https://www.cockroachlabs.com/"><img src="https://img.shields.io/badge/CockroachDB-6933FF?style=for-the-badge&logo=Cockroach%20Labs&logoColor=white"/></a>
  <a href="https://neo4j.com/"><img src="https://img.shields.io/badge/Neo4j-018bff?style=for-the-badge&logo=neo4j&logoColor=white"/></a>
  <a href="https://www.docker.com/"><img src="https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white"/></a>
</p>

---

<p align="center">
  <!-- Personal Links -->
  <a href="https://www.axel-rodriguez.com/"><img src="https://img.shields.io/badge/Website-000000?style=for-the-badge&logo=About.me&logoColor=white"/></a>
  <a href="https://www.instagram.com/axel___rodriguez/"><img src="https://img.shields.io/badge/Instagram-E4405F?style=for-the-badge&logo=instagram&logoColor=white"/></a>
  <a href="https://www.linkedin.com/in/axelrm/"><img src="https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white"/></a>
  <a href="https://github.com/EzioAARM"><img src="https://img.shields.io/badge/GitHub-100000?style=for-the-badge&logo=github&logoColor=white"/></a>
</p>

---

<h3 align="center">🚀 Multi-Database • 🧠 Go + Cobra</h3>

---

A simple yet powerful **command-line To-Do application** built in **Go**, designed to help you manage your daily tasks efficiently — while showcasing the flexibility, scalability, and performance of the Go programming language.

This project demonstrates practical usage of **Cobra CLI**, **modular design**, and **multi-database support**, making it a great foundation for future extensible CLI tools.

---

<h3 align="center">🚀 Features</h3>

- ✅ **Add new tasks** quickly and easily
- 📋 **List all tasks** with filtering options (pending, completed, all)
- 🗑️ **Delete tasks** by ID
- ✏️ **Mark tasks as done**
- 🧠 **Persistent storage** using multiple database backends
- ⚙️ **Cross-platform support** (Windows, macOS, Linux)
- 🧩 **Extensible command architecture** built with [Cobra](https://github.com/spf13/cobra)
- 💬 **Readable CLI output** with colorized text
- 🧰 **Configurable via JSON/YAML** for custom preferences

---

<h3 align="center">🧱 Technical Features</h3>

| Category                | Description                                                         |
| ----------------------- | ------------------------------------------------------------------- |
| **Language**            | Go (Golang)                                                         |
| **CLI Framework**       | [Cobra](https://github.com/spf13/cobra) for command structure       |
| **Data Layer**          | Abstracted repository pattern for interchangeable database backends |
| **Supported Databases** | PostgreSQL, CockroachDB, Neo4j (configurable at runtime)            |
| **Config System**       | Environment variables and config files (using `viper`)              |
| **Architecture**        | Modular, layered (Commands → Services → Repositories → DB)          |
| **Testing**             | Unit tests for commands and repositories                            |
| **Error Handling**      | Structured error responses with contextual logging                  |
| **Build System**        | Cross-compilation with Go toolchain (`go build -o todo.exe`)        |
| **CI/CD Ready**         | Designed for integration with GitHub Actions or similar pipelines   |

---

<h3 align="center">⚙️ Installation</h3>

### 1. Clone the repository

```bash
git clone https://github.com/<your-username>/golang-todo-cli-app.git
cd golang-todo-cli-app
```

### 2. Build the binary

```bash
go build -o todo.exe .\main.go   # Windows
go build -o todo main.go         # macOS/Linux
```

### 3. Run the CLI

```bash
.\todo.exe help     # Windows
.\todo help         # macOS/Linux
```

<h3 align="center">🗃️ Database Configuration</h3>

You can choose your preferred database backend by editing the config file (`config.yaml`) or setting an environment variable.

Example `config.yaml`:

```yaml
database:
  type: postgres # options: postgres | cockroach | neo4j
  connection_string: "postgres://user:password@localhost:5432/todo_app?sslmode=disable"
```

The app will automatically detect and use the appropriate repository implementation.

<h3 align="center">🧰 Example Usage</h3>

```bash
# Add a new task
todo add "Finish Golang CLI project"

# List all tasks
todo list

# Mark a task as done
todo done 1

# Delete a task
todo delete 1

```

<h3 align="center">🧪 Roadmap / Future Enhancements</h3>

- [ ] Add due dates and task prioritization
- [ ] Export tasks to JSON or CSV
- [ ] Add support for SQLite (embedded mode)
- [ ] Implement search and filters
- [ ] Add REST API wrapper (bonus: expose CLI as service)
- [ ] Docker container for local testing

<h3 align="center">🧑‍💻 About the Author</h3>

**Axel Rodríguez**
Software Engineer | Cloud Architect | Backend Developer

Passionate about building scalable, maintainable, and efficient systems.
Specialized in .NET, Golang, PostgreSQL, and Cloud Solutions (AWS & Azure).

📫 [LinkedIn](https://www.linkedin.com/in/axelrm/)
🐙 [GitHub](https://github.com/EzioAARM)

<h3 align="center">📄 License</h3>

This project is licensed under the [MIT License](https://mit-license.org/)

<h3 align="center">⭐️ Support</h3>

If you like this project, consider giving it a ⭐ on GitHub — it helps more than you think!
