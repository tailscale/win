// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build windows

package win

import "encoding/binary"

const (
	sizeTASKDIALOGCONFIG  = 0x60
	sizeTASKDIALOG_BUTTON = 8
)

type packPtr uint32

var packByteOrder = binary.LittleEndian
