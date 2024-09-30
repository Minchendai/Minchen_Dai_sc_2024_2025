package folder_test

import (
	"encoding/json"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
)

func Test_folder_MoveFolder(t *testing.T) {
	// TODO: your tests here
	t.Parallel()
	tests := [...]struct {
		testId  int
		name    string
		dst     string
		err     string
		folders []byte
		want    []byte
	}{
		// TODO: your tests here
		{1, "alive-tsunami", "national-screwball", "", []byte(`[{
			"name": "central-red-ghost",
			"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
			"paths": "steady-insect.endless-red-hulk.premium-shriek.enabled-scarlet-spider.central-red-ghost"
		},
		{
			"name": "national-screwball",
			"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
			"paths": "steady-insect.national-screwball"
		},
		{
			"name": "sacred-lady-shiva",
			"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
			"paths": "steady-insect.national-screwball.sacred-lady-shiva"
		},
		{
			"name": "quick-cyber",
			"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
			"paths": "steady-insect.national-screwball.sacred-lady-shiva.quick-cyber"
		},
		{
			"name": "alive-tsunami",
				"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.national-screwball.sacred-lady-shiva.quick-cyber.alive-tsunami"
		}]`),
			[]byte(`[{
			"name": "central-red-ghost",
			"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
			"paths": "steady-insect.endless-red-hulk.premium-shriek.enabled-scarlet-spider.central-red-ghost"
		},
		{
			"name": "national-screwball",
			"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
			"paths": "steady-insect.national-screwball"
		},
		{
			"name": "sacred-lady-shiva",
			"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
			"paths": "steady-insect.national-screwball.sacred-lady-shiva"
		},
		{
			"name": "quick-cyber",
			"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
			"paths": "steady-insect.national-screwball.sacred-lady-shiva.quick-cyber"
		},
		{
			"name": "alive-tsunami",
			"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
			"paths": "steady-insect.national-screwball.alive-tsunami"
		}]`)},
		{2, "not-existing-folder", "national-screwball", "source folder does not exist", []byte(`[{
			"name": "national-screwball",
			"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
			"paths": "steady-insect.national-screwball"
		}]`),
			[]byte(`[{
			"name": "national-screwball",
			"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
			"paths": "steady-insect.national-screwball"
		}]`)},
		{3, "national-screwball", "not-existing-folder", "destination folder does not exist", []byte(`[{
			"name": "national-screwball",
			"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
			"paths": "steady-insect.national-screwball"
		}]`),
			[]byte(`[{
			"name": "national-screwball",
			"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
			"paths": "steady-insect.national-screwball"
		}]`)},
		{4, "national-screwball", "national-screwball", "cannot move a folder to itself", []byte(`[{
			"name": "national-screwball",
			"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
			"paths": "steady-insect.national-screwball"
		}]`),
			[]byte(`[{
			"name": "national-screwball",
			"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
			"paths": "steady-insect.national-screwball"
		}]`)},
		{5, "national-screwball", "glowing-elongated", "cannot move a folder to a different organization", []byte(`[{
			"name": "national-screwball",
			"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
			"paths": "steady-insect.national-screwball"
		},{
			"name": "glowing-elongated",
			"org_id": "c1556e17-b7c0-45a3-a6ae-9546248fb17a",
			"paths": "stunning-horridus.sacred-moonstar.nearby-maestro.enabled-professor-monster.glowing-elongated"
		}]`),
			[]byte(`[{
			"name": "national-screwball",
			"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
			"paths": "steady-insect.national-screwball"
		},{
			"name": "glowing-elongated",
			"org_id": "c1556e17-b7c0-45a3-a6ae-9546248fb17a",
			"paths": "stunning-horridus.sacred-moonstar.nearby-maestro.enabled-professor-monster.glowing-elongated"
		}]`)},
		{6, "national-screwball", "sacred-lady-shiva", "cannot move a folder to a child of itself", []byte(`[{
			"name": "national-screwball",
			"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
			"paths": "steady-insect.national-screwball"
		},
		{
			"name": "sacred-lady-shiva",
			"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
			"paths": "steady-insect.national-screwball.sacred-lady-shiva"
		}]`),
			[]byte(`[{
			"name": "national-screwball",
			"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
			"paths": "steady-insect.national-screwball"
		},
		{
			"name": "sacred-lady-shiva",
			"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
			"paths": "steady-insect.national-screwball.sacred-lady-shiva"
		}]`)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			folders := []folder.Folder{}
			json.Unmarshal(tt.want, &folders)
			f := folder.NewDriver(folders)
			newFolders, err := f.MoveFolder(tt.name, tt.dst)
			wantFolderMap := make(map[folder.Folder]bool)
			wanted := []folder.Folder{}
			json.Unmarshal(tt.want, &wanted)
			if tt.err != "" && err != nil && err.Error() != tt.err {
				t.Errorf("Error not match for test  %d, get: %s, expected: %s ", tt.testId, err.Error(), tt.err)
			}

			for _, curFolder := range wanted {
				wantFolderMap[curFolder] = true
			}
			for _, curFolder := range newFolders {
				if !wantFolderMap[curFolder] {
					t.Errorf("Test result not match for test  %d ", tt.testId)
					break
				} else {
					delete(wantFolderMap, curFolder)
				}
			}
		})
	}
}
