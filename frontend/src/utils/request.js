import axios from 'axios'
import { ElMessage } from 'element-plus'

const service = axios.create({
    baseURL: import.meta.env.VITE_BASE_URL,
    timeout: 5000 // request timeout
})

// request interceptor
service.interceptors.request.use(
    config => {
        let token = localStorage.getItem("token")
        if (token !== undefined) {
            config.headers['AccessToken'] = token
        }
        return config
    },
    error => {
        console.log(error)
        return Promise.reject(error)
    }
)

// response interceptor
service.interceptors.response.use(
    response => {
        const res = response.data

        if (res.code !== 200) {
            ElMessage({
                message: res.msg || 'Error',
                type: 'error',
                duration: 5 * 1000
            })
            return Promise.reject(new Error(res.msg || 'Error'))
        } else {
            return res
        }
    },
    error => {
        console.log('err' + error) // for debug
        ElMessage({
            message: error.message,
            type: 'error',
            duration: 5 * 1000
        })
        return Promise.reject(error)
    }
)

export default service