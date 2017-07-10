// This Source Code Form is subject to the terms of the Mozilla Public License,
// v. 2.0. If a copy of the MPL was not distributed with this file, you can
// obtain one at https://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2017 Tim Heckman

package twiml

import "encoding/xml"

// ConfRegion specifies the region where Twilio should mix the conference.
// Specifying a value for region overrides Twilio's automatic region selection
// logic and should only be used if you are confident you understand where your
// conferences should be mixed. Twilio sets the region parameter from the first
// participant that specifies the parameter and will ignore the parameter from
// subsequent participants.
type ConfRegion uint16

const (
	// ConfRegionAustralia sets ConfRegion to Australia.
	ConfRegionAustralia ConfRegion = 1 << iota

	// ConfRegionBrazil sets ConfRegion to Brazil.
	ConfRegionBrazil

	// ConfRegionIreland sets ConfRegion to Ireland.
	ConfRegionIreland

	// ConfRegionJapan sets ConfRegion to Japan.
	ConfRegionJapan

	// ConfRegionSingapore sets ConfRegion to Singapore.
	ConfRegionSingapore

	// ConfRegionUS sets ConfRegion to the United States.
	ConfRegionUS
)

// MarshalXMLAttr implements the xml.MarshalerAttr interface.
func (r ConfRegion) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{
		Name:  name,
		Value: r.String(),
	}

	return attr, nil
}

func (r ConfRegion) String() string {
	switch r {
	case ConfRegionAustralia:
		return "au1"
	case ConfRegionBrazil:
		return "br1"
	case ConfRegionIreland:
		return "ie1"
	case ConfRegionJapan:
		return "jp1"
	case ConfRegionSingapore:
		return "sg1"
	case ConfRegionUS:
		return "us1"
	default:
		return ""
	}
}
