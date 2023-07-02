//go:build !js
// +build !js

package qtls

import (
	"runtime"

	"golang.org/x/sys/cpu"
)

var (
	hasGCMAsmAMD64 = cpu.X86.HasAES && cpu.X86.HasPCLMULQDQ
	hasGCMAsmARM64 = cpu.ARM64.HasAES && cpu.ARM64.HasPMULL
	// Keep in sync with crypto/aes/cipher_s390x.go.
	hasGCMAsmS390X = cpu.S390X.HasAES && cpu.S390X.HasAESCBC && cpu.S390X.HasAESCTR &&
		(cpu.S390X.HasGHASH || cpu.S390X.HasAESGCM)

	HasAESGCMHardwareSupport = runtime.GOARCH == "amd64" && hasGCMAsmAMD64 ||
		runtime.GOARCH == "arm64" && hasGCMAsmARM64 ||
		runtime.GOARCH == "s390x" && hasGCMAsmS390X
)

type CPU_profile struct {
	HasAESGCMHardwareSupport bool
}

func (c *CPU_profile) HasAESGCMSP(b bool) bool {
	if b {
		HasAESGCMHardwareSupport = runtime.GOARCH == "amd64" && hasGCMAsmAMD64 ||
			runtime.GOARCH == "arm64" && hasGCMAsmARM64 ||
			runtime.GOARCH == "s390x" && hasGCMAsmS390X
	} else {
		HasAESGCMHardwareSupport = b
	}
	return HasAESGCMHardwareSupport

}
