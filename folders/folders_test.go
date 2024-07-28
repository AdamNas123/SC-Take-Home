package folders_test

import (
	"reflect"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	// "github.com/georgechieng-sc/interns-2022/folders"
	// "github.com/stretchr/testify/assert"
)

func Test_GetAllFolders(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		// your test/s here

		/*Will just extract sample data from the sample.json for testing
		  There are 666 folders in the sample.json with the below testOrgId.
		  If not using sample.json -> Could generate data similar to the GenerateData() function in static.go.
		*/

		//Test 1 - Check that the number of returned folders is 666 as expected.
		testOrgId1 := "c1556e17-b7c0-45a3-a6ae-9546248fb17a"
		req := &folders.FetchFolderRequest{
			OrgID: uuid.FromStringOrNil(testOrgId1),
		}
		res, err := folders.GetAllFolders(req)
		if err != nil {
			t.Errorf("Get All Folders failed. Error: %d", err)
		}
		numFolders := len(res.Folders)
		if numFolders != 666 {
			t.Errorf("Get All Folders failed. Expected 666 folders, got %d folders", numFolders)
		}

		//Test 2 - Use different Org ID and check that the returned folder is correct.
		testOrgId2 := "4212d618-66ff-468a-862d-ea49fef5e183"
		expectedReturn := &folders.FetchFolderResponse{}
		expectedReturn.Folders = append(expectedReturn.Folders, &folders.Folder{
			Id:      uuid.FromStringOrNil("1167c1ac-911b-4a1f-b460-a98f724c7289"),
			Name:    "heroic-bella",
			OrgId:   uuid.FromStringOrNil("4212d618-66ff-468a-862d-ea49fef5e183"),
			Deleted: true,
		})
		req = &folders.FetchFolderRequest{
			OrgID: uuid.FromStringOrNil(testOrgId2),
		}
		res, err = folders.GetAllFolders(req)
		if err != nil {
			t.Errorf("Get All Folders failed. Error: %d", err)
		}
		numFolders = len(res.Folders)
		if numFolders != 1 {
			t.Errorf("Get All Folders failed. Expected 1 folder, got %d folders", numFolders)
		}

		//Can compare 2 objects using the reflect package
		if !reflect.DeepEqual(res, expectedReturn) {
			t.Errorf("Get All Folders failed. Expected folder with name %v got folder with name %v", expectedReturn.Folders[0].Name, res.Folders[0].Name)
		}

		//Can also compare each individual field of the response to the expected return
		//If more than 1 folder was expected, this could be placed within a for loop
		if res.Folders[0].Id != expectedReturn.Folders[0].Id {
			t.Errorf("Get All Folders failed. Expected folder with id %v got folder with id %v", expectedReturn.Folders[0].Id, res.Folders[0].Id)
		}
		if res.Folders[0].Name != expectedReturn.Folders[0].Name {
			t.Errorf("Get All Folders failed. Expected folder with name %v got folder with name %v", expectedReturn.Folders[0].Name, res.Folders[0].Name)
		}
		if res.Folders[0].OrgId != expectedReturn.Folders[0].OrgId {
			t.Errorf("Get All Folders failed. Expected folder with org id %v got folder with org id %v", expectedReturn.Folders[0].OrgId, res.Folders[0].OrgId)
		}
		if res.Folders[0].Deleted != expectedReturn.Folders[0].Deleted {
			t.Errorf("Get All Folders failed. Expected folder with deleted status %v got folder with deleted status %v", expectedReturn.Folders[0].Deleted, res.Folders[0].Deleted)
		}
	})
}
