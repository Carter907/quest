package graph

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// ArchiveGraph packages all the markdown guides into a .kng zip archive.
func ArchiveGraph(dirPath string, outputPath string) error {
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	w := zip.NewWriter(outFile)
	defer w.Close()

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		// Only archive markdown files
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".md" {
			continue
		}

		path := filepath.Join(dirPath, entry.Name())

		f, err := w.Create(entry.Name())
		if err != nil {
			return err
		}

		fileContent, err := os.Open(path)
		if err != nil {
			return err
		}

		_, err = io.Copy(f, fileContent)
		fileContent.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
