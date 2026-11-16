//go:build linux && (amd64 || arm64 || arm)

package libc // import "modernc.org/libc"

import "testing"

func TestUTPre411Kernel(t *testing.T) {
	tls := NewTLS()
	defer tls.Close()
	got := X__syscall5(tls, long(SYS_statx), -100, 0, 0, 0, 0)
	if utPre411Kernel(long(SYS_statx)) != (got == -long(ENOSYS)) {
		t.Fatalf("filter and X__syscall5 disagree: filter=%v got=%d", utPre411Kernel(long(SYS_statx)), got)
	}
}
