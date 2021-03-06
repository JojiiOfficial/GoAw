package gaw

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

//GetCurrentDir gets the current directory the user is is
func GetCurrentDir() string {
	exec, err := os.Getwd()
	if err != nil {
		log.Fatalln(err.Error())
		return ""
	}
	return exec
}

//GetHome returns the home directory of the current user
func GetHome() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err.Error())
		return ""
	}
	return home
}

// ResolveFullPath resolve path
func ResolveFullPath(fPath string) string {
	fPath = filepath.Clean(fPath)

	if strings.HasPrefix(fPath, "~") {
		fPath = filepath.Join(GetHome(), fPath[1:])
	}

	if strings.HasPrefix(fPath, ".") {
		fPath = filepath.Join(GetCurrentDir(), fPath[1:])
	}

	return fPath
}

// DirAbs returns the absolute path from human input (./ or ~/) and if it exists
func DirAbs(scriptPath string) (string, bool) {
	s, err := os.Stat(scriptPath)
	if err != nil || s == nil || !s.IsDir() {
		return scriptPath, false
	}
	if strings.HasPrefix(scriptPath, "/") {
		return scriptPath, true
	}

	if strings.HasPrefix(scriptPath, "./") {
		return filepath.Join(GetCurrentDir(), scriptPath[2:]), true
	}

	if strings.HasPrefix(scriptPath, "~/") {
		return filepath.Join(GetHome(), scriptPath[2:]), true
	}

	return filepath.Join(GetCurrentDir(), scriptPath), true
}

// ListDir get all files in a directory
func ListDir(dir string, filesOnly bool) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if filesOnly && info.IsDir() {
			return nil
		}

		files = append(files, path)
		return nil
	})

	return files, err
}

// GetDirSize get size of all files in a directory
func GetDirSize(dir string) (int64, error) {
	var size int64

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		size += info.Size()
		return nil
	})

	return size, err
}
