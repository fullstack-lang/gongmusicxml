// generated from ng_file_enum.ts.go
export enum Staff_type {
	// insertion point	
	Staff_typeOssia = "ossia",
	Staff_typeEditorial = "editorial",
	Staff_typeCue = "cue",
	Staff_typeAlternate = "alternate",
	Staff_typeRegular = "regular",
}

export interface Staff_typeSelect {
	value: string;
	viewValue: string;
}

export const Staff_typeList: Staff_typeSelect[] = [ // insertion point	
	{ value: Staff_type.Staff_typeOssia, viewValue: "ossia" },
	{ value: Staff_type.Staff_typeEditorial, viewValue: "editorial" },
	{ value: Staff_type.Staff_typeCue, viewValue: "cue" },
	{ value: Staff_type.Staff_typeAlternate, viewValue: "alternate" },
	{ value: Staff_type.Staff_typeRegular, viewValue: "regular" },
];
