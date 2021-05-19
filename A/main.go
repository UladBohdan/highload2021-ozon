package main

import (
	"fmt"
	"time"
)

// Brief solution description (in Russian):
//
// Короткое описание решения (с не самой оптимальной константой, но быстрое и надёжное в реализации):
// * делаем перебор посекундно от "00:00:00" до "11:59:59";
// * для каждого времени по формулам считаем углы стрелок;
// * сравниваем с заданными углами, учитывая погрешность в 3 градуса;
// * возвращаем ответ как только нашли.
//
// Сложность решения: O(1) (от константы)

func main() {
	var hGrad, mGrad, sGrad float64
	fmt.Scanf("%f %f %f", &hGrad, &mGrad, &sGrad)

	t := time.Time{}
	for {
		h := float64(t.Hour())
		m := float64(t.Minute())
		s := float64(t.Second())

		s1 := 6 * s // 360 * (s / 60)
		m1 := 360 * (m/60 + s/(60*60))
		h1 := 360 * (h/12 + m/(12*60) + s/(12*60*60))

		closeEnough := func(a, b float64) bool {
			return (a-b < 3 && a-b > -3) ||
				(b-a < 3 && b-a > -3) ||
				(a-b < 363 && a-b > 357) ||
				(b-a < 363 && b-a > 357)
		}

		if closeEnough(hGrad, h1) && closeEnough(mGrad, m1) && closeEnough(sGrad, s1) {
			fmt.Println(t.Format("15:04:05"))
			return
		}

		if h == 11 && m == 59 && s == 59 {
			break
		}

		t = t.Add(time.Second)
	}

	fmt.Println("answer not found")
}
