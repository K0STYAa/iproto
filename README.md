# RPC API VK iproto
[![golang-pipeline](https://github.com/K0STYAa/vk_iproto/actions/workflows/push.yml/badge.svg?branch=main)](https://github.com/K0STYAa/vk_iproto/actions/workflows/push.yml)

Реализован `iproto`-сервер, который осуществляет операции над стораджем.

### Схема протокола сервера:
```
<packet> ::= <request> | <response>
<request> ::= <header><body>
<response> ::= <header><return_code><body>
<header> ::= <func_id><body_length><request_id>
<func_id> ::= <uint32> - идентификатор вызываемой функции (см. раздел API)
<body_length> ::= <uint32> - длина тела запроса
<request_id> ::= <uint32> - идентификатор запроса, возвращается в ответе, нужен для асинхронной работы с сервером
<return_code> ::= <uint32> - код ответа (см. раздел API)
<body> ::= <byte>... - закодированная в MsgPack (см. раздел API) последовательность байт длиной <body_length>
```

### Формат стораджа:
```
Data            [1000]string
StorageState    uint8
Mutex           sync.RWMutex
```
Сторадж может находится в следующих состояних:
- `READ_ONLY` - доступен только на чтение
- `READ_WRITE` - доступен на чтение и запись
- `MAINTENANCE` - сторадж недоступен

`return_code`:
- `0` - успех, в этом случае клиент получает в `body` результат, описанный в таблице ниже
- `!= 0` - ошибка, в этом случае клиент получает в `body` msgpack-encoded строку - текстовое описание ошибки.

Тела запросов и ответов закодированы в формате MsgPack. Спецификация формата:
https://github.com/msgpack/msgpack/blob/master/spec.md

### Разобранны следующие операции над стораджем:
`func_id`    | Имя обработчика                  | Схема тела запроса | Схема тела ответа | Описание
------------ | -------------------------------- | ------------------ | ----------------- | --------
`0x00010001` | `ADM_STORAGE_SWITCH_READONLY`    | `<nil>`            | `<nil>`           | переводит сторадж в состояние `READ_ONLY`
`0x00010002` | `ADM_STORAGE_SWITCH_READWRITE`   | `<nil>`            | `<nil>`           | переводит сторадж в состояние `READ_WRITE`
`0x00010003` | `ADM_STORAGE_SWITCH_MAINTENANCE` | `<nil>`            | `<nil>`           | переводит сторадж в состояние `MAINTENANCE`
`0x00020001` | `STORAGE_REPLACE`                | `<int><string>`    | `<nil>`           | записывает в сторадж строку по индексу
`0x00020002` | `STORAGE_READ`                   | `<int>`            | `<string>`        | возвращает строку из стораджа по индексу

### LOG_LEVEL:
- DEBUG
- INFO
- WARNING
- ERROR
- FATAL
- PANIC

### Connection timeout = 30 sec

## Набор метрик:
### Встроенные в prometheus метрики:
1. Мониторинг горутин: количество запущенных горутин (go_goroutines)
2. Метрики производительности: использование ресурсов (например, памяти и CPU) (process_cpu_seconds_total, process_open_fds, process_start_time_seconds, process_virtual_memory_bytes, process_virtual_memory_max_bytes, etc)
### Собственные бизнес-метрики:
3. Метрики профилирования: количество успешно отработанных запросов для чтения/записи (storage_successful_reads_total, storage_successful_writes_total)
4. Метрики стабильности: количество ошибок и сбоев для чтения/записи (storage_error_reads_total, storage_error_writes_total)
5. Метрики распределения запросов: распределение запросов по эндпоинтам (storage_reads_total, storage_writes_total, change_storage_state_on_read_write_total, change_storage_state_on_read_only_total, change_storage_state_on_maintenance_total)
6. Метрики количества подключений (active_connections)

P.S. Я считаю подсчет изменений состояния отдельно для ReadOnly, ReadWrite, Maintenance, т.к. в случае реализации метрики как флага состояния, если между сборами prometheus при больших значениях скраппинг интервала в API состояние изменилось и потом вернулось обратно, то prometheus не зафиксирует изменение состояния.

Полный набор метрик можно посмотреть http://localhost:8088/metrics

## Для запуска сервера:
Build image and run container:
```
docker build -t vk_iproto .
docker run -p 80:8080 -e "LOG_LEVEL=<LOG_LEVEL>" vk_iproto
```
Or pull image from docker hub and than run:
```
docker pull k0styaa/vk_iproto
docker run -p 80:8080 -e "LOG_LEVEL=<LOG_LEVEL>" vk_iproto
```

### Вызов удаленной процедуры у стораджа:
```
type Header struct {
    FuncID     uint32
    BodyLength uint32
    RequestID  uint32
}

type Request struct {
    Header Header
    Body   []byte
}

type Response struct {
    Header     Header
    ReturnCode uint32
    Body       []byte
}

req := Request {...}
var resp models.Response

client, err := rpc.Dial("tcp", "localhost:80")
err := client.Call("MyService.MainHandler", request, &response)
```
