package folders

import (
	"github.com/gofrs/uuid"
)

const PAGINATION_LIMIT = 5

// it seems that token-based pagination is preferred, where a token points to the start of the data for the next paginated part to occur
// however the current system architecture doesn't allow for this nicely as this is to be used in databases, where you can quickly get to some start point (e.g. by its primary key or date of creation)
// since we're reading an entire JSON file every time (which we could check last modified time for, but really a proper database is the best solution), we provide an intermediary solution
// a start token and constant limit will be passed into the FetchAllFoldersByOrgID function (where, in the future, the function can be changed to paginate properly)
// the start token is the ID of the starting folder - an empty strings represents the first folder. from there, we add more folders under the organisation ID specified by the limit, stopping when we reach the end
// if we reach the end, the next provided token is the empty string.

// gets all folders given parameters provided via a request object
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	// get folders in specified organisation
	resFolder, nextToken, err := FetchAllFoldersByOrgIDPaginated(req.OrgID, uuid.FromStringOrNil(req.Token), PAGINATION_LIMIT)

	// convert UUID to string, using empty string for nil UUID
	var nextTokenString string
	if !nextToken.IsNil() {
		nextTokenString = nextToken.String()
	}

	// construct response object containing fetched folders (note a slice is provided not the entire array)
	fetchFolderResponse := &FetchFolderResponse{Folders: resFolder, Token: nextTokenString}

	return fetchFolderResponse, err
}

// get array of folders which has a particular argument-provided organisation ID
func FetchAllFoldersByOrgIDPaginated(orgID uuid.UUID, start uuid.UUID, limit int) ([]*Folder, uuid.UUID, error) {
	// source sample data of folders
	folders := GetSampleData()

	// loop through folders and collate those with a particular orgID
	resFolder := []*Folder{}
	pageArrived := false
	nextToken := uuid.Nil

	for _, folder := range folders {
		// start collecting folders once particular folder ID is reached
		if start == uuid.Nil || folder.Id == start {
			pageArrived = true
		}

		// only collect folders after start token with specified orgID
		if pageArrived && folder.OrgId == orgID {
			// hard stop after <limit> folders reached AND one more folder with specified orgID is reached
			if limit == 0 {
				nextToken = folder.Id
				break
			}
			
			limit--
			resFolder = append(resFolder, folder)
		}
	}

	// return the folder and report no error
	return resFolder, nextToken, nil
}
