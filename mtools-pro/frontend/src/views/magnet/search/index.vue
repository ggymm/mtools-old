<script setup>
import { NButton, NSpace } from 'naive-ui'
import { useClipboard } from '@vueuse/core'

import { useWindowResize } from '@/hooks/event/window-resize.js'

import { GetRuleOptions, RunSearchMagnet } from '@/api/magnet.js'

const copy = useClipboard().copy

const columns = [
  {
    title: '名称',
    key: 'name',
    ellipsis: {
      tooltip: true
    }
  },
  {
    title: '大小',
    key: 'size',
    width: 160,
    align: 'center'
  },
  {
    title: '热度',
    key: 'downloadHot',
    width: 160,
    align: 'center'
  },
  {
    title: '下载时间',
    key: 'downloadTime',
    width: 200,
    align: 'center'
  },
  {
    title: '操作',
    key: 'actions',
    width: 280,
    align: 'center',
    fixed: 'right',
    render(row) {
      return h(
        NSpace,
        { justify: 'center' },
        {
          default: () => [
            h(
              NButton,
              {
                size: 'small',
                type: 'primary',
                onClick: async () => {
                  await handleDownload(row)
                }
              },
              { default: () => '迅雷下载' }
            ),
            h(
              NButton,
              {
                size: 'small',
                type: 'primary',
                onClick: async () => {
                  await handleCopyToClipboard(row)
                }
              },
              { default: () => '复制磁力链接' }
            )
          ]
        }
      )
    }
  }
]

const tableData = ref([])
const pagination = reactive({
  page: 1,
  pageSize: 10,
  pageCount: 0,
  showSizePicker: false
})

const ruleOptions = ref([])

const tableSearch = ref({
  site: null,
  keywords: 'SSIS'
})
const tableLoading = ref(false)

const fetchTableData = async () => {
  tableLoading.value = true
  const { msg, data, success } = await RunSearchMagnet({
    page: pagination.page,
    size: pagination.pageSize,
    site: tableSearch.value.site,
    keywords: tableSearch.value.keywords
  })
  if (!success) {
    window.$message.error(msg)
    return
  }
  // table 赋值
  const { list, total } = data
  if (pagination.pageSize === 1 && total > 1) {
    // 如果是第一页，且总页数大于1
    // 设置每一页大小为列表长度
    pagination.pageSize = list.length
  }
  tableData.value = list
  pagination.pageCount = total
  tableLoading.value = false
}

const handleTableSearch = () => {
  // 判断是否选择了站点
  if (!tableSearch.value.site) {
    window.$message.error('请选择站点')
    return
  }
  fetchTableData()
}

const handleTablePageChange = (page) => {
  pagination.page = page
  fetchTableData()
}

const handleDownload = async (row) => {
  await copy(row['magnet'])
  window.$message.success('复制成功')
}

const handleCopyToClipboard = async (row) => {
  await copy(row['magnet'])
  window.$message.success('复制成功')
}

const height = ref(200)
const container = ref()
const tableContainer = ref()

const resetTableHeight = () => {
  if (!container.value || !tableContainer.value) return
  const padding = 40
  const { top: containerTop } = container.value.getBoundingClientRect()
  const { top: tableContainerTop } = tableContainer.value.getBoundingClientRect()
  height.value = container.value.clientHeight - (tableContainerTop - containerTop) - padding
}
useWindowResize(resetTableHeight, 30)

const fetchRuleOptions = async () => {
  const { msg, data, success } = await GetRuleOptions()

  if (!success) {
    window.$message.error(msg)
  }
  ruleOptions.value = data
  tableSearch.value.site = data[0].value
}

onMounted(() => {
  fetchRuleOptions()
  nextTick(() => {
    resetTableHeight()
  })
})
</script>

<template>
  <div ref="container" flex flex-col flex-gap-20 wh-full overflow-auto p-20>
    <div flex flex-row justify-between flex-gap-20>
      <div w-180>
        <n-select v-model:value="tableSearch.site" :options="ruleOptions" />
      </div>
      <div flex-1>
        <n-input v-model:value="tableSearch.keywords" type="text" />
      </div>
      <div>
        <n-button type="primary" @click="handleTableSearch" p-x-40>搜索</n-button>
      </div>
    </div>
    <div flex h-full p-20 bg-white>
      <div ref="tableContainer" bg-white>
        <n-data-table
          remote
          flex-height
          :scroll-x="580"
          :data="tableData"
          :columns="columns"
          :pagination="pagination"
          :loading="tableLoading"
          :style="{ width: '100%', height: `${height}px` }"
          @update:page="handleTablePageChange"
        />
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss"></style>
