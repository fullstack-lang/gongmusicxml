// generated from ng_file_enum.ts.go
export enum Line_end {
	// insertion point	
	Line_endUp = "up",
	Line_endDown = "down",
	Line_endBoth = "both",
	Line_endArrow = "arrow",
	Line_endNone = "none",
}

export interface Line_endSelect {
	value: string;
	viewValue: string;
}

export const Line_endList: Line_endSelect[] = [ // insertion point	
	{ value: Line_end.Line_endUp, viewValue: "up" },
	{ value: Line_end.Line_endDown, viewValue: "down" },
	{ value: Line_end.Line_endBoth, viewValue: "both" },
	{ value: Line_end.Line_endArrow, viewValue: "arrow" },
	{ value: Line_end.Line_endNone, viewValue: "none" },
];
