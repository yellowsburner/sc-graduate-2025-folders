package folders

import (
	"github.com/gofrs/uuid"
)

// gets all folders given parameters provided via a request object
// NOTE: renamed by suffxing OLD so conflict with folders_pagination.go file doesn't exist
func GetAllFolders_OLD(req *FetchFolderRequest_OLD) (*FetchFolderResponse_OLD, error) {
	// fix: unused variables stop Go from compiling, so remove
	/*
		var (
			err error
			f1  Folder
			fs  []*Folder
		)
	*/

	// f := []Folder{}
	// get folders in specified organisation
	// fix: just use the r variable directly (which is already a slice of folder pointers) instead of copying twice into slices
	// fix: descriptive variable names
	// suggestion: no error is being thrown and error handling doesn't actually occur, so either consider removing
	// fix: for now the error value is actually raised up
	resFolder, err := FetchAllFoldersByOrgID_OLD(req.OrgID)

	// fix: unused index variable stops Go from compiling, so use underscore
	/*
		// reference each folder pointer and append into slice
		for _, v := range r {
			f = append(f, *v)
		}

		// dereference each folder and append back pointers into a different slice
		var fp []*Folder
		for _, v1 := range f {
			fp = append(fp, &v1)
		}
	*/
	// var ffr *FetchFolderResponse

	// construct response object containing fetched folders (note a slice is provided not the entire array)
	fetchFolderResponse := &FetchFolderResponse_OLD{Folders: resFolder}
	return fetchFolderResponse, err
}

// get array of folders which has a particular argument-provided organisation ID
// NOTE: renamed by suffxing OLD so conflict with folders_pagination.go file doesn't exist
func FetchAllFoldersByOrgID_OLD(orgID uuid.UUID) ([]*Folder, error) {
	// source sample data of folders
	folders := GetSampleData()

	// loop through folders and collate those with a particular orgID
	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}

	// return the folder and report no error
	return resFolder, nil
}
