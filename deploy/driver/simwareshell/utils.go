package simwareshell

import "os/exec"

func getCmdPath(cmd string)(string, error) {
	path, err := exec.LookPath(cmd)
	if err != nil {
		return "", nil
	}
	return path, nil
}
