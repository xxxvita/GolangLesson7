package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	EnterLens = "Введите размеры сторон фигуры через запятую:"
)

// Структура: окружность
type Circle struct {
	// Диаметр окружности
	d float64
}

func (c *Circle) CalcArea() float64 {
	return math.Pi * math.Pow(c.d/2.0, 2)
}

func (c *Circle) String() string {
	return "ОКРУЖНОСТЬ"
}

func (c *Circle) FillData(data []float64) (bool, error) {
	if data == nil {
		return false, fmt.Errorf("Не передан слайс данных в структуру %s\n", c)
	}

	if len(data) < 1 {
		return false, fmt.Errorf("Число значений для заполнения структуры '%s' менее одного\n", c)
	}

	c.d = data[0]

	return true, nil
}

// Структура: прямоугольник
type Rectangle struct {
	// Размер стороны один
	sideA float64
	// Размер стороны два
	sideB float64
}

func (r *Rectangle) CalcArea() float64 {
	return r.sideA * r.sideB
}

func (r *Rectangle) FillData(data []float64) (bool, error) {
	if data == nil {
		return false, fmt.Errorf("Не передан слайс данных в структуру %s\n", r)
	}

	if len(data) < 2 {
		return false, fmt.Errorf("Число значений для заполнения структуры '%s' менее двух\n", r)
	}

	r.sideA = data[0]
	r.sideB = data[0]

	return true, nil
}

func (c *Rectangle) String() string {
	return "ПРЯМОУГОЛЬНИК"
}

// Заполнить структуру прямоугольник данными
func FillAndCalcArea(figure Arear, data []float64) (result float64, err error) {
	if _, testOk := figure.(*Circle); !testOk {
		if _, testOk = figure.(*Rectangle); !testOk {
			return 0, fmt.Errorf("Не передана структура данных")
		}
	}

	var ok bool

	if ok, err = figure.FillData(data); ok {
		result = figure.CalcArea()
	} else {
		return 0, fmt.Errorf("Ошибка при заполнении структуры %s (%s)", figure, err)
	}

	return
}

type Arear interface {
	CalcArea() float64
	FillData(data []float64) (bool, error)
}

func main() {
	fmt.Println("Программы вычисляет площадь фигуры, используя механику интерфейсов Go")

	fmt.Println(EnterLens)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		bEnterOk := true

		str := scanner.Text()

		if str == "q" {
			return
		}

		aStr := strings.Split(str, ",")

		// Проверка и инициализация введённых результатов
		aLens := make([]float64, 0)
		for idx, item := range aStr {
			flItem, err := strconv.ParseFloat(strings.Trim(item, " "), 64)
			if err != nil {
				fmt.Printf("Для числа номер %d задан неверный формат.\n", idx+1)
				bEnterOk = false

				break
			}

			aLens = append(aLens, flItem)
		}

		if !bEnterOk {
			fmt.Println(EnterLens)
			continue
		}

		var obj Arear

		switch len(aLens) {
		case 0:
			fmt.Println("Не введено никаких чисел")
			fmt.Println(EnterLens)
			continue
		case 1:
			// Создаётся объект окружность
			obj = &Circle{}
			fmt.Printf("Создана фигура %s\n", obj)
		case 2:
			// Создаётся объект окружность
			obj = &Rectangle{}
			fmt.Printf("Создана фигура %s\n", obj)
		case 3:
			// Создаётся объект "ошибочный"
			obj = nil
			fmt.Printf("Создана фигура %s\n", obj)
		default:
			fmt.Println("Введено слишком много длин сторон.")
			fmt.Println(EnterLens)
			continue
		}

		// Вычисляются площади фигур и распечатывается результат
		area, err := FillAndCalcArea(obj, aLens)
		if err != nil {
			fmt.Printf("Ошибка: %s\n", err)
		} else {
			fmt.Printf("Площадь фигуры : %s равна: %f\n", obj, area)
		}

		fmt.Println(EnterLens)
	}

}
