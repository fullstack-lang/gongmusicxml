// generated from ng_file_enum.ts.go
export enum Wedge_type {
	// insertion point	
	Wedge_typeCrescendo = "crescendo",
	Wedge_typeDiminuendo = "diminuendo",
	Wedge_typeStop = "stop",
	Wedge_typeContinue_ = "continue",
}

export interface Wedge_typeSelect {
	value: string;
	viewValue: string;
}

export const Wedge_typeList: Wedge_typeSelect[] = [ // insertion point	
	{ value: Wedge_type.Wedge_typeCrescendo, viewValue: "crescendo" },
	{ value: Wedge_type.Wedge_typeDiminuendo, viewValue: "diminuendo" },
	{ value: Wedge_type.Wedge_typeStop, viewValue: "stop" },
	{ value: Wedge_type.Wedge_typeContinue_, viewValue: "continue" },
];
