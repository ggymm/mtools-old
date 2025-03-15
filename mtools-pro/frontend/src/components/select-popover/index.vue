<script setup>
import { useVModel } from '@vueuse/core'

const props = defineProps({
  cols: {
    type: Number,
    default: 1
  },
  value: {
    type: Array,
    default: () => []
  },
  items: {
    type: Array,
    default: () => []
  },
  multiple: {
    type: Boolean,
    default: false
  },
  selection: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['update:value'])

const value = useVModel(props, 'value', emit)

const show = ref(false)

const has = (val) => {
  return value.value.includes(val)
}

const select = (val) => {
  if (props.multiple) {
    if (has(val)) {
      value.value = value.value.filter((item) => item !== val)
    } else {
      value.value = [...value.value, val]
    }
  } else {
    value.value = [val]
  }
}

const selectedStr = computed(() => {
  if (value.value.length === 0) {
    return props.selection
  }

  let index = -1
  let keyFormat = []
  let labelFormat = []

  const itemMap = props.items.reduce((acc, cur) => {
    acc[cur.key] = cur.label
    return acc
  }, {})

  for (const v of value.value) {
    const key = v
    const label = itemMap[key]
    if (keyFormat.length === 0) {
      index++
      keyFormat.push(key)
      labelFormat.push(label)
    } else {
      const last = keyFormat[index]
      if ((last + '').indexOf('-') > 0) {
        const [first, second] = last.split('-')
        if (parseInt(second) === key - 1) {
          const lastLabel = labelFormat[index]
          const [firstLabel] = lastLabel.split('-')

          keyFormat[index] = `${first}-${key}`
          labelFormat[index] = `${firstLabel}-${label}`
        } else {
          index++
          keyFormat.push(key)
          labelFormat.push(label)
        }
      } else {
        if (parseInt(last) === key - 1) {
          keyFormat[index] = `${last}-${key}`
          labelFormat[index] = `${labelFormat[index]}-${label}`
        } else {
          index++
          keyFormat.push(key)
          labelFormat.push(label)
        }
      }
    }
  }

  return labelFormat.join(',')
})
</script>

<template>
  <n-popover v-model:show="show" trigger="click">
    <template #trigger>
      <n-button>{{ selectedStr }}</n-button>
    </template>

    <n-grid :cols="cols">
      <n-grid-item v-for="(item, i) in items" :key="i" text-center>
        <n-tag :checked="has(item.key)" @updateChecked="select(item.key)" checkable>
          {{ item.label }}
        </n-tag>
      </n-grid-item>
    </n-grid>
  </n-popover>
</template>

<style scoped lang="scss"></style>
