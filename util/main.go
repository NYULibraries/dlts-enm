package util

import(
	"path/filepath"
	"os"
)

// TODO:
//   * Figure out if there's a better place to put this
//   * Allow user to specify cache path
// Tried using os.TempDir(), but it was returning
// /var/folders/dh/48wd7vnj3xqd1w_f126tcnvh0000gn/T/, which was not as convenient.
var Cache = "/tmp"

// Source: http://stackoverflow.com/questions/33450980/golang-remove-all-contents-of-a-directory
// Author: http://stackoverflow.com/users/221700/peterso
func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
