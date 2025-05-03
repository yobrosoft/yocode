//go:build darwin

package editor

import (
	"fmt"
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

// DarwinWindow is the macOS implementation of the Window interface
type DarwinWindow struct {
	webView unsafe.Pointer
}

// createNativeWindow creates a native window for Linux
func createNativeWindow(title string, width, height int) (Window, error) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	// Create a window with a WebView
	browser := C.createWindow(cTitle, C.int(width), C.int(height))
	if browser == nil {
		return nil, fmt.Errorf("failed to create macOS window")
	}

	return &DarwinWindow{webView: browser}, nil
}

// LoadHTML loads HTML content into the WebView
func (w *DarwinWindow) LoadHTML(html, baseURL string) error {
	cHTML := C.CString(html)
	defer C.free(unsafe.Pointer(cHTML))

	cBaseURL := C.CString(baseURL)
	defer C.free(unsafe.Pointer(cBaseURL))

	C.loadHTML(w.webView, cHTML, cBaseURL)
	return nil
}

// Run starts the main application loop
func (w *DarwinWindow) Run() {
	C.runApp()
}
