package constraints

const HeaderContentType = "Content-Type"
const ContentTypeAppJsonUtf8 = "application/json; charset=utf-8"

const reqRespHeaderPrefix = "X-HttpBin-"
const SettingProxyTo = reqRespHeaderPrefix + "Proxy-To"
const SettingResponseStatus = reqRespHeaderPrefix + "Status"
const SettingSleep = reqRespHeaderPrefix + "Sleep"
const SettingContentType = reqRespHeaderPrefix + "ContentType"
