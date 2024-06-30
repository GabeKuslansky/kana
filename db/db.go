package db

import (
	"bytes"
	"encoding/gob"
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

const (
	DB_FILE = "db.kana"
)

var db *Database

type Database struct {
	DefaultDeckId int
	DefaultDeckName string
}

func UpdateDefaultDeck(id int, name string) error {
	db := Open()

	db.DefaultDeckId = id
	db.DefaultDeckName = name

	err := Save(db)
	if err != nil {
		return err
	}

	return nil
}

func Create() error {
	path, err := GetDatabasePath()
	if err != nil {
		return err
	}
	fileBytes := []byte{}
	os.WriteFile(path, fileBytes, 0644)

	return nil
}

func GetDatabasePath() (string, error) {
	dir, err := GetDataDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, DB_FILE), nil

}

func DbFileExists() bool {
	path, err := GetDatabasePath()
	if err != nil {
		panic(err)
	}

	if _, statErr := os.Stat(path); os.IsNotExist(statErr) {
		return false
	} else if statErr != nil {
		panic(statErr)
	}

	return true
}

func GetDataDir() (string, error) {
	_os := runtime.GOOS

	var dir string
	switch _os {
	// discovered from `go tool dist list`
	case "darwin":
		homeDir := os.Getenv("HOME")
		dir = filepath.Join(homeDir, "Library", "Application Support") // /Users/Alice/Library/Application Support
	case "linux", "freebsd", "openbsd", "netbsd", "dragonfly", "solaris", "illumos":
		dataDir := os.Getenv("XDG_DATA_HOME")
		if dataDir == "" {
			homeDir := os.Getenv("HOME")
			dataDir = filepath.Join(homeDir, ".local", "share")
		}
		dir = dataDir // /home/alice/.local/share
	case "windows":
		dir = os.Getenv("LOCALAPPDATA") // C:\Users\Alice\AppData\Local
	default:
		return "", errors.New("unsupported os")
	}

	dir = filepath.Join(dir, "kana")

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0777)
	} else if err != nil {
		panic(err)
	}

	return dir, nil
}

func Save(_db *Database) error {
	path, err := GetDatabasePath()
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	encoder.Encode(_db)
	os.WriteFile(path, buffer.Bytes(), 0777)
	db = _db

	return nil
}

func Open() *Database {
	_db, err := open()
	if err != nil {
		// db.kana not present, instantiate default values
		_db = &Database{
			DefaultDeckId: -1,
		}

		Save(_db)
		return _db
	}

	db = _db
	return _db
}

func open() (*Database, error) {
	if db != nil {
		return db, nil
	}

	path, err := GetDatabasePath()
	if err != nil {
		return nil, err
	}

	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(fileBytes)
	decoder := gob.NewDecoder(buf)

	var db Database
	err = decoder.Decode(&db)
	if err != nil {
		return nil, err
	}

	return &db, nil
}
