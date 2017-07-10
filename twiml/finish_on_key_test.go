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

func TestFinishOnKey_String(t *testing.T) {
	tests := []struct {
		desc string
		in   FinishOnKey
		out  string
	}{
		{"Invalid FinishOnKey should be empty string", FinishOnKey(0), ""},
		{"FinishKeyNumber0 should be 0", FinishKeyNumber0, "0"},
		{"FinishKeyNumber1 should be 1", FinishKeyNumber1, "1"},
		{"FinishKeyNumber2 should be 2", FinishKeyNumber2, "2"},
		{"FinishKeyNumber3 should be 3", FinishKeyNumber3, "3"},
		{"FinishKeyNumber4 should be 4", FinishKeyNumber4, "4"},
		{"FinishKeyNumber5 should be 5", FinishKeyNumber5, "5"},
		{"FinishKeyNumber6 should be 6", FinishKeyNumber6, "6"},
		{"FinishKeyNumber7 should be 7", FinishKeyNumber7, "7"},
		{"FinishKeyNumber8 should be 8", FinishKeyNumber8, "8"},
		{"FinishKeyNumber9 should be 9", FinishKeyNumber9, "9"},
		{"FinishKeyStar should be *", FinishKeyStar, "*"},
		{"FinishKeyPound should be #", FinishKeyPound, "#"},
		{"FinishKeyNone should be empty-string", FinishKeyNone, ""},
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

func TestFinishOnKey_MarshalXMLAttr(t *testing.T) {
	attrName := xml.Name{Local: "finishOnKey"}

	tests := []struct {
		desc     string
		in       FinishOnKey
		inName   xml.Name
		outName  xml.Name
		outValue string
	}{
		{"Default (Zero Vault) voice should be Alice", FinishOnKey(0), attrName, attrName, ""},
		{"FinishKeyNumber0 should be 0", FinishKeyNumber0, attrName, attrName, "0"},
		{"FinishKeyNumber9 should be 9", FinishKeyNumber9, attrName, attrName, "9"},
		{"FinishKeyStar should be *", FinishKeyStar, attrName, attrName, "*"},
		{"FinishKeyPound should be #", FinishKeyPound, attrName, attrName, "#"},
		{"FinishKeyNone should be empty-string", FinishKeyNone, attrName, attrName, ""},
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
