package editor

import (
	"fmt"
	"path/filepath"
	"strings"
	"unsafe"
)

// #cgo darwin CFLAGS: -x objective-c
// #cgo darwin LDFLAGS: -framework Cocoa -framework WebKit
// #include <stdlib.h>
// #import <Cocoa/Cocoa.h>
// #import <WebKit/WebKit.h>
//
// void* createWindow(const char* title, int width, int height) {
//     NSAutoreleasePool *pool = [[NSAutoreleasePool alloc] init];
//     
//     [NSApplication sharedApplication];
//     [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
//     
//     // Create the menu bar
//     NSMenu *menubar = [NSMenu new];
//     NSMenuItem *appMenuItem = [NSMenuItem new];
//     [menubar addItem:appMenuItem];
//     [NSApp setMainMenu:menubar];
//     
//     // Create the application menu
//     NSMenu *appMenu = [NSMenu new];
//     NSMenuItem *quitMenuItem = [[NSMenuItem alloc] initWithTitle:@"Quit"
//                                                          action:@selector(terminate:)
//                                                   keyEquivalent:@"q"];
//     [quitMenuItem setKeyEquivalentModifierMask:NSEventModifierFlagCommand];
//     [appMenu addItem:quitMenuItem];
//     [appMenuItem setSubmenu:appMenu];
//     
//     NSWindow* window = [[NSWindow alloc] 
//         initWithContentRect:NSMakeRect(0, 0, width, height)
//         styleMask:NSWindowStyleMaskTitled|NSWindowStyleMaskClosable|NSWindowStyleMaskResizable|NSWindowStyleMaskMiniaturizable
//         backing:NSBackingStoreBuffered
//         defer:NO];
//     
//     NSString* nsTitle = [NSString stringWithUTF8String:title];
//     [window setTitle:nsTitle];
//     [window center];
//     
//     WKWebViewConfiguration *config = [[WKWebViewConfiguration alloc] init];
//     WKWebView *webView = [[WKWebView alloc] initWithFrame:NSMakeRect(0, 0, width, height) configuration:config];
//     
//     [window setContentView:webView];
//     [window makeKeyAndOrderFront:nil];
//     
//     [NSApp activateIgnoringOtherApps:YES];
//     
//     return (void*)webView;
// }
//
// void loadHTML(void* webView, const char* html, const char* baseURL) {
//     WKWebView* view = (WKWebView*)webView;
//     NSString* nsHTML = [NSString stringWithUTF8String:html];
//     NSString* nsBaseURL = [NSString stringWithUTF8String:baseURL];
//     NSURL* url = [NSURL URLWithString:nsBaseURL];
//     [view loadHTMLString:nsHTML baseURL:url];
// }
//
// void runApp() {
//     [NSApp run];
// }
import "C"

// Editor represents a Vim-like code editor instance
type Editor struct {
	webView unsafe.Pointer
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
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	// Create a window with a WebView and menu
	webView := C.createWindow(cTitle, C.int(config.Width), C.int(config.Height))

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

	// Convert to C string
	cHTML := C.CString(htmlWithJS)
	defer C.free(unsafe.Pointer(cHTML))

	// Set a dummy base URL since all resources are embedded
	cBaseURL := C.CString("about:blank")
	defer C.free(unsafe.Pointer(cBaseURL))

	// Load the HTML
	C.loadHTML(webView, cHTML, cBaseURL)

	return &Editor{
		webView: webView,
		filePath: filePath,
	}, nil
}

// Run starts the editor's main loop
func (e *Editor) Run() {
	// Run the application
	C.runApp()
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
