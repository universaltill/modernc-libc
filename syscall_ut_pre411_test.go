//go:build linux && (amd64 || arm64 || arm)

package libc // import "modernc.org/libc"

import (
	"runtime"
	"testing"
)

func TestUTPre411Kernel(t *testing.T) {
	tls := NewTLS()
	defer tls.Close()
	got := X__syscall5(tls, long(SYS_statx), -100, 0, 0, 0, 0)
	if utPre411Kernel(long(SYS_statx)) != (got == -long(ENOSYS)) {
		t.Fatalf("filter and X__syscall5 disagree: filter=%v got=%d", utPre411Kernel(long(SYS_statx)), got)
	}
}

func TestUTPre411KernelRange(t *testing.T) {
	arm := runtime.GOARCH == "arm"
	for _, c := range []struct {
		n    long
		want bool
	}{
		{396, false},      // pkey_free, last pre-4.11 number: reaches the kernel
		{397, arm},        // statx
		{403, arm},        // clock_gettime64
		{439, arm},        // faccessat2
		{0x0effff, arm},   // last ordinary number
		{0x0f0002, false}, // __ARM_NR_cacheflush
		{0x0f0005, false}, // __ARM_NR_set_tls
	} {
		if got := utPre411Kernel(c.n); got != c.want {
			t.Errorf("utPre411Kernel(%#x) = %v, want %v", c.n, got, c.want)
		}
	}
}
