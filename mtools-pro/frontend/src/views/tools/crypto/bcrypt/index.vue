<script setup>
import { computedAsync } from '@vueuse/core'
import { BcryptEncode, BcryptCompare } from '@/api/crypto.js'

import InputCopy from '@/components/input-copy/index.vue'

const input = ref('')
const saltCount = ref(10)
const output = computedAsync(async () => {
  const resp = await BcryptEncode({
    input: input.value,
    saltCount: saltCount.value
  })
  return resp.data
})

const origin = ref('')
const encode = ref('')
const result = computedAsync(async () => {
  const resp = await BcryptCompare({
    origin: origin.value,
    encode: encode.value
  })
  return resp.data
})
</script>

<template>
  <div ref="container" flex flex-col flex-gap-20 wh-full overflow-auto p-20 class="cus-scroll">
    <div>加密</div>
    <n-input-group>
      <n-input-group-label w-120>原始内容</n-input-group-label>
      <n-input type="text" v-model:value="input" />
    </n-input-group>
    <n-input-group>
      <n-input-group-label w-120>加密盐的长度</n-input-group-label>
      <n-input-number v-model:value="saltCount" clearable w-full />
    </n-input-group>
    <n-input-group>
      <n-input-group-label w-120>加密结果</n-input-group-label>
      <input-copy v-model:value="output" />
    </n-input-group>
    <div>匹配</div>
    <n-input-group>
      <n-input-group-label w-120>原始内容</n-input-group-label>
      <n-input type="text" v-model:value="origin" />
    </n-input-group>
    <n-input-group>
      <n-input-group-label w-120>加密内容</n-input-group-label>
      <n-input type="text" v-model:value="encode" />
    </n-input-group>
    <div>
      匹配结果:
      <span m-l-16>{{ result }}</span>
    </div>
  </div>
</template>

<style scoped lang="scss"></style>
