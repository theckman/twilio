// This Source Code Form is subject to the terms of the Mozilla Public License,
// v. 2.0. If a copy of the MPL was not distributed with this file, you can
// obtain one at https://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2017 Tim Heckman

// Package twiml provides support for generating Twilio Markup Language (TwiML).
// TwiML is an XML document that uses tags and attributes to instruct Twilio on
// how to process a phone call. The TwiML instructions can be applied to
// outbound calls (calls you've dialed), or inbound calls (calls that dialed
// your Twilio number).
//
// Most of the contents of this package are custom types (structs, or other) to
// allow us to generate documents that meet the requirements of a TwiML document
// in a Go-like fashion. One of the biggest features of this package is to
// provide types and constants that enable certain functionality or render
// certain attributes. This means you can use the client, and have code
// completion work for you, without needing to remember the string value that
// the TwiML parser expects. One example is the Trim type and its constants,
// where we translate the value to the string representation in TwiML (e.g.,
// DoNotTrim becomes "do-not-trim").
//
// It's worth noting that this package does not do deep validation of TwiML
// documents you are attempting to render. In other words if you try to render
// an invalid TwiML document, by trying to place a Redirect verb within a Gather
// verb for example, this package will happily render the document. However,
// Twilio will fail to parse this document as it is invalid per the spec.
//
// There are a few functions available to you for rendering out TwiML, with the
// main being EncodeResponse(). All of the other functions end up calling
// EncodeResponse() under the hood. If you aren't sending the data to an
// io.Writer interface, which is what EncodeResponse() expects, there is a
// MarshalResponse() function which returns a byte slice instead. There are also
// functions to work with slices of verbs instead of a full instance of
// *Response.
//
// Here is a simple example of using EncodeResponse():
//
// 		buf := &bytes.Buffer{}
// 		say := &twiml.Say{Message: "Hi there!"}
// 		resp := &twiml.Response{Verbs: []interface{}{say}}
// 		if err := twiml.EncodeResponse(buf, resp); err != nil {
// 			panic(err) // handle this better, though
// 		}
// 		fmt.Println(buf.String())
//
// The above code should render the following XML document:
//
// 		<?xml version="1.0" encoding="UTF-8"?>
// 		<Response>
// 		  <Say>Hi there!</Say>
// 		</Response>
package twiml
