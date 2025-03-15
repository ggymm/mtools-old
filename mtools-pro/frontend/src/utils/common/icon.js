import { h } from 'vue'

import { Icon } from '@iconify/vue'
import { NIcon } from 'naive-ui'

import SvgIcon from '@/components/icon/SvgIcon.vue'

export function renderIcon(icon, props = { size: 12 }) {
  return () => {
    return h(NIcon, props, {
      default: () => {
        return h(Icon, { icon })
      }
    })
  }
}

export function renderCustomIcon(icon, props = { size: 12 }) {
  return () => {
    return h(NIcon, props, {
      default: () => {
        return h(SvgIcon, { icon })
      }
    })
  }
}
