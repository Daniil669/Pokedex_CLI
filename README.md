# Pokedex_CLI

Welcome to **Pokedex_CLI**, a lightweight command-line tool built to explore the world of Pokémon directly from your terminal. 

I built this project as a hands-on way to master **Golang**, specifically focusing on handling JSON data, making HTTP requests, and managing state in a CLI environment. It interfaces with the excellent [PokeAPI](https://pokeapi.co/) to fetch real-time data.

---

### 🚀 Features

* **Fast & Efficient:** Written in Go for compiled performance.
* **Interactive REPL:** A smooth "Read-Eval-Print Loop" experience.
* **Live API Integration:** Fetches up-to-date data from PokeAPI.
* **Extensible:** Built with a command pattern that makes adding new features easy.

---

### 🛠️ Commands

Currently, the Pokedex supports the following core commands:

| Command | Description |
| :--- | :--- |
| `help` | Displays a help message describing how to use the Pokedex. |
| `exit` | Exits the Pokedex CLI. |
| *(More coming soon)* | Map, Explore, Catch, and Inspect features are in development! |

---

### ⚙️ Installation & Usage

To run this project locally, ensure you have [Go](https://go.dev/doc/install) installed on your machine.

1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/your-username/Pokedex_CLI.git](https://github.com/your-username/Pokedex_CLI.git)
    cd Pokedex_CLI
    ```

2.  **Build the application:**
    ```bash
    go build -o pokedex
    ```

3.  **Run the CLI:**
    ```bash
    ./pokedex
    ```

### 🤝 Contributing

Since this is a practice project, feedback is always welcome! If you have suggestions on how to make the Go code more idiomatic or efficient, feel free to open an issue or a pull request.

---

> **Note:** This project is not affiliated with Nintendo, Game Freak, or The Pokémon Company. It is an educational tool powered by PokeAPI.