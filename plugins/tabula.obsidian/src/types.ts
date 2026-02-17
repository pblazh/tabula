export type TabulaSettings = {
  autoExecution: boolean
  executablePath: string
  autoFormat: boolean
  tableIndex: boolean
}

export type Chunk = {
  type: 'text' | 'csv' | 'code' | 'error'
  content: string
}

export type Match = Chunk & {
  start: number
  end: number
}

export const DEFAULT_SETTINGS: TabulaSettings = {
  autoExecution: true,
  executablePath: 'tabula',
  autoFormat: true,
  tableIndex: true,
}
