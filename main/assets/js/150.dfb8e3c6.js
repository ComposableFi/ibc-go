(window.webpackJsonp=window.webpackJsonp||[]).push([[150],{713:function(e,a,l){"use strict";l.r(a);var t=l(1),c=Object(t.a)({},(function(){var e=this,a=e.$createElement,l=e._self._c||a;return l("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[l("h1",{attrs:{id:"integration"}},[l("a",{staticClass:"header-anchor",attrs:{href:"#integration"}},[e._v("#")]),e._v(" Integration")]),e._v(" "),l("p",{attrs:{synopsis:""}},[e._v("Learn how to integrate the callbacks middleware with IBC applications. The following document is intended for developers building on top of the Cosmos SDK and only applies for Cosmos SDK chains.")]),e._v(" "),l("p",[e._v("The callbacks middleware is a minimal and stateless implementation of the IBC middleware interface. It does not have a keeper, nor does it store any state. It simply routes IBC middleware messages to the appropriate callback function, which is implemented by the secondary application. Therefore, it doesn't need to be registered as a module, nor does it need to be added to the module manager. It only needs to be added to the IBC application stack.")]),e._v(" "),l("h2",{attrs:{id:"pre-requisite-readings"}},[l("a",{staticClass:"header-anchor",attrs:{href:"#pre-requisite-readings"}},[e._v("#")]),e._v(" Pre-requisite Readings")]),e._v(" "),l("ul",[l("li",{attrs:{prereq:""}},[l("RouterLink",{attrs:{to:"/ibc/middleware/develop.html"}},[e._v("IBC middleware development")])],1),e._v(" "),l("li",{attrs:{prereq:""}},[l("RouterLink",{attrs:{to:"/ibc/middleware/integration.html"}},[e._v("IBC middleware integration")])],1)]),e._v(" "),l("p",[e._v("The callbacks middleware, as the name suggests, plays the role of an IBC middleware and as such must be configured by chain developers to route and handle IBC messages correctly.\nFor Cosmos SDK chains this setup is done via the "),l("code",[e._v("app/app.go")]),e._v(" file, where modules are constructed and configured in order to bootstrap the blockchain application.")]),e._v(" "),l("h2",{attrs:{id:"configuring-an-application-stack-with-the-callbacks-middleware"}},[l("a",{staticClass:"header-anchor",attrs:{href:"#configuring-an-application-stack-with-the-callbacks-middleware"}},[e._v("#")]),e._v(" Configuring an application stack with the callbacks middleware")]),e._v(" "),l("p",[e._v("As mentioned in "),l("RouterLink",{attrs:{to:"/ibc/middleware/develop.html"}},[e._v("IBC middleware development")]),e._v(" an application stack may be composed of many or no middlewares that nest a base application.\nThese layers form the complete set of application logic that enable developers to build composable and flexible IBC application stacks.\nFor example, an application stack may just be a single base application like "),l("code",[e._v("transfer")]),e._v(", however, the same application stack composed with "),l("code",[e._v("29-fee")]),e._v(" and "),l("code",[e._v("callbacks")]),e._v(" will nest the "),l("code",[e._v("transfer")]),e._v(" base application twice by wrapping it with the Fee Middleware module and then callbacks middleware.")],1),e._v(" "),l("p",[e._v("The callbacks middleware also "),l("strong",[e._v("requires")]),e._v(" a secondary application that will receive the callbacks to implement the "),l("a",{attrs:{href:"https://github.com/cosmos/ibc-go/blob/v7.3.0/modules/apps/callbacks/types/expected_keepers.go#L11-L83",target:"_blank",rel:"noopener noreferrer"}},[l("code",[e._v("ContractKeeper")]),l("OutboundLink")],1),e._v(". Since the wasm module does not yet support the callbacks middleware, we will use the "),l("code",[e._v("mockContractKeeper")]),e._v(" module in the examples below. You should replace this with a module that implements "),l("code",[e._v("ContractKeeper")]),e._v(".")]),e._v(" "),l("h3",{attrs:{id:"transfer"}},[l("a",{staticClass:"header-anchor",attrs:{href:"#transfer"}},[e._v("#")]),e._v(" Transfer")]),e._v(" "),l("p",[e._v("See below for an example of how to create an application stack using "),l("code",[e._v("transfer")]),e._v(", "),l("code",[e._v("29-fee")]),e._v(", and "),l("code",[e._v("callbacks")]),e._v(". Feel free to omit the "),l("code",[e._v("29-fee")]),e._v(" middleware if you do not want to use it.\nThe following "),l("code",[e._v("transferStack")]),e._v(" is configured in "),l("code",[e._v("app/app.go")]),e._v(" and added to the IBC "),l("code",[e._v("Router")]),e._v(".\nThe in-line comments describe the execution flow of packets between the application stack and IBC core.")]),e._v(" "),l("tm-code-block",{staticClass:"codeblock",attrs:{language:"go",base64:"Ly8gQ3JlYXRlIFRyYW5zZmVyIFN0YWNrCi8vIFNlbmRQYWNrZXQsIHNpbmNlIGl0IGlzIG9yaWdpbmF0aW5nIGZyb20gdGhlIGFwcGxpY2F0aW9uIHRvIGNvcmUgSUJDOgovLyB0cmFuc2ZlcktlZXBlci5TZW5kUGFja2V0IC0mZ3Q7IGNhbGxiYWNrcy5TZW5kUGFja2V0IC0mZ3Q7IGZlZS5TZW5kUGFja2V0IC0mZ3Q7IGNoYW5uZWwuU2VuZFBhY2tldAoKLy8gUmVjdlBhY2tldCwgbWVzc2FnZSB0aGF0IG9yaWdpbmF0ZXMgZnJvbSBjb3JlIElCQyBhbmQgZ29lcyBkb3duIHRvIGFwcCwgdGhlIGZsb3cgaXMgdGhlIG90aGVyIHdheQovLyBjaGFubmVsLlJlY3ZQYWNrZXQgLSZndDsgY2FsbGJhY2tzLk9uUmVjdlBhY2tldCAtJmd0OyBmZWUuT25SZWN2UGFja2V0IC0mZ3Q7IHRyYW5zZmVyLk9uUmVjdlBhY2tldAoKLy8gdHJhbnNmZXIgc3RhY2sgY29udGFpbnMgKGZyb20gdG9wIHRvIGJvdHRvbSk6Ci8vIC0gSUJDIENhbGxiYWNrcyBNaWRkbGV3YXJlCi8vIC0gSUJDIEZlZSBNaWRkbGV3YXJlCi8vIC0gVHJhbnNmZXIKCi8vIGNyZWF0ZSBJQkMgbW9kdWxlIGZyb20gYm90dG9tIHRvIHRvcCBvZiBzdGFjawp2YXIgdHJhbnNmZXJTdGFjayBwb3J0dHlwZXMuSUJDTW9kdWxlCnRyYW5zZmVyU3RhY2sgPSB0cmFuc2Zlci5OZXdJQkNNb2R1bGUoYXBwLlRyYW5zZmVyS2VlcGVyKQp0cmFuc2ZlclN0YWNrID0gaWJjZmVlLk5ld0lCQ01pZGRsZXdhcmUodHJhbnNmZXJTdGFjaywgYXBwLklCQ0ZlZUtlZXBlcikKLy8gbWF4Q2FsbGJhY2tHYXMgaXMgYSBoYXJkLWNvZGVkIHZhbHVlIHRoYXQgaXMgcGFzc2VkIHRvIHRoZSBjYWxsYmFja3MgbWlkZGxld2FyZQp0cmFuc2ZlclN0YWNrID0gaWJjY2FsbGJhY2tzLk5ld0lCQ01pZGRsZXdhcmUodHJhbnNmZXJTdGFjaywgYXBwLklCQ0ZlZUtlZXBlciwgYXBwLk1vY2tDb250cmFjdEtlZXBlciwgbWF4Q2FsbGJhY2tHYXMpCi8vIFNpbmNlIHRoZSBjYWxsYmFja3MgbWlkZGxld2FyZSBpdHNlbGYgaXMgYW4gaWNzNHdyYXBwZXIsIGl0IG5lZWRzIHRvIGJlIHBhc3NlZCB0byB0aGUgdHJhbnNmZXIga2VlcGVyCmFwcC5UcmFuc2ZlcktlZXBlci5XaXRoSUNTNFdyYXBwZXIodHJhbnNmZXJTdGFjay4ocG9ydHR5cGVzLklDUzRXcmFwcGVyKSkKCi8vIEFkZCB0cmFuc2ZlciBzdGFjayB0byBJQkMgUm91dGVyCmliY1JvdXRlci5BZGRSb3V0ZShpYmN0cmFuc2ZlcnR5cGVzLk1vZHVsZU5hbWUsIHRyYW5zZmVyU3RhY2spCg=="}}),e._v(" "),l("div",{staticClass:"custom-block warning"},[l("p",[e._v("The usage of "),l("code",[e._v("WithICS4Wrapper")]),e._v(" after "),l("code",[e._v("transferStack")]),e._v("'s configuration is critical! It allows the callbacks middleware to do "),l("code",[e._v("SendPacket")]),e._v(" callbacks and asynchronous "),l("code",[e._v("ReceivePacket")]),e._v(" callbacks. You must do this regardless of whether you are using the "),l("code",[e._v("29-fee")]),e._v(" middleware or not.")])]),e._v(" "),l("h3",{attrs:{id:"interchain-accounts-controller"}},[l("a",{staticClass:"header-anchor",attrs:{href:"#interchain-accounts-controller"}},[e._v("#")]),e._v(" Interchain Accounts Controller")]),e._v(" "),l("tm-code-block",{staticClass:"codeblock",attrs:{language:"go",base64:"Ly8gQ3JlYXRlIEludGVyY2hhaW4gQWNjb3VudHMgU3RhY2sKLy8gU2VuZFBhY2tldCwgc2luY2UgaXQgaXMgb3JpZ2luYXRpbmcgZnJvbSB0aGUgYXBwbGljYXRpb24gdG8gY29yZSBJQkM6Ci8vIGljYUNvbnRyb2xsZXJLZWVwZXIuU2VuZFR4IC0mZ3Q7IGNhbGxiYWNrcy5TZW5kUGFja2V0IC0mZ3Q7IGZlZS5TZW5kUGFja2V0IC0mZ3Q7IGNoYW5uZWwuU2VuZFBhY2tldAoKdmFyIGljYUNvbnRyb2xsZXJTdGFjayBwb3J0dHlwZXMuSUJDTW9kdWxlCmljYUNvbnRyb2xsZXJTdGFjayA9IGljYWNvbnRyb2xsZXIuTmV3SUJDTWlkZGxld2FyZShuaWwsIGFwcC5JQ0FDb250cm9sbGVyS2VlcGVyKQppY2FDb250cm9sbGVyU3RhY2sgPSBpYmNmZWUuTmV3SUJDTWlkZGxld2FyZShpY2FDb250cm9sbGVyU3RhY2ssIGFwcC5JQkNGZWVLZWVwZXIpCi8vIG1heENhbGxiYWNrR2FzIGlzIGEgaGFyZC1jb2RlZCB2YWx1ZSB0aGF0IGlzIHBhc3NlZCB0byB0aGUgY2FsbGJhY2tzIG1pZGRsZXdhcmUKaWNhQ29udHJvbGxlclN0YWNrID0gaWJjY2FsbGJhY2tzLk5ld0lCQ01pZGRsZXdhcmUoaWNhQ29udHJvbGxlclN0YWNrLCBhcHAuSUJDRmVlS2VlcGVyLCBhcHAuTW9ja0NvbnRyYWN0S2VlcGVyLCBtYXhDYWxsYmFja0dhcykKLy8gU2luY2UgdGhlIGNhbGxiYWNrcyBtaWRkbGV3YXJlIGl0c2VsZiBpcyBhbiBpY3M0d3JhcHBlciwgaXQgbmVlZHMgdG8gYmUgcGFzc2VkIHRvIHRoZSBpY2EgY29udHJvbGxlciBrZWVwZXIKYXBwLklDQUNvbnRyb2xsZXJLZWVwZXIuV2l0aElDUzRXcmFwcGVyKGljYUNvbnRyb2xsZXJTdGFjay4ocG9ydHR5cGVzLklDUzRXcmFwcGVyKSkKCi8vIFJlY3ZQYWNrZXQsIG1lc3NhZ2UgdGhhdCBvcmlnaW5hdGVzIGZyb20gY29yZSBJQkMgYW5kIGdvZXMgZG93biB0byBhcHAsIHRoZSBmbG93IGlzOgovLyBjaGFubmVsLlJlY3ZQYWNrZXQgLSZndDsgY2FsbGJhY2tzLk9uUmVjdlBhY2tldCAtJmd0OyBmZWUuT25SZWN2UGFja2V0IC0mZ3Q7IGljYUhvc3QuT25SZWN2UGFja2V0Cgp2YXIgaWNhSG9zdFN0YWNrIHBvcnR0eXBlcy5JQkNNb2R1bGUKaWNhSG9zdFN0YWNrID0gaWNhaG9zdC5OZXdJQkNNb2R1bGUoYXBwLklDQUhvc3RLZWVwZXIpCmljYUhvc3RTdGFjayA9IGliY2ZlZS5OZXdJQkNNaWRkbGV3YXJlKGljYUhvc3RTdGFjaywgYXBwLklCQ0ZlZUtlZXBlcikKCi8vIEFkZCBJQ0EgaG9zdCBhbmQgY29udHJvbGxlciB0byBJQkMgcm91dGVyIGliY1JvdXRlci4KQWRkUm91dGUoaWNhY29udHJvbGxlcnR5cGVzLlN1Yk1vZHVsZU5hbWUsIGljYUNvbnRyb2xsZXJTdGFjaykuCkFkZFJvdXRlKGljYWhvc3R0eXBlcy5TdWJNb2R1bGVOYW1lLCBpY2FIb3N0U3RhY2spLgo="}}),e._v(" "),l("div",{staticClass:"custom-block warning"},[l("p",[e._v("The usage of "),l("code",[e._v("WithICS4Wrapper")]),e._v(" here is also critical!")])])],1)}),[],!1,null,null,null);a.default=c.exports}}]);