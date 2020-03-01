package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-ps"
)

func main() {
	getProcessesInfo()
	getProcessInfo()
}

// get list of information of Golang Processes
func getProcessesInfo() {
	pList, _ := ps.Processes()

	for i, process := range pList {
		fmt.Printf("#No.%d = %+v\n", i, process)
	}

}

// get information of this application's process info itself
func getProcessInfo() {
	pid := os.Getpid()
	pInfo, _ := ps.FindProcess(pid)
	fmt.Printf("ProcessInfo:%v", pInfo)
	fmt.Printf("# PID   : %d\n", pInfo.Pid())
	fmt.Printf("# PPID  : %d\n", pInfo.PPid())
	fmt.Printf("# PNAME : %s\n", pInfo.Executable())

	ppidInfo, _ := ps.FindProcess(pInfo.PPid())
	fmt.Printf("\n PARENT NAME : %s\n", ppidInfo.Executable())
}
