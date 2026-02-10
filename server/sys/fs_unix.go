//go:build !windows

package sys

import (
	"github.com/marcopiovanello/yt-dlp-web-ui/v3/server/config"
	"golang.org/x/sys/unix"
)

// FreeSpace gets the available Bytes writable to download directory
func FreeSpace() (uint64, error) {
	var stat unix.Statfs_t
	if err := unix.Statfs(config.Instance().DownloadPath, &stat); err != nil {
		return 0, err
	}
	return stat.Bavail * uint64(stat.Bsize), nil
}
