package testutil

import "io/fs"

// DirEnt create a fs.DirEntry
type DirEnt struct {
	Nam string
	Dir bool
	Typ fs.FileMode
	Fi  fs.FileInfo
	Err error
}

// Name get
func (d *DirEnt) Name() string {
	return d.Nam
}

// IsDir get
func (d *DirEnt) IsDir() bool {
	return d.Dir
}

// Type get
func (d *DirEnt) Type() fs.FileMode {
	return d.Typ
}

// Info get
func (d *DirEnt) Info() (fs.FileInfo, error) {
	return d.Fi, d.Err
}
