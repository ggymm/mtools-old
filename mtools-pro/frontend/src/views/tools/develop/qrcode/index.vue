<script setup>
import { downloadImage } from '@/utils'
import { useWindowResize } from '@/hooks/event/window-resize.js'

import { LoadLocalDecode, LoadClipboardDecode, GenerateContentEncode } from '@/api/qrcode.js'

const loadLocalImage = async () => {
  const { msg, data, success } = await LoadLocalDecode()

  if (!success) {
    window.$message.error(msg)
    return
  }
  imageSrc.value = data['base64']
  qrcodeContent.value = data['content']
  await nextTick(() => {
    resetImageSize()
  })
}

const loadClipboardImage = async () => {
  const { msg, data, success } = await LoadClipboardDecode()

  if (!success) {
    window.$message.error(msg)
    return
  }
  imageSrc.value = data['base64']
  qrcodeContent.value = data['content']
  await nextTick(() => {
    resetImageSize()
  })
}

const saveImage = () => {
  downloadImage(imageSrc.value, `qrcode-${Date.now()}.png`)
}

const generateQrcode = async () => {
  let size = 256
  if (image.value) {
    const width = image.value.clientWidth
    const height = image.value.clientHeight
    size = width > height ? height : width
  }
  const { msg, data, success } = await GenerateContentEncode({
    size: size,
    content: qrcodeContent.value
  })

  if (!success) {
    window.$message.error(msg)
    return
  }
  imageSrc.value = data['base64']
  await nextTick(() => {
    resetImageSize()
  })
}

const image = ref()
const imageSrc = ref('')
const qrcodeContent = ref('')

const resetImageSize = () => {
  const width = image.value.clientWidth
  const height = image.value.clientHeight

  const img = image.value.querySelector('img')
  if (img) {
    if (width > height) {
      img.style.width = `${height}px`
      img.style.height = `${height}px`
    } else {
      img.style.width = `${width}px`
      img.style.height = `${width}px`
    }
  }
}
useWindowResize(resetImageSize, 30)

onMounted(() => {
  nextTick(() => {
    resetImageSize()
  })
})
</script>

<template>
  <div ref="container" flex flex-col flex-gap-20 wh-full overflow-auto p-20 class="cus-scroll">
    <div flex flex-gap-20 items-center justify-center min-h-36>
      <n-button type="info" @click="loadLocalImage">加载本地图片</n-button>
      <n-button type="info" @click="loadClipboardImage">加载剪切板图片</n-button>
      <n-button type="success" @click="saveImage">保存为图片</n-button>
    </div>
    <div ref="image" flex items-center justify-center flex-1 class="max-h-[calc(100%-36px-150px-40px)]">
      <img :src="imageSrc" wh-full alt="" v-if="imageSrc" />
      <icon-mdi:file-image-box wh-full v-if="!imageSrc" />
    </div>
    <div flex items-center justify-center min-h-150>
      <n-input
        type="textarea"
        v-model:value="qrcodeContent"
        :rows="6"
        :autosize="{
          minRows: 6,
          maxRows: 6
        }"
        @change="generateQrcode"
      />
    </div>
  </div>
</template>

<style scoped lang="scss"></style>
