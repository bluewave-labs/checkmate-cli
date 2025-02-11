package fs

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CreateTar(sourceDir, destinationFile string) error {
	tarFile, err := os.Create(destinationFile)
	if err != nil {
		return err
	}
	defer tarFile.Close()

	tarWriter := tar.NewWriter(tarFile)
	defer tarWriter.Close()

	err = filepath.Walk(sourceDir, func(filePath string, fi os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return fmt.Errorf("error accessing file %s: %w", filePath, walkErr)
		}

		// Generate tar header for the file
		header, err := tar.FileInfoHeader(fi, fi.Name())
		if err != nil {
			return fmt.Errorf("error creating tar header for %s: %w", filePath, err)
		}

		// Update the header name to be relative to the source directory
		relativePath, err := filepath.Rel(sourceDir, filePath)
		if err != nil {
			return fmt.Errorf("error calculating relative path for %s: %w", filePath, err)
		}
		header.Name = filepath.ToSlash(relativePath)

		// Write the header to the tar archive
		if err := tarWriter.WriteHeader(header); err != nil {
			return fmt.Errorf("error writing tar header for %s: %w", filePath, err)
		}

		// If the file is not a regular file, skip content writing
		if !fi.Mode().IsRegular() {
			return nil
		}

		// Open the file for reading
		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("error opening file %s: %w", filePath, err)
		}
		defer file.Close()

		// Copy file content to the tar archive
		if _, err := io.Copy(tarWriter, file); err != nil {
			return fmt.Errorf("error writing file content for %s: %w", filePath, err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("error creating tar archive: %w", err)
	}

	return err
}

func ExtractTarToVolume(sourceFile, volumeMountPoint string) error {
	log.Printf("Starting extraction from %s to %s", sourceFile, volumeMountPoint)

	// Verify the source file exists and is readable
	sourceInfo, err := os.Stat(sourceFile)
	if err != nil {
		return fmt.Errorf("source file error: %w", err)
	}
	if sourceInfo.Size() == 0 {
		return fmt.Errorf("source file is empty: %s", sourceFile)
	}

	// Verify volume mount point exists and is writable
	mountInfo, err := os.Stat(volumeMountPoint)
	if err != nil {
		return fmt.Errorf("volume mount point error: %w", err)
	}
	if !mountInfo.IsDir() {
		return fmt.Errorf("volume mount point is not a directory: %s", volumeMountPoint)
	}

	// Test write permission
	testFile := filepath.Join(volumeMountPoint, ".write_test")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		return fmt.Errorf("volume mount point is not writable: %w", err)
	}
	os.Remove(testFile)

	log.Printf("Cleaning directory: %s", volumeMountPoint)
	if err := cleanDirectory(volumeMountPoint); err != nil {
		return fmt.Errorf("failed to clean volume directory: %w", err)
	}

	tarFile, err := os.Open(sourceFile)
	if err != nil {
		return fmt.Errorf("failed to open tar file: %w", err)
	}
	defer tarFile.Close()

	tarReader := tar.NewReader(tarFile)
	extractedFiles := 0

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error reading tar archive: %w", err)
		}

		targetPath := filepath.Join(volumeMountPoint, filepath.Clean(header.Name))
		log.Printf("Extracting: %s", header.Name)

		if !isWithinDirectory(volumeMountPoint, targetPath) {
			return fmt.Errorf("attempted path traversal blocked: %s", header.Name)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(targetPath, os.FileMode(header.Mode)); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", targetPath, err)
			}
			log.Printf("Created directory: %s", targetPath)

		case tar.TypeReg:
			if err := extractFile(targetPath, header.Mode, tarReader); err != nil {
				return fmt.Errorf("failed to extract file %s: %w", targetPath, err)
			}
			extractedFiles++
			log.Printf("Extracted file: %s", targetPath)

		default:
			log.Printf("Skipping unsupported file type: %s", header.Name)
		}
	}

	if extractedFiles == 0 {
		return fmt.Errorf("no files were extracted from the archive")
	}

	log.Printf("Successfully extracted %d files to %s", extractedFiles, volumeMountPoint)
	return nil
}

// Helper functions to support extractTarToVolume

func cleanDirectory(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		path := filepath.Join(dir, entry.Name())
		if err := os.RemoveAll(path); err != nil {
			return fmt.Errorf("failed to remove %s: %w", path, err)
		}
	}
	return nil
}

func extractFile(targetPath string, mode int64, reader io.Reader) error {
	if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
		return fmt.Errorf("failed to create parent directories: %w", err)
	}

	file, err := os.OpenFile(targetPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.FileMode(mode))
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	written, err := io.Copy(file, reader)
	if err != nil {
		return fmt.Errorf("failed to write file contents: %w", err)
	}
	if written == 0 {
		return fmt.Errorf("no data written to file: %s", targetPath)
	}

	return nil
}

func isWithinDirectory(dir, target string) bool {
	rel, err := filepath.Rel(dir, target)
	if err != nil {
		return false
	}
	return !strings.HasPrefix(rel, ".."+string(filepath.Separator)) && rel != ".."
}
