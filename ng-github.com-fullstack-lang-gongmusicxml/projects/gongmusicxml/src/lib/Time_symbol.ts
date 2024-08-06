// generated from ng_file_enum.ts.go
export enum Time_symbol {
	// insertion point	
	Time_symbolCommon = "common",
	Time_symbolCut = "cut",
	Time_symbolSingle_number = "single-number",
	Time_symbolNote = "note",
	Time_symbolDotted_note = "dotted-note",
	Time_symbolNormal = "normal",
}

export interface Time_symbolSelect {
	value: string;
	viewValue: string;
}

export const Time_symbolList: Time_symbolSelect[] = [ // insertion point	
	{ value: Time_symbol.Time_symbolCommon, viewValue: "common" },
	{ value: Time_symbol.Time_symbolCut, viewValue: "cut" },
	{ value: Time_symbol.Time_symbolSingle_number, viewValue: "single-number" },
	{ value: Time_symbol.Time_symbolNote, viewValue: "note" },
	{ value: Time_symbol.Time_symbolDotted_note, viewValue: "dotted-note" },
	{ value: Time_symbol.Time_symbolNormal, viewValue: "normal" },
];
