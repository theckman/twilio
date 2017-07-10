// This Source Code Form is subject to the terms of the Mozilla Public License,
// v. 2.0. If a copy of the MPL was not distributed with this file, you can
// obtain one at https://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2017 Tim Heckman

package twiml

import "encoding/xml"

// RejectReason specifies the rejection reason for a rejected call.
type RejectReason uint8

const (
	// RejectReasonRejected is the rejection reason indicating the call was
	// rejected.
	RejectReasonRejected RejectReason = 1 << iota

	// RejectReasonBusy is the rejection reason indicating the call was busy.
	RejectReasonBusy
)

// MarshalXMLAttr implements the xml.MarshalerAttr interface.
func (r RejectReason) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{
		Name:  name,
		Value: r.String(),
	}

	return attr, nil
}

func (r RejectReason) String() string {
	switch r {
	case RejectReasonRejected:
		return "rejected"
	case RejectReasonBusy:
		return "busy"
	default:
		return ""
	}
}
