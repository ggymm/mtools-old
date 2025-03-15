import { ref, nextTick, onBeforeUnmount } from 'vue'
import * as monaco from 'monaco-editor'

export function useMonacoEditor(language) {
  let editor = null
  const editorRef = ref(null)

  const createEditor = (options) => {
    if (!editorRef.value) return

    editor = monaco.editor.create(editorRef.value, {
      // 模型
      model: monaco.editor.createModel('', language),
      automaticLayout: true,
      scrollBeyondLastLine: false,
      ...options
    })
    return editor
  }

  async function format() {
    if (!editor) return
    await editor.getAction('editor.action.formatDocument').run()
  }

  const getValue = () => {
    if (!editor) return
    return editor.getValue()
  }

  const updateValue = (value) => {
    if (!editor) return
    nextTick(() => {
      editor.setValue(value)
    }).then(() => {
      setTimeout(async () => {
        await format()
      }, 10)
    })
  }

  const getOption = (name) => {
    if (!editor) return
    return editor.getOption(name)
  }

  const updateOption = (name, value) => {
    if (!editor) return
    editor.updateOptions({
      [name]: value
    })
  }

  const updateLanguage = (language) => {
    if (!editor) return
    monaco.editor.setModelLanguage(editor.getModel(), language)
  }

  onBeforeUnmount(() => {
    editor && editor.dispose()
  })

  return {
    editorRef,
    createEditor,
    getValue,
    updateValue,
    getOption,
    updateOption,
    updateLanguage
  }
}
