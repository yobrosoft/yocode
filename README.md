# Yocode - Vim-like Code Editor

A lightweight, Vim-like code editor built in Go, rendering a TypeScript/React UI in a native window.

## Features

- Opens one file at a time in a floating window
- Vim keybindings for efficient editing
- TypeScript/React UI rendered via WebView
- Simple command-line usage

## Setup

1. Install dependencies:
   ```bash
   go mod tidy
   ```
2. Build the web UI and binary:
   ```bash
   make build
   ```
3. Run the editor:
   ```bash
   ./yocode /path/to/your/file
   ```

## Development

The editor uses WebView to render a native window with web content. The UI is built with React and TypeScript, utilizing Monaco Editor for the editing experience with Vim mode enabled.
