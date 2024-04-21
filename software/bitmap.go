package wuhoo

// #include <stdlib.h>
// #include "wuhoo.h"
import "C"
import "unsafe"

type RGBA C.WuhooRGBA

func (r RGBA) Red() byte {
	return byte(r.r)
}

func (r RGBA) Green() byte {
	return byte(r.g)
}

func (r RGBA) Blue() byte {
	return byte(r.b)
}

func (r RGBA) Alpha() byte {
	return byte(r.a)
}

func NewColorRGBA(r, g, b, a byte) RGBA {
	var ret RGBA

	ret.r = C.uchar(r)
	ret.g = C.uchar(g)
	ret.b = C.uchar(b)
	ret.a = C.uchar(a)

	return ret
}

func NewColorRGB(r, g, b byte) RGBA {
	return NewColorRGBA(r, g, b, 255)
}

type Bitmap struct {
	data   *C.WuhooRGBA
	width  C.uint
	height C.uint
}

func (b *Bitmap) Width() Size {
	return Size(b.width)
}

func (b *Bitmap) Height() Size {
	return Size(b.height)
}

func (b *Bitmap) Init(pixels []RGBA, width, height Size) {
	b.width = C.uint(width)
	b.height = C.uint(height)
	pointer := C.malloc(C.ulong(len(pixels) * C.sizeof_WuhooRGBA))
	b.data = (*C.WuhooRGBA)(pointer)
	for _, px := range pixels {
		*(*C.WuhooRGBA)(pointer) = C.WuhooRGBA(px)
		// XXX: pointer arithmetic in Golang sucks - it's
		// always by byte (because unsafe.Pointer is always
		// void* I guess).  So we can't just increment by 1,
		// instead we need to actually sizeof the pointed
		// type.
		pointer = unsafe.Add(pointer, C.sizeof_WuhooRGBA)
	}
}

func (b *Bitmap) InitBlank(width, height Size) {
	b.width = C.uint(width)
	b.height = C.uint(height)
	b.data = (*C.WuhooRGBA)(
		C.calloc(C.ulong(width*height), C.ulong(C.sizeof_WuhooRGBA)))
}

func (b *Bitmap) Delete() {
	C.free(unsafe.Pointer(b.data))
}

func (b *Bitmap) SetIndex(index int, px RGBA) {
	pointer := unsafe.Add(unsafe.Pointer(b.data), index*C.sizeof_WuhooRGBA)
	*(*C.WuhooRGBA)(pointer) = C.WuhooRGBA(px)
}

func (b *Bitmap) SetXY(x, y int, px RGBA) {
	b.SetIndex(x+y*int(b.width), px)
}

func (b *Bitmap) XYToIndex(x, y int) int {
	return x + y*int(b.width)
}

func (b *Bitmap) IndexToXY(index int) (int, int) {
	width := int(b.width)
	y := index / width
	x := index - y*width

	return x, y
}

func (b *Bitmap) ForEachPixel(f func(b *Bitmap, index, x, y int, px RGBA) RGBA) {
	width := int(b.Width())
	height := int(b.Height())
	pointer := unsafe.Pointer(b.data)

	var x, y int
	for x = 0; x < width; x++ {
		for y = 0; y < height; y++ {
			px := *(*C.WuhooRGBA)(pointer)

			*(*C.WuhooRGBA)(pointer) = C.WuhooRGBA(
				f(b, b.XYToIndex(x, y), x, y, RGBA(px)))

			pointer = unsafe.Add(pointer, C.sizeof_WuhooRGBA)
		}
	}
}
