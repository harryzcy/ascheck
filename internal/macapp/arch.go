package macapp

import "debug/macho"

func resolveArch(arch macho.Cpu) string {
	switch arch {
	case macho.Cpu386:
		return "Intel 32"
	case macho.CpuAmd64:
		return "Intel 64"
	case macho.CpuArm:
		return "Arm 32"
	case macho.CpuArm64:
		return "Arm 64"
	case macho.CpuPpc:
		return "PowerPC 32"
	case macho.CpuPpc64:
		return "PowerPC 64"
	default:
		return "Unknown"
	}
}
