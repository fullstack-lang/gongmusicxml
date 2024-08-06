// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	FoAPIs []*FoAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {
	// insertion point for slices copies
	for _, foDB := range backRepo.BackRepoFo.Map_FoDBID_FoDB {

		var foAPI FoAPI
		foAPI.ID = foDB.ID
		foAPI.FoPointersEncoding = foDB.FoPointersEncoding
		foDB.CopyBasicFieldsToFo_WOP(&foAPI.Fo_WOP)

		backRepoData.FoAPIs = append(backRepoData.FoAPIs, &foAPI)
	}

}
