package folders

// Copy over the `GetFolders` and `FetchAllFoldersByOrgID` to get started

// this func gets all folder for a given orignaiton id with paginaiton
func GetAllFoldersWithPagination(req *FetchFolderRequest, pageSize, pageNumber int) (*FetchFolderResponse, error) {
	// get all folders by organization ID
	folders, err := FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err
	}

	// get start and end indices for pagination
	start := pageSize * (pageNumber - 1)
	if start >= len(folders) {
		return &FetchFolderResponse{Folders: []*Folder{}}, nil
	}

	end := start + pageSize
	if end > len(folders) {
		end = len(folders)
	}

	// create the response with the paginated folders
	response := &FetchFolderResponse{Folders: folders[start:end]}
	return response, nil
}
