# go-project-template

An opinionated template for Go projects.

* vscode
* devbox
* aider

https://github.com/user-attachments/assets/5737e02a-55ce-4120-b184-8b2966bdce2d

## One-time setup

Install the following if not already on your machine:

* [devbox](https://www.jetify.com/docs/devbox/)
* [aider](https://aider.chat)
* vscode extensions from `.vscode/extensions.json`

### Aider Setup

To use an LLM chatbot in the terminal, configure the following:

* Add API keys to a `.env` file (see `.env.example`)
* Add [model](https://aider.chat/docs/llms.html) to an `.aider-model` file (see `.aider-model.example`)

## Usage

1) Open your project in vscode
2) cmd+shift+p, `Devbox: Reopen in Devbox shell environment`

## Updating Go Version

* `devbox update go`
* Update go version in `go.mod`

## Helpful Docs

* https://www.jetify.com/docs/devbox/quickstart/
