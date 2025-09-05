import { parse, ParseConfig, unparse } from "papaparse";
import { Settings, TableComments, TableData } from "../types";

export interface ParseResult {
  data: TableData;
  comments: TableComments;
  errors?: string[];
}

export function parseCSVContent(
  settings: Settings,
  csvString: string,
): ParseResult {
  console.log("parse", csvString, settings);
  if (!csvString || typeof csvString !== "string") {
    return { data: [], comments: {}, errors: ["Invalid CSV content"] };
  }

  const config: ParseConfig = {
    header: false,
    dynamicTyping: false,
    quoteChar: settings.quote || '"',
    delimiter: settings.delimiter || ",",
    skipEmptyLines: false,
  };

  const comments: TableComments = {};
  const data: TableData = [];
  const errors: string[] = [];

  const lines = csvString.split("\n");
  let lineNumber = 0;

  console.log("lines --->", lines);

  for (const line of lines) {
    if (line.startsWith(settings.comment || "#")) {
      comments[lineNumber] = line;
    } else if (line.trim()) {
      try {
        const parsed = parse<string[]>(line, config);
        if (parsed.errors && parsed.errors.length > 0) {
          errors.push(
            ...parsed.errors.map(
              (err) => `Line ${lineNumber + 1}: ${err.message}`,
            ),
          );
        }
        if (parsed.data && parsed.data.length > 0) {
          data.push(parsed.data[0].map((str) => String(str).trim()));
        }
      } catch (error) {
        errors.push(
          `Line ${lineNumber + 1}: Failed to parse - ${error instanceof Error ? error.message : "Unknown error"}`,
        );
      }
    }
    lineNumber++;
  }

  console.log("--->", data);

  return {
    data: data.length > 0 ? data : [],
    comments,
    errors: errors.length > 0 ? errors : undefined,
  };
}

export function unparseCSVContent(
  settings: Settings,
  data: TableData,
  comments: TableComments,
): string {
  if (!data || !Array.isArray(data)) return "";

  const lines: string[] = [];
  let lineNumber = 0;

  for (const row of data) {
    if (comments[lineNumber]) {
      lines.push(comments[lineNumber]);
    }

    if (Array.isArray(row)) {
      const line = unparse([row], {
        delimiter: settings.delimiter || ",",
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
}

