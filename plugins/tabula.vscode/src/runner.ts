import * as vscode from 'vscode'
import { exec } from 'child_process'

export const runScript = (path: vscode.Uri): Promise<void> =>
  new Promise((resolve, reject) => {
    const editor = vscode.window.activeTextEditor
    if (!editor) return

    const pathParts = path.path.split('/')
    const fileName = pathParts[pathParts.length - 1]

    // Get configuration settings
    const config = vscode.workspace.getConfiguration('tabula')
    const tabulaPath = config.get<string>('executablePath', 'tabula')
    const autoFormat = config.get<boolean>('autoFormat', true)

    const isMarkdown = editor.document.languageId.toLowerCase() === 'md'

    const autoFormatFlag = autoFormat ? '-a' : ''
    const markDownFlag = isMarkdown ? '-m' : ''

    const command = `"${tabulaPath}" ${markDownFlag} ${autoFormatFlag} -u "${path.path}"`

    execute(command, fileName, resolve, reject)
  })

const execute = (
  command: string,
  fileName: string,
  resolve: (value: void | PromiseLike<void>) => void,
  reject: (reason: Error) => void,
) =>
  exec(command, (error, _stdout, stderr) => {
    if (error) {
      vscode.window.showErrorMessage(
        `Run script error: ${error.message}. Check that tabula is installed and the path is correct in settings.`,
      )
      console.error(`exec error: ${error}`)
      return reject(error)
    }
    if (stderr) {
      console.error(`stderr: ${stderr}`)
    }
    vscode.window.showInformationMessage(`Tabula updated ${fileName}.`)
    resolve()
  })
