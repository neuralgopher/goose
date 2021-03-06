package kbd

import (
	"asm"
	"idt"
	"regs"
	"unsafe"
	"video"
)

var kbdus [128]uint8 = [128]uint8{
	0, 27, '1', '2', '3', '4', '5', '6', '7', '8', /* 9 */
	'9', '0', '-', '=', '\b', /* Backspace */
	'\t',               /* Tab */
	'q', 'w', 'e', 'r', /* 19 */
	't', 'y', 'u', 'i', 'o', 'p', '[', ']', '\n', /* Enter key */
	0,                                                /* 29   - Control */
	'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l', ';', /* 39 */
	'\'', '`', 0, /* Left shift */
	'\\', 'z', 'x', 'c', 'v', 'b', 'n', /* 49 */
	'm', ',', '.', '/', 0, /* Right shift */
	'*',
	0,   /* Alt */
	' ', /* Space bar */
	0,   /* Caps lock */
	0,   /* 59 - F1 key ... > */
	0, 0, 0, 0, 0, 0, 0, 0,
	0, /* < ... F10 */
	0, /* 69 - Num lock*/
	0, /* Scroll Lock */
	0, /* Home key */
	0, /* Up Arrow */
	0, /* Page Up */
	'-',
	0, /* Left Arrow */
	0,
	0, /* Right Arrow */
	'+',
	0, /* 79 - End key*/
	0, /* Down Arrow */
	0, /* Page Down */
	0, /* Insert Key */
	0, /* Delete Key */
	0, 0, 0,
	0, /* F11 Key */
	0, /* F12 Key */
	0, /* All other keys are undefined */
}

func handler(r *regs.Regs) {
	scancode := asm.InportB(0x60)

	if scancode&0x80 == 0 {
		switch scancode {
		case 0x4B:
			video.MoveCursor(-1, 0)
		case 0x48:
			video.MoveCursor(0, -1)
		case 0x50:
			video.MoveCursor(0, 1)
		case 0x4D:
			video.MoveCursor(1, 0)
		default:
			video.PutChar(rune(kbdus[scancode]))
		}
	} else {
	}
}

func Init() {
	dummy := handler
	idt.AddIRQ(1, **(**uintptr)(unsafe.Pointer(&dummy)))
}
