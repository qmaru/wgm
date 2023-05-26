declare global {
    interface Window {
        api: string
        version: string
        messageDefault: any
    }
}

// notistack 默认属性
window.messageDefault = {
    variant: "success",
    autoHideDuration: 2000,
    preventDuplicate: true,
    anchorOrigin: {
        vertical: 'top',
        horizontal: 'center',
    }
}

const marjorVer = "2"
const minorVer = "0"
const dateVer = "20230526"
// API 地址
window.api = "http://127.0.0.1:30000/api/v1"
window.version = marjorVer + "." + minorVer + "-" + dateVer

export { }
