// generated from ng_file_enum.ts.go
export enum Fan {
	// insertion point	
	FanAccel = "accel",
	FanRit = "rit",
	FanNone = "none",
}

export interface FanSelect {
	value: string;
	viewValue: string;
}

export const FanList: FanSelect[] = [ // insertion point	
	{ value: Fan.FanAccel, viewValue: "accel" },
	{ value: Fan.FanRit, viewValue: "rit" },
	{ value: Fan.FanNone, viewValue: "none" },
];
