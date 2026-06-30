/*
Collection represents a collection of media files and its associated metadata.
It contains the name of the collection, the path to the collection directory,
and the manifest that describes the collection's properties.
*/
package collection

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// Create creates a new collection with the specified name in the given root directory.
func Create(rootDir, name string) (Collection, error) {
	path := filepath.Join(rootDir, name)

	if _, err := os.Stat(path); err == nil {
		return Collection{}, fmt.Errorf("collection %q already exists", name)
	} else if !errors.Is(err, os.ErrNotExist) {
		return Collection{}, err
	}

	if err := os.MkdirAll(path, 0755); err != nil {
		return Collection{}, err
	}

	manifest := Manifest{
		Version: ManifestVersion,
		Movie:   name,
	}

	data, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return Collection{}, err
	}

	if err := os.WriteFile(
		filepath.Join(path, "manifest.json"),
		data,
		0644,
	); err != nil {
		return Collection{}, err
	}

	return Collection{
		Name:     name,
		Path:     path,
		Manifest: manifest,
	}, nil
}

// Open opens an existing collection from the specified path by reading its manifest file.
func Open(path string) (Collection, error) {
	data, err := os.ReadFile(filepath.Join(path, "manifest.json"))
	if err != nil {
		return Collection{}, err
	}

	var manifest Manifest

	if err := json.Unmarshal(data, &manifest); err != nil {
		return Collection{}, err
	}

	return Collection{
		Name:     manifest.Movie,
		Path:     path,
		Manifest: manifest,
	}, nil
}

// List lists all collections in the specified root directory.
func List(rootDir string) ([]Collection, error) {
	entries, err := os.ReadDir(rootDir)
	if err != nil {
		return nil, err
	}

	var collections []Collection

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		c, err := Open(filepath.Join(rootDir, entry.Name()))
		if err != nil {
			continue
		}

		collections = append(collections, c)
	}

	sort.Slice(collections, func(i, j int) bool {
		return collections[i].Name < collections[j].Name
	})

	return collections, nil
}

// Delete deletes the collection at the specified path.
func Delete(path string) error {
	return os.RemoveAll(path)
}
