// Copyright 2016 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package win

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// Button control messages
const (
	BCM_FIRST            = 0x1600
	BCM_GETIDEALSIZE     = BCM_FIRST + 0x0001
	BCM_SETIMAGELIST     = BCM_FIRST + 0x0002
	BCM_GETIMAGELIST     = BCM_FIRST + 0x0003
	BCM_SETTEXTMARGIN    = BCM_FIRST + 0x0004
	BCM_GETTEXTMARGIN    = BCM_FIRST + 0x0005
	BCM_SETDROPDOWNSTATE = BCM_FIRST + 0x0006
	BCM_SETSPLITINFO     = BCM_FIRST + 0x0007
	BCM_GETSPLITINFO     = BCM_FIRST + 0x0008
	BCM_SETNOTE          = BCM_FIRST + 0x0009
	BCM_GETNOTE          = BCM_FIRST + 0x000A
	BCM_GETNOTELENGTH    = BCM_FIRST + 0x000B
	BCM_SETSHIELD        = BCM_FIRST + 0x000C
)

const (
	CCM_FIRST            = 0x2000
	CCM_LAST             = CCM_FIRST + 0x200
	CCM_SETBKCOLOR       = 8193
	CCM_SETCOLORSCHEME   = 8194
	CCM_GETCOLORSCHEME   = 8195
	CCM_GETDROPTARGET    = 8196
	CCM_SETUNICODEFORMAT = 8197
	CCM_GETUNICODEFORMAT = 8198
	CCM_SETVERSION       = 0x2007
	CCM_GETVERSION       = 0x2008
	CCM_SETNOTIFYWINDOW  = 0x2009
	CCM_SETWINDOWTHEME   = 0x200b
	CCM_DPISCALE         = 0x200c
)

// Common controls styles
const (
	CCS_TOP           = 1
	CCS_NOMOVEY       = 2
	CCS_BOTTOM        = 3
	CCS_NORESIZE      = 4
	CCS_NOPARENTALIGN = 8
	CCS_ADJUSTABLE    = 32
	CCS_NODIVIDER     = 64
	CCS_VERT          = 128
	CCS_LEFT          = 129
	CCS_NOMOVEX       = 130
	CCS_RIGHT         = 131
)

// InitCommonControlsEx flags
const (
	ICC_LISTVIEW_CLASSES   = 1
	ICC_TREEVIEW_CLASSES   = 2
	ICC_BAR_CLASSES        = 4
	ICC_TAB_CLASSES        = 8
	ICC_UPDOWN_CLASS       = 16
	ICC_PROGRESS_CLASS     = 32
	ICC_HOTKEY_CLASS       = 64
	ICC_ANIMATE_CLASS      = 128
	ICC_WIN95_CLASSES      = 255
	ICC_DATE_CLASSES       = 256
	ICC_USEREX_CLASSES     = 512
	ICC_COOL_CLASSES       = 1024
	ICC_INTERNET_CLASSES   = 2048
	ICC_PAGESCROLLER_CLASS = 4096
	ICC_NATIVEFNTCTL_CLASS = 8192
	INFOTIPSIZE            = 1024
	ICC_STANDARD_CLASSES   = 0x00004000
	ICC_LINK_CLASS         = 0x00008000
)

// WM_NOTITY messages
const (
	NM_FIRST              = 0
	NM_OUTOFMEMORY        = ^uint32(0)  // NM_FIRST - 1
	NM_CLICK              = ^uint32(1)  // NM_FIRST - 2
	NM_DBLCLK             = ^uint32(2)  // NM_FIRST - 3
	NM_RETURN             = ^uint32(3)  // NM_FIRST - 4
	NM_RCLICK             = ^uint32(4)  // NM_FIRST - 5
	NM_RDBLCLK            = ^uint32(5)  // NM_FIRST - 6
	NM_SETFOCUS           = ^uint32(6)  // NM_FIRST - 7
	NM_KILLFOCUS          = ^uint32(7)  // NM_FIRST - 8
	NM_CUSTOMDRAW         = ^uint32(11) // NM_FIRST - 12
	NM_HOVER              = ^uint32(12) // NM_FIRST - 13
	NM_NCHITTEST          = ^uint32(13) // NM_FIRST - 14
	NM_KEYDOWN            = ^uint32(14) // NM_FIRST - 15
	NM_RELEASEDCAPTURE    = ^uint32(15) // NM_FIRST - 16
	NM_SETCURSOR          = ^uint32(16) // NM_FIRST - 17
	NM_CHAR               = ^uint32(17) // NM_FIRST - 18
	NM_TOOLTIPSCREATED    = ^uint32(18) // NM_FIRST - 19
	NM_LAST               = ^uint32(98) // NM_FIRST - 99
	TRBN_THUMBPOSCHANGING = 0xfffffa22  // TRBN_FIRST - 1
)

// ProgressBar messages
const (
	PBM_SETPOS      = WM_USER + 2
	PBM_DELTAPOS    = WM_USER + 3
	PBM_SETSTEP     = WM_USER + 4
	PBM_STEPIT      = WM_USER + 5
	PBM_SETMARQUEE  = WM_USER + 10
	PBM_SETRANGE32  = 1030
	PBM_GETRANGE    = 1031
	PBM_GETPOS      = 1032
	PBM_SETBARCOLOR = 1033
	PBM_SETBKCOLOR  = CCM_SETBKCOLOR
)

// ProgressBar states
const (
	PBST_NORMAL = 0x0001
	PBST_ERROR  = 0x0002
	PBST_PAUSED = 0x0003
)

// ProgressBar styles
const (
	PBS_SMOOTH   = 0x01
	PBS_VERTICAL = 0x04
	PBS_MARQUEE  = 0x08
)

// TrackBar (Slider) messages
const (
	TBM_GETPOS      = WM_USER
	TBM_GETRANGEMIN = WM_USER + 1
	TBM_GETRANGEMAX = WM_USER + 2
	TBM_SETPOS      = WM_USER + 5
	TBM_SETRANGEMIN = WM_USER + 7
	TBM_SETRANGEMAX = WM_USER + 8
	TBM_SETPAGESIZE = WM_USER + 21
	TBM_GETPAGESIZE = WM_USER + 22
	TBM_SETLINESIZE = WM_USER + 23
	TBM_GETLINESIZE = WM_USER + 24
)

// TrackBar (Slider) styles
const (
	TBS_VERT     = 0x002
	TBS_TOOLTIPS = 0x100
)

// ImageList creation flags
const (
	ILC_MASK          = 0x00000001
	ILC_COLOR         = 0x00000000
	ILC_COLORDDB      = 0x000000FE
	ILC_COLOR4        = 0x00000004
	ILC_COLOR8        = 0x00000008
	ILC_COLOR16       = 0x00000010
	ILC_COLOR24       = 0x00000018
	ILC_COLOR32       = 0x00000020
	ILC_PALETTE       = 0x00000800
	ILC_MIRROR        = 0x00002000
	ILC_PERITEMMIRROR = 0x00008000
)

// ImageList_Draw[Ex] flags
const (
	ILD_NORMAL      = 0x00000000
	ILD_TRANSPARENT = 0x00000001
	ILD_BLEND25     = 0x00000002
	ILD_BLEND50     = 0x00000004
	ILD_MASK        = 0x00000010
	ILD_IMAGE       = 0x00000020
	ILD_SELECTED    = ILD_BLEND50
	ILD_FOCUS       = ILD_BLEND25
	ILD_BLEND       = ILD_BLEND50
)

// LoadIconMetric flags
const (
	LIM_SMALL = 0
	LIM_LARGE = 1
)

const (
	CDDS_PREPAINT      = 0x00000001
	CDDS_POSTPAINT     = 0x00000002
	CDDS_PREERASE      = 0x00000003
	CDDS_POSTERASE     = 0x00000004
	CDDS_ITEM          = 0x00010000
	CDDS_ITEMPREPAINT  = CDDS_ITEM | CDDS_PREPAINT
	CDDS_ITEMPOSTPAINT = CDDS_ITEM | CDDS_POSTPAINT
	CDDS_ITEMPREERASE  = CDDS_ITEM | CDDS_PREERASE
	CDDS_ITEMPOSTERASE = CDDS_ITEM | CDDS_POSTERASE
	CDDS_SUBITEM       = 0x00020000
)

const (
	CDIS_SELECTED         = 0x0001
	CDIS_GRAYED           = 0x0002
	CDIS_DISABLED         = 0x0004
	CDIS_CHECKED          = 0x0008
	CDIS_FOCUS            = 0x0010
	CDIS_DEFAULT          = 0x0020
	CDIS_HOT              = 0x0040
	CDIS_MARKED           = 0x0080
	CDIS_INDETERMINATE    = 0x0100
	CDIS_SHOWKEYBOARDCUES = 0x0200
	CDIS_NEARHOT          = 0x0400
	CDIS_OTHERSIDEHOT     = 0x0800
	CDIS_DROPHILITED      = 0x1000
)

const (
	CDRF_DODEFAULT         = 0x00000000
	CDRF_NEWFONT           = 0x00000002
	CDRF_SKIPDEFAULT       = 0x00000004
	CDRF_DOERASE           = 0x00000008
	CDRF_NOTIFYPOSTPAINT   = 0x00000010
	CDRF_NOTIFYITEMDRAW    = 0x00000020
	CDRF_NOTIFYSUBITEMDRAW = 0x00000020
	CDRF_NOTIFYPOSTERASE   = 0x00000040
	CDRF_SKIPPOSTPAINT     = 0x00000100
)

const (
	LVIR_BOUNDS       = 0
	LVIR_ICON         = 1
	LVIR_LABEL        = 2
	LVIR_SELECTBOUNDS = 3
)

const (
	LPSTR_TEXTCALLBACK = ^uintptr(0)
	I_CHILDRENCALLBACK = -1
	I_IMAGECALLBACK    = -1
	I_IMAGENONE        = -2
)

type HIMAGELIST HANDLE

type INITCOMMONCONTROLSEX struct {
	DwSize, DwICC uint32
}

type NMCUSTOMDRAW struct {
	Hdr         NMHDR
	DwDrawStage uint32
	Hdc         HDC
	Rc          RECT
	DwItemSpec  uintptr
	UItemState  uint32
	LItemlParam uintptr
}

var (
	// Functions
	imageList_Add         *windows.LazyProc
	imageList_AddMasked   *windows.LazyProc
	imageList_Create      *windows.LazyProc
	imageList_Destroy     *windows.LazyProc
	imageList_DrawEx      *windows.LazyProc
	imageList_ReplaceIcon *windows.LazyProc
	initCommonControlsEx  *windows.LazyProc
	loadIconMetric        *windows.LazyProc
	loadIconWithScaleDown *windows.LazyProc
)

func init() {
	// Functions
	imageList_Add = modcomctl32.NewProc("ImageList_Add")
	imageList_AddMasked = modcomctl32.NewProc("ImageList_AddMasked")
	imageList_Create = modcomctl32.NewProc("ImageList_Create")
	imageList_Destroy = modcomctl32.NewProc("ImageList_Destroy")
	imageList_DrawEx = modcomctl32.NewProc("ImageList_DrawEx")
	imageList_ReplaceIcon = modcomctl32.NewProc("ImageList_ReplaceIcon")
	initCommonControlsEx = modcomctl32.NewProc("InitCommonControlsEx")
	loadIconMetric = modcomctl32.NewProc("LoadIconMetric")
	loadIconWithScaleDown = modcomctl32.NewProc("LoadIconWithScaleDown")
}

func ImageList_Add(himl HIMAGELIST, hbmImage, hbmMask HBITMAP) int32 {
	ret, _, _ := syscall.Syscall(imageList_Add.Addr(), 3,
		uintptr(himl),
		uintptr(hbmImage),
		uintptr(hbmMask))

	return int32(ret)
}

func ImageList_AddMasked(himl HIMAGELIST, hbmImage HBITMAP, crMask COLORREF) int32 {
	ret, _, _ := syscall.Syscall(imageList_AddMasked.Addr(), 3,
		uintptr(himl),
		uintptr(hbmImage),
		uintptr(crMask))

	return int32(ret)
}

func ImageList_Create(cx, cy int32, flags uint32, cInitial, cGrow int32) HIMAGELIST {
	ret, _, _ := syscall.Syscall6(imageList_Create.Addr(), 5,
		uintptr(cx),
		uintptr(cy),
		uintptr(flags),
		uintptr(cInitial),
		uintptr(cGrow),
		0)

	return HIMAGELIST(ret)
}

func ImageList_Destroy(hIml HIMAGELIST) bool {
	ret, _, _ := syscall.Syscall(imageList_Destroy.Addr(), 1,
		uintptr(hIml),
		0,
		0)

	return ret != 0
}

func ImageList_DrawEx(himl HIMAGELIST, i int32, hdcDst HDC, x, y, dx, dy int32, rgbBk COLORREF, rgbFg COLORREF, fStyle uint32) bool {
	ret, _, _ := syscall.Syscall12(imageList_DrawEx.Addr(), 10,
		uintptr(himl),
		uintptr(i),
		uintptr(hdcDst),
		uintptr(x),
		uintptr(y),
		uintptr(dx),
		uintptr(dy),
		uintptr(rgbBk),
		uintptr(rgbFg),
		uintptr(fStyle),
		0,
		0)

	return ret != 0
}

func ImageList_ReplaceIcon(himl HIMAGELIST, i int32, hicon HICON) int32 {
	ret, _, _ := syscall.Syscall(imageList_ReplaceIcon.Addr(), 3,
		uintptr(himl),
		uintptr(i),
		uintptr(hicon))

	return int32(ret)
}

func InitCommonControlsEx(lpInitCtrls *INITCOMMONCONTROLSEX) bool {
	ret, _, _ := syscall.Syscall(initCommonControlsEx.Addr(), 1,
		uintptr(unsafe.Pointer(lpInitCtrls)),
		0,
		0)

	return ret != 0
}

func LoadIconMetric(hInstance HINSTANCE, lpIconName *uint16, lims int32, hicon *HICON) HRESULT {
	if loadIconMetric.Find() != nil {
		return HRESULT(0)
	}
	ret, _, _ := syscall.Syscall6(loadIconMetric.Addr(), 4,
		uintptr(hInstance),
		uintptr(unsafe.Pointer(lpIconName)),
		uintptr(lims),
		uintptr(unsafe.Pointer(hicon)),
		0,
		0)

	return HRESULT(ret)
}

func LoadIconWithScaleDown(hInstance HINSTANCE, lpIconName *uint16, w int32, h int32, hicon *HICON) HRESULT {
	if loadIconWithScaleDown.Find() != nil {
		return HRESULT(0)
	}
	ret, _, _ := syscall.Syscall6(loadIconWithScaleDown.Addr(), 5,
		uintptr(hInstance),
		uintptr(unsafe.Pointer(lpIconName)),
		uintptr(w),
		uintptr(h),
		uintptr(unsafe.Pointer(hicon)),
		0)

	return HRESULT(ret)
}

type TASKDIALOG_FLAGS uint32

const (
	TDF_ENABLE_HYPERLINKS           TASKDIALOG_FLAGS = 0x0001
	TDF_USE_HICON_MAIN              TASKDIALOG_FLAGS = 0x0002
	TDF_USE_HICON_FOOTER            TASKDIALOG_FLAGS = 0x0004
	TDF_ALLOW_DIALOG_CANCELLATION   TASKDIALOG_FLAGS = 0x0008
	TDF_USE_COMMAND_LINKS           TASKDIALOG_FLAGS = 0x0010
	TDF_USE_COMMAND_LINKS_NO_ICON   TASKDIALOG_FLAGS = 0x0020
	TDF_EXPAND_FOOTER_AREA          TASKDIALOG_FLAGS = 0x0040
	TDF_EXPANDED_BY_DEFAULT         TASKDIALOG_FLAGS = 0x0080
	TDF_VERIFICATION_FLAG_CHECKED   TASKDIALOG_FLAGS = 0x0100
	TDF_SHOW_PROGRESS_BAR           TASKDIALOG_FLAGS = 0x0200
	TDF_SHOW_MARQUEE_PROGRESS_BAR   TASKDIALOG_FLAGS = 0x0400
	TDF_CALLBACK_TIMER              TASKDIALOG_FLAGS = 0x0800
	TDF_POSITION_RELATIVE_TO_WINDOW TASKDIALOG_FLAGS = 0x1000
	TDF_RTL_LAYOUT                  TASKDIALOG_FLAGS = 0x2000
	TDF_NO_DEFAULT_RADIO_BUTTON     TASKDIALOG_FLAGS = 0x4000
	TDF_CAN_BE_MINIMIZED            TASKDIALOG_FLAGS = 0x8000
	TDF_NO_SET_FOREGROUND           TASKDIALOG_FLAGS = 0x00010000
	TDF_SIZE_TO_CONTENT             TASKDIALOG_FLAGS = 0x01000000
)

const (
	TDM_NAVIGATE_PAGE                       = WM_USER + 101
	TDM_CLICK_BUTTON                        = WM_USER + 102
	TDM_SET_MARQUEE_PROGRESS_BAR            = WM_USER + 103
	TDM_SET_PROGRESS_BAR_STATE              = WM_USER + 104
	TDM_SET_PROGRESS_BAR_RANGE              = WM_USER + 105
	TDM_SET_PROGRESS_BAR_POS                = WM_USER + 106
	TDM_SET_PROGRESS_BAR_MARQUEE            = WM_USER + 107
	TDM_SET_ELEMENT_TEXT                    = WM_USER + 108
	TDM_CLICK_RADIO_BUTTON                  = WM_USER + 110
	TDM_ENABLE_BUTTON                       = WM_USER + 111
	TDM_ENABLE_RADIO_BUTTON                 = WM_USER + 112
	TDM_CLICK_VERIFICATION                  = WM_USER + 113
	TDM_UPDATE_ELEMENT_TEXT                 = WM_USER + 114
	TDM_SET_BUTTON_ELEVATION_REQUIRED_STATE = WM_USER + 115
	TDM_UPDATE_ICON                         = WM_USER + 116
)

const (
	TDN_CREATED                = 0
	TDN_NAVIGATED              = 1
	TDN_BUTTON_CLICKED         = 2
	TDN_HYPERLINK_CLICKED      = 3
	TDN_TIMER                  = 4
	TDN_DESTROYED              = 5
	TDN_RADIO_BUTTON_CLICKED   = 6
	TDN_DIALOG_CONSTRUCTED     = 7
	TDN_VERIFICATION_CLICKED   = 8
	TDN_HELP                   = 9
	TDN_EXPANDO_BUTTON_CLICKED = 10
)

// TASKDIALOG_BUTTON_UNPACKED is nearly identical to TASKDIALOG_BUTTON in the Windows
// SDK, except for the fact that that the SDK version is packed. Since Go cannot
// grok that, we implement a Pack method on this struct to encode its contents
// correctly for Windows.
type TASKDIALOG_BUTTON_UNPACKED struct {
	ButtonID   int32
	ButtonText *uint16
}

// TASKDIALOG_BUTTON is opaque because it is a packed data structure.
// Use (*TASKDIALOG_BUTTON_UNPACKED).Pack() to write one to a bytes.Buffer.
type TASKDIALOG_BUTTON struct {
	_ [sizeTASKDIALOG_BUTTON]byte
}

// Pack writes the contents of u to buf in a packed format. We use a Buffer
// because walk encodes TASKDIALOG_BUTTONs from a slice.
func (u *TASKDIALOG_BUTTON_UNPACKED) Pack(buf *bytes.Buffer) error {
	if err := binary.Write(buf, packByteOrder, u.ButtonID); err != nil {
		return err
	}
	return binary.Write(buf, packByteOrder, packPtr(uintptr(unsafe.Pointer(u.ButtonText))))
}

type TASKDIALOG_ELEMENTS int32

const (
	TDE_CONTENT              TASKDIALOG_ELEMENTS = 0
	TDE_EXPANDED_INFORMATION TASKDIALOG_ELEMENTS = 1
	TDE_FOOTER               TASKDIALOG_ELEMENTS = 2
	TDE_MAIN_INSTRUCTION     TASKDIALOG_ELEMENTS = 3
)

type TASKDIALOG_ICON_ELEMENTS int32

const (
	TDIE_ICON_MAIN   TASKDIALOG_ICON_ELEMENTS = 0
	TDIE_ICON_FOOTER TASKDIALOG_ICON_ELEMENTS = 1
)

const (
	TD_WARNING_ICON     uintptr = 0xFFFF
	TD_ERROR_ICON       uintptr = 0xFFFE
	TD_INFORMATION_ICON uintptr = 0xFFFD
	TD_SHIELD_ICON      uintptr = 0xFFFC
)

type TASKDIALOG_COMMON_BUTTON_FLAGS uint32

const (
	TDCBF_OK_BUTTON     TASKDIALOG_COMMON_BUTTON_FLAGS = 0x0001
	TDCBF_YES_BUTTON    TASKDIALOG_COMMON_BUTTON_FLAGS = 0x0002
	TDCBF_NO_BUTTON     TASKDIALOG_COMMON_BUTTON_FLAGS = 0x0004
	TDCBF_CANCEL_BUTTON TASKDIALOG_COMMON_BUTTON_FLAGS = 0x0008
	TDCBF_RETRY_BUTTON  TASKDIALOG_COMMON_BUTTON_FLAGS = 0x0010
	TDCBF_CLOSE_BUTTON  TASKDIALOG_COMMON_BUTTON_FLAGS = 0x0020
)

// TASKDIALOGCONFIG_UNPACKED is nearly identical to TASKDIALOGCONFIG in the
// Windows SDK, except that the cbSize field is omitted and for the fact that
// the SDK version is packed. Since Go cannot grok packed structs, we implement
// a Pack method on this struct to encode its contents correctly for Windows.
type TASKDIALOGCONFIG_UNPACKED struct {
	HWNDParent           HWND
	HInstance            HINSTANCE
	Flags                TASKDIALOG_FLAGS
	CommonButtons        TASKDIALOG_COMMON_BUTTON_FLAGS
	WindowTitle          *uint16
	MainIcon             uintptr
	MainInstruction      *uint16
	Content              *uint16
	CButtons             uint32
	PButtons             *TASKDIALOG_BUTTON
	DefaultButton        int32
	CRadioButtons        uint32
	PRadioButtons        *TASKDIALOG_BUTTON
	DefaultRadioButton   int32
	VerificationText     *uint16
	ExpandedInformation  *uint16
	ExpandedControlText  *uint16
	CollapsedControlText *uint16
	FooterIcon           uintptr
	Footer               *uint16
	Callback             uintptr
	CallbackData         uintptr
	Width                uint32
}

// TASKDIALOGCONFIG is opaque because it is a packed data structure.
// Use (*TASKDIALOGCONFIG_UNPACKED).Pack() to obtain one.
type TASKDIALOGCONFIG struct {
	_ [sizeTASKDIALOGCONFIG]byte
}

// Pack takes the contents of u and writes them in a packed format to a
// TASKDIALOGCONFIG, which is returned as the result.
func (u *TASKDIALOGCONFIG_UNPACKED) Pack() *TASKDIALOGCONFIG {
	buf := bytes.NewBuffer(make([]byte, 0, sizeTASKDIALOGCONFIG))

	v := reflect.ValueOf(u).Elem()

	// Write the actual packed size as the cbSize field.
	binary.Write(buf, packByteOrder, uint32(sizeTASKDIALOGCONFIG))

	for i, n := 0, v.NumField(); i < n; i++ {
		fv := v.Field(i)
		switch fv.Type().Kind() {
		case reflect.Pointer:
			if err := binary.Write(buf, packByteOrder, packPtr(uintptr(fv.UnsafePointer()))); err != nil {
				panic(err)
			}
		case reflect.Uintptr:
			if err := binary.Write(buf, packByteOrder, packPtr(fv.Uint())); err != nil {
				panic(err)
			}
		default:
			if err := binary.Write(buf, packByteOrder, fv.Interface()); err != nil {
				panic(err)
			}
		}
	}

	bb := buf.Bytes()
	if len(bb) != sizeTASKDIALOGCONFIG {
		panic(fmt.Sprintf("(*TASKDIALOGCONFIG_UNPACKED).Pack() invalid size: got %d, want %d", len(bb), sizeTASKDIALOGCONFIG))
	}

	return (*TASKDIALOGCONFIG)(unsafe.Pointer(&bb[0]))
}

//sys TaskDialogIndirect(pTaskConfig *TASKDIALOGCONFIG, pnButton *int32, pnRadioButton *int32, pfVerificationFlagChecked *BOOL) (ret HRESULT) = comctl32.TaskDialogIndirect
