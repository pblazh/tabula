import { getIndexLetter, parseCSV } from './renderTable'

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
})
