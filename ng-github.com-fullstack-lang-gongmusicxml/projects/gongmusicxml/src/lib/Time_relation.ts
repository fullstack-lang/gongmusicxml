// generated from ng_file_enum.ts.go
export enum Time_relation {
	// insertion point	
	Time_relationParentheses = "parentheses",
	Time_relationBracket = "bracket",
	Time_relationEquals = "equals",
	Time_relationSlash = "slash",
	Time_relationSpace = "space",
	Time_relationHyphen = "hyphen",
}

export interface Time_relationSelect {
	value: string;
	viewValue: string;
}

export const Time_relationList: Time_relationSelect[] = [ // insertion point	
	{ value: Time_relation.Time_relationParentheses, viewValue: "parentheses" },
	{ value: Time_relation.Time_relationBracket, viewValue: "bracket" },
	{ value: Time_relation.Time_relationEquals, viewValue: "equals" },
	{ value: Time_relation.Time_relationSlash, viewValue: "slash" },
	{ value: Time_relation.Time_relationSpace, viewValue: "space" },
	{ value: Time_relation.Time_relationHyphen, viewValue: "hyphen" },
];
