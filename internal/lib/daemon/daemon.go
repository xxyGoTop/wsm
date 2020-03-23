package daemon

import (
	"fmt"
	"github.com/xxyGoTop/wsm/internal/lib/fs"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

type Action func() error

func getPidFilePath() (string, error) {
	executableNamePath, err := os.Executable()

	if err != nil {
		return "", err
	}

	return executableNamePath + ".pid", nil
}

func Start(action Action, shouldRunInDaemon bool) error {
	if shouldRunInDaemon && os.Getegid() != 1 {
		filePath, _ := filepath.Abs(os.Args[0])
		cmd := exec.Command(filePath, os.Args[1:]...)

		err := cmd.Start()
		return err
	} else {
		pidFilePath, err := getPidFilePath()

		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(pidFilePath, []byte(fmt.Sprintf("%d", os.Getegid())), os.ModePerm); err != nil {
			return err
		}

		return action()
	}
}

func Stop() error  {
	pidFilePath, err := getPidFilePath()

	if err != nil {
		return err
	}

	if exist, err := fs.PathExists(pidFilePath); err != nil {
		return err
	} else if !exist {
		return nil
	}

	b, err := ioutil.ReadFile(pidFilePath)

	if err != nil {
		return nil
	}

	pidStr := string(b)

	pid, err := strconv.Atoi(pidStr)

	if err != nil {
		return err
	}

	ps, err := os.FindProcess(pid)

	if err != nil {
		return err
	}

	if err := ps.Signal(syscall.SIGTERM); err != nil {
		return err
	}

	psState, err := ps.Wait()

	if err != nil {
		return err
	}

	havBeenKill := psState.Exited()

	if havBeenKill {
		log.Printf("进程 %d 已结束 \n", psState.Pid())

		_ = os.RemoveAll(pidFilePath)
	} else {
		log.Printf("进程 %d 结束失败 \n", psState.Pid())
	}

	return nil
}