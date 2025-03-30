# Setting Up a Go-specific VSCode Profile

## Step 1: Create the Profile

1. Open VSCode
2. Click on the profile icon in the bottom-left corner (or press `Ctrl+Shift+P` and type "Profiles: Create Profile")
3. Select "Create Profile..."
4. Choose "New Empty Profile"
5. Name it "Go Development"

## Step 2: Install Essential Go Extensions

In your new profile, install these extensions:

- **Go** (by Go Team at Google) - The official Go extension
- **Go Test Explorer** - For running/debugging tests
- **Go Doc** - Documentation view for Go packages
- **Go Outliner** - Code structure visualization
- **gopls** - Go language server (will be installed automatically with the Go extension)

## Step 3: Configure Go Settings

Open settings (File > Preferences > Settings) and add these Go-specific settings:

```json
{
  "go.useLanguageServer": true,
  "go.toolsManagement.autoUpdate": true,
  "go.lintOnSave": "package",
  "go.vetOnSave": "package",
  "go.formatTool": "goimports",
  "go.lintTool": "golangci-lint",
  "go.testOnSave": false,
  "go.testFlags": ["-v"],
  "go.coverOnSave": false,
  "go.coverageDecorator": {
    "type": "highlight",
    "coveredHighlightColor": "rgba(64,128,128,0.5)",
    "uncoveredHighlightColor": "rgba(128,64,64,0.25)"
  },
  "editor.formatOnSave": true,
  "editor.codeActionsOnSave": {
    "source.organizeImports": true
  },
  "editor.snippetSuggestions": "inline",
  "editor.suggest.snippetsPreventQuickSuggestions": false,
  "editor.tabSize": 4,
  "editor.insertSpaces": false,
  "[go]": {
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  },
  "[go.mod]": {
    "editor.formatOnSave": true
  }
}
```

## Step 4: Create a Go Workspace

1. Create a new folder for your Go projects (e.g., `GoProjects`)
2. Inside, create a file named `go.code-workspace` with this content:

```json
{
  "folders": [
    {
      "path": "."
    }
  ],
  "settings": {
    "go.useLanguageServer": true,
    "go.toolsManagement.autoUpdate": true,
    "go.testOnSave": false,
    "go.formatTool": "goimports",
    "go.lintTool": "golangci-lint",
    "go.lintOnSave": "package",
    "go.vetOnSave": "package"
  },
  "extensions": {
    "recommendations": [
      "golang.go",
      "premparihar.gotestexplorer",
      "766b.go-outliner"
    ]
  }
}
```

## Step 5: Configure .vscode Folder for Go Projects

For each Go project, create a `.vscode` folder with these files:

### settings.json

```json
{
  "go.inferGopath": false,
  "go.testFlags": ["-v", "-race"],
  "go.testTimeout": "30s",
  "go.coverOnSave": true,
  "go.coverageDecorator": {
    "type": "highlight"
  }
}
```

### launch.json

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Launch Package",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${fileDirname}",
      "env": {},
      "args": []
    },
    {
      "name": "Debug Test",
      "type": "go",
      "request": "launch",
      "mode": "test",
      "program": "${fileDirname}",
      "args": ["-test.v"]
    }
  ]
}
```

### tasks.json

```json
{
  "version": "2.0.0",
  "tasks": [
    {
      "label": "go: build current package",
      "type": "shell",
      "command": "go build ./...",
      "group": {
        "kind": "build",
        "isDefault": true
      },
      "problemMatcher": ["$go"]
    },
    {
      "label": "go: test current package",
      "type": "shell",
      "command": "go test -v ./...",
      "group": {
        "kind": "test",
        "isDefault": true
      },
      "problemMatcher": ["$go"]
    },
    {
      "label": "go: lint project",
      "type": "shell",
      "command": "golangci-lint run",
      "problemMatcher": ["$go"]
    }
  ]
}
```

## Step 6: Install Go Tools

After installing the Go extension, it will prompt you to install several Go tools. Make sure to install:

1. Ctrl + Shift + P
2. Search for `Go: Install/update toolbox.`
3. Select all and click OK.

- gopls (language server)
- dlv (debugger)
- golangci-lint (linter)
- goimports (import organization)
- gotests (test generation)

You can also manually install these tools by running the "Go: Install/Update Tools" command from the command palette.

## Step 7: Create a Sample Project Structure

For a new Go project, use this structure:

```sh
project-name/
├── .vscode/
│   ├── settings.json
│   ├── launch.json
│   └── tasks.json
├── cmd/
│   └── main.go
├── internal/
│   └── app/
│       ├── app.go
│       └── app_test.go
├── pkg/
│   └── [shared packages]
├── go.mod
└── go.sum
```

## Step 8: Set Up Go Modules

Initialize a new module:

```bash
go mod init github.com/yourusername/project-name
```

This creates the `go.mod` file which manages dependencies.

## Step 9: Test Your Setup

Create a hello world program:

```go
// cmd/main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go with VSCode!")
}
```

Try these features:

- Format the file (Shift+Alt+F)
- Run the program (F5)
- Set a breakpoint and debug
- Use the Go Test Explorer

Your Go development environment is now ready!
