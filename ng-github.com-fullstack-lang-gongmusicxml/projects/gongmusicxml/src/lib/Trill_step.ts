// generated from ng_file_enum.ts.go
export enum Trill_step {
	// insertion point	
	Trill_stepWhole = "whole",
	Trill_stepHalf = "half",
	Trill_stepUnison = "unison",
}

export interface Trill_stepSelect {
	value: string;
	viewValue: string;
}

export const Trill_stepList: Trill_stepSelect[] = [ // insertion point	
	{ value: Trill_step.Trill_stepWhole, viewValue: "whole" },
	{ value: Trill_step.Trill_stepHalf, viewValue: "half" },
	{ value: Trill_step.Trill_stepUnison, viewValue: "unison" },
];
