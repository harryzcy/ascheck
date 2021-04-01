package localcheck

import (
	"debug/macho"
	"strings"
)

// Architectures represents all supported architecture of an app
type Architectures struct {
	Intel   uint
	Arm     uint
	PowerPC uint
}

// Load loads the architectures from macho.Cpu
func (arch *Architectures) Load(cpu macho.Cpu) {
	switch cpu {
	case macho.Cpu386:
		arch.Intel |= 0b01
	case macho.CpuAmd64:
		arch.Intel |= 0b10
	case macho.CpuArm:
		arch.Arm |= 0b01
	case macho.CpuArm64:
		arch.Arm |= 0b10
	case macho.CpuPpc:
		arch.PowerPC |= 0b01
	case macho.CpuPpc64:
		arch.PowerPC |= 0b10
	}
}

// LoadFromFat loads the architectures from []macho.FatArch
func (arch *Architectures) LoadFromFat(src []macho.FatArch) {
	for _, fat := range src {
		arch.Load(fat.Cpu)
	}
}

// String returns the architecture in string format
func (arch *Architectures) String() string {
	var list []string

	if arch.Intel > 0 {
		list = append(list, "Intel "+getBitString(arch.Intel))
	}
	if arch.Arm > 0 {
		list = append(list, "Arm "+getBitString(arch.Arm))
	}
	if arch.PowerPC > 0 {
		list = append(list, "PowerPC "+getBitString(arch.PowerPC))
	}

	if len(list) > 0 {
		return strings.Join(list, ", ")
	}

	return "Unknown"
}

func getBitString(mask uint) string {
	switch mask {
	case 0b11:
		return "32/64"
	case 0b01:
		return "32"
	case 0b10:
		return "64"
	default:
		return ""
	}
}
