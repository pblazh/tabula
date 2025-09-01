export type Settings = {
  tabula: string;
  delimiter: string;
  quote: string;
  comment: string;
};

export type TableData = string[][];

export type TableComments = Record<number, string>;
