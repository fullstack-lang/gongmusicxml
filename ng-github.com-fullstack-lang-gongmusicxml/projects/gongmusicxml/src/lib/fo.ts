// generated code - do not edit

import { FoAPI } from './fo-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Fo {

	static GONGSTRUCT_NAME = "Fo"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyFoToFoAPI(fo: Fo, foAPI: FoAPI) {

	foAPI.CreatedAt = fo.CreatedAt
	foAPI.DeletedAt = fo.DeletedAt
	foAPI.ID = fo.ID

	// insertion point for basic fields copy operations
	foAPI.Name = fo.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyFoAPIToFo update basic, pointers and slice of pointers fields of fo
// from respectively the basic fields and encoded fields of pointers and slices of pointers of foAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFoAPIToFo(foAPI: FoAPI, fo: Fo, frontRepo: FrontRepo) {

	fo.CreatedAt = foAPI.CreatedAt
	fo.DeletedAt = foAPI.DeletedAt
	fo.ID = foAPI.ID

	// insertion point for basic fields copy operations
	fo.Name = foAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
