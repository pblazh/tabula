export interface Settings {
  tabula: string;
  delimiter: string;
  quote: string;
  comment: string;
}

export type TableData = string[][];

export type TableComments = Record<number, string>;

export interface ValidationError {
  field: string;
  message: string;
}

export interface OperationResult<T = void> {
  success: boolean;
  data?: T;
  error?: string;
  warnings?: string[];
}

export interface ViewState {
  isLoading: boolean;
  hasError: boolean;
  errorMessage?: string;
}

export interface TableViewData {
  data: TableData;
  comments: TableComments;
  lastModified?: number;
}
