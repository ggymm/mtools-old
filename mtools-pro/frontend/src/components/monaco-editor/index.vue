<script setup>
import { useMonacoEditor } from '@/hooks/comp/monaco-editor.js'

const props = defineProps({
  width: {
    types: String | Number,
    default: '100%'
  },
  height: {
    types: String | Number,
    default: '100%'
  },
  language: {
    type: String,
    default: 'javascript'
  },
  modelValue: {
    type: String,
    default: ''
  }
})

const emits = defineEmits(['blur', 'update:modelValue'])

const editorStyle = computed(() => ({
  width: typeof props.width === 'number' ? `${props.width}px` : props.width,
  height: typeof props.height === 'number' ? `${props.height}px` : props.height
}))

const monacoEditor = useMonacoEditor(props.language)

const editorRef = monacoEditor.editorRef
const createEditor = monacoEditor.createEditor

const getValue = monacoEditor.getValue
const updateValue = monacoEditor.updateValue
const updateLanguage = monacoEditor.updateLanguage

onMounted(() => {
  const editor = createEditor({})
  updateModelValue(props.modelValue)

  editor.onDidChangeModelContent(() => {
    emits('update:modelValue', editor.getValue())
  })
  editor.onDidBlurEditorText(() => {
    emits('blur')
  })
})

watch(
  () => props.language,
  () => {
    updateLanguage(props.language)
  }
)

watch(
  () => props.modelValue,
  () => {
    updateModelValue(props.modelValue)
  }
)

const updateModelValue = (val) => {
  if (val !== getValue()) {
    updateValue(val)
  }
}
</script>

<template>
  <div ref="editorRef" :style="editorStyle" border-1></div>
</template>

<style scoped lang="scss"></style>
