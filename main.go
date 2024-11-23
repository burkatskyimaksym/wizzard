package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

func MyCustomError(message string) error {
	return fmt.Errorf("Custom Error: %s", message)
}

func changeWallpaper() error {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return err
	}

	wallpaperPath := fmt.Sprintf("%s\\wallpaper.jpg", dir)

	user32 := syscall.NewLazyDLL("user32.dll")
	setWallpaper := user32.NewProc("SystemParametersInfoW")

	pathPtr, err := syscall.UTF16PtrFromString(wallpaperPath)
	if err != nil {
		fmt.Println("Error converting string:", err)
		return err
	}

	const SPI_SETDESKWALLPAPER = 0x0014
	const SPIF_UPDATEINIFILE = 0x01
	const SPIF_SENDCHANGE = 0x02

	ret, _, _ := setWallpaper.Call(
		uintptr(SPI_SETDESKWALLPAPER),
		0,
		uintptr(unsafe.Pointer(pathPtr)),
		uintptr(SPIF_UPDATEINIFILE|SPIF_SENDCHANGE),
	)

	if ret == 0 {
		return fmt.Errorf("Failed to set wallpaper")
	} else {
		return nil
	}
}

func main() {
	err := changeWallpaper()
	if err != nil {
		fmt.Printf("Error setting wallpaper %s\n", err)
	}
}
