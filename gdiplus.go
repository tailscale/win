// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package win

import (
	"math"
	"syscall"
	"unsafe"
)

type GpStatus int32

const (
	Ok                        GpStatus = 0
	GenericError              GpStatus = 1
	InvalidParameter          GpStatus = 2
	OutOfMemory               GpStatus = 3
	ObjectBusy                GpStatus = 4
	InsufficientBuffer        GpStatus = 5
	NotImplemented            GpStatus = 6
	Win32Error                GpStatus = 7
	WrongState                GpStatus = 8
	Aborted                   GpStatus = 9
	FileNotFound              GpStatus = 10
	ValueOverflow             GpStatus = 11
	AccessDenied              GpStatus = 12
	UnknownImageFormat        GpStatus = 13
	FontFamilyNotFound        GpStatus = 14
	FontStyleNotFound         GpStatus = 15
	NotTrueTypeFont           GpStatus = 16
	UnsupportedGdiplusVersion GpStatus = 17
	GdiplusNotInitialized     GpStatus = 18
	PropertyNotFound          GpStatus = 19
	PropertyNotSupported      GpStatus = 20
	ProfileNotFound           GpStatus = 21
)

func (s GpStatus) String() string {
	switch s {
	case Ok:
		return "Ok"

	case GenericError:
		return "GenericError"

	case InvalidParameter:
		return "InvalidParameter"

	case OutOfMemory:
		return "OutOfMemory"

	case ObjectBusy:
		return "ObjectBusy"

	case InsufficientBuffer:
		return "InsufficientBuffer"

	case NotImplemented:
		return "NotImplemented"

	case Win32Error:
		return "Win32Error"

	case WrongState:
		return "WrongState"

	case Aborted:
		return "Aborted"

	case FileNotFound:
		return "FileNotFound"

	case ValueOverflow:
		return "ValueOverflow"

	case AccessDenied:
		return "AccessDenied"

	case UnknownImageFormat:
		return "UnknownImageFormat"

	case FontFamilyNotFound:
		return "FontFamilyNotFound"

	case FontStyleNotFound:
		return "FontStyleNotFound"

	case NotTrueTypeFont:
		return "NotTrueTypeFont"

	case UnsupportedGdiplusVersion:
		return "UnsupportedGdiplusVersion"

	case GdiplusNotInitialized:
		return "GdiplusNotInitialized"

	case PropertyNotFound:
		return "PropertyNotFound"

	case PropertyNotSupported:
		return "PropertyNotSupported"

	case ProfileNotFound:
		return "ProfileNotFound"
	}

	return "Unknown Status Value"
}

type GdiplusStartupInput struct {
	GdiplusVersion           uint32
	DebugEventCallback       uintptr
	SuppressBackgroundThread BOOL
	SuppressExternalCodecs   BOOL
}

type GdiplusStartupOutput struct {
	NotificationHook   uintptr
	NotificationUnhook uintptr
}

type GpBitmap GpImage
type GpBrush struct{}
type GpFont struct{}
type GpFontCollection struct{}
type GpFontFamily struct{}
type GpGraphics struct{}
type GpImage struct{}
type GpImageAttributes struct{}
type GpPath struct{}
type GpSolidFill GpBrush
type GpStringFormat struct{}

type GpRect struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
}

type GpRectF struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

type ARGB uint32

type CombineMode int32

const (
	CombineModeReplace    = CombineMode(0)
	CombineModeIntersect  = CombineMode(1)
	CombineModeUnion      = CombineMode(2)
	CombineModeXor        = CombineMode(3)
	CombineModeExclude    = CombineMode(4)
	CombineModeCompliment = CombineMode(5)
)

type CompositingMode int32

const (
	CompositingModeSourceOver = CompositingMode(0)
	CompositingModeSourceCopy = CompositingMode(1)
)

type CompositingQuality int32

const (
	CompositingQualityInvalid        = CompositingQuality(-1)
	CompositingQualityDefault        = CompositingQuality(0)
	CompositingQualityHighSpeed      = CompositingQuality(1)
	CompositingQualityHighQuality    = CompositingQuality(2)
	CompositingQualityGammaCorrected = CompositingQuality(3)
	CompositingQualityAssumeLinear   = CompositingQuality(4)
)

type FillMode int32

const (
	FillModeAlternate = FillMode(0)
	FillModeWinding   = FillMode(1)
)

type FontStyle int32

const (
	FontStyleRegular    = FontStyle(0)
	FontStyleBold       = FontStyle(1)
	FontStyleItalic     = FontStyle(2)
	FontStyleBoldItalic = FontStyle(3)
	FontStyleUnderline  = FontStyle(4)
	FontStyleStrikeout  = FontStyle(8)
)

type InterpolationMode int32

const (
	InterpolationModeInvalid             = InterpolationMode(-1)
	InterpolationModeDefault             = InterpolationMode(0)
	InterpolationModeLowQuality          = InterpolationMode(1)
	InterpolationModeHighQuality         = InterpolationMode(2)
	InterpolationModeBilinear            = InterpolationMode(3)
	InterpolationModeBicubic             = InterpolationMode(4)
	InterpolationModeNearestNeighbor     = InterpolationMode(5)
	InterpolationModeHighQualityBilinear = InterpolationMode(6)
	InterpolationModeHighQualityBicubic  = InterpolationMode(7)
)

type PixelFormat int32

const (
	PixelFormatIndexed   = 0x00010000
	PixelFormatGDI       = 0x00020000
	PixelFormatAlpha     = 0x00040000
	PixelFormatPAlpha    = 0x00080000
	PixelFormatExtended  = 0x00100000
	PixelFormatCanonical = 0x00200000

	PixelFormatUndefined = 0
	PixelFormatDontCare  = 0

	PixelFormat1bppIndexed    = (1 | (1 << 8) | PixelFormatIndexed | PixelFormatGDI)
	PixelFormat4bppIndexed    = (2 | (4 << 8) | PixelFormatIndexed | PixelFormatGDI)
	PixelFormat8bppIndexed    = (3 | (8 << 8) | PixelFormatIndexed | PixelFormatGDI)
	PixelFormat16bppGrayScale = (4 | (16 << 8) | PixelFormatExtended)
	PixelFormat16bppRGB555    = (5 | (16 << 8) | PixelFormatGDI)
	PixelFormat16bppRGB565    = (6 | (16 << 8) | PixelFormatGDI)
	PixelFormat16bppARGB1555  = (7 | (16 << 8) | PixelFormatAlpha | PixelFormatGDI)
	PixelFormat24bppRGB       = (8 | (24 << 8) | PixelFormatGDI)
	PixelFormat32bppRGB       = (9 | (32 << 8) | PixelFormatGDI)
	PixelFormat32bppARGB      = (10 | (32 << 8) | PixelFormatAlpha | PixelFormatGDI | PixelFormatCanonical)
	PixelFormat32bppPARGB     = (11 | (32 << 8) | PixelFormatAlpha | PixelFormatPAlpha | PixelFormatGDI)
	PixelFormat48bppRGB       = (12 | (48 << 8) | PixelFormatExtended)
	PixelFormat64bppARGB      = (13 | (64 << 8) | PixelFormatAlpha | PixelFormatCanonical | PixelFormatExtended)
	PixelFormat64bppPARGB     = (14 | (64 << 8) | PixelFormatAlpha | PixelFormatPAlpha | PixelFormatExtended)
	PixelFormat32bppCMYK      = (15 | (32 << 8))
)

type PixelOffsetMode int32

const (
	PixelOffsetModeInvalid     = PixelOffsetMode(-1)
	PixelOffsetModeDefault     = PixelOffsetMode(0)
	PixelOffsetModeHighSpeed   = PixelOffsetMode(1)
	PixelOffsetModeHighQuality = PixelOffsetMode(2)
	PixelOffsetModeNone        = PixelOffsetMode(3)
	PixelOffsetModeHalf        = PixelOffsetMode(4)
)

type SmoothingMode int32

const (
	SmoothingModeInvalid      = SmoothingMode(-1)
	SmoothingModeDefault      = SmoothingMode(0)
	SmoothingModeHighSpeed    = SmoothingMode(1)
	SmoothingModeHighQuality  = SmoothingMode(2)
	SmoothingModeNone         = SmoothingMode(3)
	SmoothingModeAntiAlias8x4 = SmoothingMode(4)
	SmoothingModeAntiAlias    = SmoothingModeAntiAlias8x4
	SmoothingModeAntiAlias8x8 = SmoothingMode(5)
)

type StringAlignment int32

const (
	StringAlignmentNear   = StringAlignment(0)
	StringAlignmentCenter = StringAlignment(1)
	StringAlignmentFar    = StringAlignment(2)
)

type StringFormatFlags int32

const (
	StringFormatFlagsDirectionRightToLeft  = StringFormatFlags(0x00000001)
	StringFormatFlagsDirectionVertical     = StringFormatFlags(0x00000002)
	StringFormatFlagsNoFitBlackBox         = StringFormatFlags(0x00000004)
	StringFormatFlagsDisplayFormatControl  = StringFormatFlags(0x00000020)
	StringFormatFlagsNoFontFallback        = StringFormatFlags(0x00000400)
	StringFormatFlagsMeasureTrailingSpaces = StringFormatFlags(0x00000800)
	StringFormatFlagsNoWrap                = StringFormatFlags(0x00001000)
	StringFormatFlagsLineLimit             = StringFormatFlags(0x00002000)
	StringFormatFlagsNoClip                = StringFormatFlags(0x00004000)
	StringFormatFlagsBypassGDI             = StringFormatFlags(-((0x80000000 ^ 0xFFFFFFFF) + 1))
)

type TextRenderingHint int32

const (
	TextRenderingHintSystemDefault            = TextRenderingHint(0)
	TextRenderingHintSingleBitPerPixelGridFit = TextRenderingHint(1)
	TextRenderingHintSingleBitPerPixel        = TextRenderingHint(2)
	TextRenderingHintAntiAliasGridFit         = TextRenderingHint(3)
	TextRenderingHintAntiAlias                = TextRenderingHint(4)
	TextRenderingHintClearTypeGridFit         = TextRenderingHint(5)
)

type Unit int32

const (
	UnitWorld      = Unit(0)
	UnitDisplay    = Unit(1)
	UnitPixel      = Unit(2)
	UnitPoint      = Unit(3)
	UnitInch       = Unit(4)
	UnitDocument   = Unit(5)
	UnitMillimeter = Unit(6)
)

var (
	gdipBitmapSetResolution = modgdiplus.NewProc("GdipBitmapSetResolution")
	gdipCreateFont          = modgdiplus.NewProc("GdipCreateFont")
)

func GdipBitmapSetResolution(bitmap *GpBitmap, xdpi float32, ydpi float32) GpStatus {
	ret, _, _ := syscall.SyscallN(gdipBitmapSetResolution.Addr(),
		uintptr(unsafe.Pointer(bitmap)),
		uintptr(math.Float32bits(xdpi)),
		uintptr(math.Float32bits(ydpi)),
	)

	return GpStatus(ret)
}

func GdipCreateFont(fontFamily *GpFontFamily, emSize float32, style FontStyle, unit Unit, font **GpFont) GpStatus {
	ret, _, _ := syscall.SyscallN(gdipCreateFont.Addr(),
		uintptr(unsafe.Pointer(fontFamily)),
		uintptr(math.Float32bits(emSize)),
		uintptr(style),
		uintptr(unit),
		uintptr(unsafe.Pointer(font)),
	)

	return GpStatus(ret)
}

//sys GdipAddPathEllipseI(path *GpPath, x int32, y int32, width int32, height int32) (ret GpStatus) = gdiplus.GdipAddPathEllipseI
//sys GdipCreateBitmapFromFile(filename *uint16, bitmap **GpBitmap) (ret GpStatus) = gdiplus.GdipCreateBitmapFromFile
//sys GdipCreateBitmapFromGraphics(width int32, height int32, graphics *GpGraphics, bitmap **GpBitmap) (ret GpStatus) = gdiplus.GdipCreateBitmapFromGraphics
//sys GdipCreateBitmapFromHBITMAP(hbm HBITMAP, hpal HPALETTE, bitmap **GpBitmap) (ret GpStatus) = gdiplus.GdipCreateBitmapFromHBITMAP
//sys GdipCreateBitmapFromHICON(hicon HICON, bitmap **GpBitmap) (ret GpStatus) = gdiplus.GdipCreateBitmapFromHICON
//sys GdipCreateBitmapFromScan0(width int32, height int32, stride int32, format PixelFormat, scan0 *byte, bitmap **GpBitmap) (ret GpStatus) = gdiplus.GdipCreateBitmapFromScan0
//sys GdipCreateBitmapFromStream(stream *com.IStreamABI, bitmap **GpBitmap) (ret GpStatus) = gdiplus.GdipCreateBitmapFromStream
//sys GdipCreateFontFamilyFromName(name *uint16, collection *GpFontCollection, family **GpFontFamily) (ret GpStatus) = gdiplus.GdipCreateFontFamilyFromName
//sys GdipCreateFromHDC(hdc HDC, graphics **GpGraphics) (ret GpStatus) = gdiplus.GdipCreateFromHDC
//sys GdipCreateHBITMAPFromBitmap(bitmap *GpBitmap, hbmReturn *HBITMAP, background ARGB) (ret GpStatus) = gdiplus.GdipCreateHBITMAPFromBitmap
//sys GdipCreateHICONFromBitmap(bitmap *GpBitmap, hbmReturn *HICON) (ret GpStatus) = gdiplus.GdipCreateHICONFromBitmap
//sys GdipCreatePath(fillMode FillMode, path **GpPath) (ret GpStatus) = gdiplus.GdipCreatePath
//sys GdipCreateSolidFill(color ARGB, brush **GpSolidFill) (ret GpStatus) = gdiplus.GdipCreateSolidFill
//sys GdipCreateStringFormat(flags StringFormatFlags, language LANGID, format **GpStringFormat) (ret GpStatus) = gdiplus.GdipCreateStringFormat
//sys GdipDeleteBrush(brush *GpBrush) (ret GpStatus) = gdiplus.GdipDeleteBrush
//sys GdipDeleteFont(font *GpFont) (ret GpStatus) = gdiplus.GdipDeleteFont
//sys GdipDeleteFontFamily(family *GpFontFamily) (ret GpStatus) = gdiplus.GdipDeleteFontFamily
//sys GdipDeleteGraphics(graphics *GpGraphics) (ret GpStatus) = gdiplus.GdipDeleteGraphics
//sys GdipDeletePath(path *GpPath) (ret GpStatus) = gdiplus.GdipDeletePath
//sys GdipDeleteStringFormat(format *GpStringFormat) (ret GpStatus) = gdiplus.GdipDeleteStringFormat
//sys GdipDisposeImage(image *GpImage) (ret GpStatus) = gdiplus.GdipDisposeImage
//sys GdipDrawImageRectI(graphics *GpGraphics, image *GpImage, x int32, y int32, width int32, height int32) (ret GpStatus) = gdiplus.GdipDrawImageRectI
//sys GdipDrawImageRectRectI(graphics *GpGraphics, image *GpImage, dstX int32, dstY int32, dstWidth int32, dstHeight int32, srcX int32, srcY int32, srcWidth int32, srcHeight int32, srcUnit Unit, imgAttrs *GpImageAttributes, callback uintptr, callbackData uintptr) (ret GpStatus) = gdiplus.GdipDrawImageRectRectI
//sys GdipDrawString(graphics *GpGraphics, text *uint16, textLength int32, font *GpFont, rectf *GpRectF, strFmt *GpStringFormat, brush *GpBrush) (ret GpStatus) = gdiplus.GdipDrawString
//sys GdipFillEllipseI(graphics *GpGraphics, brush *GpBrush, x int32, y int32, width int32, height int32) (ret GpStatus) = gdiplus.GdipFillEllipseI
//sys GdipGetGenericFontFamilyMonospace(family **GpFontFamily) (ret GpStatus) = gdiplus.GdipGetGenericFontFamilyMonospace
//sys GdipGetGenericFontFamilySansSerif(family **GpFontFamily) (ret GpStatus) = gdiplus.GdipGetGenericFontFamilySansSerif
//sys GdipGetGenericFontFamilySerif(family **GpFontFamily) (ret GpStatus) = gdiplus.GdipGetGenericFontFamilySerif
//sys GdipGetImageDimension(image *GpImage, width *float32, height *float32) (ret GpStatus) = gdiplus.GdipGetImageDimension
//sys GdipGetImageGraphicsContext(image *GpImage, graphics **GpGraphics) (ret GpStatus) = gdiplus.GdipGetImageGraphicsContext
//sys GdipGetImageHeight(image *GpImage, height *uint32) (ret GpStatus) = gdiplus.GdipGetImageHeight
//sys GdipGetImageHorizontalResolution(image *GpImage, resolution *float32) (ret GpStatus) = gdiplus.GdipGetImageHorizontalResolution
//sys GdipGetImageVerticalResolution(image *GpImage, resolution *float32) (ret GpStatus) = gdiplus.GdipGetImageVerticalResolution
//sys GdipGetImageWidth(image *GpImage, width *uint32) (ret GpStatus) = gdiplus.GdipGetImageWidth
//sys GdipGetCompositingMode(graphics *GpGraphics, compositingMode *CompositingMode) (ret GpStatus) = gdiplus.GdipGetCompositingMode
//sys GdipGraphicsClear(graphics *GpGraphics, color ARGB) (ret GpStatus) = gdiplus.GdipGraphicsClear
//sys GdiplusStartup(token *uintptr, input *GdiplusStartupInput, output *GdiplusStartupOutput) (ret GpStatus) = gdiplus.GdiplusStartup
//sys GdipResetClip(graphics *GpGraphics) (ret GpStatus) = gdiplus.GdipResetClip
//sys GdipSetClipPath(graphics *GpGraphics, path *GpPath, combineMode CombineMode) (ret GpStatus) = gdiplus.GdipSetClipPath
//sys GdipSetCompositingMode(graphics *GpGraphics, compositingMode CompositingMode) (ret GpStatus) = gdiplus.GdipSetCompositingMode
//sys GdipSetCompositingQuality(graphics *GpGraphics, compositingQuality CompositingQuality) (ret GpStatus) = gdiplus.GdipSetCompositingQuality
//sys GdipSetInterpolationMode(graphics *GpGraphics, interpolationMode InterpolationMode) (ret GpStatus) = gdiplus.GdipSetInterpolationMode
//sys GdipSetPixelOffsetMode(graphics *GpGraphics, pixelOffsetMode PixelOffsetMode) (ret GpStatus) = gdiplus.GdipSetPixelOffsetMode
//sys GdipSetSmoothingMode(graphics *GpGraphics, smoothingMode SmoothingMode) (ret GpStatus) = gdiplus.GdipSetSmoothingMode
//sys GdipSetStringFormatAlign(format *GpStringFormat, align StringAlignment) (ret GpStatus) = gdiplus.GdipSetStringFormatAlign
//sys GdipSetStringFormatLineAlign(format *GpStringFormat, align StringAlignment) (ret GpStatus) = gdiplus.GdipSetStringFormatLineAlign
//sys GdipSetTextRenderingHint(graphics *GpGraphics, mode TextRenderingHint) (ret GpStatus) = gdiplus.GdipSetTextRenderingHint
