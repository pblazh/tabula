import { parse, unparse } from "papaparse";
import { Settings, TableComments, TableData } from "../types";

export interface ParseResult {
  data: TableData;
  comments: TableComments;
  errors?: string[];
}

/**
 * Parse CSV string into table data and comments
 */
export function parseCSVContent(settings: Settings, csvString: string): ParseResult {
  if (!csvString || typeof csvString !== 'string') {
    return { data: [[""]], comments: {}, errors: ["Invalid CSV content"] };
  }

  const config = {
    header: false,
    dynamicTyping: false,
    delimiter: settings.delimiter || ",",
    skipEmptyLines: false,
  };

  const comments: TableComments = {};
  const data: TableData = [];
  const errors: string[] = [];

  try {
    const lines = csvString.split("\n");
    let lineNumber = 0;

    for (const line of lines) {
      if (line.startsWith(settings.comment || "#")) {
        comments[lineNumber] = line;
      } else if (line.trim()) {
        try {
          const parsed = parse<string[]>(line, config);
          if (parsed.errors && parsed.errors.length > 0) {
            errors.push(...parsed.errors.map(err => `Line ${lineNumber + 1}: ${err.message}`));
          }
          if (parsed.data && parsed.data.length > 0) {
            data.push(parsed.data[0].map((str) => String(str).trim()));
          }
        } catch (error) {
          errors.push(`Line ${lineNumber + 1}: Failed to parse - ${error instanceof Error ? error.message : 'Unknown error'}`);
        }
      }
      lineNumber++;
    }

    return { data: data.length > 0 ? data : [[""]], comments, errors: errors.length > 0 ? errors : undefined };
  } catch (error) {
    return { 
      data: [[""]], 
      comments: {}, 
      errors: [`Parse error: ${error instanceof Error ? error.message : 'Unknown error'}`] 
    };
  }
}

/**
 * Convert table data and comments back to CSV string
 */
export function unparseCSVContent(settings: Settings, data: TableData, comments: TableComments): string {
  if (!data || !Array.isArray(data)) {
    return "";
  }

  try {
    const lines: string[] = [];
    let lineNumber = 0;

    for (const row of data) {
      if (comments[lineNumber]) {
        lines.push(comments[lineNumber]);
      }
      
      if (Array.isArray(row)) {
        const line = unparse([row], {
          delimiter: settings.delimiter || ",",
          quoteChar: settings.quote || '"',
          header: false,
        });
        lines.push(line);
      }
      lineNumber++;
    }

    // Add any remaining comments
    const remainingComments = Object.entries(comments).filter(
      ([k]) => parseInt(k) >= lineNumber,
    );

    remainingComments.sort(([a], [b]) => parseInt(a) - parseInt(b));
    for (const [, comment] of remainingComments) {
      lines.push(comment);
    }

    return lines.join("\n");
  } catch (error) {
    throw new Error(`Failed to unparse CSV: ${error instanceof Error ? error.message : 'Unknown error'}`);
  }
}

/**
 * Validate CSV settings
 */
export function validateCSVSettings(settings: Settings): string[] {
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