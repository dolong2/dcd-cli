//go:build windows

package util

import (
	"fmt"
	"os"
	"syscall"
)

func getFileKey(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var d syscall.ByHandleFileInformation
	if err := syscall.GetFileInformationByHandle(syscall.Handle(f.Fd()), &d); err != nil {
		return "", err
	}

	return fmt.Sprintf("win-%d-%d-%d", d.VolumeSerialNumber, d.FileIndexHigh, d.FileIndexLow), nil
}