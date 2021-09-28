import {
    ElButton,
    ElMessage,
    ElAlert,
    ElTable,
    ElEmpty,
    ElTabPane,
    ElTabs,
    ElContainer
} from 'element-plus'
  
// 组件写在components中
export const components: any[] = [
    ElButton,
    ElAlert,
    ElTable,
    ElEmpty,
    ElTabPane,
    ElTabs,
    ElContainer    
]
// 插件写在plugins中
export const plugins: any[] = [ElMessage]