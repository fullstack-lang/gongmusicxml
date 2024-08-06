// generated from ng_file_enum.ts.go
export enum Winged {
	// insertion point	
	WingedNone = "none",
	WingedStraight = "straight",
	WingedCurved = "curved",
	WingedDouble_straight = "double-straight",
	WingedDouble_curved = "double-curved",
}

export interface WingedSelect {
	value: string;
	viewValue: string;
}

export const WingedList: WingedSelect[] = [ // insertion point	
	{ value: Winged.WingedNone, viewValue: "none" },
	{ value: Winged.WingedStraight, viewValue: "straight" },
	{ value: Winged.WingedCurved, viewValue: "curved" },
	{ value: Winged.WingedDouble_straight, viewValue: "double-straight" },
	{ value: Winged.WingedDouble_curved, viewValue: "double-curved" },
];
