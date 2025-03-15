<script setup>
import { ChooseRuleConfig, ChooseContentFilterConfig, GetConfig, UpdateConfig } from '@/api/magnet.js'

const bak = ref()
const model = ref({
  preload: null,
  ruleConfig: null,
  cacheResult: null,
  cacheTimeout: null,
  contentFilter: null,
  proxyHost: null,
  proxyPort: null,
  proxyUser: null,
  proxyPass: null,
  enableProxy: null
})

const fetchConfig = async () => {
  const { msg, data, success } = await GetConfig()

  if (!success) {
    window.$message.error(msg)
    return
  }
  bak.value = data
  model.value = data
}

const chooseRuleConfig = async () => {
  const { msg, data, success } = await ChooseRuleConfig()

  if (!success) {
    window.$message.error(msg)
    return
  }
  model.value.ruleConfig = data
}

const chooseContentFilterConfig = async () => {
  const { msg, data, success } = await ChooseContentFilterConfig()

  if (!success) {
    window.$message.error(msg)
  }
  model.value.contentFilter = data
}

const reset = () => {
  model.value = toRaw(bak.value)
}

const update = async () => {
  const { msg, success } = await UpdateConfig(model.value)

  if (!success) {
    window.$message.error(msg)
    return
  }
  await fetchConfig()
}

onMounted(() => {
  fetchConfig()
})
</script>

<template>
  <div ref="container" flex flex-col flex-gap-20 wh-full overflow-auto p-20>
    <n-space>
      <n-button @click="reset">重置</n-button>
      <n-button type="primary" @click="update">保存</n-button>
    </n-space>
    <n-grid x-gap="20" :cols="2" h-full>
      <n-grid-item>
        <n-card title="基础配置" h-full>
          <n-form-item label="是否开启预加载">
            <n-switch v-model:value="model.preload" :round="false" :checked-value="1" :unchecked-value="0" />
          </n-form-item>
          <n-form-item label="自定义规则配置">
            <n-input-group>
              <n-input v-model:value="model.ruleConfig" readonly />
              <n-button type="primary" @click="chooseRuleConfig" ghost>选择</n-button>
            </n-input-group>
          </n-form-item>
          <n-form-item label="是否缓存搜索结果">
            <n-switch v-model:value="model.cacheResult" :round="false" :checked-value="1" :unchecked-value="0" />
          </n-form-item>
          <n-form-item label="缓存超时时间">
            <n-input-number v-model:value="model.cacheTimeout" w-full>
              <template #suffix>ms</template>
            </n-input-number>
          </n-form-item>
          <n-form-item label="内容过滤配置">
            <n-input-group>
              <n-input v-model:value="model.contentFilter" readonly />
              <n-button type="primary" @click="chooseContentFilterConfig" ghost>选择</n-button>
            </n-input-group>
          </n-form-item>
        </n-card>
      </n-grid-item>
      <n-grid-item>
        <n-card title="代理配置" h-full>
          <n-form-item label="是否启用代理">
            <n-switch v-model:value="model.enableProxy" :round="false" :checked-value="1" :unchecked-value="0" />
          </n-form-item>
          <n-form-item label="代理地址">
            <n-input v-model:value="model.proxyHost" :disabled="model.enableProxy === 0" />
          </n-form-item>
          <n-form-item label="代理端口">
            <n-input-number v-model:value="model.proxyPort" :disabled="model.enableProxy === 0" w-full />
          </n-form-item>
          <n-form-item label="代理账号">
            <n-input v-model:value="model.proxyUser" :disabled="model.enableProxy === 0" />
          </n-form-item>
          <n-form-item label="代理密码">
            <n-input v-model:value="model.proxyPass" :disabled="model.enableProxy === 0" />
          </n-form-item>
        </n-card>
      </n-grid-item>
    </n-grid>
  </div>
</template>

<style scoped lang="scss"></style>
