### Этот репозиторий представляет собой выполненное тестовое задание.

### Задание
Задан CSV-файл (comma-separated values) с заголовком, в котором перечислены названия столбцов. Строки нумеруются
целыми положительными числами, необязательно в порядке возрастания. В ячейках CSV-файла могут храниться или целые числа, или выражения вида `= ARG1 OP ARG2`.

Где `ARG1` и `ARG2` – целые числа или адреса ячеек в формате `Имя_колонки` `Номер_строки`, а `OP` – арифметическая операция
из списка: `+, -, *, /`.

Требуется написать программу, которая читает произвольную CSV-форму из файла (количество строк и столбцов может быть любым), вычисляет значения ячеек, если это необходимо, и выводит получившуюся табличку в виде CSV-представления в консоль. Решением задания будет: файл или несколько файлов с исходным кодом программы на языке Go, инструкции по
сборке и тестовые примеры (количество тестов – на усмотрение разработчика).

### Логика приложения
1. Считываю данные из CSV-файла и добавляю их в хеш-таблицу вида `map[string]map[string]string`, где первый ключ - это название столбца, а второй - название строки, а получаемое значение - это значение соответствующей ячейки. Также добавляю названия столбцов и строк в очереди для дальнейшего сохранения порядка.
2. Начинаю рекурсивно высчитывать значения каждой ячейки в порядке очереди и добавляю их в результирующий слайс `[]string` из строк таблицы.
3. Печатаю результат в консоль


### Инструкция по запуску
- добавить тестовые файлы в директорию проекта `data` 
- запустить в директории проекта `cmd/app` команду `go run . filename`, где `filename` - это имя тестируемого файла

### Пример корректной работы
```shell
ni@ni-asus:~/GolandProjects/yadroTest/cmd/app$ cat ../../data/correctTest.csv
,A,B,C,D
1,1,0,1,=A1*C30
2,2,=A1+C30,0,=D1+B2
30,0,=B1+A1,5,=B2-D1

ni@ni-asus:~/GolandProjects/yadroTest/cmd/app$ go run . correctTest.csv
,A,B,C,D
1,1,0,1,5
2,2,6,0,11
30,0,1,5,1
```

### Пример некорректной работы
* Повторяющееся название столбца
```shell
ni@ni-asus:~/GolandProjects/yadroTest/cmd/app$ cat ../../data/dublicateField.csv 
,A,B,C,D,B
1,1,0,1,=A1*C30,1
2,2,=A1+C30,0,=D1+B2,1
30,0,=B1+A1,5,=B2-D1,1

ni@ni-asus:~/GolandProjects/yadroTest/cmd/app$ go run . dublicateField.csv
2023/01/26 13:04:56 check CSV-file: duplicate field 'B' found
exit status 1
```
* Повторяющееся название строки
```shell
ni@ni-asus:~/GolandProjects/yadroTest/cmd/app$ cat ../../data/dublicateRow.csv 
,A,B,C,D,E
1,1,0,1,=A1*C30,1
2,2,=A1+C30,0,=D1+B2,1
30,0,=B1+A1,5,=B2-D1,1
2,1,1,1,1,1

ni@ni-asus:~/GolandProjects/yadroTest/cmd/app$ go run . dublicateRow.csv
2023/01/26 13:08:13 check CSV-file: duplicate row '2' found
exit status 1
```
* Несуществующий адрес ячейки
```shell
ni@ni-asus:~/GolandProjects/yadroTest/cmd/app$ cat ../../data/incorrectCell.csv 
,A,B,C,D,E
1,1,0,1,=A1*C30,=E5+7
2,2,=A1+C30,0,=D1+B2,1
30,0,=B1+A1,5,=B2-D1,1

ni@ni-asus:~/GolandProjects/yadroTest/cmd/app$ go run . incorrectCell.csv
2023/01/26 13:10:42 check CSV-file: cell [field:E row:5] does not exist
exit status 1
```
* Бесконечная рекурсия при подсчете выражения в ячейках
```shell
ni@ni-asus:~/GolandProjects/yadroTest/cmd/app$ cat ../../data/infinityRecursion.csv 
,A,B,Cell,D
1,1,0,1,=D1+D1
2,2,=A1+Cell30,0,1
30,0,=B1+A1,5,1

ni@ni-asus:~/GolandProjects/yadroTest/cmd/app$ go run . infinityRecursion.csv
2023/01/26 13:13:35 end by timeout, maybe self-pointer value in data, check you CSV-file
exit status 1
```
