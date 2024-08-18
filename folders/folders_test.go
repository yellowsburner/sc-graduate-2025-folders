package folders_test

import (
	"fmt"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllFolders_OLD(t *testing.T) {
	testID := func(orgID string, requiredCount int) {
		req := &folders.FetchFolderRequest_OLD{
			OrgID: uuid.FromStringOrNil(orgID),
		}
		res, err := folders.GetAllFolders_OLD(req)
		fmt.Println(err)
		assert.Nil(t, err)
		assert.Equal(t, len(res.Folders), requiredCount)
	}
	t.Run("Existent Organisation ID's", func(t *testing.T) {
		testID("c1556e17-b7c0-45a3-a6ae-9546248fb17a", 666)
		testID("5652b680-8d7c-49cb-b021-9b9803683504", 1)
		testID("4212d618-66ff-468a-862d-ea49fef5e183", 1)
		testID("6c63ad89-cee6-4ce5-a32a-52b5ca179c86", 1)
	})
	t.Run("Non-existent organisation ID's", func(t *testing.T) {
		testID("123", 0)
		testID("c1556e17-b7c0-45a3-a6ae-9546248fb17b", 0)
		testID("c1556e17-b7c0-45a3-a6ae-9546248fb17a ", 0)
		testID(" c1556e17-b7c0-45a3-a6ae-9546248fb17a", 0)
		testID("0", 0)
		testID("", 0)
	})
}
