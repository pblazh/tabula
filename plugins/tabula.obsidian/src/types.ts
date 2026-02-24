export type TabulaSettings = {
  autoExecution: boolean
  executablePath: string
  autoFormat: boolean
  tableIndex: boolean
}

export const DEFAULT_SETTINGS: TabulaSettings = {
  autoExecution: true,
  executablePath: 'tabula',
  autoFormat: true,
  tableIndex: true,
}
