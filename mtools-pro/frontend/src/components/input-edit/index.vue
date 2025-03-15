<script setup>
const props = defineProps({
  value: {
    type: String,
    default: ''
  }
})

const edit = ref(false)
const localValue = ref()

const emits = defineEmits(['update:value'])

const changeEdit = () => {
  edit.value = true
}

const changeUpdate = () => {
  edit.value = false
  emits('update:value', localValue.value)
}

onMounted(() => {
  localValue.value = props.value
})
</script>

<template>
  <div flex items-center h-32 font-size-14>
    <n-space v-if="!edit" align="center">
      <span>{{ value }}</span>
      <icon-material-symbols:edit-square-outline @click="changeEdit" class="icon" />
    </n-space>
    <n-space v-if="edit" align="center">
      <n-input v-model:value="localValue" type="text" @blur="changeUpdate" />
    </n-space>
  </div>
</template>

<style scoped lang="scss">
.icon {
  width: 18px;
  height: 18px;
  cursor: pointer;
}
</style>
