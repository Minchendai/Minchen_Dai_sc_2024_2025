package folder

import (
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	// Your code here...
	folders := f.GetFoldersByOrgID(orgID)
	if len(folders) == 0 {
		panic("Error: Organization does not exist")
	}

	res := []Folder{}
	isValidFolder := false
	for _, f := range folders {
		if f.Name == name {
			isValidFolder = true
			continue
		}
		path := strings.Split(f.Paths, ".")
		for _, p := range path {
			if p == name {
				res = append(res, f)
				break
			}
		}
	}

	if len(res) == 0 && !isValidFolder {
		folders = f.folders
		for _, f := range folders {
			if f.Name == name {
				panic("Error: Folder does not exist in the specified organization")
			}
		}
		panic("Error: Folder does not exist")
	}

	return res
}
