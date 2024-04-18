package pclntab

import (
	"encoding/binary"
	"io"
)

type ExecutableFileFormat string

const (
	MachOExecutableFileFormat       = "Mach-O"
	ELFExecutableFileFormat         = "ELF"
	PEExecutableFileFormat          = "PE"
	UnsupportedExecutableFileFormat = "Unsupported"

	ELFMagicNumber      uint32 = 0x7f454c46 // 0x7f E L F
	MachO32MacgicNumber uint32 = 0xfeedface
	MachO64MacgicNumber uint32 = 0xfeedfacf
)

func IdentifyExecutableFileFormat(r io.ReaderAt) ExecutableFileFormat {
	var header [4]byte
	if _, err := r.ReadAt(header[0:], 0); err != nil {
		return UnsupportedExecutableFileFormat
	}

	bigEndianHeader := binary.BigEndian.Uint32(header[0:])
	littleEndianHeader := binary.LittleEndian.Uint32(header[0:])

	switch bigEndianHeader {
	case ELFMagicNumber:
		return ELFExecutableFileFormat
	case MachO32MacgicNumber, MachO64MacgicNumber:
		return MachOExecutableFileFormat
	default:
		// do nothing
	}

	switch littleEndianHeader {
	case ELFMagicNumber:
		return ELFExecutableFileFormat
	case MachO32MacgicNumber, MachO64MacgicNumber:
		return MachOExecutableFileFormat
	default:
		// do nothing
	}

	return UnsupportedExecutableFileFormat
}
