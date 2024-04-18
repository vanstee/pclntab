package pclntab

import (
	"debug/gosym"
	"fmt"
	"io"
	"os"
)

func ParsePclntab(r io.ReaderAt) (*gosym.Table, error) {
	switch f := IdentifyExecutableFileFormat(r); f {
	case MachOExecutableFileFormat:
		return ParseMachOPclntab(r)
	case ELFExecutableFileFormat:
		return ParseELFPclntab(r)
	case PEExecutableFileFormat:
		return nil, fmt.Errorf("executable file format PE, not yet supported")
	case UnsupportedExecutableFileFormat:
		return nil, fmt.Errorf("executable file format not supported, available options: ELF, Mach-O")
	default:
		return nil, fmt.Errorf("unexpected executable file format: %v", f)
	}
}

func ParsePclntabPath(path string) (*gosym.Table, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ParsePclntab(f)
}
