//引入axios
import axios from 'axios'
//引入登录拦截
import store/*, {TOKEN_HEADER}*/ from "../store";
//引入路由
//import router from "../router";
//引入定义链接
//import {reUrl, path} from './urls'
//引入qs
//import qs from 'qs'

/*
import {
    // eslint-disable-next-line no-unused-vars
    MessageBox,
    Loading
} from 'element-plus'
*/
import message from "./message";

//axios定义
axios.defaults.baseURL = 'http://localhost:3001';
//axios.defaults.headers['Content-Type'] = 'application/x-www-form-urlencoded;charset=UTF-8'
//axios.defaults.transformRequest = [object => qs.stringify(object)]

//axios.defaults.headers['Content-Type'] = 'application/json'
axios.defaults.headers['Content-Type'] = 'text/plain'
axios.defaults.timeout = 3000;
axios.defaults.withCredentials = true;
// axios request 拦截器
axios.interceptors.request.use(
    config => {
        //判断token是否存在
        if (store.state.token !== null) {
            //将token设置成请求头
            config.headers[TOKEN_HEADER] = store.state.token;
        }
        return config;
    },
    err => {
        return Promise.reject(err);
    }
);
// http response 拦截器
/*
axios.interceptors.response.use(
    res => {
        if (res.data.code === 102) {
            store.commit('delToken');
            store.commit('delUserInfo');
            router.push(path('/login.html')).then();
        } else if (res.data.code === 500 || res.data.code === 403) {
            message.alert({title: '错误', msg: res.data.msg});
        } else {
            return res;
        }
    },
    error => {
        return Promise.reject(error);
    }
);
*/

const request = (object) => {
    let loading = null;
    if (object.loading) {
        loading = Loading.service({
            lock: true,
            text: object.loading,
            spinner: 'el-icon-loading',
            background: 'rgba(0, 0, 0, 0.6)'
        });
    }
    if (!object.timeout) {
        object.timeout = 0;
    }
    axios({
        method: object.method,
        url: object.url,
        data: object.data
    }).then((res) => {
        setTimeout(() => {
            if (object.loading) {
                loading.close();
            }
            if (!res) {
                return;
            }
            const data = res.data;
            if (data.code !== 403 && data.code !== 500) {
                if (isFunction(object.success)) {
                    object.success(data);
                }
            }
        }, object.timeout);
    }).catch((error) => {
        console.log(error);
        setTimeout(() => {
            if (object.loading) {
                loading.close();
            }
            if (isFunction(object.error)) {
                object.error(error);
            } else {
                message.alert({
                    title: '错误', msg: '请求出错'
                });
            }
        }, object.timeout);
    });
}

const isFunction = (func) => {
    return func && typeof func == 'function'
}

const executeReq = (object) => {
    if (object.confirm) {
        message.confirm({
            msg: object.confirm,
            func: () => request(object)
        });
    } else {
        request(object);
    }
}

export default {
    /**
     *
     * @param url 链接
     * @param data 数据
     * @param loading 是否开启 loading
     * @param timeout loading的关闭时间 不提供则为 0
     * @param confirm 是否需要 confirm 提示
     * @param success 成功返回
     * @param error 失败返回
     */

    get({url, data, loading, timeout, confirm, success, error}) {
        const method = 'GET';
        if (data) {
            data = qs.stringify(JSON.parse(JSON.stringify(data)));
            url = url + "?" + data;
        }
        executeReq({method, url, data, loading, timeout, confirm, success, error});
    },
    post({url, data, loading, timeout, confirm, success, error}) {
        const method = 'POST';
        executeReq({method, url, data, loading, timeout, confirm, success, error});
    },
    put({url, data, loading, timeout, confirm, success, error}) {
        const method = 'PUT';
        executeReq({method, url, data, loading, timeout, confirm, success, error});
    },
    delete({url, data, loading, timeout, confirm, success, error}) {
        const method = 'DELETE';
        executeReq({method, url, data, loading, timeout, confirm, success, error});
    }
}
