import { TabulaSettings } from './types'
import * as csv from '@fast-csv/parse'

export async function renderTable(
  settings: TabulaSettings,
  source: string,
  el: HTMLElement,
) {
  const csvData = await parseCSV(source)
  if (csvData.length === 0) return

  const guide = settings.tableIndex
  const headers = Array(csvData[0].length)
    .fill(null)
    .map((_, i) => getIndexLetter(i + 1))

  const table = el.createEl('table', {
    cls: `tabula-csv-table${guide ? ' tabula-csv-table--guide' : ''}`,
  })
  const tbody = table.createEl('tbody')

  if (guide) {
    const headerRow = tbody.createEl('tr', {
      cls: 'tabula-csv-table-guide-row',
    })
    headerRow.createEl('td', { cls: 'tabula-csv-table-guide-cell' })
    for (const h of headers) {
      headerRow.createEl('td', { text: h })
    }
  }

  for (let j = 0; j < csvData.length; j++) {
    const tr = tbody.createEl('tr')
    if (guide) {
      tr.createEl('td', {
        cls: 'tabula-csv-table-guide-cell',
        text: String(j + 1),
      })
    }
    for (const cell of csvData[j]) {
      tr.createEl('td', { text: cell })
    }
  }
}

export function parseCSV(source: string): Promise<string[][]> {
  const out: string[][] = []
  return new Promise<string[][]>((resolve, reject) => {
    csv
      .parseString(source)
      .on('error', (error) => reject(error))
      .on('data', (row: string[]) => {
        out.push(row)
      })
      .on('end', () => resolve(out))
  })
}

export function getIndexLetter(n: number) {
  n = Math.max(0, n)

  let result = ''

  while (n > 0) {
    n--
    result = String.fromCharCode(65 + (n % 26)) + result
    n = Math.floor(n / 26)
  }

  return result
}
