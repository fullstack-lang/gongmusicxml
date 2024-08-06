// generated from ng_file_enum.ts.go
export enum Pedal_type {
	// insertion point	
	Pedal_typeStart = "start",
	Pedal_typeStop = "stop",
	Pedal_typeSostenuto = "sostenuto",
	Pedal_typeChange = "change",
	Pedal_typeContinue_ = "continue",
	Pedal_typeDiscontinue = "discontinue",
	Pedal_typeResume = "resume",
}

export interface Pedal_typeSelect {
	value: string;
	viewValue: string;
}

export const Pedal_typeList: Pedal_typeSelect[] = [ // insertion point	
	{ value: Pedal_type.Pedal_typeStart, viewValue: "start" },
	{ value: Pedal_type.Pedal_typeStop, viewValue: "stop" },
	{ value: Pedal_type.Pedal_typeSostenuto, viewValue: "sostenuto" },
	{ value: Pedal_type.Pedal_typeChange, viewValue: "change" },
	{ value: Pedal_type.Pedal_typeContinue_, viewValue: "continue" },
	{ value: Pedal_type.Pedal_typeDiscontinue, viewValue: "discontinue" },
	{ value: Pedal_type.Pedal_typeResume, viewValue: "resume" },
];
