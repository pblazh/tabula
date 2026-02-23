import { getIndexLetter, parseCSV, renderTableHTML } from './renderTable'
import { TabulaSettings } from './types'

describe('render table', () => {
  describe('getIndexLetter', () => {
    test.each([
      ['-1', [-1, '']],
      ['0', [0, '']],
      ['1', [1, 'A']],
      ['22', [27, 'AA']],
      ['32', [28, 'AB']],
    ] as Array<[string, [number, string]]>)(
      '%s',
      (message, [input, output]) => {
        const letter = getIndexLetter(input)
        expect(letter).toBe(output)
      },
    )
  })

  describe('parse CSV', () => {
    test.each([
      // ['empty', ['0', [['']]]],
      ['once cell', ['0', [['0']]]],
      ['one row', ['1,2,3', [['1', '2', '3']]]],
      ['one column', ['1\n2\n3\n', [['1'], ['2'], ['3']]]],
      [
        'matrix',
        [
          '1,a\n2,b\n3,c\n',
          [
            ['1', 'a'],
            ['2', 'b'],
            ['3', 'c'],
          ],
        ],
      ],
      [
        'empty fields',
        ['1,a\n2\n3,c\n,', [['1', 'a'], ['2'], ['3', 'c'], ['', '']]],
      ],
    ] as Array<[string, [string, string[][]]]>)(
      'parseCSV -> %s',
      async (_, [input, expected]) => {
        const output = await parseCSV(input)
        expect(output).toStrictEqual(expected)
      },
    )
  })

  describe('renderTable', () => {
    test.each([
      [
        'matrix',
        [
          true,
          '1,a\n2,b\n3,c\n',
          `
<table class="tabula-csv-table tabula-csv-table--guide">
  <tbody>
    <tr class="tabula-csv-table-guide-row">
      <td class="tabula-csv-table-guide-cell"></td>
      <td>A</td>
      <td>B</td>
    <tr>
    <tr>
      <td class="tabula-csv-table-guide-cell">1</td>
      <td>1</td>
      <td>a</td>
    <tr>
    <tr>
      <td class="tabula-csv-table-guide-cell">2</td>
      <td>2</td>
      <td>b</td>
    <tr>
    <tr>
      <td class="tabula-csv-table-guide-cell">3</td>
      <td>3</td>
      <td>c</td>
    <tr>
  </tbody>
</table>`,
        ],
      ],
      [
        'matrix hidden index',
        [
          false,
          '1,a\n2,b\n3,c\n',
          `
<table class="tabula-csv-table">
  <tbody>
    <tr>
      <td>1</td>
      <td>a</td>
    <tr>
    <tr>
      <td>2</td>
      <td>b</td>
    <tr>
    <tr>
      <td>3</td>
      <td>c</td>
    <tr>
  </tbody>
</table>`,
        ],
      ],

      ['empty', ['', '', '']],
    ] as Array<[string, [boolean, string, string]]>)(
      'parseCSV %s',
      async (_, [tableIndex, input, expected]) => {
        const settings: TabulaSettings = {
          autoExecution: false,
          executablePath: '',
          autoFormat: false,
          tableIndex,
        }

        const output = await renderTableHTML(settings, input)
        expect(output).toStrictEqual(expected)
      },
    )
  })
})
