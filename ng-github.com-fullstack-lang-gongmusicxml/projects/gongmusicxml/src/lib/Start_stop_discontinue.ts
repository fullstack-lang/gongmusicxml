// generated from ng_file_enum.ts.go
export enum Start_stop_discontinue {
	// insertion point	
	Start_stop_discontinueStart = "start",
	Start_stop_discontinueStop = "stop",
	Start_stop_discontinueDiscontinue = "discontinue",
}

export interface Start_stop_discontinueSelect {
	value: string;
	viewValue: string;
}

export const Start_stop_discontinueList: Start_stop_discontinueSelect[] = [ // insertion point	
	{ value: Start_stop_discontinue.Start_stop_discontinueStart, viewValue: "start" },
	{ value: Start_stop_discontinue.Start_stop_discontinueStop, viewValue: "stop" },
	{ value: Start_stop_discontinue.Start_stop_discontinueDiscontinue, viewValue: "discontinue" },
];
