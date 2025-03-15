<script setup>
import { NTag, NSpace, NButton } from 'naive-ui'

import { renderIcon } from '@/utils/index.js'
import { useWindowResize } from '@/hooks/event/window-resize.js'

import {
  GetCatalogTree,
  CreateCatalogItem,
  GetPage,
  GetItem,
  CreateItem,
  UpdateItem,
  DeleteItem
} from '@/api/snippet.js'

import InputEdit from '@/components/input-edit/index.vue'

import languages from '@/components/monaco-editor/language.js'
import MonacoEditor from '@/components/monaco-editor/index.vue'

const catalogTreeData = ref()
const catalogTreeNode = ref()
const catalogTreeLoading = ref(false)

const snippetTableData = ref([])
const snippetTablePagination = reactive({
  page: 1,
  pageSize: 10,
  showSizePicker: true,
  pageSizes: [10, 20, 50],
  itemCount: 0
})
const snippetTableColumns = [
  {
    title: '名称',
    key: 'title',
    width: 240
  },
  {
    title: '更新时间',
    key: 'updateAt',
    width: 180,
    align: 'center'
  },
  {
    title: '操作',
    key: 'actions',
    width: 160,
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
                  await handleSnippetUpdate(row)
                }
              },
              { default: () => '编辑' }
            ),
            h(
              NButton,
              {
                size: 'small',
                type: 'primary',
                onClick: async () => {
                  await handleSnippetDelete(row)
                }
              },
              { default: () => '删除' }
            )
          ]
        }
      )
    }
  }
]

const snippetTableSearch = ref('')
const snippetTableLoading = ref(false)

const fetchCatalogTreeData = async () => {
  const { msg, data, success } = await GetCatalogTree()
  if (!success) {
    window.$message.error(msg)
    return
  }
  // tree 赋值
  catalogTreeData.value = data
}

const selectCatalogTreeNode = (keys, option, meta) => {
  if (meta['action'] === 'select') {
    catalogTreeNode.value = option[0]
    fetchSnippetTableData()
  }
}

const fetchSnippetTableData = async () => {
  snippetTableLoading.value = true
  const { msg, data, success } = await GetPage({
    page: snippetTablePagination.page,
    size: snippetTablePagination.pageSize,
    catalogId: catalogTreeNode.value?.key
  })
  snippetTableLoading.value = false

  if (!success) {
    window.$message.error(msg)
    return
  }
  // table 赋值
  const { list, total } = data
  snippetTableData.value = list
  snippetTablePagination.itemCount = total
}

const handleCreateCatalog = async () => {
  catalogModel.value = true
}

const handleSnippetTableSearch = () => {
  fetchSnippetTableData()
}

const updateSnippetTablePage = (page) => {
  snippetTablePagination.page = page
  fetchSnippetTableData()
}

const updateSnippetTablePageSize = (pageSize) => {
  snippetTablePagination.pageSize = pageSize
  fetchSnippetTableData()
}

const handleSnippetCreate = () => {
  clearSnippetModel()
  initSnippetFragment()
  snippetDrawer.value = true
}

const handleSnippetUpdate = async (row) => {
  const { msg, data, success } = await GetItem({
    id: row.id
  })

  if (!success) {
    window.$message.error(msg)
    return
  }
  snippetModel.value = {
    id: data.id,
    title: data.title,
    fragments: data.fragments
  }
  fragmentIndex.value = 0
  snippetDrawer.value = true
}

const handleSnippetDelete = async (row) => {
  window.$dialog.confirm({
    title: '提示',
    content: '确定删除该条目？',
    confirm: async () => {
      const { msg, success } = await DeleteItem({
        id: row.id
      })
      if (!success) {
        window.$message.error(msg)
        return
      }
      // 重新加载数据
      await fetchSnippetTableData()
    }
  })
}

const catalogModel = ref(false)
const catalogSaveModel = ref({
  id: null,
  name: null
})
const catalogSaveLoading = ref(false)

const saveCatalog = async () => {
  catalogSaveLoading.value = true
  let parentId = 0
  if (catalogTreeNode.value) {
    parentId = catalogTreeNode.value.key
  }
  if (!catalogSaveModel.value.id) {
    const { msg, success } = await CreateCatalogItem({
      name: catalogSaveModel.value.name,
      parentId: parentId
    })
    catalogSaveLoading.value = false

    if (!success) {
      window.$message.error(msg)
      return
    }
    // 关闭抽屉
    catalogModel.value = false
    // 刷新表格数据
    await fetchCatalogTreeData()
  } else {
    // 更新
  }
}

// fragment
// fragment.title
// fragment.content
// fragment.language
const snippetModel = ref({
  id: null,
  title: null,
  fragments: []
})
const snippetDrawer = ref(false)
const snippetSaveLoading = ref(false)

const fragmentIndex = ref(0)
const fragmentTitles = computed(() => {
  let val = []
  snippetModel.value.fragments.forEach((item) => {
    val.push(item['title'])
  })
  return val
})

const saveSnippet = async () => {
  snippetSaveLoading.value = true
  if (!snippetModel.value.id) {
    const { msg, success } = await CreateItem({
      ...snippetModel.value,
      catalogId: catalogTreeNode.value?.key
    })
    snippetSaveLoading.value = false

    if (!success) {
      window.$message.error(msg)
      return
    }
    // 关闭抽屉
    snippetDrawer.value = false
    // 刷新表格数据
    await fetchSnippetTableData()
    window.$message.success('创建成功')
  } else {
    const { msg, success } = await UpdateItem(snippetModel.value)
    snippetSaveLoading.value = false

    if (!success) {
      window.$message.error(msg)
      return
    }
    // 关闭抽屉
    snippetDrawer.value = false
    // 刷新表格数据
    await fetchSnippetTableData()
  }
}

const initSnippetFragment = () => {
  fragmentIndex.value = 0
  createSnippetFragment('untitled')
}

const createSnippetFragment = (label) => {
  snippetModel.value.fragments.push({
    title: label,
    content: '',
    language: 'plaintext'
  })
  return label
}

const renderTreeNodeIcon = () => {
  return h(renderIcon('mdi:folder', { size: 18 }))
}

const renderFragmentTitle = (tag, index) => {
  return h(
    NTag,
    {
      type: index === fragmentIndex.value ? 'info' : '',
      size: 'large',
      closable: true,
      onClose: () => {
        snippetModel.value.fragments.splice(index, 1)
        const fragments = snippetModel.value.fragments
        if (fragments.length === 0) {
          initSnippetFragment()
        } else {
          if (index === fragmentIndex.value) {
            fragmentIndex.value = Math.max(0, index - 1)
          } else if (index < fragmentIndex.value) {
            fragmentIndex.value = Math.max(0, fragmentIndex.value - 1)
          }
        }
      },
      onclick: () => {
        fragmentIndex.value = index
      },
      style: {
        cursor: 'pointer',
        padding: '0 20px'
      }
    },
    {
      default: () => tag
    }
  )
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

const clearSnippetModel = () => {
  snippetModel.value.id = null
  snippetModel.value.title = null
  snippetModel.value.fragments = []
}

onMounted(async () => {
  await fetchCatalogTreeData()
  await fetchSnippetTableData()
  await nextTick(() => {
    resetTableHeight()
  })
})
</script>

<template>
  <div ref="container" flex flex-row flex-gap-20 wh-full overflow-auto p-20>
    <div flex flex-col w-280 min-w-280 h-full overflow-auto bg-white>
      <div class="title">
        <span>文件夹</span>
        <n-button text type="primary" @click="handleCreateCatalog">添加子节点</n-button>
      </div>
      <div p-20 class="cus-scroll">
        <n-spin :show="catalogTreeLoading" min-h-120>
          <n-tree
            block-line
            default-expand-all
            :data="catalogTreeData"
            :render-prefix="renderTreeNodeIcon"
            :on-update:selected-keys="selectCatalogTreeNode"
          >
            <template #empty><span></span></template>
          </n-tree>
        </n-spin>
      </div>
    </div>
    <div flex flex-col flex-1 h-full p-20 bg-white class="max-w-[calc(100%-280px-20px)]">
      <div flex flex-row flex-justify-between p-b20>
        <div>
          <n-button type="primary" @click="handleSnippetCreate">添加</n-button>
        </div>
        <div flex flex-row flex-gap-12>
          <n-input v-model:value="snippetTableSearch" type="text" />
          <n-button type="primary" @click="handleSnippetTableSearch">搜索</n-button>
        </div>
      </div>
      <div ref="tableContainer">
        <n-data-table
          remote
          flex-height
          :scroll-x="580"
          :data="snippetTableData"
          :columns="snippetTableColumns"
          :pagination="snippetTablePagination"
          :loading="snippetTableLoading"
          :style="{ width: '100%', height: `${height}px` }"
          @update:page="updateSnippetTablePage"
          @update:pageSize="updateSnippetTablePageSize"
        />
      </div>
    </div>
    <n-modal
      preset="card"
      transform-origin="center"
      :style="{ width: '50%' }"
      :title="catalogSaveModel.id ? '编辑文件夹' : '创建文件夹'"
      :auto-focus="false"
      v-model:show="catalogModel"
    >
      <div>
        <n-input v-model:value="catalogSaveModel.name" placeholder="请输入文件夹名称" />
      </div>
      <template #footer>
        <div flex flex-row flex-gap-20 justify-end>
          <n-button @click="catalogModel = false">取消</n-button>
          <n-button type="primary" :loading="catalogSaveLoading" @click="saveCatalog">保存</n-button>
        </div>
      </template>
    </n-modal>
    <n-drawer v-model:show="snippetDrawer" :mask-closable="false" :auto-focus="false" width="75%">
      <n-drawer-content closable>
        <template #header>
          <InputEdit v-model:value="snippetModel.title" />
        </template>
        <template #footer>
          <div flex flex-justify-between w-full>
            <div w-240>
              <n-select v-model:value="snippetModel.fragments[fragmentIndex].language" :options="languages" />
            </div>
            <div flex flex-gap-12>
              <n-button @click="snippetDrawer = false">取消</n-button>
              <n-button type="primary" :loading="snippetSaveLoading" @click="saveSnippet">保存</n-button>
            </div>
          </div>
        </template>
        <div flex flex-col h-full>
          <div flex m-b-16>
            <n-dynamic-tags
              size="large"
              :value="fragmentTitles"
              :render-tag="renderFragmentTitle"
              @create="createSnippetFragment"
            />
          </div>
          <div style="height: calc(100% - 34px - 16px)">
            <MonacoEditor
              :language="snippetModel.fragments[fragmentIndex].language"
              v-model="snippetModel.fragments[fragmentIndex].content"
            />
          </div>
        </div>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>
