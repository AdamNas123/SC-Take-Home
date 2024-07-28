package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
)

func main() {
	//Creates a FetchFolderRequest object by generating a UUID from the default organisation id.
	req := &folders.FetchFolderRequest{
		OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
	}

	//Calls the get all folders function to return all folders with the given Organisation UUID (provided as FetchFolderRequest type)
	res, err := folders.GetAllFolders(req)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	//Prints the returned folders in the Pretty Print format.
	folders.PrettyPrint(res)
}
