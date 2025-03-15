<script setup>
import { useVModel, useClipboard } from '@vueuse/core'

const props = defineProps({
  value: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['update:value'])

const value = useVModel(props, 'value', emit)

const copy = useClipboard().copy

const copyValue = async () => {
  await copy(value.value)
  window.$message.success('复制成功')
}
</script>

<template>
  <n-input v-model:value="value" type="text" readonly>
    <template #suffix>
      <n-button text @click="copyValue" w-24 h-full>
        <n-icon>
          <icon-mdi:content-copy />
        </n-icon>
      </n-button>
    </template>
  </n-input>
</template>

<style scoped lang="scss"></style>
