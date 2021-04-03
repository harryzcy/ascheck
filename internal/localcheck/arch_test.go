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

func TestArchitecture_String(t *testing.T) {
	tests := []struct {
		arch     Architectures
		expected string
	}{
		{Architectures{PowerPC: 0b01}, "PowerPC 32"},
		{Architectures{PowerPC: 0b10}, "PowerPC 64"},
		{Architectures{PowerPC: 0b11}, "PowerPC 32/64"},
		{Architectures{Intel: 0b01}, "Intel 32"},
		{Architectures{Intel: 0b10}, "Intel 64"},
		{Architectures{Intel: 0b11}, "Intel 32/64"},
		{Architectures{Arm: 0b01}, "Arm 32"},
		{Architectures{Arm: 0b10}, "Arm 64"},
		{Architectures{Arm: 0b11}, "Arm 32/64"},

		{Architectures{Intel: 0b10, Arm: 0b10}, "Intel 64, Arm 64"},
		{Architectures{Intel: 0b11, Arm: 0b10}, "Intel 32/64, Arm 64"},
		{Architectures{PowerPC: 0b11, Arm: 0b11}, "PowerPC 32/64, Arm 32/64"},

		{Architectures{}, "Unknown"},
	}

	for _, test := range tests {
		actual := test.arch.String()
		assert.Equal(t, test.expected, actual)
	}
}

func TestGetBitString_EdgeCase(t *testing.T) {
	assert.Empty(t, getBitString(0))
}
