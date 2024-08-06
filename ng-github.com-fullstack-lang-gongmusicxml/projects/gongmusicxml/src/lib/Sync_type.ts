// generated from ng_file_enum.ts.go
export enum Sync_type {
	// insertion point	
	Sync_typeNone = "none",
	Sync_typeTempo = "tempo",
	Sync_typeMostly_tempo = "mostly-tempo",
	Sync_typeMostly_event = "mostly-event",
	Sync_typeEvent = "event",
	Sync_typeAlways_event = "always-event",
}

export interface Sync_typeSelect {
	value: string;
	viewValue: string;
}

export const Sync_typeList: Sync_typeSelect[] = [ // insertion point	
	{ value: Sync_type.Sync_typeNone, viewValue: "none" },
	{ value: Sync_type.Sync_typeTempo, viewValue: "tempo" },
	{ value: Sync_type.Sync_typeMostly_tempo, viewValue: "mostly-tempo" },
	{ value: Sync_type.Sync_typeMostly_event, viewValue: "mostly-event" },
	{ value: Sync_type.Sync_typeEvent, viewValue: "event" },
	{ value: Sync_type.Sync_typeAlways_event, viewValue: "always-event" },
];
