# go-webcalc
# Arithmetic Expression Calculator
A simple Go web service to evaluate arithmetic expressions. It supports addition (`+`), subtraction (`-`), multiplication (`*`), division (`/`), and parentheses().

1. Склонируйте репозиторий
Откройте терминал и выполните команду:
git clone -b main https://github.com/<username>/go-webcalc.git
cd go-webcalc
Замените <username> на имя пользователя, которому принадлежит репозиторий.

2. Перейдите в директорию с файлом webcalc
Убедитесь, что файл webcalc находится в рабочей директории. Используйте команду:
ls
Если файл есть, переходите к следующему шагу.

3. Запустите проект
Выполните команду:
go run ./webcalc
Эта команда соберёт и запустит файл webcalc.


Чтобы запустить проект для репозитория go-webcalc, ветка main, файл webcalc, выполните следующие шаги:

1. Склонируйте репозиторий
Откройте терминал и выполните команду:

bash
Copy code
git clone -b main https://github.com/<username>/go-webcalc.git
cd go-webcalc
Замените <username> на имя пользователя, которому принадлежит репозиторий.

2. Перейдите в директорию с файлом webcalc
Убедитесь, что файл webcalc находится в рабочей директории. Используйте команду:

bash
Copy code
ls
Если файл есть, переходите к следующему шагу.

3. Запустите проект
Выполните команду:

bash
Copy code
go run ./webcalc
Эта команда соберёт и запустит ваш файл webcalc.

4. Проверьте работу сервиса
Сервис запустится на порту 8080. Вы увидите сообщение:
Server is running on port 8080
Теперь можно отправлять запросы к серверу. Например, с помощью curl:
Examples
Valid Request
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
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
