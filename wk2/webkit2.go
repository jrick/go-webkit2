// Copyright (c) 2014 Josh Rickmar.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// Package wk2 provides WebKit2GTK+ bindings for Go.
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
		{glib.Type(C.webkit_cache_model_get_type()), marshalCacheModel},
		{glib.Type(C.webkit_load_event_get_type()), marshalLoadEvent},
		{glib.Type(C.webkit_process_model_get_type()), marshalProcessModel},
		{glib.Type(C.webkit_tls_errors_policy_get_type()), marshalTLSErrorsPolicy},

		// Objects/Interfaces
		{glib.Type(C.webkit_back_forward_list_get_type()), marshalBackForwardList},
		{glib.Type(C.webkit_back_forward_list_item_get_type()), marshalBackForwardListItem},
		{glib.Type(C.webkit_cookie_manager_get_type()), marshalCookieManager},
		{glib.Type(C.webkit_download_get_type()), marshalDownload},
		{glib.Type(C.webkit_favicon_database_get_type()), marshalFaviconDatabase},
		{glib.Type(C.webkit_security_manager_get_type()), marshalSecurityManager},
		{glib.Type(C.webkit_uri_request_get_type()), marshalURIRequest},
		{glib.Type(C.webkit_web_context_get_type()), marshalWebContext},
		{glib.Type(C.webkit_web_view_get_type()), marshalWebView},
		{glib.Type(C.webkit_web_view_group_get_type()), marshalWebViewGroup},

		// Boxed
		{glib.Type(C.webkit_certificate_info_get_type()), marshalCertificateInfo},
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

// CacheModel is a representation of WebKit2GTK's WebKitCacheModel.
type CacheModel int

// These constants define the values determining a WebContext's cache model.
const (
	CacheModelDocumentViewer  CacheModel = C.WEBKIT_CACHE_MODEL_DOCUMENT_VIEWER
	CacheModelWebBrowser      CacheModel = C.WEBKIT_CACHE_MODEL_WEB_BROWSER
	CacheModelDocumentBrowser CacheModel = C.WEBKIT_CACHE_MODEL_DOCUMENT_BROWSER
)

func marshalCacheModel(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return CacheModel(c), nil
}

// LoadEvent is a representation of WebKit2GTK+'s WebKitLoadEvent.
type LoadEvent int

// These constants define the different events that happen during a
// WebView's load operation.
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

// ProcessModel is a representation of WebKit2GTK+'s WebKitProcessModel.
type ProcessModel int

// These constants are used to determine the WebContext process model.
const (
	ProcessModelSharedSecondaryProcess     ProcessModel = C.WEBKIT_PROCESS_MODEL_SHARED_SECONDARY_PROCESS
	ProcessModelMultipleSecondaryProcesses ProcessModel = C.WEBKIT_PROCESS_MODEL_MULTIPLE_SECONDARY_PROCESSES
)

func marshalProcessModel(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return ProcessModel(c), nil
}

// TLSErrorsPolicy is a representation of WebKit2GTK+'s WebKitTLSErrorsPolicy.
type TLSErrorsPolicy int

// These constants define the TLS errors policy.
const (
	TLSErrorsPolicyIgnore TLSErrorsPolicy = C.WEBKIT_TLS_ERRORS_POLICY_IGNORE
	TLSErrorsPolicyFail   TLSErrorsPolicy = C.WEBKIT_TLS_ERRORS_POLICY_FAIL
)

func marshalTLSErrorsPolicy(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return TLSErrorsPolicy(c), nil
}

//
// WebKitBackForwardList
//

// BackForwardList is a representation of WebKit2GTK+'s WebKitBackForwardList.
type BackForwardList struct {
	*glib.Object
}

func marshalBackForwardList(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	return wrapBackForwardList(obj), nil
}

func wrapBackForwardList(obj *glib.Object) *BackForwardList {
	return &BackForwardList{obj}
}

// native returns a pointer to the underlying WebKitBackForwardList.
func (l *BackForwardList) native() *C.WebKitBackForwardList {
	if l == nil || l.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(l.GObject)
	return C.toWebKitBackForwardList(p)
}

// Len is a wrapper around webkit_back_forward_list_get_length().
func (l *BackForwardList) Len() uint {
	c := C.webkit_back_forward_list_get_length(l.native())
	return uint(c)
}

// CurrentItem is a wrapper around webkit_back_forward_list_get_current_item().
func (l *BackForwardList) CurrentItem() *BackForwardListItem {
	c := C.webkit_back_forward_list_get_current_item(l.native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return wrapBackForwardListItem(obj)
}

// BackItem is a wrapper around webkit_back_forward_list_get_back_item().
func (l *BackForwardList) BackItem() *BackForwardListItem {
	c := C.webkit_back_forward_list_get_back_item(l.native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return wrapBackForwardListItem(obj)
}

// ForwardItem is a wrapper around webkit_back_forward_list_get_forward_item().
func (l *BackForwardList) ForwardItem() *BackForwardListItem {
	c := C.webkit_back_forward_list_get_forward_item(l.native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return wrapBackForwardListItem(obj)
}

// NthItem is a wrapper around webkit_back_forward_list_get_nth_item().
func (l *BackForwardList) NthItem(n int) *BackForwardListItem {
	c := C.webkit_back_forward_list_get_nth_item(l.native(), C.gint(n))
	if c == nil {
		return nil
	}
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
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

// BackForwardListItem is a representation of WebKit2GTK+'s
// WebKitBackForwardListItem.
type BackForwardListItem struct {
	glib.InitiallyUnowned
}

func marshalBackForwardListItem(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	return wrapBackForwardListItem(obj), nil
}

func wrapBackForwardListItem(obj *glib.Object) *BackForwardListItem {
	return &BackForwardListItem{glib.InitiallyUnowned{Object: obj}}
}

// native returns a pointer to the underlying WebKitBackForwardListItem.
func (item *BackForwardListItem) native() *C.WebKitBackForwardListItem {
	if item == nil || item.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(item.GObject)
	return C.toWebKitBackForwardListItem(p)
}

//
// WebKitCertificateInfo
//

// CertificateInfo is a representation of WebKit2GTK+'s
// WebKitCertificateInfo.
type CertificateInfo struct {
	info *C.WebKitCertificateInfo
}

func marshalCertificateInfo(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	info := (*C.WebKitCertificateInfo)(unsafe.Pointer(c))
	wrapped := wrapCertificateInfo(info)
	runtime.SetFinalizer(wrapped, (*CertificateInfo).free)
	return wrapped, nil
}

func wrapCertificateInfo(info *C.WebKitCertificateInfo) *CertificateInfo {
	return &CertificateInfo{info}
}

// native returns a pointer to the underlying WebKitCertificateInfo.
func (i *CertificateInfo) native() *C.WebKitCertificateInfo {
	if i == nil {
		return nil
	}
	return i.info
}

// Native returns a pointer to the underlying WebKitCertificateInfo.
func (i *CertificateInfo) Native() uintptr {
	return uintptr(unsafe.Pointer(i.native()))
}

// free is a wrapper around webkit_certificate_info_free().
func (i *CertificateInfo) free() {
	C.webkit_certificate_info_free(i.native())
}

//
// WebKitCookieManager
//

// CookieManager is a representation of WebKit2GTK+'s WebKitCookieManager.
type CookieManager struct {
	*glib.Object
}

func marshalCookieManager(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	return wrapCookieManager(obj), nil
}

func wrapCookieManager(obj *glib.Object) *CookieManager {
	return &CookieManager{obj}
}

// native returns a pointer to the underlying WebKitCookieManager.
func (m *CookieManager) native() *C.WebKitCookieManager {
	if m == nil || m.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(m.GObject)
	return C.toWebKitCookieManager(p)
}

//
// WebKitDownload
//

// Download is a representation of WebKit2GTK+'s WebKitDownload.
type Download struct {
	*glib.Object
}

func marshalDownload(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	return wrapDownload(obj), nil
}

func wrapDownload(obj *glib.Object) *Download {
	return &Download{obj}
}

// native returns a pointer to the underlying WebKitDownload.
func (d *Download) native() *C.WebKitDownload {
	if d == nil || d.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(d.GObject)
	return C.toWebKitDownload(p)
}

//
// WebKitFaviconDatabase
//

// FaviconDatabase is a representation of WebKit2GTK+'s WebKitFaviconDatabase.
type FaviconDatabase struct {
	*glib.Object
}

func marshalFaviconDatabase(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	return wrapFaviconDatabase(obj), nil
}

func wrapFaviconDatabase(obj *glib.Object) *FaviconDatabase {
	return &FaviconDatabase{Object: obj}
}

// native returns a pointer to the underlying WebKitFaviconDatabase.
func (d *FaviconDatabase) native() *C.WebKitFaviconDatabase {
	if d == nil || d.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(d.GObject)
	return C.toWebKitFaviconDatabase(p)
}

//
// WebKitSecurityManager
//

// SecurityManager is a representation of WebKit2GTK+'s WebKitSecurityManager.
type SecurityManager struct {
	*glib.Object
}

func marshalSecurityManager(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	return wrapSecurityManager(obj), nil
}

func wrapSecurityManager(obj *glib.Object) *SecurityManager {
	return &SecurityManager{obj}
}

// native returns a pointer to the underlying WebKitSecurityManager.
func (s *SecurityManager) native() *C.WebKitSecurityManager {
	if s == nil || s.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(s.GObject)
	return C.toWebKitSecurityManager(p)
}

//
// WebKitURIRequest
//

// URIRequest is a representation of WebKit2GTK+'s WebKitURIRequest.
type URIRequest struct {
	*glib.Object
}

func marshalURIRequest(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	return wrapURIRequest(obj), nil
}

func wrapURIRequest(obj *glib.Object) *URIRequest {
	return &URIRequest{obj}
}

// native returns a pointer to the underlying WebKitURIRequest.
func (r *URIRequest) native() *C.WebKitURIRequest {
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
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return wrapURIRequest(obj)
}

//
// WebKitWebContext
//

// WebContext is a representation of WebKit2GTK+'s WebKitWebContext.
type WebContext struct {
	*glib.Object
}

func marshalWebContext(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	return wrapWebContext(obj), nil
}

func wrapWebContext(obj *glib.Object) *WebContext {
	return &WebContext{obj}
}

// native returns a pointer to the underlying WebKitWebContext.
func (w *WebContext) native() *C.WebKitWebContext {
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
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return wrapWebContext(obj)
}

// CacheModel is a wrapper around webkit_web_context_get_cache_model().
func (w *WebContext) CacheModel() CacheModel {
	c := C.webkit_web_context_get_cache_model(w.native())
	return CacheModel(c)
}

// SetCacheModel is a wrapper around webkit_web_context_set_cache_model().
func (w *WebContext) SetCacheModel(cm CacheModel) {
	C.webkit_web_context_set_cache_model(w.native(), C.WebKitCacheModel(cm))
}

// ClearCache is a wrapper around webkit_web_context_clear_cache().
func (w *WebContext) ClearCache() {
	C.webkit_web_context_clear_cache(w.native())
}

// DownloadURI is a wrapper around webkit_web_context_download_uri().
func (w *WebContext) DownloadURI(uri string) *Download {
	cstr := C.CString(uri)
	defer C.free(unsafe.Pointer(cstr))
	c := C.webkit_web_context_download_uri(w.native(), (*C.gchar)(cstr))
	if c == nil {
		return nil
	}
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return wrapDownload(obj)
}

// CookieManager is a wrapper around webkit_web_context_get_cookie_manager().
func (w *WebContext) CookieManager() *CookieManager {
	c := C.webkit_web_context_get_cookie_manager(w.native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return wrapCookieManager(obj)
}

// FaviconDatabase is a wrapper around webkit_web_context_get_favicon_database().
func (w *WebContext) FaviconDatabase() *FaviconDatabase {
	c := C.webkit_web_context_get_favicon_database(w.native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return wrapFaviconDatabase(obj)
}

// SetFaviconDatabaseDirectory is a wrapper around
// webkit_web_context_set_favicon_database_directory().
func (w *WebContext) SetFaviconDatabaseDirectory(dir string) {
	cstr := C.CString(dir)
	defer C.free(unsafe.Pointer(cstr))
	C.webkit_web_context_set_favicon_database_directory(w.native(),
		(*C.gchar)(cstr))
}

// FaviconDatabaseDirectory is a wrapper around
// webkit_web_context_get_favicon_database_directory().
func (w *WebContext) FaviconDatabaseDirectory() string {
	c := C.webkit_web_context_get_favicon_database_directory(w.native())
	return C.GoString((*C.char)(c))
}

// SecurityManager is a wrapper around webkit_web_context_get_security_manager().
func (w *WebContext) SecurityManager() *SecurityManager {
	c := C.webkit_web_context_get_security_manager(w.native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return wrapSecurityManager(obj)
}

// SetAdditionalPluginsDirectory is a wrapper around
// webkit_web_context_set_additional_plugins_directory().
func (w *WebContext) SetAdditionalPluginsDirectory(dir string) {
	cstr := C.CString(dir)
	defer C.free(unsafe.Pointer(cstr))
	C.webkit_web_context_set_additional_plugins_directory(w.native(),
		(*C.gchar)(cstr))
}

// TODO: webkit_web_context_get_plugins
// TODO: webkit_web_context_get_plugins_finish

// SpellCheckingEnabled is a wrapper around
// webkit_web_context_get_spell_checking_enabled().
func (w *WebContext) SpellCheckingEnabled() bool {
	c := C.webkit_web_context_get_spell_checking_enabled(w.native())
	return gobool(c)
}

// SetSpellCheckingEnabled is a wrapper around
// webkit_web_context_set_spell_checking_enabled().
func (w *WebContext) SetSpellCheckingEnabled(enabled bool) {
	C.webkit_web_context_set_spell_checking_enabled(w.native(), gbool(enabled))
}

// SpellCheckingLanguages is a wrapper around
// webkit_web_context_get_spell_checking_languages().
func (w *WebContext) SpellCheckingLanguages() []string {
	c := C.webkit_web_context_get_spell_checking_languages(w.native())
	if c == nil {
		return nil
	}

	var languages []string
	for i := 0; C.indexGCharArray(c, C.int(i)) != nil; i++ {
		cstr := C.indexGCharArray(c, C.int(i))
		languages = append(languages, C.GoString((*C.char)(cstr)))
	}
	return languages
}

// SetSpellCheckingLanguages is a wrapper around
// webkit_web_context_set_spell_checking_languages().
func (w *WebContext) SetSpellCheckingLanguages(languages []string) {
	cLanguages := C.allocGCharArray((C.size_t)(len(languages)))
	for i, s := range languages {
		cstr := C.CString(s)
		C.indexGCharArraySet(cLanguages, C.int(i), (*C.gchar)(cstr))
	}
	defer C.freeGCharArray(cLanguages)

	C.webkit_web_context_set_spell_checking_languages(w.native(), cLanguages)
}

// SetPreferredLanguages is a wrapper around
// webkit_web_context_set_preferred_languages().
func (w *WebContext) SetPreferredLanguages(languages []string) {
	cLanguages := C.allocGCharArray((C.size_t)(len(languages)))
	for i, s := range languages {
		cstr := C.CString(s)
		C.indexGCharArraySet(cLanguages, C.int(i), (*C.gchar)(cstr))
	}
	defer C.freeGCharArray(cLanguages)

	C.webkit_web_context_set_preferred_languages(w.native(), cLanguages)
}

// SetTLSErrorsPolicy is a wrapper around webkit_web_context_set_tls_errors_policy().
func (w *WebContext) SetTLSErrorsPolicy(policy TLSErrorsPolicy) {
	C.webkit_web_context_set_tls_errors_policy(w.native(),
		C.WebKitTLSErrorsPolicy(policy))
}

// TLSErrorsPolicy is a wrapper around webkit_web_context_get_tls_errors_policy().
func (w *WebContext) TLSErrorsPolicy() TLSErrorsPolicy {
	c := C.webkit_web_context_get_tls_errors_policy(w.native())
	return TLSErrorsPolicy(c)
}

// SetWebExtensionsDirectory is a wrapper around
// webkit_web_context_set_web_extensions_directory().
func (w *WebContext) SetWebExtensionsDirectory(dir string) {
	cstr := C.CString(dir)
	defer C.free(unsafe.Pointer(cstr))
	C.webkit_web_context_set_web_extensions_directory(w.native(),
		(*C.gchar)(cstr))
}

// TODO: webkit_web_context_set_web_extensions_initialization_user_data

// PrefetchDNS is a wrapper around webkit_web_context_prefetch_dns().
func (w *WebContext) PrefetchDNS(hostname string) {
	cstr := C.CString(hostname)
	defer C.free(unsafe.Pointer(cstr))
	C.webkit_web_context_prefetch_dns(w.native(), (*C.gchar)(cstr))
}

// SetDiskCacheDirectory is a wrapper around
// webkit_web_context_set_disk_cache_directory().
func (w *WebContext) SetDiskCacheDirectory(dir string) {
	cstr := C.CString(dir)
	defer C.free(unsafe.Pointer(cstr))
	C.webkit_web_context_set_disk_cache_directory(w.native(), (*C.gchar)(cstr))
}

// AllowTLSCertificateForHost is a wrapper around
// webkit_web_context_allow_tls_certificate_for_host().
func (w *WebContext) AllowTLSCertificateForHost(info *CertificateInfo, host string) {
	cstr := C.CString(host)
	defer C.free(unsafe.Pointer(cstr))
	C.webkit_web_context_allow_tls_certificate_for_host(w.native(),
		info.native(), (*C.gchar)(cstr))
}

// ProcessModel is a wrapper around webkit_web_context_get_process_model().
func (w *WebContext) ProcessModel() ProcessModel {
	c := C.webkit_web_context_get_process_model(w.native())
	return ProcessModel(c)
}

// SetProcessModel is a wrapper around webkit_web_context_set_process_model().
func (w *WebContext) SetProcessModel(model ProcessModel) {
	C.webkit_web_context_set_process_model(w.native(),
		C.WebKitProcessModel(model))
}

// TODO: webkit_web_context_register_uri_scheme

//
// WebKitWebView
//

// WebView is a representation of WebKit2GTK+'s WebKitWebView.
type WebView struct {
	gtk.Widget
}

func marshalWebView(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	return wrapWebView(obj), nil
}

func wrapWebView(obj *glib.Object) *WebView {
	return &WebView{gtk.Widget{InitiallyUnowned: glib.InitiallyUnowned{Object: obj}}}
}

// native returns a pointer to the underlying WebKitWebView.
func (w *WebView) native() *C.WebKitWebView {
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
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return wrapWebView(obj)
}

// NewWebViewWithContext is a wrapper around webkit_web_view_new_with_context().
func NewWebViewWithContext(context *WebContext) *WebView {
	c := C.webkit_web_view_new_with_context(context.native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return wrapWebView(obj)
}

// NewWebViewWithGroup is a wrapper around webkit_web_view_new_with_group().
func NewWebViewWithGroup(group *WebViewGroup) *WebView {
	c := C.webkit_web_view_new_with_group(group.native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return wrapWebView(obj)
}

// Context is a wrapper around webkit_web_view_get_context().
func (w *WebView) Context() *WebContext {
	c := C.webkit_web_view_get_context(w.native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return wrapWebContext(obj)
}

// LoadURI is a wrapper around webkit_web_view_load_uri().
func (w *WebView) LoadURI(uri string) {
	cstr := C.CString(uri)
	defer C.free(unsafe.Pointer(cstr))
	C.webkit_web_view_load_uri(w.native(), (*C.gchar)(cstr))
}

// LoadHTML is a wrapper around webkit_web_view_load_html().
func (w *WebView) LoadHTML(content, baseURI string) {
	cContent := C.CString(content)
	cBaseURI := C.CString(baseURI)
	defer C.free(unsafe.Pointer(cContent))
	defer C.free(unsafe.Pointer(cBaseURI))
	C.webkit_web_view_load_html(w.native(), (*C.gchar)(cContent),
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
	C.webkit_web_view_load_alternate_html(w.native(), (*C.gchar)(cContent),
		(*C.gchar)(cContentURI), (*C.gchar)(cBaseURI))
}

// LoadPlainText is a wrapper around webkit_web_view_load_plain_text().
func (w *WebView) LoadPlainText(plainText string) {
	cstr := C.CString(plainText)
	defer C.free(unsafe.Pointer(cstr))
	C.webkit_web_view_load_plain_text(w.native(), (*C.gchar)(cstr))
}

// LoadRequest is a wrapper around webkit_web_view_load_request().
func (w *WebView) LoadRequest(request *URIRequest) {
	C.webkit_web_view_load_request(w.native(), request.native())
}

// CanGoBack is a wrapper around webkit_web_view_can_go_back().
func (w *WebView) CanGoBack() bool {
	c := C.webkit_web_view_can_go_back(w.native())
	return gobool(c)
}

// GoBack is a wrapper around webkit_web_view_go_back().
func (w *WebView) GoBack() {
	C.webkit_web_view_go_back(w.native())
}

// GoForward is a wrapper around webkit_web_view_can_go_forward().
func (w *WebView) GoForward() {
	C.webkit_web_view_go_forward(w.native())
}

// Title is a wrapper around webkit_web_view_get_title().
func (w *WebView) Title() string {
	c := C.webkit_web_view_get_title(w.native())
	return C.GoString((*C.char)(c))
}

// PageID is a wrapper around webkit_web_view_get_page_id().
func (w *WebView) PageID() uint64 {
	c := C.webkit_web_view_get_page_id(w.native())
	return uint64(c)
}

// Reload is a wrapper around webkit_web_view_reload().
func (w *WebView) Reload() {
	C.webkit_web_view_reload(w.native())
}

// ReloadBypassCache is a wrapper around webkit_web_view_reload_bypass_cache().
func (w *WebView) ReloadBypassCache() {
	C.webkit_web_view_reload_bypass_cache(w.native())
}

// StopLoading is a wrapper around webkit_web_view_stop_loading().
func (w *WebView) StopLoading() {
	C.webkit_web_view_stop_loading(w.native())
}

// IsLoading is a wrapper around webkit_web_view_is_loading().
func (w *WebView) IsLoading() bool {
	c := C.webkit_web_view_is_loading(w.native())
	return gobool(c)
}

// EstimatedLoadProgress is a wrapper around
// webkit_web_view_get_estimated_load_progress().
func (w *WebView) EstimatedLoadProgress() float64 {
	c := C.webkit_web_view_get_estimated_load_progress(w.native())
	return float64(c)
}

// CustomCharset is a wrapper around webkit_web_view_get_custom_charset().
func (w *WebView) CustomCharset() string {
	c := C.webkit_web_view_get_custom_charset(w.native())
	return C.GoString((*C.char)(c))
}

// TODO: BackForwardList
// TODO: GoToBackForwardListItem

// URI is a wrapper around webkit_web_view_get_uri().
func (w *WebView) URI() string {
	c := C.webkit_web_view_get_uri(w.native())
	return C.GoString((*C.char)(c))
}

//
// WebKitWebViewGroup
//

// WebViewGroup is a representation of WebKit2GTK+'s WebKitWebViewGroup.
type WebViewGroup struct {
	*glib.Object
}

func marshalWebViewGroup(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	return wrapWebViewGroup(obj), nil
}

func wrapWebViewGroup(obj *glib.Object) *WebViewGroup {
	return &WebViewGroup{obj}
}

// native returns a pointer to the underlying WebKitWebViewGroup.
func (w *WebViewGroup) native() *C.WebKitWebViewGroup {
	if w == nil || w.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(w.GObject)
	return C.toWebKitWebViewGroup(p)
}
