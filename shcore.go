// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build windows
// +build windows

package win

type MONITOR_DPI_TYPE int32

const (
	MDT_EFFECTIVE_DPI MONITOR_DPI_TYPE = 0
	MDT_ANGULAR_DPI MONITOR_DPI_TYPE = 1
	MDT_RAW_DPI MONITOR_DPI_TYPE = 2
	MDT_DEFAULT = MDT_EFFECTIVE_DPI
)

type SHELL_UI_COMPONENT int32

const (
	SHELL_UI_COMPONENT_TASKBARS SHELL_UI_COMPONENT = 0
	SHELL_UI_COMPONENT_NOTIFICATIONAREA SHELL_UI_COMPONENT = 1
	SHELL_UI_COMPONENT_DESKBAND SHELL_UI_COMPONENT = 2
)

//sys GetDpiForMonitor(hmonitor HMONITOR, dpiType MONITOR_DPI_TYPE, dpiX *uint32, dpiY *uint32) (ret HRESULT) = shcore.GetDpiForMonitor
//sys GetDpiForShellUIComponent(component SHELL_UI_COMPONENT) (ret uint32) = shcore.GetDpiForShellUIComponent
//sys GetScaleFactorForMonitor(hmonitor HMONITOR, scalePercent *uint32) (ret HRESULT) = shcore.GetScaleFactorForMonitor
