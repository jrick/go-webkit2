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
// #include <webkit2.h>
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
		{glib.Type(C.webkit_uri_request_get_type()), marshalURIRequest},
		{glib.Type(C.webkit_web_context_get_type()), marshalWebContext},
		{glib.Type(C.webkit_web_view_get_type()), marshalWebView},
		{glib.Type(C.webkit_web_view_group_get_type()), marshalWebViewGroup},
	}
	glib.RegisterGValueMarshalers(tm)
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
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return wrapWebView(obj)
}

// NewWebViewWithContext is a wrapper around webkit_web_view_new_with_context().
func NewWebViewWithContext(context *WebContext) *WebView {
	c := C.webkit_web_view_new_with_context(context.Native())
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return wrapWebView(obj)
}

// NewWebViewWithGroup is a wrapper around webkit_web_view_new_with_group().
func NewWebViewWithGroup(group *WebViewGroup) *WebView {
	c := C.webkit_web_view_new_with_group(group.Native())
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return wrapWebView(obj)
}

// Context is a wrapper around webkit_web_view_get_context().
func (w *WebView) Context() *WebContext {
	c := C.webkit_web_view_get_context(w.Native())
	wc := (*WebContext)(unsafe.Pointer(c))
	// TODO: is this a good idea?
	w.Ref()
	runtime.SetFinalizer(wc, func(_ *WebContext) {
		w.Unref()
	})
	return wc
}

// LoadUri is a wrapper around webkit_web_view_load_uri().
func (w *WebView) LoadUri(uri string) {
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
