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

func TestConfStartOnEnterBool_String(t *testing.T) {
	tests := []struct {
		desc string
		in   ConfStartOnEnterBool
		out  string
	}{
		{"Default (Zero Value) ConfStartOnEnterBool should return empty-string", ConfStartOnEnterBool(0), ""},
		{"ConfStartOnEnterTrue should be true", ConfStartOnEnterTrue, "true"},
		{"ConfStartOnEnterFalse should be false", ConfStartOnEnterFalse, "false"},
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

func TestConfStartOnEnterBool_Bool(t *testing.T) {
	tests := []struct {
		desc string
		in   ConfStartOnEnterBool
		out  bool
	}{
		{"Default (Zero Value) ConfStartOnEnterBool should return true", ConfStartOnEnterBool(0), true},
		{"ConfStartOnEnterTrue should be true", ConfStartOnEnterTrue, true},
		{"ConfStartOnEnterFalse should be false", ConfStartOnEnterFalse, false},
		{"An Unknown ConfStartOnEnterBool value should return true", ConfStartOnEnterBool(^uint8(0)), true},
	}

	for _, test := range tests {
		if out := test.in.Bool(); out != test.out {
			t.Errorf(
				"\nDescription: %s\nFinishOnKey(%d).Bool() = %v; want %v",
				test.desc, test.in, out, test.out,
			)
		}
	}
}

func TestConfStartOnEnterBool_MarshalXMLAttr(t *testing.T) {
	attrName := xml.Name{Local: "trim"}

	tests := []struct {
		desc     string
		in       ConfStartOnEnterBool
		inName   xml.Name
		outName  xml.Name
		outValue string
	}{
		{"Default (Zero Value) ConfStartOnEnterBool should return empty-string", ConfStartOnEnterBool(0), attrName, attrName, ""},
		{"ConfStartOnEnterTrue should be true", ConfStartOnEnterTrue, attrName, attrName, "true"},
		{"ConfStartOnEnterFalse should be false", ConfStartOnEnterFalse, attrName, attrName, "false"},
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
