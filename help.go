// Copyright 2021 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package win

const (
	HELPINFO_WINDOW   = 1
	HELPINFO_MENUITEM = 2
)

// HELPINFO is the structure sent on WM_HELP messages.
//
// See https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-helpinfo
type HELPINFO struct {
	Size        uint32 // cbSize, the struct size in bytes
	ContextType int32  // either HELPINFO_WINDOW or HELPINFO_MENUITEM
	CtrlId      int32
	ItemHandle  HANDLE
	ContextId   uintptr
	MousePos    POINT
}
