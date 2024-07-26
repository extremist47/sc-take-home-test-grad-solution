package folders

import (
	"github.com/gofrs/uuid"
)

//THIS IS IMPORVE VERSION OF THE CODE

// this func gets all the folders for a given orgination id
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {

	//REMOVED UNUSED VARIABLES
	//IMPROVED CODE USING POINTERS
	//USED BETTER VARIABLE NAMES

	// get all folders by orgination id
	folders, err := FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err
	}
	//create reponse using the fetched olders
	response := &FetchFolderResponse{Folders: folders}
	return response, nil
}

// this func gets the folders that belong to a particular orgination id
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {

	//getting smaple data
	folders := GetSampleData()

	// filtring the folders by orgination id
	var resFolders []*Folder
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolders = append(resFolders, folder)
		}
	}
	return resFolders, nil
}
