import { spawn } from 'node:child_process'
import * as crypto from 'node:crypto'
import * as path from 'node:path'
import * as os from 'node:os'
import * as fs from 'node:fs/promises'

import { DataAdapter } from 'obsidian'
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

    const dataPath = await createSource(this.adapter, this.path, content)

    try {
      const args = [
        this.settings.autoFormat ? '-a' : '',
        '-m',
        '-i',
        dataPath,
      ].filter(Boolean)

      return await run(this.settings.executablePath, args)
    } finally {
      fs.unlink(dataPath)
        .catch((err) => {
          return {
            result: '',
            error: String(err),
          }
        })
        .catch((err) => {
          console.log('FAILED TO REMOVE', dataPath, err)
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

    child.on('close', (code: null) => {
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
): Promise<string> {
  // @ts-expect-error undocumented
  if (!adapter.basePath) {
    throw new Error('Can not determine base path of a vault')
  }
  // @ts-expect-error undocumented
  const adapterBasePath: string = adapter.basePath

  const tmpPath = path.join(
    basePath,
    `tabula_${crypto.randomBytes(6).toString('hex')}.md`,
  )

  await adapter.write(tmpPath, content)
  return path.join(adapterBasePath, tmpPath)
}

async function createTmpSource(
  _adapter: DataAdapter,
  _basePath: string,
  content: string,
): Promise<string> {
  const tmpPath = path.join(
    os.tmpdir(),
    `tabula_${crypto.randomBytes(6).toString('hex')}.md`,
  )
  await fs.writeFile(tmpPath, content)

  return tmpPath
}
