package files

import (
	"encoding/gob"
	"link-collector-bot/lib/e"
	"link-collector-bot/storage"
	"os"
	"path/filepath"
)

type Storage struct {
	basePath string
}

const (
	defaultPerm = 0774
)

func New(basePath string) *Storage {
	return &Storage{basePath: basePath}
}

func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = e.WrapIsErr("can't save page", err) }()

	fPath := filepath.Join(s.basePath, page.UserName)

	if err := os.MkdirAll(fPath, defaultPerm); err != nil {
		return err
	}

	fName, err := fileName(page)
	if err != nil {
		return err
	}

	fPath = filepath.Join(fPath, fName)

	file, err := os.Create(fPath)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	if err := gob.NewEncoder(file).Encode(page); err != nil {
		return err
	}

	return nil
}



func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}