package internal

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Storage struct {
	dataDir string
}

func NewStorage(dataDir string) *Storage {
	return &Storage{
		dataDir: dataDir,
	}
}

func (s *Storage) CreateBucket(name string) error {
	return os.Mkdir(filepath.Join(s.dataDir, name), 0644)
}

func (s *Storage) ReadBucket(name string) error {
	fileInfo, err := os.Stat(filepath.Join(s.dataDir, name))
	if fileInfo.IsDir() {
		return nil
	} else {
		return os.ErrExist
	}
	return err
}

func (s *Storage) DeleteBucket(name string) error {
	return os.Remove(filepath.Join(s.dataDir, name))
}

func (s *Storage) ListBuckets(maxResults int, pageToken string, prefix string) ([]string, error) {
	var err error
	f, err := os.Open(s.dataDir)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	var buckets []string
	buckets = make([]string, 0, maxResults)
	err = nil
	prefixLen := len(prefix)
	tokenFound := false
	for len(buckets) < maxResults {
		fileInfo, err := f.Readdir(maxResults)
		if err == io.EOF {
			if len(buckets) > 0 {
				err = nil
			}
			break
		} else if err == nil {
			for _, file := range fileInfo {
				// filter directories
				if !file.IsDir() {
					continue
				}
				fileName := file.Name()
				// filter by the start token
				if !tokenFound {
					if fileName == pageToken {
						tokenFound = true
					}
					continue
				}
				// filter by the prefix
				if prefixLen > 0 && !strings.HasPrefix(fileName, prefix) {
					continue
				}
				buckets = append(buckets, fileName)
			}
		}
	}
	return buckets, err
}
