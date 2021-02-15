# tz_backend

>Реализовано два режима работы: 
   > - cmd/allinone (сервер с тремя обработчиками-парсерами с каскадной балансировкой между ними)
   > - cmd/standalone (сервер с одним обраьботчиком-парсером, который имеет возможность передавать превышающие лимит пакеты на другой http-сервер)
>Для запуска программы необходимо перейти в соответствующую режиму работы директорию
>
>Для режима standalone для каждого сервера создан конфигурационный файл в папке cmd/standalone. При запуске программы необходимо указать в качестве параметра нужный файл
> Пример:
>    ```go run main.go -c serverconfig2.json```
>
>После запуска сервера, клиент для отправки сообщений доступен по адресу:
> localhost:8090/client/
>При открытии страницы выполнится скрипт js, который начнет генерировать пакеты сообщений и отправлять на сервер
