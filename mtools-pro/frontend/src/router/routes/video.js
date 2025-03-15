const Layout = () => import('@/layout/index.vue')

export default {
  name: 'Video',
  path: '/video',
  component: Layout,
  meta: {
    title: '视频工具箱',
    icon: 'ic:baseline-menu'
  },
  children: [
    {
      name: 'Video_Record',
      path: 'record',
      component: () => import('@/views/video/record/index.vue'),
      meta: {
        title: '录制视频',
        icon: 'ic:baseline-menu'
      }
    },
    {
      name: 'Video_Concat',
      path: 'concat',
      component: () => import('@/views/video/concat/index.vue'),
      meta: {
        title: '视频合并',
        icon: 'ic:baseline-menu'
      }
    },
    {
      name: 'Video_Division',
      path: 'division',
      component: () => import('@/views/video/division/index.vue'),
      meta: {
        title: '视频拆分',
        icon: 'ic:baseline-menu'
      }
    },
    {
      name: 'Video_Download',
      path: 'download',
      component: () => import('@/views/video/download/index.vue'),
      meta: {
        title: '视频下载',
        icon: 'ic:baseline-menu'
      }
    }
  ]
}
