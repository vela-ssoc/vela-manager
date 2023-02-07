//go:build !linux

package lumber

import "os"

func chown(string, os.FileInfo) error {
	return nil
}
