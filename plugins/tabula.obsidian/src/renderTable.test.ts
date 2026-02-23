import { getIndexLetter } from './renderTable'

describe('render table', () => {
  test.each([
    ['-1', [-1, '']],
    ['0', [0, '']],
    ['1', [1, 'A']],
    ['22', [27, 'AA']],
    ['32', [28, 'AB']],
  ] as Array<[string, [number, string]]>)(
    'getIndexLetter %s',
    (message, [input, output]) => {
      const letter = getIndexLetter(input)
      expect(letter).toBe(output)
    },
  )
})
