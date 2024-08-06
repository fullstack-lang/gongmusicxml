// generated from ng_file_enum.ts.go
export enum System_relation_number {
	// insertion point	
	System_relation_numberOnly_top = "only-top",
	System_relation_numberOnly_bottom = "only-bottom",
	System_relation_numberAlso_top = "also-top",
	System_relation_numberAlso_bottom = "also-bottom",
	System_relation_numberNone = "none",
}

export interface System_relation_numberSelect {
	value: string;
	viewValue: string;
}

export const System_relation_numberList: System_relation_numberSelect[] = [ // insertion point	
	{ value: System_relation_number.System_relation_numberOnly_top, viewValue: "only-top" },
	{ value: System_relation_number.System_relation_numberOnly_bottom, viewValue: "only-bottom" },
	{ value: System_relation_number.System_relation_numberAlso_top, viewValue: "also-top" },
	{ value: System_relation_number.System_relation_numberAlso_bottom, viewValue: "also-bottom" },
	{ value: System_relation_number.System_relation_numberNone, viewValue: "none" },
];
