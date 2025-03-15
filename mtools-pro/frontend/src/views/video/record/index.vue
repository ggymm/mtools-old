<script setup>
const video = ref()

let stream = null
let recorder = null
const start = async () => {
  try {
    stream = await navigator.mediaDevices.getDisplayMedia({ video: true })
    recorder = new MediaRecorder(stream)

    recorder.addEventListener('dataavailable', (event) => {
      if (event.data && event.data.size > 0) {
        video.value.src = video.value.srcObject = null
        video.value.src = URL.createObjectURL(event.data)
        video.value.blob = event.data
      }
    })

    recorder.start()

    video.value.srcObject = stream
    video.value.controls = true
  } catch (e) {}
}

const clear = () => {
  video.value.src = video.value.srcObject = null
}

const finish = () => {
  stream.getVideoTracks().forEach((track) => {
    track.stop()
  })
  recorder.stop()
}

const download = () => {
  if (!video.value.blob) {
    return
  }
  const a = document.createElement('a')
  a.href = URL.createObjectURL(video.value.blob)
  a.download = `video-${new Date().getTime()}.mp4`
  a.click()
}

onMounted(() => {
  if (!navigator.getDisplayMedia && (!navigator.mediaDevices || !navigator.mediaDevices.getDisplayMedia)) {
    window.$message.error('当前浏览器不支持录制')
  }
})
</script>

<template>
  <div ref="container" flex flex-col flex-gap-20 wh-full overflow-auto p-20 class="cus-scroll">
    <div flex flex-gap-20 justify-center>
      <n-button type="info" @click="start">录制视频</n-button>
      <n-button type="info" @click="clear">清空视频内容</n-button>
      <n-button type="warning" @click="finish">停止录制</n-button>
      <n-button type="success" @click="download">下载视频</n-button>
    </div>
    <video ref="video" controls autoplay playsinline style="height: calc(100% - 34px - 20px)" />
  </div>
</template>

<style scoped lang="scss"></style>
