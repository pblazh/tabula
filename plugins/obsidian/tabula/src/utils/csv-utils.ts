import { Settings, TableComments, TableData } from "../types";
import { parseCSVContent, unparseCSVContent } from "../services/csv-service";

/**
 * @deprecated Use parseCSVContent from csv-service instead
 */
export const parseCSV = (
  settings: Settings,
  csvString: string,
): { data: TableData; comments: TableComments } => {
  const result = parseCSVContent(settings, csvString);
  return { data: result.data, comments: result.comments };
};

/**
 * @deprecated Use unparseCSVContent from csv-service instead
 */
export const unparseCSV = (
  settings: Settings,
  data: TableData,
  comments: TableComments,
): string => {
  return unparseCSVContent(settings, data, comments);
};
