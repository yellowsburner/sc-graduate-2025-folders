package folders

import "github.com/gofrs/uuid"

type FetchFolderRequest_OLD struct {
	OrgID uuid.UUID
}

type FetchFolderResponse_OLD struct {
	Folders []*Folder
}

type FetchFolderRequest struct {
	OrgID uuid.UUID
	Token string
}


type FetchFolderResponse struct {
	Folders []*Folder
	Token string
}