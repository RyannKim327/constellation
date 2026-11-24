package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type runner struct {
	root     string
	buildDir string
	port     string
	cmd      *exec.Cmd
	binary   string
	buildNum int
}

func newRunner(root, buildDir, port string) *runner {
	return &runner{root: root, buildDir: buildDir, port: port}
}

func (r *runner) build() (string, error) {
	r.buildNum++
	binary := filepath.Join(r.buildDir, fmt.Sprintf("constellation-%d", r.buildNum))

	cmd := exec.Command("go", "build", "-o", binary, ".")
	cmd.Dir = r.root
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("build failed: %w", err)
	}

	return binary, nil
}

func (r *runner) start(binary string) error {
	cmd := exec.Command(binary)
	cmd.Env = append(os.Environ(), "PORT="+r.port)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return err
	}

	r.cmd = cmd
	r.binary = binary
	return nil
}

func (r *runner) stop() {
	if r.cmd == nil {
		return
	}

	_ = r.cmd.Process.Kill()
	_ = r.cmd.Wait()
	_ = os.Remove(r.binary)
	r.cmd = nil
	r.binary = ""
}

func (r *runner) restart(binary string) error {
	r.stop()
	if err := r.start(binary); err != nil {
		return err
	}
	if err := waitForPort(r.port, 3*time.Second); err != nil {
		r.stop()
		return err
	}
	return nil
}

func waitForPort(port string, timeout time.Duration) error {
	address := net.JoinHostPort("127.0.0.1", port)
	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		conn, err := net.DialTimeout("tcp", address, 250*time.Millisecond)
		if err == nil {
			_ = conn.Close()
			return nil
		}
		time.Sleep(50 * time.Millisecond)
	}

	return fmt.Errorf("server on port %s did not start in %s", port, timeout)
}
