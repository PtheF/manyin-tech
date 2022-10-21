
import { Message, Notification } from "element-ui"

export const ErrorMessage = (msg) => {
    ShowMessage(msg, "error", 5)
}

export const OkMessage = (msg) => {
    ShowMessage(msg, "success", 5)
}

export const InfoMessage = (msg) => {
    ShowMessage(msg, "info", 5)
}

export const WarnMessage = (msg) => {
    ShowMessage(msg, "warning")
}

export const ShowMessage = (msg, type, duration) => {
    Message({
        message: msg,
        type: type,
        duration: duration * 1000
    })
}

export const ShowNotify = (title, msg, type) => {
    Notification({
        title: title,
        message: msg,
        type: type
    })
}

export const OkNotify = (title, msg) => {
    ShowNotify(title, msg, "success")
}

export const ErrorNotify = (title, msg) => {
    ShowNotify(title, msg, "error")
}

export const InfoNotify = (title, msg) => {
    ShowNotify(title, msg, "info")
}