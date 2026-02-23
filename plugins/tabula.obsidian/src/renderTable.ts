import { TabulaSettings } from './types'
import * as csv from '@fast-csv/parse'

export async function renderTable(
  settings: TabulaSettings,
  source: string,
  el: HTMLElement,
) {
  const csv = await parseCSV(source)

  const table = el.createEl('table')
  table.className = 'csv-table'
  const body = table.createEl('tbody')
  body.className = 'csv-tbody'

  const row = body.createEl('tr')
  row.className = 'csv-tr'

  const columns = csv[0].length

  for (let j = 0; j <= columns; j++) {
    const td = row.createEl('td', {
      text: j === 0 ? '' : getIndexLetter(j),
    })
    td.className = `csv-td--a ${settings.tableIndex ? '' : 'disabled'}`
  }

  for (let i = 0; i < csv.length; i++) {
    const cols = csv[i]

    const row = body.createEl('tr')
    row.className = 'csv-tr'

    const td = row.createEl('td', { text: String(i + 1) })
    td.className = `csv-td--n ${settings.tableIndex ? '' : 'disabled'}`

    for (let j = 0; j < cols.length; j++) {
      const td = row.createEl('td', { text: cols[j] })
      td.className = 'csv-td'
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
