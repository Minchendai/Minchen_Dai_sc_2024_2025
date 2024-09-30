package folder

import (
	"errors"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// Your code here...

	folders := f.folders
	var dstFolder *Folder
	var err error
	var fromFolder *Folder

	if name == dst {
		err = errors.New("cannot move a folder to itself")
	}

	if err == nil {
		for i, curfolder := range folders {
			if curfolder.Name == name {
				fromFolder = &folders[i]
			}
			if curfolder.Name == dst {
				dstFolder = &folders[i]
			}
			if fromFolder != nil && dstFolder != nil {
				break
			}
		}
	}

	if fromFolder == nil && err == nil {
		err = errors.New("source folder does not exist")
	}

	if dstFolder == nil && err == nil {
		err = errors.New("destination folder does not exist")
	}

	if err == nil && fromFolder.OrgId != dstFolder.OrgId {
		err = errors.New("cannot move a folder to a different organization")
	}

	if err == nil {
		for _, folderName := range strings.Split(dstFolder.Paths, ".") {
			if folderName == name {
				err = errors.New("cannot move a folder to a child of itself")
			}
		}
	}

	if err == nil {
		for i, curfolder := range folders {
			if strings.HasPrefix(curfolder.Paths, fromFolder.Paths) && curfolder.OrgId == fromFolder.OrgId {
				folders[i].Paths = dstFolder.Paths + "." + fromFolder.Name + strings.Split(curfolder.Paths, fromFolder.Name)[1]
			}
		}
	}

	return folders, err
}
