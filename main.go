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

	// //Calls the get all folders function to return all folders with the given Organisation UUID (provided as FetchFolderRequest type)
	// res, err := folders.GetAllFolders(req)
	// if err != nil {
	// 	fmt.Printf("%v", err)
	// 	return
	// }
	// //Prints the returned folders in the Pretty Print format.
	// folders.PrettyPrint(res)

	//-----PAGINATED COMPONENT-------
	//GetAllFoldersPaginated requires a token argument or empty string. Can replace the empty string with the NextToken provided in the output to retrieve next page of results.

	//Calls the get all folders paginated function to return a maximum of 10 folders with the given Organisation Id.
	res, err := folders.GetAllFoldersPaginated(req, "")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	//Prints the returned folders and the next token in the Pretty Print format.
	folders.PrettyPrint(res)
}
