import { TabulaSettings } from './types'

export function renderCode(
  _settings: TabulaSettings,
  _source: string,
  el: HTMLElement,
) {
  const container = el.createEl('div')
  container.className = 'cm-comment tabula-code'
  const button = container.createEl('div')
  button.appendText('⚙ Tabula')
}
