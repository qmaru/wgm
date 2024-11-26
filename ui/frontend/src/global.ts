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
        vertical: 'bottom',
        horizontal: 'center',
    }
}

// API 地址
window.api = "http://127.0.0.1:30000/api/v1"
window.version = "20241126"

export { }
