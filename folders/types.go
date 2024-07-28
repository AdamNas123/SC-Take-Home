package folders

import "github.com/gofrs/uuid"

type FetchFolderRequest struct {
	OrgID uuid.UUID
}

type FetchFolderResponse struct {
	Folders []*Folder
}

// New struct created to also include the NextToken string, pointing to the next page of results.
type FetchFolderResponsePaginated struct {
	Folders   []*Folder
	NextToken string
}
