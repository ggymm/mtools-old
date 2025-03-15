<script setup>
import parser from 'cron-parser'

import InputCopy from '@/components/input-copy/index.vue'
import SelectPopover from '@/components/select-popover/index.vue'

const genItems = (min, max, labelFormat) => {
  return [...Array(max - min + 1).keys()].map((item) => {
    return {
      key: item + min,
      label: labelFormat(item + min)
    }
  })
}

const formatDate = (date) => {
  const year = date.getFullYear()
  const month = (date.getMonth() + 1).toString().padStart(2, '0')
  const day = date.getDate().toString().padStart(2, '0')
  const hours = date.getHours().toString().padStart(2, '0')
  const minutes = date.getMinutes().toString().padStart(2, '0')
  const seconds = date.getSeconds().toString().padStart(2, '0')

  const weekdays = ['周日', '周一', '周二', '周三', '周四', '周五', '周六']
  const weekday = weekdays[date.getDay()]

  return `${year}-${month}-${day} ${weekday} ${hours}:${minutes}:${seconds}`
}

const fieldMap = {
  second: {
    cols: 5,
    items: genItems(0, 59, (val) => {
      return val < 10 ? `0${val}` : val
    }),
    prefix: ':',
    selection: '每秒钟'
  },
  minute: {
    cols: 5,
    items: genItems(0, 59, (val) => {
      return val < 10 ? `0${val}` : val
    }),
    prefix: ':',
    selection: '每分钟'
  },
  hour: {
    cols: 4,
    items: genItems(0, 23, (val) => {
      return val < 10 ? `0${val}` : val
    }),
    prefix: '的',
    selection: '每小时'
  },
  dayOfWeek: {
    cols: 1,
    items: genItems(0, 6, (val) => {
      const week = ['日', '一', '二', '三', '四', '五', '六']
      return `星期${week[val]}`
    }),
    prefix: '和',
    selection: '一周的每一天'
  },
  dayOfMonth: {
    cols: 4,
    items: genItems(1, 31, (val) => {
      return val + '日'
    }),
    prefix: '的',
    selection: '每日'
  },
  month: {
    cols: 2,
    items: genItems(1, 12, (val) => {
      return val + '月'
    }),
    prefix: '的',
    selection: '每月'
  }
}

const fieldGroup = {
  second: [],
  minute: ['second'],
  hour: ['minute', 'second'],
  day: ['hour', 'minute', 'second'],
  week: ['dayOfWeek', 'hour', 'minute', 'second'],
  month: ['dayOfMonth', 'dayOfWeek', 'hour', 'minute', 'second'],
  year: ['month', 'dayOfMonth', 'dayOfWeek', 'hour', 'minute', 'second']
}

const period = ref(['year'])
const periodItems = [
  { key: 'second', label: '秒' },
  { key: 'minute', label: '分' },
  { key: 'hour', label: '时' },
  { key: 'day', label: '日' },
  { key: 'week', label: '周' },
  { key: 'month', label: '月' },
  { key: 'year', label: '年' }
]

const fields = ref([])
const crontab = ref('* * * * * *')
const crontabRunTimes = ref('')

const updatePeriod = () => {
  fields.value = []

  const field = fieldGroup[period.value[0]]
  for (const f of field) {
    fields.value.push({
      key: f,
      value: [],
      ...fieldMap[f]
    })
  }
}

watch(
  fields,
  () => {
    const res = ['*', '*', '*', '*', '*', '*']
    const match = {
      second: 0,
      minute: 1,
      hour: 2,
      dayOfMonth: 3,
      month: 4,
      dayOfWeek: 5
    }

    // 需要排序
    for (let i = 0; i < fields.value.length; i++) {
      const value = fields.value[i].value
      value.sort()
      fields.value[i].value = value
    }

    for (let i = fields.value.length - 1; i > -1; i--) {
      const key = fields.value[i].key
      const value = fields.value[i].value
      if (value.length > 0) {
        let index = -1
        let valueFormat = []
        for (const v of value) {
          if (valueFormat.length === 0) {
            index++
            valueFormat.push(v)
          } else {
            const last = valueFormat[index]
            if ((last + '').indexOf('-') > 0) {
              const [first, second] = last.split('-')
              if (parseInt(second) === v - 1) {
                valueFormat[index] = `${first}-${v}`
              } else {
                index++
                valueFormat.push(v)
              }
            } else {
              if (parseInt(last) === v - 1) {
                valueFormat[index] = `${last}-${v}`
              } else {
                index++
                valueFormat.push(v)
              }
            }
          }
        }
        res[match[key]] = valueFormat.join(',')
      }
    }
    crontab.value = res.join(' ')

    // 计算最近5次运行时间
    try {
      const interval = parser.parseExpression(crontab.value, {
        iterator: true,
        currentDate: new Date()
      })
      const times = []
      for (let i = 0; i < 5; i++) {
        const { value } = interval.next()
        times.push(formatDate(value.toDate()))
      }
      crontabRunTimes.value = times.join('\n')
    } catch (e) {
      crontabRunTimes.value = '无效的Crontab表达式'
    }
  },
  {
    deep: true
  }
)

onMounted(() => {
  updatePeriod()
})
</script>

<template>
  <div ref="container" flex flex-col flex-gap-20 wh-full overflow-auto p-20 class="cus-scroll">
    <div flex items-center>
      <span flex justify-center w-48>每</span>
      <select-popover :items="periodItems" v-model:value="period" @update:value="updatePeriod" />
      <template v-for="(item, i) in fields" :key="i">
        <span flex justify-center w-48>{{ item.prefix }}</span>
        <select-popover
          :items="item.items"
          :cols="item.cols"
          :multiple="true"
          :selection="item.selection"
          v-model:value="item.value"
          flex-1
        />
      </template>
    </div>
    <n-input-group>
      <n-input-group-label w-120>Crontab</n-input-group-label>
      <input-copy :value="crontab" />
    </n-input-group>
    <div>最近5次运行时间</div>
    <n-input type="textarea" v-model:value="crontabRunTimes" :rows="6" readonly />
  </div>
</template>

<style scoped lang="scss"></style>
