package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"unsafe"
)

func changeWallpaper() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	wallpaperPath := fmt.Sprintf("%s\\wallpaper.png", dir)

	user32 := syscall.NewLazyDLL("user32.dll")
	setWallpaper := user32.NewProc("SystemParametersInfoW")

	pathPtr, err := syscall.UTF16PtrFromString(wallpaperPath)
	if err != nil {
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

func openBrowser() error {
	url := "https://learn.logikaschool.com/login"

	var cmd *exec.Cmd

	// Check the operating system
	switch runtime.GOOS {
	case "windows":
		// On Windows, use the "start" command
		cmd = exec.Command("cmd", "/C", "start", url)
	case "darwin":
		// On macOS, use the "open" command
		cmd = exec.Command("open", url)
	default:
		// On Linux and others, use the "xdg-open" command
		cmd = exec.Command("xdg-open", url)
	}

	// Execute the command
	return cmd.Start()
}
