package util

import(
	"path/filepath"
	"os"
)

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
