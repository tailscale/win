// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build windows && !386

package win

import "encoding/binary"

const (
	sizeTASKDIALOGCONFIG  = 0xA0
	sizeTASKDIALOG_BUTTON = 12
)

type packPtr uint64

var packByteOrder = binary.LittleEndian
