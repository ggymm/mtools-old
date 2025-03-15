<script setup>
import MD5 from 'crypto-js/md5.js'
import SHA1 from 'crypto-js/sha1.js'
import SHA256 from 'crypto-js/sha256.js'
import SHA512 from 'crypto-js/sha512.js'

import hex from 'crypto-js/enc-hex.js'
import base64 from 'crypto-js/enc-base64.js'
import base64url from 'crypto-js/enc-base64url.js'

import InputCopy from '@/components/input-copy/index.vue'

const formats = [
  { label: '16进制', value: 'hex' },
  { label: 'Base64', value: 'base64' },
  { label: 'Base64URL', value: 'base64url' }
]

const hashFuncList = { MD5, SHA1, SHA256, SHA512 }
const hashFormatList = { hex, base64, base64url }

const input = ref('')
const format = ref('hex')

const funcNameList = Object.keys(hashFuncList)

const hashText = (name) => {
  return hashFormatList[format.value].stringify(hashFuncList[name](input.value))
}
</script>

<template>
  <div ref="container" flex flex-col flex-gap-20 wh-full overflow-auto p-20 class="cus-scroll">
    <n-input
      type="textarea"
      v-model:value="input"
      :autosize="{
        minRows: 3,
        maxRows: 5
      }"
    />
    <n-select v-model:value="format" :options="formats" />
    <div v-for="name in funcNameList" :key="name">
      <n-input-group>
        <n-input-group-label w-120>{{ name }}</n-input-group-label>
        <input-copy :value="hashText(name)" readonly />
      </n-input-group>
    </div>
  </div>
</template>

<style scoped lang="scss"></style>
