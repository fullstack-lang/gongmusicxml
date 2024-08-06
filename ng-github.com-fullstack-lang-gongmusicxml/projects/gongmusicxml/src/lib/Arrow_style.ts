// generated from ng_file_enum.ts.go
export enum Arrow_style {
	// insertion point	
	Arrow_styleSingle = "single",
	Arrow_styleDouble = "double",
	Arrow_styleFilled = "filled",
	Arrow_styleHollow = "hollow",
	Arrow_stylePaired = "paired",
	Arrow_styleCombined = "combined",
	Arrow_styleOther = "other",
}

export interface Arrow_styleSelect {
	value: string;
	viewValue: string;
}

export const Arrow_styleList: Arrow_styleSelect[] = [ // insertion point	
	{ value: Arrow_style.Arrow_styleSingle, viewValue: "single" },
	{ value: Arrow_style.Arrow_styleDouble, viewValue: "double" },
	{ value: Arrow_style.Arrow_styleFilled, viewValue: "filled" },
	{ value: Arrow_style.Arrow_styleHollow, viewValue: "hollow" },
	{ value: Arrow_style.Arrow_stylePaired, viewValue: "paired" },
	{ value: Arrow_style.Arrow_styleCombined, viewValue: "combined" },
	{ value: Arrow_style.Arrow_styleOther, viewValue: "other" },
];
