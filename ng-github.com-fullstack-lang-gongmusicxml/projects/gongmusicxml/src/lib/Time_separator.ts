// generated from ng_file_enum.ts.go
export enum Time_separator {
	// insertion point	
	Time_separatorNone = "none",
	Time_separatorHorizontal = "horizontal",
	Time_separatorDiagonal = "diagonal",
	Time_separatorVertical = "vertical",
	Time_separatorAdjacent = "adjacent",
}

export interface Time_separatorSelect {
	value: string;
	viewValue: string;
}

export const Time_separatorList: Time_separatorSelect[] = [ // insertion point	
	{ value: Time_separator.Time_separatorNone, viewValue: "none" },
	{ value: Time_separator.Time_separatorHorizontal, viewValue: "horizontal" },
	{ value: Time_separator.Time_separatorDiagonal, viewValue: "diagonal" },
	{ value: Time_separator.Time_separatorVertical, viewValue: "vertical" },
	{ value: Time_separator.Time_separatorAdjacent, viewValue: "adjacent" },
];
