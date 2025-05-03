package editor

// Window represents a platform-specific window implementation
type Window interface {
	// LoadHTML loads HTML content into the window
	LoadHTML(html, baseURL string) error

	// Run starts the main application loop
	Run()
}

// CreateWindow creates a platform-specific window
func CreateWindow(title string, width, height int) (Window, error) {
	return createNativeWindow(title, width, height)
}
