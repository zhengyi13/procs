package procs

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)


type proc struct {
	pid int
	ppid int
	tpid int
}

func GetParentPid(pid int) (ppid int, err error) {
	path := fmt.Sprintf("/proc/%s/status", strconv.Itoa(pid))
 	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		t := scanner.Text()
 		if matched, _ := regexp.MatchString(`^PPid`, t); matched != false {
			s := strings.Split(t, ":")
			ppid, err = strconv.Atoi(strings.TrimSpace(s[1]))
			if err != nil {
				return 0, err
			}
		}
 	}

	return ppid, nil
}

func GetMyPid() (pid int, err error) {
 	f, err := os.Open("/proc/self/status")
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		t := scanner.Text()
 		if matched, _ := regexp.MatchString(`^Pid`, t); matched != false {
			s := strings.Split(t, ":")
			pid, err = strconv.Atoi(strings.TrimSpace(s[1]))
			if err != nil {
				return 0, err
			}
		}
	}

	return pid, nil
}
