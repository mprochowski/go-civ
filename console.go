package main

import (
	"syscall"
	"unsafe"
)

type Console struct {
	Width int
	Height int
}

func getConsole() Console {
	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		panic(errno)
	}

	return Console{int(ws.Col), int(ws.Row) }
}


type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

//
//func getWidth2() (uint, uint) {
//	cmd := exec.Command("stty", "size")
//	cmd.Stdin = os.Stdin
//	out, err := cmd.Output()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	slice := strings.Split(string(out), " ")
//
//	cols,_ := strconv.ParseUint(slice[0], 10, 32)
//	rows,_ :=  strconv.ParseUint(slice[1], 10, 32)
//
//	return uint(cols),uint(rows)
//}