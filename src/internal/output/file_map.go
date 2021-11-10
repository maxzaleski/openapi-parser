package output

import "sync"

type (
	// FileMap represents a map of files.
	//
	// @returns map[string]*File
	FileMap interface {
		// Add safely adds a file to the map.
		Add(f *File)
		// Get safely gets a file from the map.
		Get(name string) *File
		// Range returns a slice of keys.
		Range() []string
	}

	fileMap struct {
		src map[string]*File
		mu  sync.Mutex
	}
)

// NewFileMap returns a new implementation of FileMap.
func NewFileMap() FileMap {
	return &fileMap{
		src: make(map[string]*File),
		mu:  sync.Mutex{},
	}
}

func (l *fileMap) Add(f *File) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.src[f.Name] = f
}

func (l *fileMap) Get(name string) *File {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.src[name]
}

func (l *fileMap) Range() []string {
	l.mu.Lock()
	defer l.mu.Unlock()

	keys := make([]string, 0, len(l.src))
	for k, _ := range l.src {
		keys = append(keys, k)
	}
	return keys
}
