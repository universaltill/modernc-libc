// Universal Till patch (universaltill/universal-till#6300, ut-docs ADR-0188).
// Not part of upstream modernc.org/libc.

//go:build linux && (amd64 || arm64 || loong64 || ppc64le || s390x || riscv64 || 386 || arm)

package libc // import "modernc.org/libc"

import "runtime"

// utFirstNewSyscallArm is the first 32-bit ARM syscall number added after
// Linux 4.10: statx is 397, and every time64 variant (clock_gettime64 403 ..
// futex_time64 422), pidfd_* and faccessat2 follow it.
const utFirstNewSyscallArm = 397

// utARMPrivateBase is __ARM_NR_BASE: the ARM-private calls (breakpoint,
// cacheflush, usr26, usr32, set_tls) live at 0x0f0001.. and are known to every
// kernel and allowed by Android's app filter, so they are never gated.
const utARMPrivateBase = 0x0f0000

// utPre411Kernel reports whether syscall n must be answered with -ENOSYS
// without ever reaching the kernel. musl's 32-bit code tries the newest
// syscall first (statx, clock_gettime64, ...) and falls back to the legacy
// one on -ENOSYS. Android 8-10 app seccomp filters do not know the new numbers
// and kill the process with SIGSYS instead of returning -ENOSYS, so the
// fallback is never reached. Answering from user space gives musl the exact
// answer a 4.9 kernel would, and the fallback runs. 32-bit ARM only; every
// other GOARCH compiles this to a constant false.
func utPre411Kernel(n long) bool {
	return runtime.GOARCH == "arm" && n >= utFirstNewSyscallArm && n < utARMPrivateBase
}
