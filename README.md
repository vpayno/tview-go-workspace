# tview-go-workspace

[![actionlint](https://github.com/vpayno/tview-go-workspace/actions/workflows/gh-actions.yml/badge.svg?branch=main)](https://github.com/vpayno/tview-go-workspace/actions/workflows/gh-actions.yml)
[![go](https://github.com/vpayno/tview-go-workspace/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/vpayno/tview-go-workspace/actions/workflows/go.yml)
[![spellcheck](https://github.com/vpayno/tview-go-workspace/actions/workflows/spellcheck.yml/badge.svg?branch=main)](https://github.com/vpayno/tview-go-workspace/actions/workflows/spellcheck.yml)

Personal workspace for learning tview & Go.

## Links

- [tview GitHub](https://github.com/rivo/tview)
- [Go](https://www.go.dev/)

## RunMe Playbook

This and other readme files in this repo are RunMe Plabooks.

Use this playbook step/task to update the [RunMe](https://runme.dev) cli.

If you don't have runme installed, you'll need to copy/paste the command. :)

```bash { background=false category=runme closeTerminalOnSuccess=true excludeFromRunAll=true interactive=true interpreter=bash name=setup-install-runme promptEnv=true terminalRows=10 }
go install github.com/stateful/runme/v3@v3
```

Install Playbook dependencies:

```bash { background=false category=runme closeTerminalOnSuccess=true excludeFromRunAll=true interactive=true interpreter=bash name=setup-runme-deps promptEnv=true terminalRows=10 }
go install github.com/charmbracelet/gum@latest
```
