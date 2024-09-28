package constraints

const ContentTypeKey = "Content-Type"
const ContentTypeValueJson = "application/json; charset=utf-8"

const reqRespHeaderPrefix = "X-HttpBin-"
const SettingProxyTo = reqRespHeaderPrefix + "Proxy-To"
const SettingResponseStatus = reqRespHeaderPrefix + "Status"
const SettingSleep = reqRespHeaderPrefix + "Sleep"
