package builder

import (
	"os/exec"
	"path/filepath"
)

type DirTree struct {
	DirName    string
	InnerDirs  []DirTree
	InnerFiles []string
}

func (dt *DirTree) BuildDirTree(root string) error {
	dirPath := filepath.Join(root, dt.DirName)
	cmd := exec.Command("mkdir", dirPath)
	if err := cmd.Run(); err != nil {
		return err
	}

	for _, file := range dt.InnerFiles {
    file := filepath.Join(dirPath,file)
    mkfile := exec.Command("touch",file)
		if err := mkfile.Run(); err != nil {
			return err
		}
	}

	for _, dir := range dt.InnerDirs {
		err := dir.BuildDirTree(dirPath)
		if err != nil {
			return err
		}
	}
	return nil
}

var DIR_TEST = DirTree{
	DirName:    "test",
	InnerDirs:  []DirTree{dir1, dir2},
	InnerFiles: []string{"example0.txt"},
}
var dir1 = DirTree{
	DirName:    "dir1",
	InnerDirs:  []DirTree{dir3},
	InnerFiles: []string{"example1.txt"},
}
var dir2 = DirTree{
	DirName:    "dir2",
	InnerDirs:  []DirTree{},
	InnerFiles: []string{"example2.txt"},
}
var dir3 = DirTree{
	DirName:    "dir3",
	InnerDirs:  []DirTree{},
	InnerFiles: []string{"exmaple3.txt"},
}
