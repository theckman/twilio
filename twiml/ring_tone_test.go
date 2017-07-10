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

func TestRingTone_String(t *testing.T) {
	tests := []struct {
		desc string
		in   RingTone
		out  string
	}{
		{"Default (Zero Value) RingTone should return automatic", RingTone(0), "automatic"},
		{"RingToneAutomatic should return automatic", RingToneAutomatic, "automatic"},
		{"RingToneAustralia should return au", RingToneAustralia, "au"},
		{"RingToneAustria should return at", RingToneAustria, "at"},
		{"RingToneBelgium should return be", RingToneBelgium, "be"},
		{"RingToneBulgaria should return bg", RingToneBulgaria, "bg"},
		{"RingToneBrazil should return br", RingToneBrazil, "br"},
		{"RingToneChile should return cl", RingToneChile, "cl"},
		{"RingToneChina should return cn", RingToneChina, "cn"},
		{"RingToneCzechia should return cz", RingToneCzechia, "cz"},
		{"RingToneDenmark should return dk", RingToneDenmark, "dk"},
		{"RingToneEstonia should return ee", RingToneEstonia, "ee"},
		{"RingToneFinland should return fi", RingToneFinland, "fi"},
		{"RingToneFrance should return fr", RingToneFrance, "fr"},
		{"RingToneGreece should return gr", RingToneGreece, "gr"},
		{"RingToneGermany should return de", RingToneGermany, "de"},
		{"RingToneHungary should return hu", RingToneHungary, "hu"},
		{"RingToneIsrael should return il", RingToneIsrael, "il"},
		{"RingToneIndia should return in", RingToneIndia, "in"},
		{"RingToneItaly should return it", RingToneItaly, "it"},
		{"RingToneLithuania should return lt", RingToneLithuania, "lt"},
		{"RingToneJapan should return jp", RingToneJapan, "jp"},
		{"RingToneMexico should return mx", RingToneMexico, "mx"},
		{"RingToneMalaysia should return my", RingToneMalaysia, "my"},
		{"RingToneNetherlands should return nl", RingToneNetherlands, "nl"},
		{"RingToneNorway should return no", RingToneNorway, "no"},
		{"RingToneNewZealand should return nz", RingToneNewZealand, "nz"},
		{"RingTonePhilippines should return ph", RingTonePhilippines, "ph"},
		{"RingTonePoland should return pl", RingTonePoland, "pl"},
		{"RingTonePortugal should return pt", RingTonePortugal, "pt"},
		{"RingToneRussia should return ru", RingToneRussia, "ru"},
		{"RingToneSingapore should return sg", RingToneSingapore, "sg"},
		{"RingToneSpain should return es", RingToneSpain, "es"},
		{"RingToneSweden should return se", RingToneSweden, "se"},
		{"RingToneSwitzerland should return ch", RingToneSwitzerland, "ch"},
		{"RingToneTaiwan should return tw", RingToneTaiwan, "tw"},
		{"RingToneThailand should return th", RingToneThailand, "th"},
		{"RingToneUK should return uk", RingToneUK, "uk"},
		{"RingToneUS should return us", RingToneUS, "us"},
		{"RingToneUSOld should return us-old", RingToneUSOld, "us-old"},
		{"RingToneVenezuela should return ve", RingToneVenezuela, "ve"},
		{"RingToneSouthAfrica should return za", RingToneSouthAfrica, "za"},
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

func TestRingTone_MarshalXMLAttr(t *testing.T) {
	attrName := xml.Name{Local: "trim"}

	tests := []struct {
		desc     string
		in       RingTone
		inName   xml.Name
		outName  xml.Name
		outValue string
	}{
		{"Default (Zero Value) RingTone should return automatic", RingTone(0), attrName, attrName, "automatic"},
		{"RingToneJapan should return jp", RingToneJapan, attrName, attrName, "jp"},
		{"RingToneUS should return us", RingToneUS, attrName, attrName, "us"},
		{"RingToneUSOld should return us-old", RingToneUSOld, attrName, attrName, "us-old"},
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
