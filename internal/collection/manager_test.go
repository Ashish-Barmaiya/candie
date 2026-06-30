package collection

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestCreate(t *testing.T) {
	root := t.TempDir()

	c, err := Create(root, "The Batman")
	if err != nil {
		t.Fatalf("Create() returned error: %v", err)
	}

	if c.Name != "The Batman" {
		t.Errorf("expected name %q, got %q", "The Batman", c.Name)
	}

	if _, err := os.Stat(c.Path); err != nil {
		t.Fatalf("collection directory does not exist: %v", err)
	}

	manifestPath := filepath.Join(c.Path, "manifest.json")

	if _, err := os.Stat(manifestPath); err != nil {
		t.Fatalf("manifest.json was not created: %v", err)
	}

	if c.Manifest.Version != ManifestVersion {
		t.Errorf(
			"expected manifest version %d, got %d",
			ManifestVersion,
			c.Manifest.Version,
		)
	}

	if c.Manifest.Movie != "The Batman" {
		t.Errorf(
			"expected movie %q, got %q",
			"The Batman",
			c.Manifest.Movie,
		)
	}
}

func TestOpen(t *testing.T) {
	root := t.TempDir()

	created, err := Create(root, "Oppenheimer")
	if err != nil {
		t.Fatalf("Create() returned error: %v", err)
	}

	opened, err := Open(created.Path)
	if err != nil {
		t.Fatalf("Open() returned error: %v", err)
	}

	if opened.Name != created.Name {
		t.Errorf("expected %q, got %q", created.Name, opened.Name)
	}

	if opened.Path != created.Path {
		t.Errorf("expected %q, got %q", created.Path, opened.Path)
	}

	if opened.Manifest.Movie != "Oppenheimer" {
		t.Errorf(
			"expected movie %q, got %q",
			"Oppenheimer",
			opened.Manifest.Movie,
		)
	}

	if opened.Manifest.Version != ManifestVersion {
		t.Errorf(
			"expected version %d, got %d",
			ManifestVersion,
			opened.Manifest.Version,
		)
	}
}

func TestList(t *testing.T) {
	root := t.TempDir()

	_, _ = Create(root, "Movie B")
	_, _ = Create(root, "Movie A")
	_, _ = Create(root, "Movie C")

	collections, err := List(root)
	if err != nil {
		t.Fatalf("List() returned error: %v", err)
	}

	if len(collections) != 3 {
		t.Fatalf("expected 3 collections, got %d", len(collections))
	}

	expected := []string{
		"Movie A",
		"Movie B",
		"Movie C",
	}

	for i, name := range expected {
		if collections[i].Name != name {
			t.Errorf(
				"expected collection %d to be %q, got %q",
				i,
				name,
				collections[i].Name,
			)
		}
	}
}

func TestDelete(t *testing.T) {
	root := t.TempDir()

	c, err := Create(root, "Dune")
	if err != nil {
		t.Fatalf("Create() returned error: %v", err)
	}

	if err := Delete(c.Path); err != nil {
		t.Fatalf("Delete() returned error: %v", err)
	}

	_, err = os.Stat(c.Path)

	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("expected directory to be deleted")
	}
}

func TestCreateAlreadyExists(t *testing.T) {
	root := t.TempDir()

	_, err := Create(root, "Interstellar")
	if err != nil {
		t.Fatalf("Create() returned error: %v", err)
	}

	_, err = Create(root, "Interstellar")
	if err == nil {
		t.Fatal("expected error when creating an existing collection")
	}
}

func TestOpenMissingManifest(t *testing.T) {
	root := t.TempDir()

	path := filepath.Join(root, "Interstellar")

	if err := os.Mkdir(path, 0755); err != nil {
		t.Fatal(err)
	}

	_, err := Open(path)
	if err == nil {
		t.Fatal("expected error when manifest.json is missing")
	}
}
