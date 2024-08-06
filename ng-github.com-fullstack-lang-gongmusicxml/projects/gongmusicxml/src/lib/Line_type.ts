// generated from ng_file_enum.ts.go
export enum Line_type {
	// insertion point	
	Line_typeSolid = "solid",
	Line_typeDashed = "dashed",
	Line_typeDotted = "dotted",
	Line_typeWavy = "wavy",
}

export interface Line_typeSelect {
	value: string;
	viewValue: string;
}

export const Line_typeList: Line_typeSelect[] = [ // insertion point	
	{ value: Line_type.Line_typeSolid, viewValue: "solid" },
	{ value: Line_type.Line_typeDashed, viewValue: "dashed" },
	{ value: Line_type.Line_typeDotted, viewValue: "dotted" },
	{ value: Line_type.Line_typeWavy, viewValue: "wavy" },
];
