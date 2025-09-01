import { parse, unparse } from "papaparse";
import { Settings, TableComments, TableData } from "../types";

export const parseCSV = (
  settings: Settings,
  csvString: string,
): { data: TableData; comments: TableComments } => {
  const config = {
    header: false,
    dynamicTyping: false,
    delimiter: ",",
  };

  const comments: TableComments = {};
  const data: TableData = [];

  const lines = csvString.split("\n");

  let lineNumber = 0;
  for (let line of lines) {
    if (line.startsWith(settings.comment)) {
      comments[lineNumber] = line;
    } else {
      const parsed = parse<string[]>(line, config);
      if (parsed.data.length === 0) continue;
      data.push(parsed.data[0].map((str) => str.trim()));
    }
    lineNumber++;
  }

  return { data, comments };
};

export const unparseCSV = (
  settings: Settings,
  data: TableData,
  comments: TableComments,
): string => {
  const lines: string[] = [];
  let lineNumber = 0;
  for (let row of data) {
    if (comments[lines.length]) {
      lines.push(comments[lineNumber]);
    }
    const line = unparse([row], settings);
    lines.push(line);
  }
  const rest = Object.entries(comments).filter(
    ([k]) => parseInt(k) >= lines.length,
  );

  rest.sort(([k], [j]) => parseInt(k) - parseInt(j));
  for (let [, comment] of rest) {
    lines.push(comment);
  }
  return lines.join("\n");
};
