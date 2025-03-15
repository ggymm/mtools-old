<script setup>
import { NumberBase } from '@/api/conv.js'

import InputCopy from '@/components/input-copy/index.vue'

const input = ref('100')
const inputBase = ref(10)

const output = ref({
  binary: '',
  octal: '',
  decimal: '',
  hexadecimal: ''
})

watch(
  [input, inputBase],
  async () => {
    const { msg, data, success } = await NumberBase({
      input: input.value,
      inputBase: inputBase.value
    })

    if (!success) {
      window.$message.error(msg)
      return
    }
    output.value = data
  },
  {
    immediate: true
  }
)
</script>

<template>
  <div ref="container" flex flex-col flex-gap-20 wh-full overflow-auto p-20 class="cus-scroll">
    <n-input-group>
      <n-input-group-label w-120>输入内容</n-input-group-label>
      <n-input type="text" v-model:value="input" />
    </n-input-group>
    <n-input-group>
      <n-input-group-label w-120>输入进制</n-input-group-label>
      <n-input-number v-model:value="inputBase" clearable w-full />
    </n-input-group>
    <div>转换结果</div>
    <n-input-group>
      <n-input-group-label w-120>2进制</n-input-group-label>
      <input-copy :value="output.binary" />
    </n-input-group>
    <n-input-group>
      <n-input-group-label w-120>8进制</n-input-group-label>
      <input-copy :value="output.octal" />
    </n-input-group>
    <n-input-group>
      <n-input-group-label w-120>10进制</n-input-group-label>
      <input-copy :value="output.decimal" />
    </n-input-group>
    <n-input-group>
      <n-input-group-label w-120>16进制</n-input-group-label>
      <input-copy :value="output.hexadecimal" />
    </n-input-group>
  </div>
</template>

<style scoped lang="scss"></style>
