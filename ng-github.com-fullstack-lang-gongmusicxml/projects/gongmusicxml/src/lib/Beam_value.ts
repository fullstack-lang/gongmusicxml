// generated from ng_file_enum.ts.go
export enum Beam_value {
	// insertion point	
	Beam_valueBegin = "begin",
	Beam_valueContinue_ = "continue",
	Beam_valueEnd = "end",
	Beam_valueForwardhook = "forward hook",
	Beam_valueBackwardhook = "backward hook",
}

export interface Beam_valueSelect {
	value: string;
	viewValue: string;
}

export const Beam_valueList: Beam_valueSelect[] = [ // insertion point	
	{ value: Beam_value.Beam_valueBegin, viewValue: "begin" },
	{ value: Beam_value.Beam_valueContinue_, viewValue: "continue" },
	{ value: Beam_value.Beam_valueEnd, viewValue: "end" },
	{ value: Beam_value.Beam_valueForwardhook, viewValue: "forward hook" },
	{ value: Beam_value.Beam_valueBackwardhook, viewValue: "backward hook" },
];
