<script setup>
const encodeInput = ref('')
const encodeOutput = computed(() => {
  try {
    return encodeURIComponent(encodeInput.value)
  } catch (error) {
    return ''
  }
})

const decodeInput = ref('')
const decodeOutput = computed(() => {
  try {
    return decodeURIComponent(decodeInput.value)
  } catch (error) {
    return ''
  }
})
</script>

<template>
  <div ref="container" flex flex-col flex-gap-20 wh-full overflow-auto p-20 class="cus-scroll">
    <div>编码</div>
    <n-input
      type="textarea"
      v-model:value="encodeInput"
      :autosize="{
        minRows: 3,
        maxRows: 5
      }"
    />
    <n-input
      type="textarea"
      v-model:value="encodeOutput"
      :autosize="{
        minRows: 3,
        maxRows: 5
      }"
    />
    <div>解码</div>
    <n-input
      type="textarea"
      v-model:value="decodeInput"
      :autosize="{
        minRows: 3,
        maxRows: 5
      }"
    />
    <n-input
      type="textarea"
      v-model:value="decodeOutput"
      :autosize="{
        minRows: 3,
        maxRows: 5
      }"
    />
  </div>
</template>

<style scoped lang="scss"></style>
