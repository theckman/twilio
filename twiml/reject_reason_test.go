// This Source Code Form is subject to the terms of the Mozilla Public License,
// v. 2.0. If a copy of the MPL was not distributed with this file, you can
// obtain one at https://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2017 Tim Heckman

package twiml

import (
	"encoding/xml"
	"testing"
)

func TestRejectReason_String(t *testing.T) {
	tests := []struct {
		desc string
		in   RejectReason
		out  string
	}{
		{"Default (Zero Value) RejectReason should be nothing", RejectReason(0), ""},
		{"RejectReasonRejected should set the rejection as rejected", RejectReasonRejected, "rejected"},
		{"RejectReasonBusy should set the rejection as busy", RejectReasonBusy, "busy"},
	}

	for _, test := range tests {
		if out := test.in.String(); out != test.out {
			t.Errorf(
				"\nDescription: %s\nFinishOnKey(%d).String() = %q; want %q",
				test.desc, test.in, out, test.out,
			)
		}
	}
}

func TestRejectReason_MarshalXMLAttr(t *testing.T) {
	attrName := xml.Name{Local: "trim"}

	tests := []struct {
		desc     string
		in       RejectReason
		inName   xml.Name
		outName  xml.Name
		outValue string
	}{
		{"Default (Zero Value) RejectReason should be nothing", RejectReason(0), attrName, attrName, ""},
		{"RejectReasonRejected should set the rejection as rejected", RejectReasonRejected, attrName, attrName, "rejected"},
		{"RejectReasonBusy should set the rejection as busy", RejectReasonBusy, attrName, attrName, "busy"},
	}

	var out xml.Attr
	var err error

	for _, test := range tests {
		out, err = test.in.MarshalXMLAttr(test.inName)

		if err != nil {
			t.Errorf(
				"\nDescription: %s\nVoice(%d).MarshalAttr(%#v) Error: %s",
				test.desc, test.in, test.inName, err,
			)
		}

		if out.Name.Space != test.outName.Space || out.Name.Local != test.outName.Local {
			t.Errorf(
				"\nDescription: %s\nVoice(%d).MarshalAttr(%#v).Name = %#v; want %#v",
				test.desc, test.in, test.inName, out.Name, test.outName,
			)
		}

		if out.Value != test.outValue {
			t.Errorf(
				"\nDescription: %s\nVoice(%d).MarshalAttr(%#v).Value = %q; want %q",
				test.desc, test.in, test.inName, out.Value, test.outValue,
			)
		}
	}
}
