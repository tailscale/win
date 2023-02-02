// Code generated by 'go generate'; DO NOT EDIT.

package win

import (
	"syscall"
	"unsafe"

	"github.com/dblohm7/wingoes/com"
	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modgdiplus = windows.NewLazySystemDLL("gdiplus.dll")

	procGdipAddPathEllipseI               = modgdiplus.NewProc("GdipAddPathEllipseI")
	procGdipCreateBitmapFromFile          = modgdiplus.NewProc("GdipCreateBitmapFromFile")
	procGdipCreateBitmapFromGraphics      = modgdiplus.NewProc("GdipCreateBitmapFromGraphics")
	procGdipCreateBitmapFromHBITMAP       = modgdiplus.NewProc("GdipCreateBitmapFromHBITMAP")
	procGdipCreateBitmapFromScan0         = modgdiplus.NewProc("GdipCreateBitmapFromScan0")
	procGdipCreateBitmapFromStream        = modgdiplus.NewProc("GdipCreateBitmapFromStream")
	procGdipCreateFontFamilyFromName      = modgdiplus.NewProc("GdipCreateFontFamilyFromName")
	procGdipCreateFromHDC                 = modgdiplus.NewProc("GdipCreateFromHDC")
	procGdipCreateHBITMAPFromBitmap       = modgdiplus.NewProc("GdipCreateHBITMAPFromBitmap")
	procGdipCreatePath                    = modgdiplus.NewProc("GdipCreatePath")
	procGdipCreateSolidFill               = modgdiplus.NewProc("GdipCreateSolidFill")
	procGdipCreateStringFormat            = modgdiplus.NewProc("GdipCreateStringFormat")
	procGdipDeleteBrush                   = modgdiplus.NewProc("GdipDeleteBrush")
	procGdipDeleteFont                    = modgdiplus.NewProc("GdipDeleteFont")
	procGdipDeleteFontFamily              = modgdiplus.NewProc("GdipDeleteFontFamily")
	procGdipDeleteGraphics                = modgdiplus.NewProc("GdipDeleteGraphics")
	procGdipDeletePath                    = modgdiplus.NewProc("GdipDeletePath")
	procGdipDeleteStringFormat            = modgdiplus.NewProc("GdipDeleteStringFormat")
	procGdipDisposeImage                  = modgdiplus.NewProc("GdipDisposeImage")
	procGdipDrawImageRectI                = modgdiplus.NewProc("GdipDrawImageRectI")
	procGdipDrawImageRectRectI            = modgdiplus.NewProc("GdipDrawImageRectRectI")
	procGdipDrawString                    = modgdiplus.NewProc("GdipDrawString")
	procGdipFillEllipseI                  = modgdiplus.NewProc("GdipFillEllipseI")
	procGdipGetCompositingMode            = modgdiplus.NewProc("GdipGetCompositingMode")
	procGdipGetGenericFontFamilyMonospace = modgdiplus.NewProc("GdipGetGenericFontFamilyMonospace")
	procGdipGetGenericFontFamilySansSerif = modgdiplus.NewProc("GdipGetGenericFontFamilySansSerif")
	procGdipGetGenericFontFamilySerif     = modgdiplus.NewProc("GdipGetGenericFontFamilySerif")
	procGdipGetImageDimension             = modgdiplus.NewProc("GdipGetImageDimension")
	procGdipGetImageGraphicsContext       = modgdiplus.NewProc("GdipGetImageGraphicsContext")
	procGdipGetImageHeight                = modgdiplus.NewProc("GdipGetImageHeight")
	procGdipGetImageHorizontalResolution  = modgdiplus.NewProc("GdipGetImageHorizontalResolution")
	procGdipGetImageVerticalResolution    = modgdiplus.NewProc("GdipGetImageVerticalResolution")
	procGdipGetImageWidth                 = modgdiplus.NewProc("GdipGetImageWidth")
	procGdipGraphicsClear                 = modgdiplus.NewProc("GdipGraphicsClear")
	procGdipResetClip                     = modgdiplus.NewProc("GdipResetClip")
	procGdipSetClipPath                   = modgdiplus.NewProc("GdipSetClipPath")
	procGdipSetCompositingMode            = modgdiplus.NewProc("GdipSetCompositingMode")
	procGdipSetCompositingQuality         = modgdiplus.NewProc("GdipSetCompositingQuality")
	procGdipSetInterpolationMode          = modgdiplus.NewProc("GdipSetInterpolationMode")
	procGdipSetPixelOffsetMode            = modgdiplus.NewProc("GdipSetPixelOffsetMode")
	procGdipSetSmoothingMode              = modgdiplus.NewProc("GdipSetSmoothingMode")
	procGdipSetStringFormatAlign          = modgdiplus.NewProc("GdipSetStringFormatAlign")
	procGdipSetStringFormatLineAlign      = modgdiplus.NewProc("GdipSetStringFormatLineAlign")
	procGdipSetTextRenderingHint          = modgdiplus.NewProc("GdipSetTextRenderingHint")
	procGdiplusStartup                    = modgdiplus.NewProc("GdiplusStartup")
)

func GdipAddPathEllipseI(path *GpPath, x int32, y int32, width int32, height int32) (ret GpStatus) {
	r0, _, _ := syscall.Syscall6(procGdipAddPathEllipseI.Addr(), 5, uintptr(unsafe.Pointer(path)), uintptr(x), uintptr(y), uintptr(width), uintptr(height), 0)
	ret = GpStatus(r0)
	return
}

func GdipCreateBitmapFromFile(filename *uint16, bitmap **GpBitmap) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipCreateBitmapFromFile.Addr(), 2, uintptr(unsafe.Pointer(filename)), uintptr(unsafe.Pointer(bitmap)), 0)
	ret = GpStatus(r0)
	return
}

func GdipCreateBitmapFromGraphics(width int32, height int32, graphics *GpGraphics, bitmap **GpBitmap) (ret GpStatus) {
	r0, _, _ := syscall.Syscall6(procGdipCreateBitmapFromGraphics.Addr(), 4, uintptr(width), uintptr(height), uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(bitmap)), 0, 0)
	ret = GpStatus(r0)
	return
}

func GdipCreateBitmapFromHBITMAP(hbm HBITMAP, hpal HPALETTE, bitmap **GpBitmap) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipCreateBitmapFromHBITMAP.Addr(), 3, uintptr(hbm), uintptr(hpal), uintptr(unsafe.Pointer(bitmap)))
	ret = GpStatus(r0)
	return
}

func GdipCreateBitmapFromScan0(width int32, height int32, stride int32, format PixelFormat, scan0 *byte, bitmap **GpBitmap) (ret GpStatus) {
	r0, _, _ := syscall.Syscall6(procGdipCreateBitmapFromScan0.Addr(), 6, uintptr(width), uintptr(height), uintptr(stride), uintptr(format), uintptr(unsafe.Pointer(scan0)), uintptr(unsafe.Pointer(bitmap)))
	ret = GpStatus(r0)
	return
}

func GdipCreateBitmapFromStream(stream *com.IStreamABI, bitmap **GpBitmap) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipCreateBitmapFromStream.Addr(), 2, uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(bitmap)), 0)
	ret = GpStatus(r0)
	return
}

func GdipCreateFontFamilyFromName(name *uint16, collection *GpFontCollection, family **GpFontFamily) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipCreateFontFamilyFromName.Addr(), 3, uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(collection)), uintptr(unsafe.Pointer(family)))
	ret = GpStatus(r0)
	return
}

func GdipCreateFromHDC(hdc HDC, graphics **GpGraphics) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipCreateFromHDC.Addr(), 2, uintptr(hdc), uintptr(unsafe.Pointer(graphics)), 0)
	ret = GpStatus(r0)
	return
}

func GdipCreateHBITMAPFromBitmap(bitmap *GpBitmap, hbmReturn *HBITMAP, background ARGB) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipCreateHBITMAPFromBitmap.Addr(), 3, uintptr(unsafe.Pointer(bitmap)), uintptr(unsafe.Pointer(hbmReturn)), uintptr(background))
	ret = GpStatus(r0)
	return
}

func GdipCreatePath(fillMode FillMode, path **GpPath) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipCreatePath.Addr(), 2, uintptr(fillMode), uintptr(unsafe.Pointer(path)), 0)
	ret = GpStatus(r0)
	return
}

func GdipCreateSolidFill(color ARGB, brush **GpSolidFill) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipCreateSolidFill.Addr(), 2, uintptr(color), uintptr(unsafe.Pointer(brush)), 0)
	ret = GpStatus(r0)
	return
}

func GdipCreateStringFormat(flags StringFormatFlags, language LANGID, format **GpStringFormat) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipCreateStringFormat.Addr(), 3, uintptr(flags), uintptr(language), uintptr(unsafe.Pointer(format)))
	ret = GpStatus(r0)
	return
}

func GdipDeleteBrush(brush *GpBrush) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipDeleteBrush.Addr(), 1, uintptr(unsafe.Pointer(brush)), 0, 0)
	ret = GpStatus(r0)
	return
}

func GdipDeleteFont(font *GpFont) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipDeleteFont.Addr(), 1, uintptr(unsafe.Pointer(font)), 0, 0)
	ret = GpStatus(r0)
	return
}

func GdipDeleteFontFamily(family *GpFontFamily) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipDeleteFontFamily.Addr(), 1, uintptr(unsafe.Pointer(family)), 0, 0)
	ret = GpStatus(r0)
	return
}

func GdipDeleteGraphics(graphics *GpGraphics) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipDeleteGraphics.Addr(), 1, uintptr(unsafe.Pointer(graphics)), 0, 0)
	ret = GpStatus(r0)
	return
}

func GdipDeletePath(path *GpPath) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipDeletePath.Addr(), 1, uintptr(unsafe.Pointer(path)), 0, 0)
	ret = GpStatus(r0)
	return
}

func GdipDeleteStringFormat(format *GpStringFormat) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipDeleteStringFormat.Addr(), 1, uintptr(unsafe.Pointer(format)), 0, 0)
	ret = GpStatus(r0)
	return
}

func GdipDisposeImage(image *GpImage) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipDisposeImage.Addr(), 1, uintptr(unsafe.Pointer(image)), 0, 0)
	ret = GpStatus(r0)
	return
}

func GdipDrawImageRectI(graphics *GpGraphics, image *GpImage, x int32, y int32, width int32, height int32) (ret GpStatus) {
	r0, _, _ := syscall.Syscall6(procGdipDrawImageRectI.Addr(), 6, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(image)), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
	ret = GpStatus(r0)
	return
}

func GdipDrawImageRectRectI(graphics *GpGraphics, image *GpImage, dstX int32, dstY int32, dstWidth int32, dstHeight int32, srcX int32, srcY int32, srcWidth int32, srcHeight int32, srcUnit Unit, imgAttrs *GpImageAttributes, callback uintptr, callbackData uintptr) (ret GpStatus) {
	r0, _, _ := syscall.Syscall15(procGdipDrawImageRectRectI.Addr(), 14, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(image)), uintptr(dstX), uintptr(dstY), uintptr(dstWidth), uintptr(dstHeight), uintptr(srcX), uintptr(srcY), uintptr(srcWidth), uintptr(srcHeight), uintptr(srcUnit), uintptr(unsafe.Pointer(imgAttrs)), uintptr(callback), uintptr(callbackData), 0)
	ret = GpStatus(r0)
	return
}

func GdipDrawString(graphics *GpGraphics, text *uint16, textLength int32, font *GpFont, rectf *GpRectF, strFmt *GpStringFormat, brush *GpBrush) (ret GpStatus) {
	r0, _, _ := syscall.Syscall9(procGdipDrawString.Addr(), 7, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(text)), uintptr(textLength), uintptr(unsafe.Pointer(font)), uintptr(unsafe.Pointer(rectf)), uintptr(unsafe.Pointer(strFmt)), uintptr(unsafe.Pointer(brush)), 0, 0)
	ret = GpStatus(r0)
	return
}

func GdipFillEllipseI(graphics *GpGraphics, brush *GpBrush, x int32, y int32, width int32, height int32) (ret GpStatus) {
	r0, _, _ := syscall.Syscall6(procGdipFillEllipseI.Addr(), 6, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(brush)), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
	ret = GpStatus(r0)
	return
}

func GdipGetCompositingMode(graphics *GpGraphics, compositingMode *CompositingMode) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipGetCompositingMode.Addr(), 2, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(compositingMode)), 0)
	ret = GpStatus(r0)
	return
}

func GdipGetGenericFontFamilyMonospace(family **GpFontFamily) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipGetGenericFontFamilyMonospace.Addr(), 1, uintptr(unsafe.Pointer(family)), 0, 0)
	ret = GpStatus(r0)
	return
}

func GdipGetGenericFontFamilySansSerif(family **GpFontFamily) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipGetGenericFontFamilySansSerif.Addr(), 1, uintptr(unsafe.Pointer(family)), 0, 0)
	ret = GpStatus(r0)
	return
}

func GdipGetGenericFontFamilySerif(family **GpFontFamily) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipGetGenericFontFamilySerif.Addr(), 1, uintptr(unsafe.Pointer(family)), 0, 0)
	ret = GpStatus(r0)
	return
}

func GdipGetImageDimension(image *GpImage, width *float32, height *float32) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipGetImageDimension.Addr(), 3, uintptr(unsafe.Pointer(image)), uintptr(unsafe.Pointer(width)), uintptr(unsafe.Pointer(height)))
	ret = GpStatus(r0)
	return
}

func GdipGetImageGraphicsContext(image *GpImage, graphics **GpGraphics) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipGetImageGraphicsContext.Addr(), 2, uintptr(unsafe.Pointer(image)), uintptr(unsafe.Pointer(graphics)), 0)
	ret = GpStatus(r0)
	return
}

func GdipGetImageHeight(image *GpImage, height *uint32) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipGetImageHeight.Addr(), 2, uintptr(unsafe.Pointer(image)), uintptr(unsafe.Pointer(height)), 0)
	ret = GpStatus(r0)
	return
}

func GdipGetImageHorizontalResolution(image *GpImage, resolution *float32) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipGetImageHorizontalResolution.Addr(), 2, uintptr(unsafe.Pointer(image)), uintptr(unsafe.Pointer(resolution)), 0)
	ret = GpStatus(r0)
	return
}

func GdipGetImageVerticalResolution(image *GpImage, resolution *float32) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipGetImageVerticalResolution.Addr(), 2, uintptr(unsafe.Pointer(image)), uintptr(unsafe.Pointer(resolution)), 0)
	ret = GpStatus(r0)
	return
}

func GdipGetImageWidth(image *GpImage, width *uint32) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipGetImageWidth.Addr(), 2, uintptr(unsafe.Pointer(image)), uintptr(unsafe.Pointer(width)), 0)
	ret = GpStatus(r0)
	return
}

func GdipGraphicsClear(graphics *GpGraphics, color ARGB) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipGraphicsClear.Addr(), 2, uintptr(unsafe.Pointer(graphics)), uintptr(color), 0)
	ret = GpStatus(r0)
	return
}

func GdipResetClip(graphics *GpGraphics) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipResetClip.Addr(), 1, uintptr(unsafe.Pointer(graphics)), 0, 0)
	ret = GpStatus(r0)
	return
}

func GdipSetClipPath(graphics *GpGraphics, path *GpPath, combineMode CombineMode) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipSetClipPath.Addr(), 3, uintptr(unsafe.Pointer(graphics)), uintptr(unsafe.Pointer(path)), uintptr(combineMode))
	ret = GpStatus(r0)
	return
}

func GdipSetCompositingMode(graphics *GpGraphics, compositingMode CompositingMode) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipSetCompositingMode.Addr(), 2, uintptr(unsafe.Pointer(graphics)), uintptr(compositingMode), 0)
	ret = GpStatus(r0)
	return
}

func GdipSetCompositingQuality(graphics *GpGraphics, compositingQuality CompositingQuality) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipSetCompositingQuality.Addr(), 2, uintptr(unsafe.Pointer(graphics)), uintptr(compositingQuality), 0)
	ret = GpStatus(r0)
	return
}

func GdipSetInterpolationMode(graphics *GpGraphics, interpolationMode InterpolationMode) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipSetInterpolationMode.Addr(), 2, uintptr(unsafe.Pointer(graphics)), uintptr(interpolationMode), 0)
	ret = GpStatus(r0)
	return
}

func GdipSetPixelOffsetMode(graphics *GpGraphics, pixelOffsetMode PixelOffsetMode) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipSetPixelOffsetMode.Addr(), 2, uintptr(unsafe.Pointer(graphics)), uintptr(pixelOffsetMode), 0)
	ret = GpStatus(r0)
	return
}

func GdipSetSmoothingMode(graphics *GpGraphics, smoothingMode SmoothingMode) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipSetSmoothingMode.Addr(), 2, uintptr(unsafe.Pointer(graphics)), uintptr(smoothingMode), 0)
	ret = GpStatus(r0)
	return
}

func GdipSetStringFormatAlign(format *GpStringFormat, align StringAlignment) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipSetStringFormatAlign.Addr(), 2, uintptr(unsafe.Pointer(format)), uintptr(align), 0)
	ret = GpStatus(r0)
	return
}

func GdipSetStringFormatLineAlign(format *GpStringFormat, align StringAlignment) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipSetStringFormatLineAlign.Addr(), 2, uintptr(unsafe.Pointer(format)), uintptr(align), 0)
	ret = GpStatus(r0)
	return
}

func GdipSetTextRenderingHint(graphics *GpGraphics, mode TextRenderingHint) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdipSetTextRenderingHint.Addr(), 2, uintptr(unsafe.Pointer(graphics)), uintptr(mode), 0)
	ret = GpStatus(r0)
	return
}

func GdiplusStartup(token *uintptr, input *GdiplusStartupInput, output *GdiplusStartupOutput) (ret GpStatus) {
	r0, _, _ := syscall.Syscall(procGdiplusStartup.Addr(), 3, uintptr(unsafe.Pointer(token)), uintptr(unsafe.Pointer(input)), uintptr(unsafe.Pointer(output)))
	ret = GpStatus(r0)
	return
}
