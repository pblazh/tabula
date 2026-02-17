import { defineConfig } from '@vscode/test-cli'

export default defineConfig({
  files: 'out/**/*.test.js',
  version: 'stable',
  mocha: {
    ui: 'tdd',
    timeout: 20000,
  },
  launchArgs: ['--disable-dev-shm-usage', '--disable-gpu'],
  extensionDevelopmentPath: process.cwd(),
})
