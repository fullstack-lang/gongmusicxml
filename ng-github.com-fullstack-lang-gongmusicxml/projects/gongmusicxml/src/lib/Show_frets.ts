// generated from ng_file_enum.ts.go
export enum Show_frets {
	// insertion point	
	Show_fretsNumbers = "numbers",
	Show_fretsLetters = "letters",
}

export interface Show_fretsSelect {
	value: string;
	viewValue: string;
}

export const Show_fretsList: Show_fretsSelect[] = [ // insertion point	
	{ value: Show_frets.Show_fretsNumbers, viewValue: "numbers" },
	{ value: Show_frets.Show_fretsLetters, viewValue: "letters" },
];
