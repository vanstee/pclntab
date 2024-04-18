package pclntab

import (
	"debug/elf"
	"debug/gosym"
	"fmt"
	"io"
)

func ParseELFPclntab(r io.ReaderAt) (*gosym.Table, error) {
	f, err := elf.NewFile(r)
	if err != nil {
		return nil, err
	}

	gopclntabSection := f.Section(".gopclntab")
	if gopclntabSection == nil {
		// TODO: Support executables built with -buildmode=pie
		return nil, fmt.Errorf("failed to retrieve .gopclntab section")
	}

	gopclntabSectionData, err := gopclntabSection.Data()
	if err != nil {
		return nil, err
	}

	textSection := f.Section(".text")
	if textSection == nil {
		return nil, fmt.Errorf("failed to retrieve .text section")
	}
	textSectionAddr := textSection.Addr

	lineTable := gosym.NewLineTable(gopclntabSectionData, textSectionAddr)
	symbolTable, err := gosym.NewTable([]byte{}, lineTable)
	if err != nil {
		return nil, err
	}

	return symbolTable, nil
}
