package folders

import "github.com/gofrs/uuid"

// Copy over the `GetFolders` and `FetchAllFoldersByOrgID` to get started

// Set pagination limit as 10 records per response
const pageLimit = 10

//Set token as the ID of the folder

func GetAllFoldersPaginated(req *FetchFolderRequest, token string) (*FetchFolderResponsePaginated, error) {
	var (
		err error
		fs  []*Folder
	)
	r, nextToken, err := FetchAllFoldersByOrgIDPaginated(req.OrgID, token)
	if err != nil {
		return nil, err
	}

	fs = append(fs, r...)
	ffr := &FetchFolderResponsePaginated{Folders: fs, nextToken: nextToken}
	return ffr, nil
}

func FetchAllFoldersByOrgIDPaginated(orgID uuid.UUID, token string) ([]*Folder, string, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	//If a token is provided, find the corresponding folder and set it as the start position.
	//Otherwise, start position will be the first element in folders.
	startFolder := 0
	if token != "" {
		for pos, folder := range folders {
			if folder.Id == uuid.FromStringOrNil(token) {
				startFolder = pos
			}
		}
	}

	//Start searching from startPage and retrieve 10 pages (page limit)
	currentFolder := startFolder
	for i := startFolder; i < len(folders); i++ {
		if folders[i].OrgId == orgID {
			resFolder = append(resFolder, folders[i])
		}
		//If the page limit is hit, break out of the loop
		if len(resFolder) == pageLimit {
			currentFolder = i
			break
		}
	}

	//Need to find the token of the next folder, which will be set as the start page index for the next response.
	nextToken := ""
	for i := currentFolder; i < len(folders); i++ {
		if folders[i].OrgId == orgID {
			nextToken = folders[i].Id.String()
		}
	}
	return resFolder, nextToken, nil
}
