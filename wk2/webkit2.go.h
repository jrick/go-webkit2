/*
 * Copyright (c) 2014 Josh Rickmar.
 * Use of this source code is governed by an ISC
 * license that can be found in the LICENSE file.
 */

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
