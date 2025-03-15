import request from './req'

export function GetCatalogTree() {
  const params = {}
  return request.post('', {
    method: 'SnippetService.GetCatalogTree',
    params: JSON.stringify(params)
  })
}

export function CreateCatalogItem(params) {
  return request.post('', {
    method: 'SnippetService.CreateCatalogItem',
    params: JSON.stringify(params)
  })
}

export function GetPage(params) {
  return request.post('', {
    method: 'SnippetService.GetPage',
    params: JSON.stringify(params)
  })
}

export function GetItem(params) {
  return request.post('', {
    method: 'SnippetService.GetItem',
    params: JSON.stringify(params)
  })
}

export function CreateItem(params) {
  return request.post('', {
    method: 'SnippetService.CreateItem',
    params: JSON.stringify(params)
  })
}

export function UpdateItem(params) {
  return request.post('', {
    method: 'SnippetService.UpdateItem',
    params: JSON.stringify(params)
  })
}

export function DeleteItem(params) {
  return request.post('', {
    method: 'SnippetService.DeleteItem',
    params: JSON.stringify(params)
  })
}
