// generated from ng_file_enum.ts.go
export enum Note_size_type {
	// insertion point	
	Note_size_typeCue = "cue",
	Note_size_typeGrace = "grace",
	Note_size_typeGrace_cue = "grace-cue",
	Note_size_typeLarge = "large",
}

export interface Note_size_typeSelect {
	value: string;
	viewValue: string;
}

export const Note_size_typeList: Note_size_typeSelect[] = [ // insertion point	
	{ value: Note_size_type.Note_size_typeCue, viewValue: "cue" },
	{ value: Note_size_type.Note_size_typeGrace, viewValue: "grace" },
	{ value: Note_size_type.Note_size_typeGrace_cue, viewValue: "grace-cue" },
	{ value: Note_size_type.Note_size_typeLarge, viewValue: "large" },
];
