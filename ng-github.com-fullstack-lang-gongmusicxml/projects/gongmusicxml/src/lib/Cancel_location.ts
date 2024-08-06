// generated from ng_file_enum.ts.go
export enum Cancel_location {
	// insertion point	
	Cancel_locationLeft = "left",
	Cancel_locationRight = "right",
	Cancel_locationBefore_barline = "before-barline",
}

export interface Cancel_locationSelect {
	value: string;
	viewValue: string;
}

export const Cancel_locationList: Cancel_locationSelect[] = [ // insertion point	
	{ value: Cancel_location.Cancel_locationLeft, viewValue: "left" },
	{ value: Cancel_location.Cancel_locationRight, viewValue: "right" },
	{ value: Cancel_location.Cancel_locationBefore_barline, viewValue: "before-barline" },
];
