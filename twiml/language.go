// This Source Code Form is subject to the terms of the Mozilla Public License,
// v. 2.0. If a copy of the MPL was not distributed with this file, you can
// obtain one at https://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2017 Tim Heckman

package twiml

import "encoding/xml"

// Language represents a language as understood by the TwiML. The language
// selected depends on the voice used to speak. By default this package uses the
// alice voice, as it allows more language support. If you wish to use the "man"
// or "woman" voice, you'll need to use one of the legacy languages.
type Language uint16

const (
	//LangDefault is the default value for language and causes the field to not be set.
	LangDefault Language = iota

	// LangEnglishUS is the English language as spoken in the United States.
	LangEnglishUS

	// LangCatalanSpain is the Catalan language as spoken in Spain.
	LangCatalanSpain

	// LangChineseCantonese is the Chinese (Cantonese) language.
	LangChineseCantonese

	// LangChineseMandarin is the Chinese (Mandarin) language.
	LangChineseMandarin

	// LangChineseTaiwaneseMandarin is the Chinese (Taiwanese Mandarin) language.
	LangChineseTaiwaneseMandarin

	// LangDanishDenmark is the Danish language as spoken in Denmark.
	LangDanishDenmark

	// LangDutchNetherlands is the Dutch language as spoken in the Netherlands.
	LangDutchNetherlands

	// LangEnglishAustralia is the English language as spoken in Australia.
	LangEnglishAustralia

	// LangEnglishCanada is what you think it is, bud. English as spoken in Canada.
	// No surprise there, eh?
	LangEnglishCanada

	// LangEnglishUK is the English language as spoken in the United Kingdom.
	LangEnglishUK

	// LangFinnishFinland is the Finnish language as spoken in Finland.
	LangFinnishFinland

	// LangFrenchCanada is the French language as spoken in Canada.
	LangFrenchCanada

	// LangFrenchFrance is the French language as spoken in France.
	LangFrenchFrance

	// LangGermanGermany is the German language as spoken in Germany.
	LangGermanGermany

	// LangItalianItaly is the Italian language as spoken in Italy.
	LangItalianItaly

	// LangJapaneseJapan is the Japanese language as spoken in Japan.
	// ありがとうございます
	LangJapaneseJapan

	// LangKoreanKorea is the Korean language as spoken in Korea.
	LangKoreanKorea

	// LangNorwegianNorway is the Norwegian language as spoken in Norway.
	LangNorwegianNorway

	// LangPolishPoland is the Polish language as spoken in Poland.
	LangPolishPoland

	// LangPortugeseBrazil is the Portugese language as spoken in Brazil.
	LangPortugeseBrazil

	// LangPortugesePortugal is the Portugese language as spoken on Portugal
	LangPortugesePortugal

	// LangRussianRussia is the Russian language as spoken in Russia.
	LangRussianRussia

	// LangSpanishMexico is the Spanish language as spoken in Mexico.
	LangSpanishMexico

	// LangSpanishSpain is the Spanish language as spoken in Spain.
	LangSpanishSpain

	// LangSwedishSweden is the Swedish language as spoken in Sweden.
	LangSwedishSweden
)

// MarshalXMLAttr implements the xml.MarshalerAttr interface.
func (l Language) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{
		Name:  name,
		Value: l.String(),
	}

	return attr, nil
}

func (l Language) String() string {
	switch l {
	case LangDefault:
		return ""
	case LangEnglishUS:
		return "en-US"
	case LangCatalanSpain:
		return "ca-ES"
	case LangChineseCantonese:
		return "zh-HK"
	case LangChineseMandarin:
		return "zh-CN"
	case LangChineseTaiwaneseMandarin:
		return "zh-TW"
	case LangDanishDenmark:
		return "da-DK"
	case LangDutchNetherlands:
		return "nl-NL"
	case LangEnglishAustralia:
		return "en-AU"
	case LangEnglishCanada:
		return "en-CA"
	case LangEnglishUK:
		return "en-GB"
	case LangFinnishFinland:
		return "fi-FI"
	case LangFrenchCanada:
		return "fr-CA"
	case LangFrenchFrance:
		return "fr-FR"
	case LangGermanGermany:
		return "de-DE"
	case LangItalianItaly:
		return "it-IT"
	case LangJapaneseJapan:
		return "ja-JP"
	case LangKoreanKorea:
		return "ko-KR"
	case LangNorwegianNorway:
		return "nb-NO"
	case LangPolishPoland:
		return "pl-PL"
	case LangPortugeseBrazil:
		return "pt-BR"
	case LangPortugesePortugal:
		return "pt-PT"
	case LangRussianRussia:
		return "ru-RU"
	case LangSpanishMexico:
		return "es-MX"
	case LangSpanishSpain:
		return "es-ES"
	case LangSwedishSweden:
		return "sv-SE"
	default:
		return "unknown"
	}
}
