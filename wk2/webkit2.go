// Copyright (c) 2014 Josh Rickmar.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// Package wk2 provides webkitgtk2 bindings for Go.
package wk2

// #cgo pkg-config: webkit2gtk-3.0
//
// #include <stdint.h>
// #include <stdlib.h>
// #include <string.h>
//
// #include <webkit2/webkit2.h>
//
// #include "webkit2.go.h"
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/conformal/gotk3/glib"
	"github.com/conformal/gotk3/gtk"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums
		{glib.Type(C.webkit_load_event_get_type()), marshalLoadEvent},

		// Objects/Interfaces
		{glib.Type(C.webkit_back_forward_list_get_type()), marshalBackForwardList},
		{glib.Type(C.webkit_back_forward_list_item_get_type()), marshalBackForwardListItem},
		{glib.Type(C.webkit_uri_request_get_type()), marshalURIRequest},
		{glib.Type(C.webkit_web_context_get_type()), marshalWebContext},
		{glib.Type(C.webkit_web_view_get_type()), marshalWebView},
		{glib.Type(C.webkit_web_view_group_get_type()), marshalWebViewGroup},
	}
	glib.RegisterGValueMarshalers(tm)
}

func gbool(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}

func gobool(b C.gboolean) bool {
	if b != 0 {
		return true
	}
	return false
}

//
// Constants
//

// LoadEvent is a representation of WebKit's WebKitLoadEvent.
type LoadEvent int

const (
	LoadStarted    LoadEvent = C.WEBKIT_LOAD_STARTED
	LoadRedirected LoadEvent = C.WEBKIT_LOAD_REDIRECTED
	LoadCommitted  LoadEvent = C.WEBKIT_LOAD_COMMITTED
	LoadFinished   LoadEvent = C.WEBKIT_LOAD_FINISHED
)

func marshalLoadEvent(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return LoadEvent(c), nil
}

//
// WebKitBackForwardList
//

// BackForwardList is a representation of WebKit's WebKitBackForwardList.
type BackForwardList struct {
	*glib.Object
}

func marshalBackForwardList(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapBackForwardList(obj), nil
}

func wrapBackForwardList(obj *glib.Object) *BackForwardList {
	return &BackForwardList{obj}
}

// Native returns a pointer to the underlying WebKitBackForwardList.
func (l *BackForwardList) Native() *C.WebKitBackForwardList {
	if l == nil || l.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(l.GObject)
	return C.toWebKitBackForwardList(p)
}

// Len is a wrapper around webkit_back_forward_list_get_length().
func (l *BackForwardList) Len() uint {
	c := C.webkit_back_forward_list_get_length(l.Native())
	return uint(c)
}

// CurrentItem is a wrapper around webkit_back_forward_list_get_current_item().
func (l *BackForwardList) CurrentItem() *BackForwardListItem {
	c := C.webkit_back_forward_list_get_current_item(l.Native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return wrapBackForwardListItem(obj)
}

// BackItem is a wrapper around webkit_back_forward_list_get_back_item().
func (l *BackForwardList) BackItem() *BackForwardListItem {
	c := C.webkit_back_forward_list_get_back_item(l.Native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return wrapBackForwardListItem(obj)
}

// ForwardItem is a wrapper around webkit_back_forward_list_get_forward_item().
func (l *BackForwardList) ForwardItem() *BackForwardListItem {
	c := C.webkit_back_forward_list_get_forward_item(l.Native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return wrapBackForwardListItem(obj)
}

// NthItem is a wrapper around webkit_back_forward_list_get_nth_item().
func (l *BackForwardList) NthItem(n int) *BackForwardListItem {
	c := C.webkit_back_forward_list_get_nth_item(l.Native(), C.gint(n))
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return wrapBackForwardListItem(obj)
}

// TODO: BackForwardList GList methods.  These probably should be
// returning slices, or make a wrapper data type to dereference and
// wrap around the GList uintptr.

//
// WebKitBackForwardListItem
//

// BackForwardListItem is a representation of WebKit's
// WebKitBackForwardListItem.
type BackForwardListItem struct {
	glib.InitiallyUnowned
}

func marshalBackForwardListItem(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapBackForwardListItem(obj), nil
}

func wrapBackForwardListItem(obj *glib.Object) *BackForwardListItem {
	return &BackForwardListItem{glib.InitiallyUnowned{obj}}
}

// Native returns a pointer to the underlying WebKitBackForwardListItem.
func (item *BackForwardListItem) Native() *C.WebKitBackForwardListItem {
	if item == nil || item.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(item.GObject)
	return C.toWebKitBackForwardListItem(p)
}

//
// WebKitURIRequest
//

// URIRequest is a representation of WebKit's WebKitURIRequest.
type URIRequest struct {
	*glib.Object
}

func marshalURIRequest(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapURIRequest(obj), nil
}

func wrapURIRequest(obj *glib.Object) *URIRequest {
	return &URIRequest{obj}
}

// Native returns a pointer to the underlying WebKitURIRequest.
func (r *URIRequest) Native() *C.WebKitURIRequest {
	if r == nil || r.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(r.GObject)
	return C.toWebKitURIRequest(p)
}

// NewURIRequest is a wrapper around webkit_uri_request_new().
func NewURIRequest(uri string) *URIRequest {
	cstr := C.CString(uri)
	defer C.free(unsafe.Pointer(cstr))
	c := C.webkit_uri_request_new((*C.gchar)(cstr))
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return wrapURIRequest(obj)
}

//
// WebKitWebContext
//

// WebContext is a representation of WebKit's WebKitWebContext.
type WebContext struct {
	*glib.Object
}

func marshalWebContext(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapWebContext(obj), nil
}

func wrapWebContext(obj *glib.Object) *WebContext {
	return &WebContext{obj}
}

// Native returns a pointer to the underlying WebKitWebContext.
func (w *WebContext) Native() *C.WebKitWebContext {
	if w == nil || w.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(w.GObject)
	return C.toWebKitWebContext(p)
}

// DefaultWebContext is a wrapper around webkit_web_context_get_default().
func DefaultWebContext() *WebContext {
	c := C.webkit_web_context_get_default()
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return wrapWebContext(obj)
}

//
// WebKitWebView
//

// WebView is a representation of WebKit's WebKitWebView.
type WebView struct {
	gtk.Widget
}

func marshalWebView(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapWebView(obj), nil
}

func wrapWebView(obj *glib.Object) *WebView {
	return &WebView{gtk.Widget{glib.InitiallyUnowned{obj}}}
}

// Native returns a pointer to the underlying WebKitWebView.
func (w *WebView) Native() *C.WebKitWebView {
	if w == nil || w.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(w.GObject)
	return C.toWebKitWebView(p)
}

// NewWebView is a wrapper around webkit_web_view_new().
func NewWebView() *WebView {
	c := C.webkit_web_view_new()
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return wrapWebView(obj)
}

// NewWebViewWithContext is a wrapper around webkit_web_view_new_with_context().
func NewWebViewWithContext(context *WebContext) *WebView {
	c := C.webkit_web_view_new_with_context(context.Native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return wrapWebView(obj)
}

// NewWebViewWithGroup is a wrapper around webkit_web_view_new_with_group().
func NewWebViewWithGroup(group *WebViewGroup) *WebView {
	c := C.webkit_web_view_new_with_group(group.Native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return wrapWebView(obj)
}

// Context is a wrapper around webkit_web_view_get_context().
func (w *WebView) Context() *WebContext {
	c := C.webkit_web_view_get_context(w.Native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return wrapWebContext(obj)
}

// LoadURI is a wrapper around webkit_web_view_load_uri().
func (w *WebView) LoadURI(uri string) {
	cstr := C.CString(uri)
	defer C.free(unsafe.Pointer(cstr))
	C.webkit_web_view_load_uri(w.Native(), (*C.gchar)(cstr))
}

// LoadHTML is a wrapper around webkit_web_view_load_html().
func (w *WebView) LoadHTML(content, baseURI string) {
	cContent := C.CString(content)
	cBaseURI := C.CString(baseURI)
	defer C.free(unsafe.Pointer(cContent))
	defer C.free(unsafe.Pointer(cBaseURI))
	C.webkit_web_view_load_html(w.Native(), (*C.gchar)(cContent),
		(*C.gchar)(cBaseURI))
}

// LoadAlternateHTML is a wrapper around webkit_web_view_load_alternate_html().
func (w *WebView) LoadAlternateHTML(content, contentURI, baseURI string) {
	cContent := C.CString(content)
	cContentURI := C.CString(contentURI)
	cBaseURI := C.CString(baseURI)
	defer C.free(unsafe.Pointer(cContent))
	defer C.free(unsafe.Pointer(cContentURI))
	defer C.free(unsafe.Pointer(cBaseURI))
	C.webkit_web_view_load_alternate_html(w.Native(), (*C.gchar)(cContent),
		(*C.gchar)(cContentURI), (*C.gchar)(cBaseURI))
}

// LoadPlainText is a wrapper around webkit_web_view_load_plain_text().
func (w *WebView) LoadPlainText(plainText string) {
	cstr := C.CString(plainText)
	defer C.free(unsafe.Pointer(cstr))
	C.webkit_web_view_load_plain_text(w.Native(), (*C.gchar)(cstr))
}

// LoadRequest is a wrapper around webkit_web_view_load_request().
func (w *WebView) LoadRequest(request *URIRequest) {
	C.webkit_web_view_load_request(w.Native(), request.Native())
}

// CanGoBack is a wrapper around webkit_web_view_can_go_back().
func (w *WebView) CanGoBack() bool {
	c := C.webkit_web_view_can_go_back(w.Native())
	return gobool(c)
}

// GoBack is a wrapper around webkit_web_view_go_back().
func (w *WebView) GoBack() {
	C.webkit_web_view_go_back(w.Native())
}

// GoForward is a wrapper around webkit_web_view_can_go_forward().
func (w *WebView) GoForward() {
	C.webkit_web_view_go_forward(w.Native())
}

// Title is a wrapper around webkit_web_view_get_title().
func (w *WebView) Title() string {
	c := C.webkit_web_view_get_title(w.Native())
	return C.GoString((*C.char)(c))
}

// PageID is a wrapper around webkit_web_view_get_page_id().
func (w *WebView) PageID() uint64 {
	c := C.webkit_web_view_get_page_id(w.Native())
	return uint64(c)
}

// Reload is a wrapper around webkit_web_view_reload().
func (w *WebView) Reload() {
	C.webkit_web_view_reload(w.Native())
}

// ReloadBypassCache is a wrapper around webkit_web_view_reload_bypass_cache().
func (w *WebView) ReloadBypassCache() {
	C.webkit_web_view_reload_bypass_cache(w.Native())
}

// StopLoading is a wrapper around webkit_web_view_stop_loading().
func (w *WebView) StopLoading() {
	C.webkit_web_view_stop_loading(w.Native())
}

// IsLoading is a wrapper around webkit_web_view_is_loading().
func (w *WebView) IsLoading() bool {
	c := C.webkit_web_view_is_loading(w.Native())
	return gobool(c)
}

// EstimatedLoadProgress is a wrapper around
// webkit_web_view_get_estimated_load_progress().
func (w *WebView) EstimatedLoadProgress() float64 {
	c := C.webkit_web_view_get_estimated_load_progress(w.Native())
	return float64(c)
}

// CustomCharset is a wrapper around webkit_web_view_get_custom_charset().
func (w *WebView) CustomCharset() string {
	c := C.webkit_web_view_get_custom_charset(w.Native())
	return C.GoString((*C.char)(c))
}

// TODO: BackForwardList
// TODO: GoToBackForwardListItem

// URI is a wrapper around webkit_web_view_get_uri().
func (w *WebView) URI() string {
	c := C.webkit_web_view_get_uri(w.Native())
	return C.GoString((*C.char)(c))
}

//
// WebKitWebViewGroup
//

// WebViewGroup is a representation of WebKit's WebKitWebViewGroup.
type WebViewGroup struct {
	*glib.Object
}

func marshalWebViewGroup(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapWebViewGroup(obj), nil
}

func wrapWebViewGroup(obj *glib.Object) *WebViewGroup {
	return &WebViewGroup{obj}
}

// Native returns a pointer to the underlying WebKitWebViewGroup.
func (w *WebViewGroup) Native() *C.WebKitWebViewGroup {
	if w == nil || w.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(w.GObject)
	return C.toWebKitWebViewGroup(p)
}
