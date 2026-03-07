// Copyright 2023 Tailscale Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package win

const (
	BT_IMAGEFILE  = 0
	BT_BORDERFILL = 1
	BT_NONE       = 2
)

const (
	IL_VERTICAL   = 0
	IL_HORIZONTAL = 1
)

const (
	BT_RECT      = 0
	BT_ROUNDRECT = 1
	BT_ELLIPSE   = 2
)

const (
	FT_SOLID          = 0
	FT_VERTGRADIENT   = 1
	FT_HORZGRADIENT   = 2
	FT_RADIALGRADIENT = 3
	FT_TILEIMAGE      = 4
)

const (
	ST_TRUESIZE = 0
	ST_STRETCH  = 1
	ST_TILE     = 2
)

const (
	HA_LEFT   = 0
	HA_CENTER = 1
	HA_RIGHT  = 2
)

const (
	CA_LEFT   = 0
	CA_CENTER = 1
	CA_RIGHT  = 2
)

const (
	VA_TOP    = 0
	VA_CENTER = 1
	VA_BOTTOM = 2
)

const (
	OT_TOPLEFT           = 0
	OT_TOPRIGHT          = 1
	OT_TOPMIDDLE         = 2
	OT_BOTTOMLEFT        = 3
	OT_BOTTOMRIGHT       = 4
	OT_BOTTOMMIDDLE      = 5
	OT_MIDDLELEFT        = 6
	OT_MIDDLERIGHT       = 7
	OT_LEFTOFCAPTION     = 8
	OT_RIGHTOFCAPTION    = 9
	OT_LEFTOFLASTBUTTON  = 10
	OT_RIGHTOFLASTBUTTON = 11
	OT_ABOVELASTBUTTON   = 12
	OT_BELOWLASTBUTTON   = 13
)

const (
	ICE_NONE   = 0
	ICE_GLOW   = 1
	ICE_SHADOW = 2
	ICE_PULSE  = 3
	ICE_ALPHA  = 4
)

const (
	TST_NONE       = 0
	TST_SINGLE     = 1
	TST_CONTINUOUS = 2
)

const (
	GT_NONE       = 0
	GT_IMAGEGLYPH = 1
	GT_FONTGLYPH  = 2
)

const (
	IST_NONE = 0
	IST_SIZE = 1
	IST_DPI  = 2
)

const (
	TSST_NONE = 0
	TSST_SIZE = 1
	TSST_DPI  = 2
)

const (
	GFST_NONE = 0
	GFST_SIZE = 1
	GFST_DPI  = 2
)
