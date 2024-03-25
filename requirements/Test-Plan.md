# Тест план

## Базовый тест-план ActDeviceApi

## 1. Создание устройства

### Тест 1 Регистрация устройства GRPC

Предварительные условия:
1) Система запущена.
2) Проходит Readiness probe

| Шаг | Действие                                                                                                  | Ожидаемый результат                                                                                                                                                                                                                            |
|-----|-----------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 1.  | Отправьте запрос CreateDeviceV1Request со следующими параметрами:<br/> `user_id ='666', platform = 'ios'` | Пришел ответ CreateDeviceV1Response с deviceId > 0 (**device_ID_1**)"<br/> * В БД появилась 1 запись: platform='iOS', user_id='666',enteredAt = 2022-10-14T12:07:30.397025Z removed =false, created_at >= entered_at |
| 2.  | Отправить запрос DescribeDeviceV1Request c device_id = **device_ID_1**                                    | Пришел ответ DescribeDeviceV1Response с device = {id = **device_ID_1**, platform='ios', user_id ='666', eteredAt = "2022-10-14T12:07:30.397025Z" }                                                                                             |                                                                                                                                
| 3.  | Отправить запрос ListDeviceV1Request c page = 1 perPage = 1                                               | Пришел ответ ListDeviceV1Response c телом {"items": [{"id": "11","platform": "iOS","userId": "666","enteredAt": "2022-10-14T12:07:30.397025Z"}]}                                                                                               |                                                                                                                                      

### Тест 2 Редактирование устройства GRPC

Предварительные условия:
1) Система запущена.
2) Проходит Readiness probe

| Шаг | Действие                                                                                                          | Ожидаемый результат                                                                                                                                                                                                                                                                                                                                                                 |
|-----|-------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 1.  | Отправьте запрос UpdateDeviceV1Request c device_id = **device_ID_1**,<br/> `user_id ='777', platform = 'Android'` | Получен ответ UpdateDeviceV1Response с Device = {success = True}  В БД появилась запись: platform='iOS', user_id='666',enteredAt = 2022-10-14T12:07:30.397025Z removed =false, created_at >= entered_at  Заменилась на platform='Android', user_id='777',enteredAt = 2022-10-14T12:07:30.397025Z removed =false, created_at >= entered_at, updated_at = Время обновления устройства |
| 2.  | Отправить запрос DescribeDeviceV1Request c device_id = **device_ID_1**                                            | Пришел ответ DescribeDeviceV1Response с Device = {id = **device_ID_1**, platform='Android', user_id ='777', etered_at = "2022-10-14T12:07:30.397025Z" }                                                                                                                                                                                                                             |
| 3.  | Отправить запрос ListDeviceV1Request c page = 1 perPage = 1                                                       | Пришел ответ ListDeviceV1Response c телом {"items": [{"id": "11","platform": "Android","userId": "777","enteredAt": "2022-10-14T12:07:30.397025Z"}]}                                                                                                                                                                                                                                | 

### Тест 3 Удаление устройства GRPC

Предварительные условия:
1) Система запущена.
2) Проходит Readiness probe

| Шаг | Действие                                                                | Ожидаемый результат                                                                                                    |
|-----|-------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------|
| 1.  | Отправьте запрос RemoveDeviceV1Request c device_id = **device_ID_1**    | Получен ответ RemoteDeviceV1Response с Device = {found = true} В БД поле removed значение false сменилось на true      |
| 2.  | ОТправить запрос DescribeDeviceV1Request c device_id = **device_ID_1**  | Пришел ответ DescribeDeviceV1Response с ошибкой 404 Not Found {"code": 5, "message": "device not found","details": []} |

### Тест 4 Список устройств

| Шаг | Действие                                                    | Ожидаемый результат                                                                                                                                 |
|-----|-------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|
| 1.  | Отправить запрос ListDeviceV1Request c page = 1 perPage = 1 | Пришел ответ ListDeviceV1Response c телом {"items": [{"id": "11","platform": "Android","userId": "777","enteredAt": "2022-10-14T12:07:30.397025Z"}] |
| 1.  | Отправить запрос ListDeviceV1Request c page = 2 perPage = 1 | Пришел ответ ListDeviceV1Response c телом {"items": [{"id": "10","platform": ...}]                                                                  |
| 1.  | Отправить запрос ListDeviceV1Request c page = 1 perPage = 2 | Пришел ответ ListDeviceV1Response c телом {"items": [{"id": "11","platform": ...}, {"id": "10", "platform": ...}]                                   |

## Базовый тест-план ActNotificationApi

### Таблица PairWise

Предварительные условия:
1) Система запущена.
2) Проходит Readiness probe

| Platform    | Language   | Period       |
|-------------|------------|--------------|
| Win         | EN         | Morning      |
| Win         | RU         | Day          |
| Win         | ES         | Everning     |
| Win         | IT         | Night        |
| Android     | RU         | Everning     |
| Android     | ES         | Night        |
| Android     | IT         | Morning      |
| Android     | EN         | Day          |
| IOS         | ES         | Morning      |
| IOS         | IT         | Day          |
| IOS         | EN         | Everning     |
| IOS         | RU         | Night        |
| Linux       | IT         | Everning     |
| Linux       | EN         | Night        |
| Linux       | RU         | Morning      |
| Linux       | ES         | Day          |


### Тест 1. Подписка на  получение уведомлений

| Шаг | Действие                                                                                                                                                                                                    | Ожидаемый результат                                                                                                                                                          |
|-----|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 1.  | Отправьте запрос SubscribeNotificationV1Request c deviceId = **10**                                                                                                                                         | Ожидание ответа от SendNotificationV1quest                                                                                                                                   |
| 2.  | Отправить запрос SendNotificationV1quest c {"notification": {"notificationId": "2","deviceId": "10","username": "Adscs","message": "Andrew","lang": "LANG_ENGLISH","notificationStatus": "STATUS_CREATED"}} | Пришел ответ SendNotificationV1Response с  телом {"notificationId": "2"} , Пришел ответ SubscribeNotificationV1Response с телом { "notificationId": "2", "message": "Andrew" |

### Тест 2. Получение списка уведомлениц по устройству

| Шаг | Действие                                                      | Ожидаемый результат                                                                                                          |
|-----|---------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------|
| 1.  | Отправить запрос GetNotificationV1Request c deviceId = **10** | Получен ответ GetNotificationV1Response с телом {"notification": [{"notificationId": "2","message": "Good morning Andrew"}]} |


### Тест 3. Отправка уведомлений на существующее устройство

| Шаг  | Действие                                                                                                                                                                                                    | Ожидаемый результат                                                                                                                                                                               |
|------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 1.   | Отправить запрос SendNotificationV1quest c {"notification": {"notificationId": "2","deviceId": "2","username": "ABBA","message": "Steve T","lang": "LANG_ENGLISH","notificationStatus": "STATUS_CREATED"}}  | Получен ответ GetNotificationV1Response с телом {"notification": [{"notificationId": "2","message": "Good morning Steve T"}]} Уведоление помещено в базу "notificationStatus": "STATUS_CREATED"   |
| 2.   | Отправить запрос SendNotificationV1quest c {"notification": {"notificationId": "3","deviceId": "2","username": "Aha","message": "Steve Vai","lang": "LANG_ENGLISH","notificationStatus": "STATUS_CREATED"}} | Получен ответ GetNotificationV1Response с телом {"notification": [{"notificationId": "3","message": "Good morning Steve Vai"}]} Уведоление помещено в базу "notificationStatus": "STATUS_CREATED" |
| 3.   | Активировать сессию SubscribeNotificationV1Request c deviceId = **2**                                                                                                                                       | Получить уведомления "notificationId": "2" и "notificationId": "3". В БД статус STATUS_DELIVERED                                                                                                  |


### Тест 3. Отправка уведомлений на несуществующее устройство

| Шаг  | Действие                                                                                                                                                                                                            | Ожидаемый результат         |
|------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------|
| 1.   | Отправить запрос SendNotificationV1quest c {"notification": {"notificationId": "2","deviceId": "not exist","username": "uri7","message": "Satriani","lang": "LANG_ENGLISH","notificationStatus": "STATUS_CREATED"}} | Получен сообщение об ошибке |

