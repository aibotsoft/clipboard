package clipboard

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

const (
	cfUnicodetext = 13
)

var (
	user32         = syscall.MustLoadDLL("user32")
	openClipboard  = user32.MustFindProc("OpenClipboard")
	closeClipboard = user32.MustFindProc("CloseClipboard")
	//emptyClipboard   = user32.MustFindProc("EmptyClipboard")
	getClipboardData = user32.MustFindProc("GetClipboardData")
	//setClipboardData = user32.MustFindProc("SetClipboardData")

	kernel32 = syscall.NewLazyDLL("kernel32")
	//globalAlloc  = kernel32.NewProc("GlobalAlloc")
	//globalFree   = kernel32.NewProc("GlobalFree")
	globalLock   = kernel32.NewProc("GlobalLock")
	globalUnlock = kernel32.NewProc("GlobalUnlock")
	//lstrcpy      = kernel32.NewProc("lstrcpyW")
)

func readAll() (string, error) {
	r, _, err := openClipboard.Call(0)
	if r == 0 {
		return "", fmt.Errorf("openClipboard: %w", err)
	}
	defer closeClipboard.Call()

	h, _, err := getClipboardData.Call(cfUnicodetext)
	if h == 0 {
		return "", fmt.Errorf("getClipboardData: %w", err)
	}

	l, _, err := globalLock.Call(h)
	if l == 0 {
		return "", fmt.Errorf("globalLock: %w", err)
	}
	defer globalUnlock.Call(h)

	text := syscall.UTF16ToString((*[1 << 20]uint16)(unsafe.Pointer(l))[:])

	return text, nil
}

func Get() (text string, err error) {
	for i := 0; i < 10; i++ {
		text, err = readAll()
		if err == nil {
			return
		}
		//fmt.Println(err)
		time.Sleep(time.Millisecond)
	}
	return
}
