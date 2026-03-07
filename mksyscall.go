// Copyright (c) Tailscale Inc & AUTHORS
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package win

//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output zsyscall_windows.go comctl32.go dwmapi.go gdiplus.go kernel32.go shcore.go user32.go uxtheme.go
//go:generate go run golang.org/x/tools/cmd/goimports -w zsyscall_windows.go
