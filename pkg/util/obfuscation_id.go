package util

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/er-azh/fortnitemanifest"
)

func GetLaunchCommand(gamePath string) (string, error) {
	var manifestPath string

	filepath.WalkDir(gamePath+"\\.egstore", func(path string, d fs.DirEntry, err error) error {
		if err != nil || manifestPath != "" || d.IsDir() {
			return err
		}
		if strings.HasSuffix(path, ".manifest") {
			manifestPath = path
		}
		return nil
	})

	f, err := os.Open(manifestPath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	manifest, err := fortnitemanifest.ParseManifest(f)
	if err != nil {
		return "", err
	}
	return manifest.Metadata.LaunchCommand, nil
}
