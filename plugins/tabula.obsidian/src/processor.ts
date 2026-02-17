import { Executer } from './executer'
import { Chunk, Match } from './types'

export function extractChunks(content: string): Chunk[] {
  const chunks: Chunk[] = []

  // Find all CSV blocks and Tabula comments with their positions
  const matches: Match[] = []

  // Find CSV blocks
  const csvBlockRegex = /```csv\n*(?<content>[\s\S]*?)\n```\n/g
  let match
  while ((match = csvBlockRegex.exec(content)) !== null) {
    matches.push({
      type: 'csv',
      start: match.index,
      end: match.index + match[0].length,
      content: match.groups?.content ?? '',
    })
  }

  // Find Tabula comments
  const commentRegex = /```tabula\n*(?<content>[\s\S]*?)\n```\n/g
  while ((match = commentRegex.exec(content)) !== null) {
    const commentContent = match.groups?.content?.trim() ?? ''

    matches.push({
      type: 'code',
      start: match.index,
      end: match.index + match[0].length,
      content: commentContent,
    })
  }

  // Find error comments
  const errorRegex = /<!--\s*error:(?<content>[\s\S]*?)-->\n/g
  while ((match = errorRegex.exec(content)) !== null) {
    const errorContent = match.groups?.content?.trim() ?? ''

    matches.push({
      type: 'error',
      start: match.index,
      end: match.index + match[0].length,
      content: errorContent,
    })
  }

  // Sort matches by position
  matches.sort((a, b) => a.start - b.start)

  // Build chunks array
  let lastIndex = 0
  for (const m of matches) {
    // Add text chunk before this match (if any)
    if (m.start > lastIndex) {
      chunks.push({
        type: 'text',
        content: content.substring(lastIndex, m.start),
      })
    }

    // Add the CSV or comment chunk
    if (m.type !== 'error') {
      chunks.push(m)
    }

    lastIndex = m.end
  }

  // Add remaining text after last match (if any)
  if (lastIndex < content.length) {
    chunks.push({
      type: 'text',
      content: content.substring(lastIndex),
    })
  }

  return chunks
}

const formatChunk = (chunk: Chunk): string => {
  switch (chunk.type) {
    case 'csv':
      return ['```csv', chunk.content, '```', ''].join('\n')
    case 'code':
      return ['```tabula', chunk.content, '```', ''].join('\n')
    case 'error':
      return `<!-- error:\n${chunk.content}-->\n`
    default:
      return chunk.content
  }
}

export function outputChunks(chunks: Chunk[]): string {
  let output = ''
  chunks.forEach((chunk) => {
    output += formatChunk(chunk)
  })

  return output
}

export async function processChunks(
  executer: Executer,
  chunks: Chunk[],
): Promise<Chunk[]> {
  const processed: Promise<Chunk[]>[] = []
  const executeChunk = createProcessor(executer)

  for (let i = 0; i < chunks.length; i++) {
    const chunk = chunks[i]

    let nextChunk: Chunk | undefined

    // Skip chunks with empty lines or spaces
    for (let j = i + 1; j < chunks.length; j++) {
      if (!isEmptyChunk(chunks[j])) {
        nextChunk = chunks[j]
        break
      }
    }

    // If tabula code is immediately after the csv
    if (chunk.type === 'csv' && nextChunk?.type === 'code') {
      processed.push(executeChunk(chunk, nextChunk))
      continue
    }

    processed.push(Promise.resolve([chunk]))
  }
  const resolved = await Promise.all(processed)
  return resolved.flat(1)
}

const createProcessor =
  (executer: Executer) =>
  async (csv: Chunk, code: Chunk): Promise<Chunk[]> => {
    const out: Chunk[] = []
    try {
      out.push({
        ...csv,
        content: await executer.execute(csv.content, code.content),
      })
    } catch (err) {
      if (err) {
        out.push(csv)
        out.push({ type: 'error', content: String(err) })
      }
    }

    return out
  }

function isEmptyChunk(chunk: Chunk): boolean {
  if (chunk.type !== 'text') return false
  return chunk.content.replace(/[\s\n]*/, '').length === 0
}
