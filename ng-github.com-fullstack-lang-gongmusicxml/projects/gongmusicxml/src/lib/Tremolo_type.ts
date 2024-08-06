// generated from ng_file_enum.ts.go
export enum Tremolo_type {
	// insertion point	
	Tremolo_typeStart = "start",
	Tremolo_typeStop = "stop",
	Tremolo_typeSingle = "single",
	Tremolo_typeUnmeasured = "unmeasured",
}

export interface Tremolo_typeSelect {
	value: string;
	viewValue: string;
}

export const Tremolo_typeList: Tremolo_typeSelect[] = [ // insertion point	
	{ value: Tremolo_type.Tremolo_typeStart, viewValue: "start" },
	{ value: Tremolo_type.Tremolo_typeStop, viewValue: "stop" },
	{ value: Tremolo_type.Tremolo_typeSingle, viewValue: "single" },
	{ value: Tremolo_type.Tremolo_typeUnmeasured, viewValue: "unmeasured" },
];
