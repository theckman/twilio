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

func TestConfRegion_String(t *testing.T) {
	tests := []struct {
		desc string
		in   ConfRegion
		out  string
	}{
		{"Default (Zero Value) ConfRegion should return empty-string", ConfRegion(0), ""},
		{"ConfRegionAustralia should return Australia", ConfRegionAustralia, "au1"},
		{"ConfRegionBrazil should return Brazil", ConfRegionBrazil, "br1"},
		{"ConfRegionIreland should return Ireland", ConfRegionIreland, "ie1"},
		{"ConfRegionJapan should return Japan", ConfRegionJapan, "jp1"},
		{"ConfRegionSingapore should return Singapore", ConfRegionSingapore, "sg1"},
		{"ConfRegionUS should return the United States", ConfRegionUS, "us1"},
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

func TestConfRegion_MarshalXMLAttr(t *testing.T) {
	attrName := xml.Name{Local: "trim"}

	tests := []struct {
		desc     string
		in       ConfRegion
		inName   xml.Name
		outName  xml.Name
		outValue string
	}{
		{"Default (Zero Value) ConfRegion should return empty-string", ConfRegion(0), attrName, attrName, ""},
		{"ConfRegionAustralia should return Australia", ConfRegionAustralia, attrName, attrName, "au1"},
		{"ConfRegionBrazil should return Brazil", ConfRegionBrazil, attrName, attrName, "br1"},
		{"ConfRegionIreland should return Ireland", ConfRegionIreland, attrName, attrName, "ie1"},
		{"ConfRegionJapan should return Japan", ConfRegionJapan, attrName, attrName, "jp1"},
		{"ConfRegionSingapore should return Singapore", ConfRegionSingapore, attrName, attrName, "sg1"},
		{"ConfRegionUS should return the United States", ConfRegionUS, attrName, attrName, "us1"},
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
