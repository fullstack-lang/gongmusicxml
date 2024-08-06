// generated from ng_file_enum.ts.go
export enum Syllabic {
	// insertion point	
	SyllabicSingle = "single",
	SyllabicBegin = "begin",
	SyllabicEnd = "end",
	SyllabicMiddle = "middle",
}

export interface SyllabicSelect {
	value: string;
	viewValue: string;
}

export const SyllabicList: SyllabicSelect[] = [ // insertion point	
	{ value: Syllabic.SyllabicSingle, viewValue: "single" },
	{ value: Syllabic.SyllabicBegin, viewValue: "begin" },
	{ value: Syllabic.SyllabicEnd, viewValue: "end" },
	{ value: Syllabic.SyllabicMiddle, viewValue: "middle" },
];
