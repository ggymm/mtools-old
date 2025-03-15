const Layout = () => import('@/layout/index.vue')

export default {
  name: 'Tools',
  path: '/tools',
  component: Layout,
  meta: {
    title: '工具库',
    icon: 'ic:baseline-menu'
  },
  children: [
    {
      name: 'Tools_Conv',
      path: 'conv',
      component: () => import('@/views/tools/conv/index.vue'),
      meta: {
        title: '转换',
        icon: 'ic:baseline-menu'
      },
      children: [
        {
          name: 'Tools_Conv_Url',
          path: 'url',
          component: () => import('@/views/tools/conv/url/index.vue'),
          meta: {
            title: 'Url编码',
            icon: 'ic:baseline-menu'
          }
        },
        {
          name: 'Tools_Conv_Binary',
          path: 'binary',
          component: () => import('@/views/tools/conv/binary/index.vue'),
          meta: {
            title: '进制转换',
            icon: 'ic:baseline-menu'
          }
        },
        {
          name: 'Tools_Conv_Base64',
          path: 'base64',
          component: () => import('@/views/tools/conv/base64/index.vue'),
          meta: {
            title: 'Base64编码',
            icon: 'ic:baseline-menu'
          }
        },
        {
          name: 'Tools_Conv_Datetime',
          path: 'datetime',
          component: () => import('@/views/tools/conv/datetime/index.vue'),
          meta: {
            title: 'Datetime转换',
            icon: 'ic:baseline-menu'
          }
        }
      ]
    },
    {
      name: 'Tools_Crypto',
      path: 'crypto',
      component: () => import('@/views/tools/crypto/index.vue'),
      meta: {
        title: '加密解密',
        icon: 'ic:baseline-menu'
      },
      children: [
        {
          name: 'Tools_Crypto_Hash',
          path: 'hash',
          component: () => import('@/views/tools/crypto/hash/index.vue'),
          meta: {
            title: 'Hash编码',
            icon: 'ic:baseline-menu'
          }
        },
        {
          name: 'Tools_Crypto_Bcrypt',
          path: 'bcrypt',
          component: () => import('@/views/tools/crypto/bcrypt/index.vue'),
          meta: {
            title: 'Bcrypt加密',
            icon: 'ic:baseline-menu'
          }
        }
      ]
    },
    {
      name: 'Tools_Develop',
      path: 'develop',
      component: () => import('@/views/tools/develop/index.vue'),
      meta: {
        title: '开发工具',
        icon: 'ic:baseline-menu'
      },
      children: [
        {
          name: 'Tools_Develop_Cron',
          path: 'cron',
          component: () => import('@/views/tools/develop/cron/index.vue'),
          meta: {
            title: 'Cron表达式',
            icon: 'ic:baseline-menu'
          }
        },
        {
          name: 'Tools_Develop_Qrcode',
          path: 'qrcode',
          component: () => import('@/views/tools/develop/qrcode/index.vue'),
          meta: {
            title: '二维码工具',
            icon: 'ic:baseline-menu'
          }
        }
      ]
    }
  ]
}
