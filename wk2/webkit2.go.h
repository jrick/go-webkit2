/*
 * Copyright (c) 2014 Josh Rickmar.
 * Use of this source code is governed by an ISC
 * license that can be found in the LICENSE file.
 */

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
