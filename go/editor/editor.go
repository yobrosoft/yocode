package editor

import (
	"fmt"
	"path/filepath"
	"strings"
)

// Editor represents a Vim-like code editor instance
type Editor struct {
	window   Window
	filePath string
}

// Config holds configuration options for creating a new editor
type Config struct {
	Width  int
	Height int
}

// DefaultConfig returns a default configuration for the editor
func DefaultConfig() Config {
	return Config{
		Width:  800,
		Height: 600,
	}
}

// New creates a new editor instance for the specified file
func New(filePath string, config Config) (*Editor, error) {
	// Create a window title with the filename
	title := fmt.Sprintf("Yocode - %s", filepath.Base(filePath))

	// Create a platform-specific native window
	window, err := CreateWindow(title, config.Width, config.Height)
	if err != nil {
		return nil, fmt.Errorf("failed to create native window: %w", err)
	}

	// Load the embedded UI assets
	htmlContent, err := Assets.ReadFile("assets/index.html")
	if err != nil {
		return nil, fmt.Errorf("failed to read embedded index.html: %w", err)
	}

	// Load the JavaScript files to inject into the HTML
	mainJS, err := Assets.ReadFile("assets/main.js")
	if err != nil {
		return nil, fmt.Errorf("failed to read embedded main.js: %w", err)
	}

	vimJS, err := Assets.ReadFile("assets/vim.js")
	if err != nil {
		return nil, fmt.Errorf("failed to read embedded vim.js: %w", err)
	}

	// Inject the JavaScript directly into the HTML
	htmlWithJS := injectScriptsIntoHTML(string(htmlContent), map[string]string{
		"main.js": string(mainJS),
		"vim.js":  string(vimJS),
	})

	// Set a dummy base URL since all resources are embedded
	baseURL := "about:blank"

	// Load the HTML content into the WebView
	if err := window.LoadHTML(htmlWithJS, baseURL); err != nil {
		return nil, fmt.Errorf("failed to load HTML content: %w", err)
	}

	return &Editor{
		window:   window,
		filePath: filePath,
	}, nil
}

// Run starts the editor's main loop
func (e *Editor) Run() {
	// Run the platform-specific main loop
	e.window.Run()
}

// injectScriptsIntoHTML injects JavaScript directly into the HTML content
func injectScriptsIntoHTML(html string, scripts map[string]string) string {
	// Simple replacement of script tags with inline scripts
	for scriptName, scriptContent := range scripts {
		scriptTag := `<script src="` + scriptName + `"></script>`
		inlineScript := `<script>` + scriptContent + `</script>`
		html = strings.Replace(html, scriptTag, inlineScript, 1)
	}
	return html
}
