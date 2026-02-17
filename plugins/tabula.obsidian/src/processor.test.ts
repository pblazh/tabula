import { extractChunks, outputChunks, processChunks } from './processor'
import { Executer } from './executer'
import { DataAdapter } from 'obsidian'

// Mock executer that returns "processed" for all executions
class MockExecuter extends Executer {
  constructor() {
    super(
      {
        autoExecute: false,
        executablePath: '',
        autoFormat: false,
        tableIndex: false,
      },
      null as unknown as DataAdapter,
      '',
    )
  }

  execute(_data: string, _code: string): Promise<string> {
    return Promise.resolve('processed')
  }
}

const mockExecuter = new MockExecuter()

describe('extractChunks', () => {
  describe('Empty and text-only files', () => {
    test('should handle empty file', () => {
      const chunks = extractChunks('')
      expect(chunks).toEqual([])
    })

    test('should handle text-only file', () => {
      const content = '# Heading\n\nSome text content.'
      const chunks = extractChunks(content)

      expect(chunks).toHaveLength(1)
      expect(chunks[0]).toMatchObject({
        type: 'text',
        content: content,
      })
    })
  })

  describe('CSV blocks', () => {
    test('should extract single CSV block', () => {
      const content = `
# Title

\`\`\`csv
A,B,C
1,2,3
\`\`\`

End text
`
      const chunks = extractChunks(content)

      expect(chunks).toHaveLength(3)
      expect(chunks[0].type).toBe('text')
      expect(chunks[1].type).toBe('csv')
      expect(chunks[1].content).toBe('A,B,C\n1,2,3')
      expect(chunks[2].type).toBe('text')
    })

    test('should extract multiple CSV blocks', () => {
      const content = `
\`\`\`csv
A,B
1,2
\`\`\`

Text

\`\`\`csv
X,Y
3,4
\`\`\`
`
      const chunks = extractChunks(content)

      const csvChunks = chunks.filter((c) => c.type === 'csv')
      expect(csvChunks).toHaveLength(2)
      expect(csvChunks[0].content).toBe('A,B\n1,2')
      expect(csvChunks[1].content).toBe('X,Y\n3,4')
    })
  })

  describe('Code chunks', () => {
    test('should extract code chunks', () => {
      const content = `
\`\`\`tabula
let D1 = "Total";
let D2 = A2 + B2;
\`\`\`

\`\`\`csv
A,B,C
1,2,3
\`\`\`
`
      const chunks = extractChunks(content)

      const codeChunks = chunks.filter((c) => c.type === 'code')
      expect(codeChunks).toHaveLength(1)
      expect(codeChunks[0].content).toContain('let D1')
      expect(codeChunks[0].content).toContain('let D2')
    })
  })

  describe('Mixed content', () => {
    test('should handle mixed text, code and CSV', () => {
      const content = `
# Document

\`\`\`tabula
let D1 = "Total"
\`\`\`
\`\`\`csv
A,B,C
1,2,3
\`\`\`

More text

\`\`\`tabula
let E1 = "Sum";
fmt E1 = "%.2f";
\`\`\`
\`\`\`csv
X,Y,Z
4,5,6
\`\`\`
`
      const chunks = extractChunks(content)

      expect(chunks.filter((c) => c.type === 'text').length).toBeGreaterThan(0)
      expect(chunks.filter((c) => c.type === 'csv').length).toBe(2)
      expect(chunks.filter((c) => c.type === 'code').length).toBe(2)
    })
  })

  describe('Error comments', () => {
    test('should handle document with error and regular comments', () => {
      const content = `
\`\`\`tabula
let A1 = "test"
\`\`\`
<!-- error: Something went wrong -->
\`\`\`csv
A,B
\`\`\`
`
      const chunks = extractChunks(content)

      const errorChunks = chunks.filter((c) => c.type === 'error')
      const codeChunks = chunks.filter((c) => c.type === 'code')

      expect(errorChunks).toHaveLength(0)
      expect(codeChunks).toHaveLength(1)
    })
  })

  describe('Edge cases', () => {
    test('should handle CSV block at start of file', () => {
      const content = `
\`\`\`csv
A,B
\`\`\`
Text
`

      const chunks = extractChunks(content)

      expect(chunks[1].type).toBe('csv')
    })

    test('should handle CSV block at end of file', () => {
      const content = `
Text
\`\`\`csv
A,B
\`\`\`
`

      const chunks = extractChunks(content)

      expect(chunks[chunks.length - 1].type).toBe('csv')
    })

    test('should handle empty CSV block', () => {
      const content = `
\`\`\`csv
\`\`\`
`
      const chunks = extractChunks(content)

      const csvChunks = chunks.filter((c) => c.type === 'csv')
      expect(csvChunks).toHaveLength(1)
      expect(csvChunks[0].content).toBe('')
    })
  })

  describe('processChunks', () => {
    test('should keep text chunks untouched', async () => {
      const chunks = extractChunks('Some text')
      const processed = await processChunks(mockExecuter, chunks)

      expect(processed).toEqual(chunks)
    })

    test("should replace CSV content with 'processed' when immediately followed by code", async () => {
      const content = `
\`\`\`csv
A,B,C
1,2,3
\`\`\`
\`\`\`tabula
let D1 = "Total"
\`\`\`
`

      const chunks = extractChunks(content)
      const processed = await processChunks(mockExecuter, chunks)

      const csvChunk = processed.find((c) => c.type === 'csv')
      expect(csvChunk?.content).toBe('processed')
    })

    test('should keep CSV content unchanged when not followed by code', async () => {
      const content = `
\`\`\`csv
A,B,C
1,2,3
\`\`\`

Some text
`
      const chunks = extractChunks(content)
      const processed = await processChunks(mockExecuter, chunks)

      const csvChunk = processed.find((c) => c.type === 'csv')
      expect(csvChunk?.content).toBe('A,B,C\n1,2,3')
    })

    test('should keep CSV at end of file unchanged', async () => {
      const content = `
\`\`\`csv
A,B,C
1,2,3
\`\`\`
`
      const chunks = extractChunks(content)
      const processed = await processChunks(mockExecuter, chunks)

      const csvChunk = processed.find((c) => c.type === 'csv')
      expect(csvChunk?.content).toBe('A,B,C\n1,2,3')
    })

    test('should handle multiple CSV blocks correctly', async () => {
      const content = `
\`\`\`csv
A,B
1,2
\`\`\`
\`\`\`tabula
let C1 = "X"
\`\`\`

\`\`\`csv
X,Y
3,4
\`\`\`

More text
`

      const chunks = extractChunks(content)
      const processed = await processChunks(mockExecuter, chunks)

      const csvChunks = processed.filter((c) => c.type === 'csv')
      expect(csvChunks[0].content).toBe('processed') // immediately followed by code
      expect(csvChunks[1].content).toBe('X,Y\n3,4') // followed by text
    })

    test('should keep CSV unchanged when followed by another CSV', async () => {
      const content = `
\`\`\`csv
A,B
1,2
\`\`\`

\`\`\`csv
X,Y
3,4
\`\`\`
`
      const chunks = extractChunks(content)
      const processed = await processChunks(mockExecuter, chunks)

      const csvChunks = processed.filter((c) => c.type === 'csv')
      expect(csvChunks[0].content).toBe('A,B\n1,2')
      expect(csvChunks[1].content).toBe('X,Y\n3,4')
    })

    test('should keep CSV unchanged when followed by non-whitespace text', async () => {
      const content = `
\`\`\`csv
A,B
1,2
\`\`\`
Some regular text here
`

      const chunks = extractChunks(content)
      const processed = await processChunks(mockExecuter, chunks)

      const csvChunk = processed.find((c) => c.type === 'csv')
      expect(csvChunk?.content).toBe('A,B\n1,2')
    })

    test('should handle empty CSV block immediately followed by code', async () => {
      const content = `
\`\`\`csv

\`\`\`
\`\`\`tabula
let A1 = "test"
\`\`\`
`

      const chunks = extractChunks(content)
      const processed = await processChunks(mockExecuter, chunks)

      const csvChunk = processed.find((c) => c.type === 'csv')
      expect(csvChunk?.content).toBe('processed')
    })

    test('should handle code followed by CSV (order matters)', async () => {
      const content = `
\`\`\`tabula
let D1 = "Total"
\`\`\`
\`\`\`csv
A,B,C
1,2,3
\`\`\`
`
      const chunks = extractChunks(content)
      const processed = await processChunks(mockExecuter, chunks)

      const csvChunk = processed.find((c) => c.type === 'csv')
      expect(csvChunk?.content).toBe('A,B,C\n1,2,3') // not followed by code
    })

    test('should handle complex document with mixed CSV/code patterns', async () => {
      const content = `
# Header

\`\`\`csv
A,B
1,2
\`\`\`
\`\`\`tabula
let C1 = SUM(A:B)
\`\`\`

Middle text

\`\`\`csv
X,Y
3,4
\`\`\`

More text

\`\`\`tabula
let Z1 = "Total";
\`\`\`
\`\`\`csv
M,N
5,6
\`\`\`
`
      const chunks = extractChunks(content)
      const processed = await processChunks(mockExecuter, chunks)

      const csvChunks = processed.filter((c) => c.type === 'csv')
      expect(csvChunks[0].content).toBe('processed') // CSV immediately followed by code
      expect(csvChunks[1].content).toBe('X,Y\n3,4') // CSV followed by text
      expect(csvChunks[2].content).toBe('M,N\n5,6') // CSV at end, not followed by code
    })

    test('should preserve chunk count (not add or remove chunks)', async () => {
      const content = `
\`\`\`csv
A,B
\`\`\`
\`\`\`tabula
let x = 1;
\`\`\`
Text
\`\`\`csv
C,D
\`\`\`
`
      const chunks = extractChunks(content)
      const processed = await processChunks(mockExecuter, chunks)

      expect(processed.length).toBe(chunks.length)
    })

    test('should only modify content field for processed CSV chunks', async () => {
      const content = `
\`\`\`csv
A,B,C
1,2,3
\`\`\`
\`\`\`tabula
let D1 = "X";
\`\`\`
`
      const chunks = extractChunks(content)
      const processed = await processChunks(mockExecuter, chunks)

      const originalCsv = chunks.find((c) => c.type === 'csv')
      const processedCsv = processed.find((c) => c.type === 'csv')

      expect(processedCsv?.type).toBe(originalCsv?.type)
      expect(processedCsv?.content).toBe('processed')
      expect(processedCsv?.content).not.toBe(originalCsv?.content)
    })
  })

  describe('Round-trip: extract -> output', () => {
    test('should preserve text-only content', () => {
      const input = '# Heading\n\nSome text content.'
      const chunks = extractChunks(input)
      const output = outputChunks(chunks)

      expect(output).toBe(input)
    })

    test('should preserve single CSV block', () => {
      const input = `
# Title

\`\`\`csv
A,B,C
1,2,3
\`\`\`

End text
`
      const chunks = extractChunks(input)
      const output = outputChunks(chunks)

      expect(output).toBe(input)
    })

    test('should preserve multiple CSV blocks', () => {
      const input = `
\`\`\`csv
A,B
1,2
\`\`\`

Text

\`\`\`csv
X,Y
3,4
\`\`\`
`
      const chunks = extractChunks(input)
      const output = outputChunks(chunks)

      expect(output).toBe(input)
    })

    test('should preserve single-line code', () => {
      const input = `
\`\`\`tabula
let D1 = "Total";
\`\`\`

\`\`\`csv
A,B,C
1,2,3
\`\`\`
`
      const chunks = extractChunks(input)
      const output = outputChunks(chunks)

      expect(output).toBe(input)
    })

    test('should preserve multi-line code', () => {
      const input = `
\`\`\`tabula
let D1 = "Total";
let D2 = A2 + B2;
\`\`\`

\`\`\`csv
A,B,C
1,2,3
\`\`\`
`
      const chunks = extractChunks(input)
      const output = outputChunks(chunks)

      expect(output).toBe(input)
    })

    test('should remove error comments in round-trip', () => {
      const input = `
<!-- error: Something went wrong -->

\`\`\`csv
A,B,C
1,2,3
\`\`\`
`

      const expected = `

\`\`\`csv
A,B,C
1,2,3
\`\`\`
`
      const chunks = extractChunks(input)
      const output = outputChunks(chunks)

      expect(output).toBe(expected)
    })

    test('should preserve mixed content with multiple chunks', () => {
      const input = `
# Document

\`\`\`tabula
let D1 = "Total";
\`\`\`
\`\`\`csv
A,B,C
1,2,3
\`\`\`

More text

\`\`\`tabula
let E1 = "Sum";
fmt E1 = "%.2f";
\`\`\`
\`\`\`csv
X,Y,Z
4,5,6
\`\`\`
`
      const chunks = extractChunks(input)
      const output = outputChunks(chunks)

      expect(output).toBe(input)
    })

    test('should preserve CSV block at start', () => {
      const input = `\`\`\`csv
A,B
\`\`\`
Text
`
      const chunks = extractChunks(input)
      const output = outputChunks(chunks)

      expect(output).toBe(input)
    })

    test('should preserve CSV block at end', () => {
      const input = `
Text
\`\`\`csv
A,B
\`\`\``
      const chunks = extractChunks(input)
      const output = outputChunks(chunks)

      expect(output).toBe(input)
    })

    test('should preserve empty content', () => {
      const input = ''
      const chunks = extractChunks(input)
      const output = outputChunks(chunks)

      expect(output).toBe(input)
    })

    test('should preserve complex mixed content', () => {
      const input = `\`\`\`csv
A
\`\`\`
\`\`\`tabula
let x = 9;
\`\`\`
\`\`\`csv
B
\`\`\``
      const chunks = extractChunks(input)
      const output = outputChunks(chunks)

      expect(output).toBe(input)
    })
  })
})
