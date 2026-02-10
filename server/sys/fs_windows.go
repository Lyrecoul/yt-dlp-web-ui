//go:build windows

package sys

import (
	"github.com/marcopiovanello/yt-dlp-web-ui/v3/server/config"
	"golang.org/x/sys/windows"
)

// FreeSpace gets the available Bytes writable to download directory
func FreeSpace() (uint64, error) {
	path, err := windows.UTF16PtrFromString(config.Instance().DownloadPath)
	if err != nil {
		return 0, err
	}

	var freeBytesAvailable uint64

	err = windows.GetDiskFreeSpaceEx(
		path,
		&freeBytesAvailable,
		nil,
		nil,
	)
	if err != nil {
		return 0, err
	}

	return freeBytesAvailable, nil
}
