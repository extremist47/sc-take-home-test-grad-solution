package folders_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetAllFolders(t *testing.T) {
	t.Run("Fetch all folders for a given organization ID", func(t *testing.T) {
		// create a request with the organization ID
		req := &folders.FetchFolderRequest{
			OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
		}

		// call func GetAllFolders and check for errors
		res, err := folders.GetAllFolders(req)
		//checks no errors and response isnt nil
		assert.NoError(t, err)
		assert.NotNil(t, res)

		//checks if folders are extracted
		assert.Greater(t, len(res.Folders), 0, "Expected to fetch some folders")
	})
}
