// This Source Code Form is subject to the terms of the Mozilla Public License,
// v. 2.0. If a copy of the MPL was not distributed with this file, you can
// obtain one at https://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2017 Tim Heckman

package twiml

import "encoding/xml"

// RingTone lets you record both legs of a call within the associated Dial verb. Recordings are available in two options: mono-channel or dual-channel.
type RingTone uint8

const (
	// RingToneAutomatic is omitted from being rendered to XML, which
	// effectively tells Twilio to use the default value.
	RingToneAutomatic RingTone = iota

	// RingToneAustralia is the ringback tone from Australia (au).
	RingToneAustralia

	// RingToneAustria is the ringback tone from Austria (at).
	RingToneAustria

	// RingToneBelgium is the ringback tone from Belgium (be).
	RingToneBelgium

	// RingToneBulgaria is the ringback tone from Bulgaria (bg).
	RingToneBulgaria

	// RingToneBrazil is the ringback tone from Brazil (br).
	RingToneBrazil

	// RingToneChile is the ringback tone from Chile (cl).
	RingToneChile

	// RingToneChina is the ringback tone from China (cn).
	RingToneChina

	// RingToneCzechia is the ringback tone from Czechia (cz).
	RingToneCzechia

	// RingToneDenmark is the ringback tone from Denmark (dk).
	RingToneDenmark

	// RingToneEstonia is the ringback tone from Estonia (ee).
	RingToneEstonia

	// RingToneFinland is the ringback tone from Finland (fi).
	RingToneFinland

	// RingToneFrance is the ringback tone from France (fr).
	RingToneFrance

	// RingToneGreece is the ringback tone from Greece (gr).
	RingToneGreece

	// RingToneGermany is the ringback tone from Germany (de).
	RingToneGermany

	// RingToneHungary is the ringback tone from Hungary (hu).
	RingToneHungary

	// RingToneIsrael is the ringback tone from Israel (il).
	RingToneIsrael

	// RingToneIndia is the ringback tone from India (in).
	RingToneIndia

	// RingToneItaly is the ringback tone from Italy (it).
	RingToneItaly

	// RingToneLithuania is the ringback tone from Lithuania (lt).
	RingToneLithuania

	// RingToneJapan is the ringback tone from Japan (jp).
	RingToneJapan

	// RingToneMexico is the ringback tone from Mexico (mx).
	RingToneMexico

	// RingToneMalaysia is the ringback tone from Malaysia (my).
	RingToneMalaysia

	// RingToneNetherlands is the ringback tone from the Netherlands (nl).
	RingToneNetherlands

	// RingToneNorway is the ringback tone from Norway (no).
	RingToneNorway

	// RingToneNewZealand is the ringback tone from New Zealand (nz).
	RingToneNewZealand

	// RingTonePhilippines is the ringback tone from Philippines (ph).
	RingTonePhilippines

	// RingTonePoland is the ringback tone from Poland (pl).
	RingTonePoland

	// RingTonePortugal is the ringback tone from Portugal (pt).
	RingTonePortugal

	// RingToneRussia is the ringback tone from Russia (ru).
	RingToneRussia

	// RingToneSingapore is the ringback tone from Singapore (sg).
	RingToneSingapore

	// RingToneSpain is the ringback tone from Spain (es).
	RingToneSpain

	// RingToneSweden is the ringback tone from Sweden (se).
	RingToneSweden

	// RingToneSwitzerland is the ringback tone from Switzerland (ch).
	RingToneSwitzerland

	// RingToneTaiwan is the ringback tone from Taiwan (tw).
	RingToneTaiwan

	// RingToneThailand is the ringback tone from Thailand (th).
	RingToneThailand

	// RingToneUK is the ringback tone from the United Kingdom (uk).
	RingToneUK

	// RingToneUS is the ringback tone from the United States (us).
	RingToneUS

	// RingToneUSOld is the old ringback tone from the United States (us-old).
	// The API documentation does not clarify what the difference is between
	// RingToneUSOld and RingToneUS.
	RingToneUSOld

	// RingToneVenezuela is the ringback tone from Venezuela (ve).
	RingToneVenezuela

	// RingToneSouthAfrica is the ringback tone from South Africa (za).
	RingToneSouthAfrica
)

// MarshalXMLAttr implements the xml.MarshalerAttr interface.
func (r RingTone) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{
		Name:  name,
		Value: r.String(),
	}

	return attr, nil
}

func (r RingTone) String() string {
	switch r {
	case RingToneAustralia:
		return "au"
	case RingToneAustria:
		return "at"
	case RingToneBelgium:
		return "be"
	case RingToneBulgaria:
		return "bg"
	case RingToneBrazil:
		return "br"
	case RingToneChile:
		return "cl"
	case RingToneChina:
		return "cn"
	case RingToneCzechia:
		return "cz"
	case RingToneDenmark:
		return "dk"
	case RingToneEstonia:
		return "ee"
	case RingToneFinland:
		return "fi"
	case RingToneFrance:
		return "fr"
	case RingToneGreece:
		return "gr"
	case RingToneGermany:
		return "de"
	case RingToneHungary:
		return "hu"
	case RingToneIsrael:
		return "il"
	case RingToneIndia:
		return "in"
	case RingToneItaly:
		return "it"
	case RingToneJapan:
		return "jp"
	case RingToneLithuania:
		return "lt"
	case RingToneMexico:
		return "mx"
	case RingToneMalaysia:
		return "my"
	case RingToneNetherlands:
		return "nl"
	case RingToneNorway:
		return "no"
	case RingToneNewZealand:
		return "nz"
	case RingTonePhilippines:
		return "ph"
	case RingTonePoland:
		return "pl"
	case RingTonePortugal:
		return "pt"
	case RingToneRussia:
		return "ru"
	case RingToneSingapore:
		return "sg"
	case RingToneSpain:
		return "es"
	case RingToneSweden:
		return "se"
	case RingToneSwitzerland:
		return "ch"
	case RingToneTaiwan:
		return "tw"
	case RingToneThailand:
		return "th"
	case RingToneUK:
		return "uk"
	case RingToneUS:
		return "us"
	case RingToneUSOld:
		return "us-old"
	case RingToneVenezuela:
		return "ve"
	case RingToneSouthAfrica:
		return "za"
	case RingToneAutomatic:
		return "automatic"
	default:
		return ""
	}
}
