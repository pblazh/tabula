import { TabulaSettings } from './types'
import * as csv from '@fast-csv/parse'
import m from 'mustache'

export async function renderTable(
  settings: TabulaSettings,
  source: string,
  el: HTMLElement,
) {
  const html = await renderTableHTML(settings, source)
  el.innerHTML = html
}

const template = `
<table class="tabula-csv-table{{#guide}} tabula-csv-table--guide{{/guide}}">
  <tbody>
{{#guide}}
    <tr class="tabula-csv-table-guide-row">
      <td class="tabula-csv-table-guide-cell"></td>
{{#header}}
      <td>{{.}}</td>
{{/header}}
    <tr>
{{/guide}}
  {{#matrix}}
    <tr>
{{#guide}}
      <td class="tabula-csv-table-guide-cell">{{index}}</td>
{{/guide}}
{{#cells}}
      <td>{{value}}</td>
{{/cells}}
    <tr>
  {{/matrix}}
  </tbody>
</table>`

m.parse(template)

export async function renderTableHTML(
  settings: TabulaSettings,
  source: string,
): Promise<string> {
  const csv = await parseCSV(source)
  const guide = settings.tableIndex // ? '' : 'csv-table--hidden'

  if (csv.length === 0) {
    return ''
  }

  type Cell = {
    index: string
    value: string
  }
  type Row = { index: string; cells: Cell[] }

  const header: string[] = Array(csv[0].length)
    .fill(null)
    .map((_, i) => String(getIndexLetter(i + 1)))

  const matrix: Row[] = csv.map((row, j) => ({
    index: String(j + 1),
    cells: row.map(
      (cell) =>
        ({
          index: String(j),
          value: cell,
        }) as Cell,
    ),
  }))

  return m.render(template, { guide, header, matrix })
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
