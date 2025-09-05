import { Settings } from "../types";

export function validateSettings(settings: Settings): string[] {
  const errors: string[] = [];

  if (!settings.delimiter) {
    errors.push("Delimiter cannot be empty");
  }

  if (settings.delimiter === settings.quote) {
    errors.push("Delimiter and quote character cannot be the same");
  }

  if (settings.delimiter === settings.comment) {
    errors.push("Delimiter and comment prefix cannot be the same");
  }

  return errors;
}
