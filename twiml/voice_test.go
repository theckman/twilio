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

func TestVoice_String(t *testing.T) {
	tests := []struct {
		desc string
		in   Voice
		out  string
	}{
		{"Default (Zero Vault) voice should empty string", Voice(0), ""},
		{"VoiceDefault voice should be empty string", VoiceDefault, ""},
		{"VoiceAlice voice should be Alice", VoiceAlice, "alice"},
		{"VoiceMan voice should be Alice", VoiceMan, "man"},
		{"VoiceWoman voice should be Alice", VoiceWoman, "woman"},
	}

	var out string

	for _, test := range tests {
		out = test.in.String()

		if out != test.out {
			t.Errorf(
				"\nDescription: %s\nVoice(%d).String() = %q; want %q",
				test.desc, test.in, out, test.out,
			)
		}
	}
}

func TestVoice_MarshalXMLAttr(t *testing.T) {
	attrName := xml.Name{Local: "voice"}

	tests := []struct {
		desc     string
		in       Voice
		inName   xml.Name
		outName  xml.Name
		outValue string
	}{
		{"Default (Zero Vault) voice should empty string", Voice(0), attrName, attrName, ""},
		{"VoiceAlice voice should be Alice", VoiceAlice, attrName, attrName, "alice"},
		{"VoiceMan voice should be Alice", VoiceMan, attrName, attrName, "man"},
		{"VoiceWoman voice should be Alice", VoiceWoman, attrName, attrName, "woman"},
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
