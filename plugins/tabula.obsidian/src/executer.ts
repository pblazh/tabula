import { TabulaSettings } from './types'
import { spawn } from 'node:child_process'
import * as crypto from 'node:crypto'
import * as path from 'node:path'
import * as os from 'node:os'
import * as fs from 'node:fs/promises'
import { DataAdapter } from 'obsidian'

export class Executer {
  constructor(
    private settings: TabulaSettings,
    private adapter: DataAdapter,
    private path: string,
  ) {}

  async execute(data: string, code: string): Promise<string> {
    // @ts-expect-error undocumented
    const createSource = this.adapter.basePath
      ? createVaultSource
      : createTmpSource

    const { code: codePath, data: dataPath } = await createSource(
      this.adapter,
      this.path,
      code,
      data,
    )

    try {
      const args = [
        this.settings.autoFormat ? '-a' : '',
        '-s',
        codePath,
        '-i',
        dataPath,
      ].filter(Boolean)

      return await run(this.settings.executablePath, args, data)
    } finally {
      await Promise.all([
        this.adapter.remove(codePath),
        this.adapter.remove(dataPath),
      ]).catch((err) => {
        return {
          result: '',
          error: String(err),
        }
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
  code: string,
  data: string,
): Promise<{ code: string; data: string }> {
  // @ts-expect-error undocumented
  if (!adapter.basePath) {
    throw new Error('Can not determine base path of a vault')
  }
  // @ts-expect-error undocumented
  const adapterBasePath: string = adapter.basePath

  const codePath = path.join(
    basePath,
    `tabula_${crypto.randomBytes(6).toString('hex')}.tbl`,
  )
  const codePathAbsolute = path.join(adapterBasePath, codePath)

  const dataPath = path.join(
    basePath,
    `tabula_${crypto.randomBytes(6).toString('hex')}.csv`,
  )
  const dataPathAbsolute = path.join(adapterBasePath, dataPath)

  await Promise.all([
    adapter.write(codePath, code),
    adapter.write(dataPath, data),
  ])

  return { data: dataPathAbsolute, code: codePathAbsolute }
}

async function createTmpSource(
  _adapter: DataAdapter,
  _basePath: string,
  code: string,
  data: string,
): Promise<{ code: string; data: string }> {
  const codePath = path.join(
    os.tmpdir(),
    `tabula_${crypto.randomBytes(6).toString('hex')}.tbl`,
  )

  const dataPath = path.join(
    os.tmpdir(),
    `tabula_${crypto.randomBytes(6).toString('hex')}.csv`,
  )
  await Promise.all([
    fs.writeFile(codePath, code),
    fs.writeFile(dataPath, data),
  ])

  return { data: dataPath, code: codePath }
}
