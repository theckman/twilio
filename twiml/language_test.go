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

func TestLanguage_String(t *testing.T) {
	tests := []struct {
		desc string
		in   Language
		out  string
	}{
		{"Default (Zero Value) should be empty string", Language(0), ""},
		{"LangDefault should be empty string", LangDefault, ""},
		{"DanishDenmark should be Danish, Denmark", LangDanishDenmark, "da-DK"},
		{"GermanGermany should be German, Germany", LangGermanGermany, "de-DE"},
		{"EnglishAustralia should be English, Australia", LangEnglishAustralia, "en-AU"},
		{"EnglishCanada should be English, Canada", LangEnglishCanada, "en-CA"},
		{"EnglishUK should be English, United Kingdom", LangEnglishUK, "en-GB"},
		{"EnglishUS should be English, United States", LangEnglishUS, "en-US"},
		{"CatalanSpain should be Catalan, Spain", LangCatalanSpain, "ca-ES"},
		{"SpanishSpain should be Spanish, Spain", LangSpanishSpain, "es-ES"},
		{"SpanishMexico should be Spanish, Mexico", LangSpanishMexico, "es-MX"},
		{"FinnishFinland should be Finnish, Finland", LangFinnishFinland, "fi-FI"},
		{"FrenchCanada should be French, Canada", LangFrenchCanada, "fr-CA"},
		{"FrenchFrance should be French, France", LangFrenchFrance, "fr-FR"},
		{"ItalianItaly should be Italian, Italy", LangItalianItaly, "it-IT"},
		{"JapaneseJapan should be Japanese, Japan", LangJapaneseJapan, "ja-JP"},
		{"KoreaKorean should be Korea, Korean", LangKoreanKorea, "ko-KR"},
		{"NorwegianNorway should be Norwegian, Norway", LangNorwegianNorway, "nb-NO"},
		{"DutchNetherlands should be Dutch, Netherlands", LangDutchNetherlands, "nl-NL"},
		{"PolishPoland should be Polish, Poland", LangPolishPoland, "pl-PL"},
		{"PortugeseBrazil should be Portugese, Brazil", LangPortugeseBrazil, "pt-BR"},
		{"PortugesePortugal should be Portugese, Portugal", LangPortugesePortugal, "pt-PT"},
		{"RussianRussia should be Russia, Russian", LangRussianRussia, "ru-RU"},
		{"SwedishSweden should be Swedish, Sweden", LangSwedishSweden, "sv-SE"},
		{"ChineseMandarin should be Chinese (Mandarin)", LangChineseMandarin, "zh-CN"},
		{"ChineseCantonese should be Chinese (Cantonese)", LangChineseCantonese, "zh-HK"},
		{"ChineseTaiwaneseMandarin should be Chinese (Taiwanese Mandarin)", LangChineseTaiwaneseMandarin, "zh-TW"},
	}

	var out string

	for _, test := range tests {
		out = test.in.String()

		if out != test.out {
			t.Errorf(
				"\nDescription: %s\nLanguage(%d).String() = %q; want %q",
				test.desc, test.in, out, test.out,
			)
		}
	}
}

func TestLanguage_MarshalXMLAttr(t *testing.T) {
	langName := xml.Name{Local: "language"}

	tests := []struct {
		desc     string
		in       Language
		inName   xml.Name
		outName  xml.Name
		outValue string
	}{
		{"Default (Zero Value) should be empty string", Language(0), langName, langName, ""},
		{"EnglishUS should be English, United States", LangEnglishUS, langName, langName, "en-US"},
		{"SpanishMexico should be Spanish, Mexico", LangSpanishMexico, langName, langName, "es-MX"},
		{"JapaneseJapan should be Japanese, Japan", LangJapaneseJapan, langName, langName, "ja-JP"},
		{"KoreaKorean should be Korea, Korean", LangKoreanKorea, langName, langName, "ko-KR"},
		{"NorwegianNorway should be Norwegian, Norway", LangNorwegianNorway, langName, langName, "nb-NO"},
		{"DutchNetherlands should be Dutch, Netherlands", LangDutchNetherlands, langName, langName, "nl-NL"},
		{
			"ChineseCantonese should be Chinese (Cantonese)",
			LangChineseCantonese,
			xml.Name{Space: "blah", Local: "Test"},
			xml.Name{Space: "blah", Local: "Test"},
			"zh-HK",
		},
	}
	var out xml.Attr
	var err error

	for _, test := range tests {
		out, err = test.in.MarshalXMLAttr(test.inName)

		if err != nil {
			t.Errorf(
				"\nDescription: %s\nLanguage(%d).MarshalAttr(%#v) Error: %s",
				test.desc, test.in, test.inName, err,
			)
		}

		if out.Name.Space != test.outName.Space || out.Name.Local != test.outName.Local {
			t.Errorf(
				"\nDescription: %s\nLanguage(%d).MarshalAttr(%#v).Name = %#v; want %#v",
				test.desc, test.in, test.inName, out.Name, test.outName,
			)
		}

		if out.Value != test.outValue {
			t.Errorf(
				"\nDescription: %s\nLanguage(%d).MarshalAttr(%#v).Value = %q; want %q",
				test.desc, test.in, test.inName, out.Value, test.outValue,
			)
		}
	}
}

func BenchmarkLanguageStringEnglishUS(b *testing.B) {
	lang := LangEnglishUS
	for i := 0; i < b.N; i++ {
		lang.String()
	}
}

func BenchmarkLanguageStringJapaneseJapan(b *testing.B) {
	lang := LangJapaneseJapan
	for i := 0; i < b.N; i++ {
		lang.String()
	}
}

func BenchmarkLanguageStringSwedishSweden(b *testing.B) {
	lang := LangSwedishSweden
	for i := 0; i < b.N; i++ {
		lang.String()
	}
}

func BenchmarkLanguageStringUnknown(b *testing.B) {
	lang := Language(100)
	for i := 0; i < b.N; i++ {
		lang.String()
	}
}
