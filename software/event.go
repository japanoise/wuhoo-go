package wuhoo

// #include <stdlib.h>
// #include "wuhoo.h"
//
// WuhooEventType getEventType(WuhooEvent *ev) {
//   return ev->type;
// }
//
// WuhooEventKey getEventDataKey(WuhooEvent *ev) {
//   return ev->data.key;
// }
//
// WuhooEventMousePress getEventDataMousePress(WuhooEvent *ev) {
//   return ev->data.mouse_press;
// }
//
// WuhooEventMouseMove getEventDataMouseMove(WuhooEvent *ev) {
//   return ev->data.mouse_move;
// }
//
// WuhooEventMouseWheel getEventDataMouseWheel(WuhooEvent *ev) {
//   return ev->data.mouse_wheel;
// }
//
// WuhooEventWindow getEventDataWindow(WuhooEvent *ev) {
//   return ev->data.window;
// }
//
// WuhooEventDrop getEventDataDrop(WuhooEvent *ev) {
//   return ev->data.drop;
// }
import "C"
import (
	"unicode/utf8"
	"unsafe"
)

type Event struct {
	handle *C.WuhooEvent
}

func (e *Event) Type() EventType {
	return EventType(C.getEventType(e.handle))
}

func (e *Event) Init() {
	e.handle = (*C.WuhooEvent)(C.malloc(C.sizeof_WuhooEvent))
}

func (e *Event) Destroy() {
	C.free(unsafe.Pointer(e.handle))
}

type EventKey C.WuhooEventKey

func (e *Event) Key() EventKey {
	return EventKey(C.getEventDataKey(e.handle))
}

func (e EventKey) Mods() KeyMods {
	return KeyMods(e.mods)
}

func (e EventKey) State() KeyState {
	return KeyState(e.state)
}

func (e EventKey) Code() KeyCode {
	return KeyCode(e.code)
}

func (e EventKey) Character() (rune, int) {
	data := make([]byte, C.WUHOO_MAX_CHARACTER_SIZE)

	for i := range data {
		data[i] = byte(e.character[i])
	}

	return utf8.DecodeRune(data)
}

type EventMousePress C.WuhooEventMousePress

func (e *Event) MousePress() EventMousePress {
	return EventMousePress(C.getEventDataMousePress(e.handle))
}

func (e EventMousePress) Mods() MouseMods {
	return MouseMods(e.mods)
}

func (e EventMousePress) State() MouseState {
	return MouseState(e.state)
}

func (e EventMousePress) ClickCount() int32 {
	return int32(e.click_count)
}

func (e EventMousePress) X() Position {
	return Position(e.x)
}

func (e EventMousePress) Y() Position {
	return Position(e.y)
}

type EventMouseMove C.WuhooEventMouseMove

func (e *Event) MouseMove() EventMouseMove {
	return EventMouseMove(C.getEventDataMouseMove(e.handle))
}

func (e EventMouseMove) Mods() MouseMods {
	return MouseMods(e.mods)
}

func (e EventMouseMove) State() MouseState {
	return MouseState(e.state)
}

func (e EventMouseMove) X() Position {
	return Position(e.x)
}

func (e EventMouseMove) Y() Position {
	return Position(e.y)
}

type EventMouseWheel C.WuhooEventMouseWheel

func (e *Event) MouseWheel() EventMouseWheel {
	return EventMouseWheel(C.getEventDataMouseWheel(e.handle))
}

func (e EventMouseWheel) Mods() MouseMods {
	return MouseMods(e.mods)
}

func (e EventMouseWheel) X() Position {
	return Position(e.x)
}

func (e EventMouseWheel) Y() Position {
	return Position(e.y)
}

func (e EventMouseWheel) DeltaX() float32 {
	return float32(e.delta_x)
}

func (e EventMouseWheel) DeltaY() float32 {
	return float32(e.delta_y)
}

type EventWindow C.WuhooEventWindow

func (e *Event) Window() EventWindow {
	return EventWindow(C.getEventDataWindow(e.handle))
}

func (e EventWindow) State() WindowState {
	return WindowState(e.state)
}

func (e EventWindow) Flags() WindowFlags {
	return WindowFlags(e.flags)
}

func (e EventWindow) Data1() int32 {
	return int32(e.data1)
}

func (e EventWindow) Data2() int32 {
	return int32(e.data2)
}

type EventDrop C.WuhooEventDrop

func (e *Event) Drop() EventDrop {
	return EventDrop(C.getEventDataDrop(e.handle))
}

type Handle C.WuhooHandle

func (e EventDrop) Handle() Handle {
	return Handle(e.context)
}

func (e EventDrop) Count() Size {
	return Size(e.count)
}

func (e EventDrop) Size() Size {
	return Size(e.size)
}
