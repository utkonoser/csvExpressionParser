package csvParser

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"yadroTest/internal/usecase/implementQueue"
)

// csvTable - хеш-таблица для хранения представления CSV-файла
var csvTable = make(map[string]map[string]string)
var queue = implementQueue.Queue{}

// HashTableOfCsvFile - функция обрабатывающая входящие значения и добавляющая их представление в хеш-таблицу
func HashTableOfCsvFile(dataPath string) {
	file, err := os.Open(dataPath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	reader := csv.NewReader(file)
	reader.Comma = ','

	counter := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if counter == 0 {
			for _, fieldName := range record[1:] {
				if queue.InFields(fieldName) {
					log.Fatalf("check CSV-file: duplicate field '%s' found", fieldName)
				}
				queue.InsertField(fieldName)
				csvTable[fieldName] = make(map[string]string)
			}
			counter++
			continue
		}

		rowName := record[0]
		values := record[1:]

		if queue.InRows(rowName) {
			log.Fatalf("check CSV-file: duplicate row '%s' found", rowName)
		}
		queue.InsertRow(rowName)

		if len(values) != len(queue.Fields) {
			log.Fatal("length of values and fields must be equal ")
		}
		for i := 0; i < len(values); i++ {
			csvTable[queue.Fields[i]][rowName] = values[i]
		}
	}
}

// Expression - структура арифметического выражения
type expression struct {
	arg1      string
	arg2      string
	operation string
}

// Action - метод expression, вычисляющий выражение
func (e *expression) Action() (int, error) {
	arg1, err := strconv.Atoi(e.arg1)
	arg2, err := strconv.Atoi(e.arg2)
	if err != nil {
		return 0, err
	}
	switch e.operation {
	case "+":
		return arg1 + arg2, nil
	case "-":
		return arg1 - arg2, nil
	case "*":
		return arg1 * arg2, nil
	case "/":
		if arg2 == 0 {
			return 0, errors.New("division by zero is forbidden, incorrect values in the CSV file")
		}
		return arg1 / arg2, nil
	}
	return 0, nil
}

// returnValueOfExpression - функция, рекурсивно вычисляющая значение выражения
func returnValueOfExpression(str string) string {

	twoArgs, err := regexp.Compile("[=+*/-]")
	operation, err := regexp.Compile("[0-9,a-zA-Z=]")
	letters, err := regexp.Compile("[0-9,=+*/-]")
	numbers, err := regexp.Compile("[A-Za-z,=+*/-]")

	if err != nil {
		log.Fatal(err)
	}

	var sliceOfArgs []string
	args := twoArgs.Split(str, -1)
	for _, i := range args {
		if i == "" {
			continue
		}
		sliceOfArgs = append(sliceOfArgs, i)
	}

	exp := expression{
		arg1:      "",
		arg2:      "",
		operation: "",
	}

	var arguments []string
	for _, i := range sliceOfArgs {
		_, err := strconv.Atoi(i)
		if err == nil {
			arguments = append(arguments, i)
			continue
		}
		letter := letters.Split(i, -1)
		var field string
		for _, j := range letter {
			if j == "" {
				continue
			}
			field += j
		}

		num := numbers.Split(i, -1)
		var row string
		for _, j := range num {
			if j == "" {
				continue
			}
			row += j
		}

		value, ok := csvTable[field][row]
		if !ok {
			log.Fatalf("check CSV-file: cell [field:%s row:%s] does not exist", field, row)
		}
		arguments = append(arguments, value)

	}

	exp.arg1 = arguments[0]
	if exp.arg1[0] == '=' {
		exp.arg1 = returnValueOfExpression(exp.arg1)
	}
	exp.arg2 = arguments[1]
	if exp.arg2[0] == '=' {
		exp.arg2 = returnValueOfExpression(exp.arg2)
	}
	exp.operation = strings.Join(operation.Split(str, -1), "")

	number, err := exp.Action()
	if err != nil {
		log.Fatal(err)
	}
	res := strconv.Itoa(number)
	return res
}

// PrintResult - функция, печатающая в консоль результат выполнения программы
func PrintResult() {
	counter := 0
	var sliceOfResultRows []string
	for len(queue.Rows) != 0 {
		var items []string
		if counter == 0 {
			row := "," + strings.Join(queue.Fields, ",")
			sliceOfResultRows = append(sliceOfResultRows, row)
			counter++
			continue
		}

		rowName := queue.RemoveRow()
		items = append(items, rowName)
		for _, fieldName := range queue.Fields {
			val := csvTable[fieldName][rowName]
			if val[0] == '=' {
				item := returnValueOfExpression(val)
				items = append(items, item)
				continue
			}
			items = append(items, val)
		}
		row := strings.Join(items, ",")
		sliceOfResultRows = append(sliceOfResultRows, row)
	}

	for _, row := range sliceOfResultRows {
		fmt.Println(row)
	}
}
