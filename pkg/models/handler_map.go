package models

func HandlerNameById(key uint32) string {
    // HandlerMap
	handlerMap := map[uint32] string {
        0x00010001: "ADM_STORAGE_SWITCH_READONLY",
        0x00010002: "ADM_STORAGE_SWITCH_READWRITE",
        0x00010003: "ADM_STORAGE_SWITCH_MAINTENANCE",
        0x00020001: "STORAGE_REPLACE",
        0x00020002: "STORAGE_READ",
    }

    return handlerMap[key]
}