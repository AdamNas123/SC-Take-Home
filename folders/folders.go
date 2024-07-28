package folders

import "github.com/gofrs/uuid"

// The GetAllFolders function returns an array of all the folders that were found with the provided UUID (Provided in the FetchFolderResponse struct)
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	//Defines 3 variables that are not currently used.
	//Improvement: Remove the f1 variable as it is not required and utilise the other 2 variables in the code.
	var (
		err error
		// f1  Folder
		fs []*Folder
	)
	//Creates an empty array of folders. Improvement: Remove this line and just use fs that was previously defined
	// f := []Folder{}
	//Calls the fetch all folders function, returning all folders with the given org id. Improvement: Use the error variable in case any errors are returned
	r, err := FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err
	}
	//Loops through and appends each returned folder into the fs list.
	//Improvement: Do not need index, so k can be changed to _. Can also replace entire loop with below line to append all the values of r to the fs list.
	// for _, v := range r {
	// 	fs = append(fs, v)
	// }
	fs = append(fs, r...)
	//Improvement: Below code is also not required, as it is just creating another list. So, can be removed.
	// var fp []*Folder
	// for k1, v1 := range f {
	// 	fp = append(fp, &v1)
	// }

	//Creates a FetchFolderResponse variable so that the requested list of folders can be returned in the correct format.
	//Improvement: Can remove declaration line and merge with assignment line.
	// var ffr *FetchFolderResponse
	ffr := &FetchFolderResponse{Folders: fs}
	return ffr, nil
}

func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
