//go:build linux || darwin

package util

import (
	"fmt"
	"os"
	"syscall"
)

func getFileKey(path string) (string, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	stat, ok := fi.Sys().(*syscall.Stat_t)
	if !ok {
		return "", fmt.Errorf("failed to get raw syscall.Stat_t")
	}

	return fmt.Sprintf("unix-%d-%d", stat.Dev, stat.Ino), nil
}
