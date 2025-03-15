import axios from 'axios'

axios.defaults.withCredentials = true

const service = axios.create({
  headers: { 'Content-Type': 'application/json' },
  baseURL: 'http://localhost:8080/api',
  timeout: 30000
})

service.interceptors.request.use(
  (config) => {
    return config
  },
  (error) => {
    console.log(error) // for debug
    return Promise.reject(error)
  }
)

service.interceptors.response.use(
  (response) => {
    return response.data
  },
  (error) => {
    return Promise.reject(error)
  }
)

export default service
