import path from 'path'

import monacoEditorPlugin from 'vite-plugin-monaco-editor'

export default () => {
  // noinspection JSUnusedGlobalSymbols
  return monacoEditorPlugin.default({
    languageWorkers: ['editorWorkerService', 'json', 'typescript'],
    customDistPath: (root, buildOutDir, base) => {
      return path.join(root, buildOutDir, base, 'assets/monaco')
    }
  })
}
