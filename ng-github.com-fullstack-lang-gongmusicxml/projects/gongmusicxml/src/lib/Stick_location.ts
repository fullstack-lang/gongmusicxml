// generated from ng_file_enum.ts.go
export enum Stick_location {
	// insertion point	
	Stick_locationCenter = "center",
	Stick_locationRim = "rim",
	Stick_locationCymbalbell = "cymbal bell",
	Stick_locationCymbaledge = "cymbal edge",
}

export interface Stick_locationSelect {
	value: string;
	viewValue: string;
}

export const Stick_locationList: Stick_locationSelect[] = [ // insertion point	
	{ value: Stick_location.Stick_locationCenter, viewValue: "center" },
	{ value: Stick_location.Stick_locationRim, viewValue: "rim" },
	{ value: Stick_location.Stick_locationCymbalbell, viewValue: "cymbal bell" },
	{ value: Stick_location.Stick_locationCymbaledge, viewValue: "cymbal edge" },
];
