// generated from ng_file_enum.ts.go
export enum Numeral_mode {
	// insertion point	
	Numeral_modeMajor = "major",
	Numeral_modeMinor = "minor",
	Numeral_modeNaturalminor = "natural minor",
	Numeral_modeMelodicminor = "melodic minor",
	Numeral_modeHarmonicminor = "harmonic minor",
}

export interface Numeral_modeSelect {
	value: string;
	viewValue: string;
}

export const Numeral_modeList: Numeral_modeSelect[] = [ // insertion point	
	{ value: Numeral_mode.Numeral_modeMajor, viewValue: "major" },
	{ value: Numeral_mode.Numeral_modeMinor, viewValue: "minor" },
	{ value: Numeral_mode.Numeral_modeNaturalminor, viewValue: "natural minor" },
	{ value: Numeral_mode.Numeral_modeMelodicminor, viewValue: "melodic minor" },
	{ value: Numeral_mode.Numeral_modeHarmonicminor, viewValue: "harmonic minor" },
];
