export const enumValueToDisplayName = (enumValue) =>
	(enumValue[0].toUpperCase() + enumValue.substr(1).toLowerCase()).replace("_", " ");
