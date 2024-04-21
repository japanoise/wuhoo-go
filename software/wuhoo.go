package wuhoo

// #cgo windows LDFLAGS: -lkernel32 -luser32 -lgdi32 -lshell32
// #cgo linux LDFLAGS: -lX11
// #cgo darwin LDFLAGS: -framework Cocoa -framework Quartz -framework Carbon
// #include <stdlib.h>
// #define WUHOO_IMPLEMENTATION
// #include "wuhoo.h"
import "C"
import (
	"errors"
	"unsafe"
)

type Flags C.WuhooFlags

type KeyMods C.WuhooKeyModifiers

type KeyCode C.WuhooKeyCode

type WindowFlags C.WuhooWindowFlags

type WindowState C.WuhooWindowState

type EventType C.WuhooEventType

type KeyState C.WuhooKeyState

type MouseState C.WuhooMouseState

type MouseMods C.WuhooMouseModifiers

type Position int32

type Size uint32

func returnValue(r C.WuhooResult) error {
	if r == C.WuhooResult(C.WuhooSuccess) {
		return nil
	} else {
		return errors.New(C.GoString(C.WuhooResultString(r)))
	}
}

type Window struct {
	handle *C.WuhooWindow
}

func (w *Window) Init() error {
	w.handle = (*C.WuhooWindow)(C.malloc(C.sizeof_WuhooWindow))
	return returnValue(C.WuhooWindowInit(w.handle))
}

func (w *Window) Destroy() error {
	ret := returnValue(C.WuhooWindowDestroy(w.handle))
	C.free(unsafe.Pointer(w.handle))

	return ret
}

func (w *Window) Show() error {
	return returnValue(C.WuhooWindowShow(w.handle))
}

func (w *Window) SetTitle(title string) error {
	cstr := C.CString(title)
	ret := returnValue(C.WuhooWindowSetTitle(w.handle, cstr))
	C.free(unsafe.Pointer(cstr))

	return ret
}

func (w *Window) Create(posX, posY Position, width, height Size, title string, flags Flags) error {
	cstr := C.CString(title)
	ret := returnValue(C.WuhooWindowCreate(w.handle, C.int(posX), C.int(posY),
		C.uint(width), C.uint(height), cstr, C.uint(flags), nil))
	C.free(unsafe.Pointer(cstr))

	return ret
}

func (w *Window) Blit(b *Bitmap, srcX, srcY, dstX, dstY, dstWidth, dstHeight Size) error {
	return returnValue(C.WuhooWindowBlit(
		w.handle, b.data, C.uint(srcX), C.uint(srcY), b.width, b.height,
		C.uint(dstX), C.uint(dstY), C.uint(dstWidth), C.uint(dstHeight)))
}

func (w *Window) EventNext(event *Event) error {
	return returnValue(C.WuhooWindowEventNext(w.handle, event.handle))
}

func (w *Window) DropContentsGet(event *Event, buffer []byte) error {
	cbytes := C.CBytes(buffer)
	ret := returnValue(C.WuhooWindowDropContentsGet(
		w.handle, event.handle, (*C.char)(cbytes), C.uint(len(buffer))))
	C.free(cbytes)
	return ret
}

func (w *Window) RegionSet(posX, posY Position, width, height Size) error {
	return returnValue(C.WuhooWindowRegionSet(
		w.handle, C.int(posX), C.int(posY),
		C.uint(width), C.uint(height)))
}

func (w *Window) RegionGet(posX, posY *Position, width, height *Size) error {
	return returnValue(C.WuhooWindowRegionGet(
		w.handle, (*C.int)(posX), (*C.int)(posY),
		(*C.uint)(width), (*C.uint)(height)))
}

func (w *Window) ClientRegionSet(posX, posY Position, width, height Size) error {
	return returnValue(C.WuhooWindowClientRegionSet(
		w.handle, C.int(posX), C.int(posY),
		C.uint(width), C.uint(height)))
}

func (w *Window) ClientRegionGet(posX, posY *Position, width, height *Size) error {
	return returnValue(C.WuhooWindowClientRegionGet(
		w.handle, (*C.int)(posX), (*C.int)(posY),
		(*C.uint)(width), (*C.uint)(height)))
}
