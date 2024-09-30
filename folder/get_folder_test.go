package folder_test

import (
	"encoding/json"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	// "github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []byte
	}{
		// TODO: your tests here
		{"alive-tsunami", uuid.Must(uuid.FromString("9b4cdb0a-cfea-4f9d-8a68-24f038fae385")), folder.GetAllFolders(), []byte(`[{
		"name": "steady-insect",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect"
	},
	{
		"name": "helped-blackheart",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.helped-blackheart"
	},
	{
		"name": "many-silver-sable",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.helped-blackheart.many-silver-sable"
	},
	{
		"name": "stable-karatecha",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.helped-blackheart.many-silver-sable.stable-karatecha"
	},
	{
		"name": "exciting-magma",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.helped-blackheart.many-silver-sable.stable-karatecha.exciting-magma"
	},
	{
		"name": "upward-the-anarchist",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.helped-blackheart.many-silver-sable.upward-the-anarchist"
	},
	{
		"name": "patient-prodigy",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.helped-blackheart.many-silver-sable.upward-the-anarchist.patient-prodigy"
	},
	{
		"name": "obliging-microchip",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.helped-blackheart.many-silver-sable.upward-the-anarchist.obliging-microchip"
	},
	{
		"name": "massive-ser-duncan",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.helped-blackheart.many-silver-sable.upward-the-anarchist.massive-ser-duncan"
	},
	{
		"name": "sacred-mystique",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.helped-blackheart.many-silver-sable.upward-the-anarchist.sacred-mystique"
	},
	{
		"name": "ideal-miss-america",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.helped-blackheart.ideal-miss-america"
	},
	{
		"name": "premium-man-wolf",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.helped-blackheart.ideal-miss-america.premium-man-wolf"
	},
	{
		"name": "shining-american-eagle",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.helped-blackheart.ideal-miss-america.premium-man-wolf.shining-american-eagle"
	},
	{
		"name": "concrete-golden-guardian",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.helped-blackheart.ideal-miss-america.concrete-golden-guardian"
	},
	{
		"name": "adapted-captain-britain",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.helped-blackheart.ideal-miss-america.concrete-golden-guardian.adapted-captain-britain"
	},
	{
		"name": "proper-dolphin",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.helped-blackheart.ideal-miss-america.concrete-golden-guardian.proper-dolphin"
	},
	{
		"name": "ruling-outlaw-kid",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.helped-blackheart.ideal-miss-america.concrete-golden-guardian.ruling-outlaw-kid"
	},
	{
		"name": "casual-spectrum",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.helped-blackheart.ideal-miss-america.concrete-golden-guardian.casual-spectrum"
	},
	{
		"name": "native-deadpool",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.helped-blackheart.ideal-miss-america.native-deadpool"
	},
	{
		"name": "merry-fantomex",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.helped-blackheart.ideal-miss-america.native-deadpool.merry-fantomex"
	},
	{
		"name": "loyal-monstress",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.helped-blackheart.ideal-miss-america.loyal-monstress"
	},
	{
		"name": "discrete-shocker",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.helped-blackheart.ideal-miss-america.loyal-monstress.discrete-shocker"
	},
	{
		"name": "curious-green-lantern",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.helped-blackheart.ideal-miss-america.loyal-monstress.curious-green-lantern"
	},
	{
		"name": "endless-red-hulk",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk"
	},
	{
		"name": "complete-aquagirl",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk.complete-aquagirl"
	},
	{
		"name": "holy-raphael",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk.complete-aquagirl.holy-raphael"
	},
	{
		"name": "adapted-warbird",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk.complete-aquagirl.holy-raphael.adapted-warbird"
	},
	{
		"name": "allowing-dust",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk.complete-aquagirl.holy-raphael.allowing-dust"
	},
	{
		"name": "adjusted-titania",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk.complete-aquagirl.holy-raphael.adjusted-titania"
	},
	{
		"name": "champion-thunder",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk.complete-aquagirl.champion-thunder"
	},
	{
		"name": "precious-arrowette",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk.complete-aquagirl.champion-thunder.precious-arrowette"
	},
	{
		"name": "prompt-dynamite",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk.complete-aquagirl.champion-thunder.prompt-dynamite"
	},
	{
		"name": "suited-king-cobra",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk.complete-aquagirl.champion-thunder.suited-king-cobra"
	},
	{
		"name": "super-hiroim",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk.super-hiroim"
	},
	{
		"name": "divine-doctor",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk.super-hiroim.divine-doctor"
	},
	{
		"name": "square-tombstone",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk.super-hiroim.divine-doctor.square-tombstone"
	},
	{
		"name": "sincere-the-hunter",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk.super-hiroim.divine-doctor.sincere-the-hunter"
	},
	{
		"name": "cuddly-lady-bullseye",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk.super-hiroim.divine-doctor.cuddly-lady-bullseye"
	},
	{
		"name": "novel-guardian",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk.super-hiroim.divine-doctor.novel-guardian"
	},
	{
		"name": "profound-hiroim",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk.super-hiroim.profound-hiroim"
	},
	{
		"name": "nice-smiling-tiger",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk.super-hiroim.profound-hiroim.nice-smiling-tiger"
	},
	{
		"name": "advanced-free",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk.super-hiroim.profound-hiroim.advanced-free"
	},
	{
		"name": "stirred-gunslinger",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk.super-hiroim.stirred-gunslinger"
	},
	{
		"name": "prepared-comedian",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk.super-hiroim.stirred-gunslinger.prepared-comedian"
	},
	{
		"name": "grown-wolverine",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk.super-hiroim.stirred-gunslinger.grown-wolverine"
	},
	{
		"name": "composed-atlas",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk.super-hiroim.stirred-gunslinger.composed-atlas"
	},
	{
		"name": "premium-shriek",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk.premium-shriek"
	},
	{
		"name": "enabled-scarlet-spider",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk.premium-shriek.enabled-scarlet-spider"
	},
	{
		"name": "giving-wolfpack",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk.premium-shriek.enabled-scarlet-spider.giving-wolfpack"
	},
	{
		"name": "adequate-master-chief",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk.premium-shriek.enabled-scarlet-spider.adequate-master-chief"
	},
	{
		"name": "true-beetle",
		"org_id": "9b4cdb0a-cfea-4f9d-8a68-24f038fae385",
		"paths": "steady-insect.endless-red-hulk.premium-shriek.enabled-scarlet-spider.true-beetle"
	},
	{
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
	}]`)},
		{"not-existing-organization", uuid.Must(uuid.FromString("9b4cdb0a-cfea-4f9d-8a68-24f038fae386")), folder.GetAllFolders(), []byte(`[]`)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(tt.orgID)
			wantFolderMap := make(map[folder.Folder]bool)
			wanted := []folder.Folder{}
			err := json.Unmarshal(tt.want, &wanted)
			if err != nil {
				panic(err)
			}
			for _, curFolder := range wanted {
				wantFolderMap[curFolder] = true
			}
			for _, curFolder := range get {
				if !wantFolderMap[curFolder] {
					t.Errorf("Test result not match for Organization %s ", tt.orgID)
					break
				} else {
					delete(wantFolderMap, curFolder)
				}
			}
		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name          string
		orgID         uuid.UUID
		folders       []folder.Folder
		expectedError string
		want          []byte
	}{
		// TODO: your tests here
		{"bursting-lionheart", uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")), folder.GetAllFolders(), "", []byte(`[{
			"name": "striking-black-panther",
			"org_id": "38b9879b-f73b-4b0e-b9d9-4fc4c23643a7",
			"paths": "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart.striking-black-panther"
		},
		{
			"name": "advanced-professor-monster",
			"org_id": "38b9879b-f73b-4b0e-b9d9-4fc4c23643a7",
			"paths": "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart.advanced-professor-monster"
		},
		{
			"name": "assuring-red-shift",
			"org_id": "38b9879b-f73b-4b0e-b9d9-4fc4c23643a7",
			"paths": "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart.assuring-red-shift"
		},
		{
			"name": "merry-mega-man",
			"org_id": "38b9879b-f73b-4b0e-b9d9-4fc4c23643a7",
			"paths": "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart.merry-mega-man"
		}]`)},
		{"striking-black-panther", uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")), folder.GetAllFolders(), "", []byte(`[]`)},
		{"topical-micromax", uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")), folder.GetAllFolders(), "", []byte(`[{
		"name": "bursting-lionheart",
		"org_id": "38b9879b-f73b-4b0e-b9d9-4fc4c23643a7",
		"paths": "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart"
	},
	{
		"name": "striking-black-panther",
		"org_id": "38b9879b-f73b-4b0e-b9d9-4fc4c23643a7",
		"paths": "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart.striking-black-panther"
	},
	{
		"name": "advanced-professor-monster",
		"org_id": "38b9879b-f73b-4b0e-b9d9-4fc4c23643a7",
		"paths": "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart.advanced-professor-monster"
	},
	{
		"name": "assuring-red-shift",
		"org_id": "38b9879b-f73b-4b0e-b9d9-4fc4c23643a7",
		"paths": "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart.assuring-red-shift"
	},
	{
		"name": "merry-mega-man",
		"org_id": "38b9879b-f73b-4b0e-b9d9-4fc4c23643a7",
		"paths": "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart.merry-mega-man"
	},
	{
		"name": "patient-red-wolf",
		"org_id": "38b9879b-f73b-4b0e-b9d9-4fc4c23643a7",
		"paths": "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf"
	},
	{
		"name": "coherent-night-nurse",
		"org_id": "38b9879b-f73b-4b0e-b9d9-4fc4c23643a7",
		"paths": "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf.coherent-night-nurse"
	},
	{
		"name": "smashing-raphael",
		"org_id": "38b9879b-f73b-4b0e-b9d9-4fc4c23643a7",
		"paths": "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf.smashing-raphael"
	},
	{
		"name": "gentle-tempest",
		"org_id": "38b9879b-f73b-4b0e-b9d9-4fc4c23643a7",
		"paths": "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf.gentle-tempest"
	},
	{
		"name": "famous-rescue",
		"org_id": "38b9879b-f73b-4b0e-b9d9-4fc4c23643a7",
		"paths": "creative-scalphunter.clear-arclight.topical-micromax.famous-rescue"
	},
	{
		"name": "crucial-mister-sinister",
		"org_id": "38b9879b-f73b-4b0e-b9d9-4fc4c23643a7",
		"paths": "creative-scalphunter.clear-arclight.topical-micromax.famous-rescue.crucial-mister-sinister"
	},
	{
		"name": "flexible-iron-man",
		"org_id": "38b9879b-f73b-4b0e-b9d9-4fc4c23643a7",
		"paths": "creative-scalphunter.clear-arclight.topical-micromax.famous-rescue.flexible-iron-man"
	}]`)},
		{"not-existing-organization", uuid.Must(uuid.FromString("9b4cdb0a-cfea-4f9d-8a68-24f038fae386")), folder.GetAllFolders(), "Error: Organization does not exist", []byte(`[]`)},
		{"not-existing-folder", uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")), folder.GetAllFolders(), "Error: Folder does not exist", []byte(`[]`)},
		{"bursting-lionheart", uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")), folder.GetAllFolders(), "Error: Folder does not exist in the specified organization", []byte(`[]`)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil && err != tt.expectedError {
					t.Errorf("Error %s not match for Organization %s Name %s", err, tt.orgID, tt.name)
				}
			}()

			f := folder.NewDriver(tt.folders)
			get := f.GetAllChildFolders(tt.orgID, tt.name)
			wantFolderMap := make(map[folder.Folder]bool)
			wanted := []folder.Folder{}
			json.Unmarshal(tt.want, &wanted)
			if tt.expectedError != "" {
				panic("Error not detected")
			}
			for _, curFolder := range wanted {
				wantFolderMap[curFolder] = true
			}
			for _, curFolder := range get {
				if !wantFolderMap[curFolder] {
					t.Errorf("Test result not match for Organization %s ", tt.orgID)
					break
				} else {
					delete(wantFolderMap, curFolder)
				}
			}
		})
	}
}
