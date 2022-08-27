export const enumValueToDisplayName = (enumValue: string) =>
  (enumValue[0].toUpperCase() + enumValue.substr(1).toLowerCase()).replace(
    "_",
    " "
  );
