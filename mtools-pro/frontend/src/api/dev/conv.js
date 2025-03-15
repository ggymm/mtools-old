import request from '@/api/dev/req.js'

export function NumberBase(params) {
  return request.post('', {
    method: 'ConvService.NumberBase',
    params: JSON.stringify(params)
  })
}
