package folders

import "github.com/gofrs/uuid"

// Copy over the `GetFolders` and `FetchAllFoldersByOrgID` to get started
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	var (
		err error

		fs []*Folder
	)
	token := ""
	r, err := FetchAllFoldersByOrgID(req.OrgID, token)
	if err != nil {
		return nil, err
	}

	fs = append(fs, r...)
	ffr := &FetchFolderResponse{Folders: fs}
	return ffr, nil
}

func FetchAllFoldersByOrgID(orgID uuid.UUID, token string) ([]*Folder, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
