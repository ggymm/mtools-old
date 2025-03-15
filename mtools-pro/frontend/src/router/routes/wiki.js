const Layout = () => import('@/layout/index.vue')

export default {
  name: 'Wiki',
  path: '/wiki',
  component: Layout,
  meta: {
    title: '知识库管理',
    icon: 'ic:baseline-menu'
  },
  children: [
    {
      name: 'Wiki_Doc',
      path: 'doc',
      component: () => import('@/views/wiki/doc/index.vue'),
      meta: {
        title: '文档集合',
        icon: 'ic:baseline-menu'
      }
    },
    {
      name: 'Wiki_Scripts',
      path: 'script',
      component: () => import('@/views/wiki/script/index.vue'),
      meta: {
        title: '脚本集合',
        icon: 'ic:baseline-menu'
      }
    },
    {
      name: 'Wiki_Snippet',
      path: 'snippet',
      component: () => import('@/views/wiki/snippet/index.vue'),
      meta: {
        title: '代码片段',
        icon: 'ic:baseline-menu'
      }
    }
  ]
}
