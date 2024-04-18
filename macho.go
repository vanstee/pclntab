package pclntab

import (
	"debug/gosym"
	"debug/macho"
	"fmt"
	"io"
)

func ParseMachOPclntab(r io.ReaderAt) (*gosym.Table, error) {
	f, err := macho.NewFile(r)
	if err != nil {
		return nil, err
	}

	gopclntabSection := f.Section("__gopclntab")
	if gopclntabSection == nil {
		return nil, fmt.Errorf("failed to retrieve __gopclntab section")
	}

	gopclntabSectionData, err := gopclntabSection.Data()
	if err != nil {
		return nil, err
	}

	textSection := f.Section("__text")
	if textSection == nil {
		return nil, fmt.Errorf("failed to retrieve __text section")
	}
	textSectionAddr := textSection.Addr

	lineTable := gosym.NewLineTable(gopclntabSectionData, textSectionAddr)
	symbolTable, err := gosym.NewTable([]byte{}, lineTable)
	if err != nil {
		return nil, err
	}

	return symbolTable, nil
}
