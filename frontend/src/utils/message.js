import {ElMessageBox, ElMessage} from 'element-plus'

const clear = () => {
    ElMessage.closeAll()
}

/**
 * 消息提示
 * @param msg 提示信息
 * @param type 消息类型
 */
const message = (msg, type) => {
    ElMessage({
        showClose: true,
        message: msg,
        type: type
    });
}
/**
 * 成功提示
 * @param msg 提示信息
 */
const success = (msg) => {
    message(msg, 'success');
}
/**
 * 消息提示
 * @param msg 提示信息
 */
const info = (msg) => {
    message(msg, 'info');
}
/**
 * 错误提示
 * @param msg 提示信息
 */
const error = (msg) => {
    message(msg, 'error');
}

/**
 * 警告提示
 * @param msg 提示信息
 */
const warning = (msg) => {
    message(msg, 'warning');
}

/**
 * alert提示框
 * @param title 标题
 * @param msg 信息
 * @param ok ok函数
 * @param okText ok按钮文字
 */
const alert = ({title, msg, ok, okText}) => {
    if (!title) {
        title = '提示';
    }
    if (!msg) {
        msg = '错误';
    }
    if (!okText) {
        okText = '确定';
    }
    ElMessageBox.alert(msg, title, {
        confirmButtonText: okText,
    }).then(ok ? ok : () => {});
}

/**
 * confirm 提示框
 * @param title 标题
 * @param msg 信息
 * @param ok ok函数
 * @param okText ok按钮文字
 * @param cancel 取消函数
 * @param cText cancel按钮文字
 */
const confirm = ({title, msg, ok, okText, cancel, cText}) => {
    ElMessageBox.confirm(msg, title ? title:'提示', {
        confirmButtonText: okText ? okText:'确定',
        cancelButtonText: cText ? cText:'取消',
    }).then(ok ? ok : () => {}).catch(cancel ? cancel : () => {});
}

export default {
    success,
    info,
    error,
    warning,
    alert,
    confirm,
    clear,
}