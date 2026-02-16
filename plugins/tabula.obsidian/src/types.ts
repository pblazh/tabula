export type TabulaSettings = {
  autoExecute: boolean;
  executablePath: string;
  autoFormat: boolean;
};

export type Chunk = {
  type: "text" | "csv" | "code" | "error";
  content: string;
};

export type Match = Chunk & {
  start: number;
  end: number;
};

export const DEFAULT_SETTINGS: TabulaSettings = {
  autoExecute: true,
  executablePath: "tabula",
  autoFormat: true,
};
