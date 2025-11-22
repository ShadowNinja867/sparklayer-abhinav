Create the foundation of a to-do list application, focusing on backend functionality and essential frontend interaction. Your task is implementing a RESTful API using Go and a simple TypeScript interface using [Svelte](https://svelte.dev/). The goal is to be able to be able to list and add todos.

You are not expected to know Go, Svelte, or OpenAPI. Part of the challenge is to see how quickly you can adapt and pick new things up!

The backend needs to meet the openapi spec which is within the backend folder. You need to create the list endpoint and the add todo endpoint. The storage system is in-memory.

The frontend already has functionality to list the todos, your task is to complete the form which submits todo's to the backend system.

Please use this repo as your base. You can click "Use this template" in the top right on GitHub, then "Create a new repository". Commit and push your code so it can reviewed by us. Please get as far as you can within 2 hours.

If you find anything in the template you would like to improve, please feel free to!


## Setup

If not already installed, please install the following:
1. Go ([install instructions](https://go.dev/doc/install))
2. NPM/NodeJS. We recommend using [NVM](https://github.com/nvm-sh/nvm)

We have tested this with Go 1.25 and Node 20. You may have issues if you try to use a different version.

A good starting point would be to look at the following files:
- `backend/main.go`
- `frontend/src/App.svelte`

## Running

Open two separate terminals - one for the Svelte app and one for the golang API.


### Golang API

1. In the first terminal, change to the backend directory (`cd backend`)
2. Run `go run main.go` to start the API server

This must be running for the frontend to work.

When you make a change, you must stop the server (`ctrl-c` in the terminal), and restart it with `go run main.go`.


### Svelte App

1. In the second terminal, change to the frontend directory (`cd frontend`)
2. Run `npm run dev` to start the Svelte app
3. If it doesn't open automatically, open [http://localhost:5173](http://localhost:5173) to view your website

Leave this running. It will automatically update when you make any changes.
