package ps

import "os"

// ps provides an API for finding and listing processes in a platform-agnostic
// way.
//
// NOTE: If you're reading these docs online via GoDocs or some other system,
// you might only see the Unix docs. This project makes heavy use of
// platform-specific implementations. We recommend reading the source if you
// are interested.

// Process is the generic interface that is implemented on every platform
// and provides common operations for processes.
type Process interface {
	// Pid is the process ID for this process.
	Pid() int

	// PPid is the parent process ID for this process.
	PPid() int

	// Executable name running this process. This is not a path to the
	// executable.
	Executable() string

	// Owner is the account name of the process owner.
	Owner() string

	// Architecture is the architecture of the process.
	Architecture() string
}

// Processes returns all processes.
//
// This of course will be a point-in-time snapshot of when this method was
// called. Some operating systems don't provide snapshot capability of the
// process table, in which case the process table returned might contain
// ephemeral entities that happened to be running when this was called.
func Processes(fullInfo bool) ([]Process, error) {
	return processes(fullInfo)
}

// FindProcess looks up a single process by pid.
//
// Process will be nil and error will be nil if a matching process is
// not found.
func FindProcess(pid int, fullInfo bool) (Process, error) {
	return findProcess(pid, fullInfo)
}

// Kill finds a process from a PID and terminates it.
func Kill(pid int) error {
	p, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	return p.Kill()
}
