// generated from ng_file_enum.ts.go
export enum Line_length {
	// insertion point	
	Line_lengthShort = "short",
	Line_lengthMedium = "medium",
	Line_lengthLong = "long",
}

export interface Line_lengthSelect {
	value: string;
	viewValue: string;
}

export const Line_lengthList: Line_lengthSelect[] = [ // insertion point	
	{ value: Line_length.Line_lengthShort, viewValue: "short" },
	{ value: Line_length.Line_lengthMedium, viewValue: "medium" },
	{ value: Line_length.Line_lengthLong, viewValue: "long" },
];
