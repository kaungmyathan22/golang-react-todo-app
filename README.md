# GoLang + React Todo App

A simple Todo App built with GoLang for the backend and React for the frontend. This project does not rely on a database and stores todo items in-memory.

## Features

- Add new todo items.
- Mark todo items as complete.
- Remove completed todo items.
- View the list of todo items.

## Prerequisites

- [Go](https://golang.org/) installed on your machine.
- [Node.js](https://nodejs.org/) and [npm](https://www.npmjs.com/) for the React frontend.

## Getting Started

1. #### Clone the repository:

   ```bash
   git clone https://github.com/your-username/your-todo-app.git
   cd your-todo-app
   ```
2. #### Start the GoLang server:

    ```bash
    cd backend
    go run main.go
    ```

The server will run on http://localhost:8080.

3. #### Install dependencies and start the React frontend:

    ```bash
    cd frontend
    npm install
    npm start
    ```

The React app will be accessible at http://localhost:3000.
