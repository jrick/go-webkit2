/*
 * Copyright (c) 2014 Josh Rickmar.
 * Use of this source code is governed by an ISC
 * license that can be found in the LICENSE file.
 */

static gchar **
allocGCharArray(size_t n)
{
	gchar **	v;

	v = calloc(n, sizeof(gchar *));
	return (v);
}

void
freeGCharArray(gchar **v)
{
	int		i;

	if (v == NULL) {
		return;
	}

	for (i = 0; v[i] != NULL; ++i) {
		free(v[i]);
	}
	free(v);
}

static gchar *
indexGCharArray(gchar **v, int n)
{
	return (v[n]);
}

static void
indexGCharArraySet(gchar **v, int n, gchar *s)
{
	v[n] = s;
}

static WebKitBackForwardList *
toWebKitBackForwardList(void *p)
{
	return (WEBKIT_BACK_FORWARD_LIST(p));
}

static WebKitBackForwardListItem *
toWebKitBackForwardListItem(void *p)
{
	return (WEBKIT_BACK_FORWARD_LIST_ITEM(p));
}

static WebKitCookieManager *
toWebKitCookieManager(void *p)
{
	return (WEBKIT_COOKIE_MANAGER(p));
}

static WebKitDownload *
toWebKitDownload(void *p)
{
	return (WEBKIT_DOWNLOAD(p));
}

static WebKitFaviconDatabase *
toWebKitFaviconDatabase(void *p)
{
	return (WEBKIT_FAVICON_DATABASE(p));
}

static WebKitSecurityManager *
toWebKitSecurityManager(void *p)
{
	return (WEBKIT_SECURITY_MANAGER(p));
}

static WebKitURIRequest *
toWebKitURIRequest(void *p)
{
	return (WEBKIT_URI_REQUEST(p));
}

static WebKitWebContext *
toWebKitWebContext(void *p)
{
	return (WEBKIT_WEB_CONTEXT(p));
}

static WebKitWebView *
toWebKitWebView(void *p)
{
	return (WEBKIT_WEB_VIEW(p));
}

static WebKitWebViewGroup *
toWebKitWebViewGroup(void *p)
{
	return (WEBKIT_WEB_VIEW_GROUP(p));
}
