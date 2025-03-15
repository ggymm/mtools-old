import request from '@/api/dev/req.js'

export function BcryptEncode(params) {
  return request.post('', {
    method: 'CryptoService.BcryptEncode',
    params: JSON.stringify(params)
  })
}

export function BcryptCompare(params) {
  return request.post('', {
    method: 'CryptoService.BcryptCompare',
    params: JSON.stringify(params)
  })
}
