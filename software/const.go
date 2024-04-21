package wuhoo

// #include "wuhoo.h"
import "C"

const PlatformAPIString string = C.WUHOO_PLATFORM_API_STRING

const (
	FlagTitled       Flags = C.WUHOO_FLAG_TITLED
	FlagCanvas       Flags = C.WUHOO_FLAG_CANVAS
	FlagResizeable   Flags = C.WUHOO_FLAG_RESIZEABLE
	FlagBorderless   Flags = C.WUHOO_FLAG_BORDERLESS
	FlagMouseCapture Flags = C.WUHOO_FLAG_MOUSE_CAPTURE
	FlagFileDrop     Flags = C.WUHOO_FLAG_FILE_DROP
	FlagCloseable    Flags = C.WUHOO_FLAG_CLOSEABLE
	FlagClientRegion Flags = C.WUHOO_FLAG_CLIENT_REGION
	FlagWindowRegion Flags = C.WUHOO_FLAG_WINDOW_REGION
)

const (
	KeyModNone     KeyMods = C.WUHOO_KMOD_NONE
	KeyModLShift   KeyMods = C.WUHOO_KMOD_LSHIFT
	KeyModRShift   KeyMods = C.WUHOO_KMOD_RSHIFT
	KeyModLCtrl    KeyMods = C.WUHOO_KMOD_LCTRL
	KeyModRCtrl    KeyMods = C.WUHOO_KMOD_RCTRL
	KeyModLAlt     KeyMods = C.WUHOO_KMOD_LALT
	KeyModRAlt     KeyMods = C.WUHOO_KMOD_RALT
	KeyModLGui     KeyMods = C.WUHOO_KMOD_LGUI
	KeyModRGui     KeyMods = C.WUHOO_KMOD_RGUI
	KeyModNum      KeyMods = C.WUHOO_KMOD_NUM
	KeyModCaps     KeyMods = C.WUHOO_KMOD_CAPS
	KeyModMode     KeyMods = C.WUHOO_KMOD_MODE
	KeyModDeadchar KeyMods = C.WUHOO_KMOD_DEADCHAR
	KeyModCtrl     KeyMods = C.WUHOO_KMOD_CTRL
	KeyModShift    KeyMods = C.WUHOO_KMOD_SHIFT
	KeyModAlt      KeyMods = C.WUHOO_KMOD_ALT
	KeyModGui      KeyMods = C.WUHOO_KMOD_GUI
	KeyModMax      KeyMods = C.WUHOO_KMOD_MAX
)

const (
	KeyUnknown       KeyCode = C.WUHOO_VKEY_UNKNOWN
	KeyBackspace     KeyCode = C.WUHOO_VKEY_BACKSPACE
	KeyForwardDelete KeyCode = C.WUHOO_VKEY_FORWARD_DELETE
	KeyTab           KeyCode = C.WUHOO_VKEY_TAB
	KeyEnter         KeyCode = C.WUHOO_VKEY_ENTER
	KeyKpadEnter     KeyCode = C.WUHOO_VKEY_KPAD_ENTER
	KeyEscape        KeyCode = C.WUHOO_VKEY_ESCAPE
	KeySpace         KeyCode = C.WUHOO_VKEY_SPACE
	KeyQuote         KeyCode = C.WUHOO_VKEY_QUOTE
	KeyComma         KeyCode = C.WUHOO_VKEY_COMMA
	KeyMinus         KeyCode = C.WUHOO_VKEY_MINUS
	KeyPeriod        KeyCode = C.WUHOO_VKEY_PERIOD
	KeyForwardSlash  KeyCode = C.WUHOO_VKEY_FORWARD_SLASH
	Key0             KeyCode = C.WUHOO_VKEY_0
	Key1             KeyCode = C.WUHOO_VKEY_1
	Key2             KeyCode = C.WUHOO_VKEY_2
	Key3             KeyCode = C.WUHOO_VKEY_3
	Key4             KeyCode = C.WUHOO_VKEY_4
	Key5             KeyCode = C.WUHOO_VKEY_5
	Key6             KeyCode = C.WUHOO_VKEY_6
	Key7             KeyCode = C.WUHOO_VKEY_7
	Key8             KeyCode = C.WUHOO_VKEY_8
	Key9             KeyCode = C.WUHOO_VKEY_9
	KeySemicolon     KeyCode = C.WUHOO_VKEY_SEMICOLON
	KeyEquals        KeyCode = C.WUHOO_VKEY_EQUALS
	KeyA             KeyCode = C.WUHOO_VKEY_A
	KeyB             KeyCode = C.WUHOO_VKEY_B
	KeyC             KeyCode = C.WUHOO_VKEY_C
	KeyD             KeyCode = C.WUHOO_VKEY_D
	KeyE             KeyCode = C.WUHOO_VKEY_E
	KeyF             KeyCode = C.WUHOO_VKEY_F
	KeyG             KeyCode = C.WUHOO_VKEY_G
	KeyH             KeyCode = C.WUHOO_VKEY_H
	KeyI             KeyCode = C.WUHOO_VKEY_I
	KeyJ             KeyCode = C.WUHOO_VKEY_J
	KeyK             KeyCode = C.WUHOO_VKEY_K
	KeyL             KeyCode = C.WUHOO_VKEY_L
	KeyM             KeyCode = C.WUHOO_VKEY_M
	KeyN             KeyCode = C.WUHOO_VKEY_N
	KeyO             KeyCode = C.WUHOO_VKEY_O
	KeyP             KeyCode = C.WUHOO_VKEY_P
	KeyQ             KeyCode = C.WUHOO_VKEY_Q
	KeyR             KeyCode = C.WUHOO_VKEY_R
	KeyS             KeyCode = C.WUHOO_VKEY_S
	KeyT             KeyCode = C.WUHOO_VKEY_T
	KeyU             KeyCode = C.WUHOO_VKEY_U
	KeyV             KeyCode = C.WUHOO_VKEY_V
	KeyW             KeyCode = C.WUHOO_VKEY_W
	KeyX             KeyCode = C.WUHOO_VKEY_X
	KeyY             KeyCode = C.WUHOO_VKEY_Y
	KeyZ             KeyCode = C.WUHOO_VKEY_Z
	KeyLeftBracket   KeyCode = C.WUHOO_VKEY_LEFT_BRACKET
	KeyBack_slash    KeyCode = C.WUHOO_VKEY_BACK_SLASH
	KeyRightBracket  KeyCode = C.WUHOO_VKEY_RIGHT_BRACKET
	KeyGrave         KeyCode = C.WUHOO_VKEY_GRAVE
	KeyTilda         KeyCode = C.WUHOO_VKEY_TILDA
	KeyDelete        KeyCode = C.WUHOO_VKEY_DELETE
	KeyF1            KeyCode = C.WUHOO_VKEY_F1
	KeyF2            KeyCode = C.WUHOO_VKEY_F2
	KeyF3            KeyCode = C.WUHOO_VKEY_F3
	KeyF4            KeyCode = C.WUHOO_VKEY_F4
	KeyF5            KeyCode = C.WUHOO_VKEY_F5
	KeyF6            KeyCode = C.WUHOO_VKEY_F6
	KeyF7            KeyCode = C.WUHOO_VKEY_F7
	KeyF8            KeyCode = C.WUHOO_VKEY_F8
	KeyF9            KeyCode = C.WUHOO_VKEY_F9
	KeyF10           KeyCode = C.WUHOO_VKEY_F10
	KeyF11           KeyCode = C.WUHOO_VKEY_F11
	KeyF12           KeyCode = C.WUHOO_VKEY_F12
	KeyF13           KeyCode = C.WUHOO_VKEY_F13
	KeyPrintscreen   KeyCode = C.WUHOO_VKEY_PRINTSCREEN
	KeyF14           KeyCode = C.WUHOO_VKEY_F14
	KeyF15           KeyCode = C.WUHOO_VKEY_F15
	KeyInsert        KeyCode = C.WUHOO_VKEY_INSERT
	KeyHelp          KeyCode = C.WUHOO_VKEY_HELP
	KeyHome          KeyCode = C.WUHOO_VKEY_HOME
	KeyPageUp        KeyCode = C.WUHOO_VKEY_PAGE_UP
	KeyEnd           KeyCode = C.WUHOO_VKEY_END
	KeyPageDown      KeyCode = C.WUHOO_VKEY_PAGE_DOWN
	KeyMenu          KeyCode = C.WUHOO_VKEY_MENU
	KeyKpad0         KeyCode = C.WUHOO_VKEY_KPAD_0
	KeyKpad1         KeyCode = C.WUHOO_VKEY_KPAD_1
	KeyKpad2         KeyCode = C.WUHOO_VKEY_KPAD_2
	KeyKpad3         KeyCode = C.WUHOO_VKEY_KPAD_3
	KeyKpad4         KeyCode = C.WUHOO_VKEY_KPAD_4
	KeyKpad5         KeyCode = C.WUHOO_VKEY_KPAD_5
	KeyKpad6         KeyCode = C.WUHOO_VKEY_KPAD_6
	KeyKpad7         KeyCode = C.WUHOO_VKEY_KPAD_7
	KeyKpad8         KeyCode = C.WUHOO_VKEY_KPAD_8
	KeyKpad9         KeyCode = C.WUHOO_VKEY_KPAD_9
	KeyKpadDivide    KeyCode = C.WUHOO_VKEY_KPAD_DIVIDE
	KeyKpadSlash     KeyCode = C.WUHOO_VKEY_KPAD_SLASH
	KeyKpadPlus      KeyCode = C.WUHOO_VKEY_KPAD_PLUS
	KeyKpadMinus     KeyCode = C.WUHOO_VKEY_KPAD_MINUS
	KeyKpadEquals    KeyCode = C.WUHOO_VKEY_KPAD_EQUALS
	KeyKpadMulitply  KeyCode = C.WUHOO_VKEY_KPAD_MULITPLY
	KeyKpadDecimal   KeyCode = C.WUHOO_VKEY_KPAD_DECIMAL
	KeyKpadNumLock   KeyCode = C.WUHOO_VKEY_KPAD_NUM_LOCK
	KeyKpadComma     KeyCode = C.WUHOO_VKEY_KPAD_COMMA
	KeyShift         KeyCode = C.WUHOO_VKEY_SHIFT
	KeyControl       KeyCode = C.WUHOO_VKEY_CONTROL
	KeyAlt           KeyCode = C.WUHOO_VKEY_ALT
	KeyCapsLock      KeyCode = C.WUHOO_VKEY_CAPS_LOCK
	KeyUp            KeyCode = C.WUHOO_VKEY_UP
	KeyDown          KeyCode = C.WUHOO_VKEY_DOWN
	KeyRight         KeyCode = C.WUHOO_VKEY_RIGHT
	KeyLeft          KeyCode = C.WUHOO_VKEY_LEFT
	KeyMax           KeyCode = C.WUHOO_VKEY_MAX
)

const (
	WindowFlagResized       WindowFlags = C.WUHOO_WINDOW_FLAG_RESIZED
	WindowFlagFullScreen    WindowFlags = C.WUHOO_WINDOW_FLAG_FULL_SCREEN
	WindowFlagFocusGained   WindowFlags = C.WUHOO_WINDOW_FLAG_FOCUS_GAINED
	WindowFlagFocusLost     WindowFlags = C.WUHOO_WINDOW_FLAG_FOCUS_LOST
	WindowFlagMinimized     WindowFlags = C.WUHOO_WINDOW_FLAG_MINIMIZED
	WindowFlagMaximized     WindowFlags = C.WUHOO_WINDOW_FLAG_MAXIMIZED
	WindowFlagMoved         WindowFlags = C.WUHOO_WINDOW_FLAG_MOVED
	WindowFlagClosed        WindowFlags = C.WUHOO_WINDOW_FLAG_CLOSED
	WindowFlagRegionUpdated WindowFlags = C.WUHOO_WINDOW_FLAG_REGION_UPDATED
	WindowFlagDropStarted   WindowFlags = C.WUHOO_WINDOW_FLAG_DROP_STARTED
)

const (
	WindowStateUnknown      WindowState = C.WUHOO_WSTATE_UNKNOWN
	WindowStateResized      WindowState = C.WUHOO_WSTATE_RESIZED
	WindowStateFullScreened WindowState = C.WUHOO_WSTATE_FULL_SCREENED
	WindowStateMoved        WindowState = C.WUHOO_WSTATE_MOVED
	WindowStateMinimized    WindowState = C.WUHOO_WSTATE_MINIMIZED
	WindowStateMaximized    WindowState = C.WUHOO_WSTATE_MAXIMIZED
	WindowStateCreated      WindowState = C.WUHOO_WSTATE_CREATED
	WindowStateClosed       WindowState = C.WUHOO_WSTATE_CLOSED
	WindowStateUnfocused    WindowState = C.WUHOO_WSTATE_UNFOCUSED
	WindowStateFocused      WindowState = C.WUHOO_WSTATE_FOCUSED
	WindowStateHidden       WindowState = C.WUHOO_WSTATE_HIDDEN
	WindowStateInvalidated  WindowState = C.WUHOO_WSTATE_INVALIDATED
	WindowStateMax          WindowState = C.WUHOO_WSTATE_MAX
)

const (
	EventTypeWindow     EventType = C.WUHOO_EVT_WINDOW
	EventTypeKey        EventType = C.WUHOO_EVT_KEY
	EventTypeMousePress EventType = C.WUHOO_EVT_MOUSE_PRESS
	EventTypeMouseMove  EventType = C.WUHOO_EVT_MOUSE_MOVE
	EventTypeMouseWheel EventType = C.WUHOO_EVT_MOUSE_WHEEL
	EventTypeDrop       EventType = C.WUHOO_EVT_DROP
	EventTypeMax        EventType = C.WUHOO_EVT_MAX
)

const (
	KeyStateUnknown KeyState = C.WUHOO_KSTATE_UNKNOWN
	KeyStateUp      KeyState = C.WUHOO_KSTATE_UP
	KeyStateDown    KeyState = C.WUHOO_KSTATE_DOWN
	KeyStateMax     KeyState = C.WUHOO_KSTATE_MAX
)

const (
	MouseStateUnknown   MouseState = C.WUHOO_MSTATE_UNKNOWN
	MouseStateLPressed  MouseState = C.WUHOO_MSTATE_LPRESSED
	MouseStateLReleased MouseState = C.WUHOO_MSTATE_LRELEASED
	MouseStateRPressed  MouseState = C.WUHOO_MSTATE_RPRESSED
	MouseStateRReleased MouseState = C.WUHOO_MSTATE_RRELEASED
	MouseStateMPressed  MouseState = C.WUHOO_MSTATE_MPRESSED
	MouseStateMReleased MouseState = C.WUHOO_MSTATE_MRELEASED
	MouseStateMax       MouseState = C.WUHOO_MSTATE_MAX
)

// Note that this is the same as key mod, just with different names
// and typing. Copy the block from above and replace Key with Mouse
// to regenerate this.
const (
	MouseModNone     MouseMods = C.WUHOO_KMOD_NONE
	MouseModLShift   MouseMods = C.WUHOO_KMOD_LSHIFT
	MouseModRShift   MouseMods = C.WUHOO_KMOD_RSHIFT
	MouseModLCtrl    MouseMods = C.WUHOO_KMOD_LCTRL
	MouseModRCtrl    MouseMods = C.WUHOO_KMOD_RCTRL
	MouseModLAlt     MouseMods = C.WUHOO_KMOD_LALT
	MouseModRAlt     MouseMods = C.WUHOO_KMOD_RALT
	MouseModLGui     MouseMods = C.WUHOO_KMOD_LGUI
	MouseModRGui     MouseMods = C.WUHOO_KMOD_RGUI
	MouseModNum      MouseMods = C.WUHOO_KMOD_NUM
	MouseModCaps     MouseMods = C.WUHOO_KMOD_CAPS
	MouseModMode     MouseMods = C.WUHOO_KMOD_MODE
	MouseModDeadchar MouseMods = C.WUHOO_KMOD_DEADCHAR
	MouseModCtrl     MouseMods = C.WUHOO_KMOD_CTRL
	MouseModShift    MouseMods = C.WUHOO_KMOD_SHIFT
	MouseModAlt      MouseMods = C.WUHOO_KMOD_ALT
	MouseModGui      MouseMods = C.WUHOO_KMOD_GUI
	MouseModMax      MouseMods = C.WUHOO_KMOD_MAX
)
