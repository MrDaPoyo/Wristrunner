package notes

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const storeVersion = 1
const JSON_FILENAME = "notes.json"

type noteFile struct {
	Version int    `json:"version"`
	Notes   []Note `json:"notes"`
}

type store struct{ path string }

func newStore(dir string) store {
	return store{path: filepath.Join(dir, JSON_FILENAME)}
}

func (s store) load() ([]Note, error) {
	data, err := os.ReadFile(s.path)
	if os.IsNotExist(err) {
		return nil, nil // first run
	}
	if err != nil {
		return nil, err
	}
	var f noteFile
	if err := json.Unmarshal(data, &f); err != nil {
		// keep the bad file for recovery instead of overwriting it
		_ = os.Rename(s.path, s.path+".corrupt")
		return nil, err
	}
	return f.Notes, nil
}

func (s store) save(notes []Note) error {
	data, err := json.MarshalIndent(noteFile{Version: storeVersion, Notes: notes}, "", "  ")
	if err != nil {
		return err
	}
	return writeAtomic(s.path, data)
}

// writeAtomic writes a sibling temp file, fsyncs, then renames into place.
func writeAtomic(path string, data []byte) error {
	tmp, err := os.CreateTemp(filepath.Dir(path), ".tmp-*")
	if err != nil {
		return err
	}
	name := tmp.Name()
	defer os.Remove(name) // no-op after a successful rename

	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Sync(); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Rename(name, path)
}
