# go-webcalc
# Arithmetic Expression Calculator
A simple Go web service to evaluate arithmetic expressions. It supports addition (`+`), subtraction (`-`), multiplication (`*`), division (`/`), and parentheses().
Для запуска проекта выполните следующие шаги:
1. Склонируйте репозиторий
Откройте терминал и выполните команду:
```git clone -b main https://github.com/egocentri/go-webcalc.git```
```cd go-webcalc```

3. Перейдите в директорию с файлом webcalc
Убедитесь, что файл webcalc находится в рабочей директории. Используйте команду:
```ls```
Если файл есть, переходите к следующему шагу.

4. Запустите проект
Выполните команду:
```go run ./webcalc```
Эта команда соберёт и запустит файл webcalc.

5. Сервис запустится на порту 8080. Вы увидите сообщение:
Server is running on port 8080
Теперь можно отправлять запросы к серверу. Например, с помощью curl:
Examples
Valid Request
curl ```--location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'```
Response:
{
  "result": "6"
}
Invalid Request (422)
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2++2"
}'
Response:
{
  "error": "Expression is not valid"
}
