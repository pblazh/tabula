import { spawn } from 'node:child_process'
import * as crypto from 'node:crypto'
import * as path from 'node:path'
import * as os from 'node:os'
import * as fs from 'node:fs/promises'

import { DataAdapter, normalizePath, Notice } from 'obsidian'
import { TabulaSettings } from './types'

export class Executer {
  constructor(
    private settings: TabulaSettings,
    private adapter: DataAdapter,
    private path: string,
  ) {}

  async execute(content: string): Promise<string> {
    // @ts-expect-error undocumented
    const createSource = this.adapter.basePath
      ? createVaultSource
      : createTmpSource

    const { dataPath, cleanup } = await createSource(
      this.adapter,
      this.path,
      content,
    )

    try {
      const args = [
        this.settings.autoFormat ? '-a' : '',
        '-m',
        '-i',
        dataPath,
      ].filter(Boolean)

      return await run(this.settings.executablePath, args)
    } finally {
      cleanup().catch((err) => {
        new Notice(`FAILED TO REMOVE ${dataPath}, ${err}`, 0)
      })
    }
  }
}

function run(
  cmd: string,
  args: string[] = [],
  input: string = '',
): Promise<string> {
  return new Promise((resolve, reject) => {
    const child = spawn(cmd, args)

    let stdout = ''
    let stderr = ''

    child.stdout.on('data', (data: string) => {
      stdout += data
    })

    child.stderr.on('data', (data: string) => {
      stderr += data
    })

    child.on('error', reject)

    child.on('close', (code: number | null) => {
      if (code === 0) {
        resolve(stdout)
      } else {
        reject(new Error(stderr))
      }
    })

    if (input) {
      child.stdin.write(input)
    }
    child.stdin.end()
  })
}

async function createVaultSource(
  adapter: DataAdapter,
  basePath: string,
  content: string,
): Promise<{ dataPath: string; cleanup: () => Promise<void> }> {
  // @ts-expect-error undocumented
  if (!adapter.basePath) {
    throw new Error('Can not determine base path of a vault')
  }
  // @ts-expect-error undocumented
  const adapterBasePath: string = adapter.basePath

  const tmpPath = normalizePath(
    path.join(basePath, `tabula_${crypto.randomBytes(6).toString('hex')}.md`),
  )

  await adapter.write(tmpPath, content)
  return {
    dataPath: path.join(adapterBasePath, tmpPath),
    cleanup: () => adapter.remove(tmpPath),
  }
}

async function createTmpSource(
  _adapter: DataAdapter,
  _basePath: string,
  content: string,
): Promise<{ dataPath: string; cleanup: () => Promise<void> }> {
  const tmpPath = path.join(
    os.tmpdir(),
    `tabula_${crypto.randomBytes(6).toString('hex')}.md`,
  )
  await fs.writeFile(tmpPath, content)

  return {
    dataPath: tmpPath,
    cleanup: () => fs.unlink(tmpPath),
  }
}
