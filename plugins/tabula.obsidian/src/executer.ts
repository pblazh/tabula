import { spawn } from 'node:child_process'
import * as crypto from 'node:crypto'
import * as path from 'node:path'
import * as os from 'node:os'
import * as fs from 'node:fs/promises'

import { Notice } from 'obsidian'
import type { TFile } from 'obsidian'
import { TabulaSettings } from './types'

export class Executer {
  constructor(
    private settings: TabulaSettings,
    private file: TFile,
  ) {}

  async execute(): Promise<string> {
    // @ts-expect-error undocumented
    const basePath = this.file.vault.adapter.basePath as string | undefined
    if (basePath) {
      return runOnFile(this.settings, path.join(basePath, this.file.path))
    }

    const { dataPath, cleanup } = await createTmpSource(
      await this.file.vault.read(this.file),
    )
    try {
      return await runOnFile(this.settings, dataPath)
    } finally {
      cleanup().catch((err) => {
        new Notice(`FAILED TO REMOVE ${dataPath}, ${err}`, 0)
      })
    }
  }
}

function runOnFile(
  settings: TabulaSettings,
  dataPath: string,
): Promise<string> {
  const args = [settings.autoFormat ? '-a' : '', '-m', '-i', dataPath].filter(
    Boolean,
  )

  return run(settings.executablePath, args)
}

function run(cmd: string, args: string[] = []): Promise<string> {
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

    child.stdin.end()
  })
}

async function createTmpSource(
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
