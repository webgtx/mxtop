package conf

import (
	"errors"
	"fmt"
	"os"

	"github.com/ssleert/ginip"
)

// mxtop conf file example
//
// [mxtop]
// ansible_dir = /usr/share/mxtop/ansible/

const ConfDir = "/etc/mxtop/"

var (
	AnsibleDir = "/etc/mxtop/ansible/"
)

var (
	ErrAnsibleDirNotExist = errors.New("ansible_dir not exist")
	ErrAnsibleDirPerm     = errors.New("ansible_dir permission denied")
	ErrAnsibleDirNotDir   = errors.New("ansible_dir is not a directory")
	ErrConfFileIsADir     = errors.New("conf file is a directory")
	ErrConfFileNotExist   = errors.New("conf file not exist")
	ErrConfFileNotRegular = errors.New("conf file is not a regular file")
	ErrConfFilePerm       = errors.New("conf file permission denied")
)

func Parse() error {
	iniPath := os.Getenv("MXTOP_CONF_FILE")
	if iniPath == "" {
		iniPath = ConfDir + "conf.ini"
	}
	iniStat, err := os.Lstat(iniPath)
	fmt.Println(err)
	if os.IsNotExist(err) {
		return ErrConfFileNotExist
	}
	if mode := iniStat.Mode(); mode.IsDir() {
		return ErrConfFileIsADir
	} else if !mode.IsRegular() {
		return ErrConfFileNotRegular
	}

	ini, err := ginip.Load(iniPath)
	if errors.Is(err, os.ErrPermission) {
		return ErrConfFilePerm
	}
	confValue, err := ini.GetValueString("", "ansible_dir")
	if err != nil {
		return err
	}
	ansibleStat, err := os.Lstat(confValue)
	if os.IsNotExist(err) {
		return ErrAnsibleDirNotExist
	}
	if mode := ansibleStat.Mode(); !mode.IsDir() {
		return ErrAnsibleDirNotDir
	}

	AnsibleDir = confValue

	return nil
}
