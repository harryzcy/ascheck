package localcheck

import (
	"debug/macho"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArchitecture_Load(t *testing.T) {
	tests := []struct {
		cpu      macho.Cpu
		expected Architectures
	}{
		{macho.CpuPpc, Architectures{PowerPC: 0b01}},
		{macho.CpuPpc64, Architectures{PowerPC: 0b10}},
		{macho.Cpu386, Architectures{Intel: 0b01}},
		{macho.CpuAmd64, Architectures{Intel: 0b10}},
		{macho.CpuArm, Architectures{Arm: 0b01}},
		{macho.CpuArm64, Architectures{Arm: 0b10}},
	}

	for _, test := range tests {
		arch := Architectures{}
		assert.Empty(t, arch)

		arch.Load(test.cpu)
		assert.Equal(t, test.expected, arch)
	}
}

func TestArchitecture_LoadFat(t *testing.T) {
	tests := []struct {
		in       []macho.FatArch
		expected Architectures
	}{
		{[]macho.FatArch{
			{FatArchHeader: macho.FatArchHeader{Cpu: macho.CpuAmd64}},
		},
			Architectures{Intel: 0b10},
		},
		{[]macho.FatArch{
			{FatArchHeader: macho.FatArchHeader{Cpu: macho.CpuAmd64}},
			{FatArchHeader: macho.FatArchHeader{Cpu: macho.CpuArm64}},
		},
			Architectures{Intel: 0b10, Arm: 0b10},
		},
	}

	for _, test := range tests {
		arch := Architectures{}
		assert.Empty(t, arch)

		arch.LoadFromFat(test.in)
		assert.Equal(t, test.expected, arch)
	}
}
