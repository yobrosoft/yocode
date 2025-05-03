//go:build linux

package editor

import (
	"fmt"
	"unsafe"
)

// #cgo linux pkg-config: gtk+-3.0 webkit2gtk-4.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
// #include <webkit2/webkit2.h>
//
// typedef struct {
//     GtkWidget *window;
//     GtkWidget *webview;
// } BrowserWindow;
//
// void* createWindow(const char* title, int width, int height) {
//     if (!gtk_init_check(NULL, NULL)) {
//         return NULL;
//     }
//
//     BrowserWindow* browser = (BrowserWindow*)malloc(sizeof(BrowserWindow));
//
//     browser->window = gtk_window_new(GTK_WINDOW_TOPLEVEL);
//     gtk_window_set_title(GTK_WINDOW(browser->window), title);
//     gtk_window_set_default_size(GTK_WINDOW(browser->window), width, height);
//     gtk_window_set_position(GTK_WINDOW(browser->window), GTK_WIN_POS_CENTER);
//
//     // Create a webview
//     browser->webview = webkit_web_view_new();
//     gtk_container_add(GTK_CONTAINER(browser->window), browser->webview);
//
//     // Connect signals
//     g_signal_connect(browser->window, "destroy", G_CALLBACK(gtk_main_quit), NULL);
//
//     // Show all widgets
//     gtk_widget_show_all(browser->window);
//
//     return browser;
// }
//
// void loadHTML(void* browserPtr, const char* html, const char* baseURL) {
//     BrowserWindow* browser = (BrowserWindow*)browserPtr;
//     webkit_web_view_load_html(WEBKIT_WEB_VIEW(browser->webview), html, baseURL);
// }
//
// void runApp() {
//     gtk_main();
// }
//
// void freeWindow(void* browserPtr) {
//     if (browserPtr != NULL) {
//         free(browserPtr);
//     }
// }
import "C"

type LinuxWindow struct {
	webView unsafe.Pointer
}

// createNativeWindow creates a native window for Linux
func createNativeWindow(title string, width, height int) (Window, error) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	// Create a window with a WebView
	browser := C.createWindow(cTitle, C.int(width), C.int(height))
	if browser == nil {
		return nil, fmt.Errorf("failed to create GTK window")
	}

	return &LinuxWindow{webView: browser}, nil
}

// loadHTMLContent loads HTML content into the WebView
func (w *LinuxWindow) LoadHTML(html, baseURL string) error {
	cHTML := C.CString(html)
	defer C.free(unsafe.Pointer(cHTML))

	cBaseURL := C.CString(baseURL)
	defer C.free(unsafe.Pointer(cBaseURL))

	C.loadHTML(w.webView, cHTML, cBaseURL)
	return nil
}

// runMainLoop runs the main application loop
func (w *LinuxWindow) Run() {
	C.runApp()
}
