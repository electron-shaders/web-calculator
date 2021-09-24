import request from '../utils/request'

export const send = () => {
    return request({
        url: 'https://yesno.wtf/api',
        method: 'get'
    })
}
