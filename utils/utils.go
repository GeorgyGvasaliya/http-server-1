package utils

import "path/filepath"

const layoutsPattern = "views/layouts/*.gohtml"

func LayoutFiles() []string {
	files, err := filepath.Glob(layoutsPattern)
	if err != nil {
		return nil
	}
	return files
}
