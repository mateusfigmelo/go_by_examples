package testutil

import (
	"os"
	"path/filepath"
)

// TempDir creates a temporary directory and returns its path.
// The directory will be automatically removed when the test completes.
func TempDir(prefix string) (string, func(), error) {
	dir, err := os.MkdirTemp("", prefix)
	if err != nil {
		return "", nil, err
	}

	cleanup := func() {
		os.RemoveAll(dir)
	}

	return dir, cleanup, nil
}

// CreateTempFile creates a temporary file with the given content.
// Returns the file path and a cleanup function.
func CreateTempFile(dir, prefix string, content []byte) (string, func(), error) {
	f, err := os.CreateTemp(dir, prefix)
	if err != nil {
		return "", nil, err
	}
	defer f.Close()

	if _, err := f.Write(content); err != nil {
		os.Remove(f.Name())
		return "", nil, err
	}

	cleanup := func() {
		os.Remove(f.Name())
	}

	return f.Name(), cleanup, nil
}

// FileExists checks if a file exists at the given path
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// DirExists checks if a directory exists at the given path
func DirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// ReadFile reads the entire file at path and returns its contents
func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// WriteFile writes data to a file at path
func WriteFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

// ListFiles returns a list of files in the given directory
func ListFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
