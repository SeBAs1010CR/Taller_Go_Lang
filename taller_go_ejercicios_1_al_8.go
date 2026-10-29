package main

import "fmt"

// Ejercicio 1: intercambio de valores.
// Recibe punteros porque Go pasa los argumentos por valor: si recibiera ints,
// la función trabajaría con copias locales y el cambio no llegaría al llamador.
func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

// Ejercicio 2: estadísticas de un arreglo.
func sum(values []int) int {
	total := 0
	for i := 0; i < len(values); i++ {
		total += values[i]
	}
	return total
}

func minimum(values []int) int {
	result := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < result {
			result = values[i]
		}
	}
	return result
}

func maximum(values []int) int {
	result := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > result {
			result = values[i]
		}
	}
	return result
}

func countNegatives(values []int) int {
	count := 0
	for i := 0; i < len(values); i++ {
		if values[i] < 0 {
			count++
		}
	}
	return count
}

// printReverse recorre el arreglo de atrás hacia adelante sin modificar la entrada.
func printReverse(values []int) {
	for i := len(values) - 1; i >= 0; i-- {
		fmt.Print(values[i])
		if i > 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

// Ejercicio 3: copia independiente.
// make reserva un arreglo nuevo con ceros; luego copiamos elemento por elemento.
// other := original NO copia: el slice es un descriptor (puntero al arreglo subyacente,
// longitud y capacidad). Ambos descriptores apuntan a los mismos elementos, por eso
// modificar other también modificaría original.
func copyArray(values []int) []int {
	result := make([]int, len(values))
	for i := 0; i < len(values); i++ {
		result[i] = values[i]
	}
	return result
}

// Ejercicio 4: ordenamiento por inserción, in place, sin sort ni slices.Sort.
// Se usa sobre una copia para conservar el arreglo original.
func sortAscending(values []int) {
	for i := 1; i < len(values); i++ {
		current := values[i]
		j := i - 1
		for j >= 0 && values[j] > current {
			values[j+1] = values[j]
			j--
		}
		values[j+1] = current
	}
}

// Ejercicio 5: dos valores que suman un objetivo.
// Requiere un slice ordenado ascendentemente y no lo modifica.
// Se usan dos índices: left al inicio y right al final, siempre con left < right
// para garantizar posiciones distintas.
// Si la suma es menor que el objetivo, values[left] no puede formar par con ningún
// valor restante: el mayor posible es values[right], y aun así la suma sería menor.
// Por eso se descarta el extremo izquierdo (left++). Si la suma es mayor, con el
// mismo razonamiento se descarta el extremo derecho (right--).
// Entradas menores de 2 elementos no tienen par posible.
func findPair(values []int, target int) bool {
	if len(values) < 2 {
		return false
	}
	left := 0
	right := len(values) - 1
	for left < right {
		sum := values[left] + values[right]
		if sum == target {
			fmt.Printf("Par válido: (%d, %d)\n", values[left], values[right])
			return true
		}
		if sum < target {
			left++
		} else {
			right--
		}
	}
	return false
}

// ==========================================
// SECCIÓN DE RAÚL (Ejercicios 6, 7 y 8)
// ==========================================

// Ejercicio 6: 3Sum
func threeSum(values []int) int {
	if len(values) < 3 {
		return 0
	}
	sorted := copyArray(values)
	sortAscending(sorted)

	count := 0
	n := len(sorted)

	for i := 0; i < n-2; i++ {
		if i > 0 && sorted[i] == sorted[i-1] {
			continue
		}
		left := i + 1
		right := n - 1

		for left < right {
			sum := sorted[i] + sorted[left] + sorted[right]
			if sum == 0 {
				fmt.Printf("  Terna %d: (%d, %d, %d)\n", count+1, sorted[i], sorted[left], sorted[right])
				count++
				left++
				right--
				for left < right && sorted[left] == sorted[left-1] {
					left++
				}
				for left < right && sorted[right] == sorted[right+1] {
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return count
}

// Ejercicio 7: Máximos desde ambos extremos
func buildLeftMax(height []int, leftMax []int) {
	if len(height) == 0 {
		return
	}
	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		if height[i] > leftMax[i-1] {
			leftMax[i] = height[i]
		} else {
			leftMax[i] = leftMax[i-1]
		}
	}
}

func buildRightMax(height []int, rightMax []int) {
	n := len(height)
	if n == 0 {
		return
	}
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		if height[i] > rightMax[i+1] {
			rightMax[i] = height[i]
		} else {
			rightMax[i] = rightMax[i+1]
		}
	}
}

// Ejercicio 8: Trapping Rain Water
func trappedWater(height []int) int64 {
	n := len(height)
	if n == 0 {
		return 0
	}
	leftMax := make([]int, n)
	rightMax := make([]int, n)

	buildLeftMax(height, leftMax)
	buildRightMax(height, rightMax)

	var total int64 = 0
	for i := 0; i < n; i++ {
		minHeight := leftMax[i]
		if rightMax[i] < minHeight {
			minHeight = rightMax[i]
		}
		waterAtPos := minHeight - height[i]
		if waterAtPos > 0 {
			total += int64(waterAtPos)
		}
	}
	return total
}

func printTrappedWaterDetail(height []int) {
	n := len(height)
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	buildLeftMax(height, leftMax)
	buildRightMax(height, rightMax)

	fmt.Println("\nDetalle de cálculo por posición (Agua atrapada):")
	fmt.Printf("%-8s %-8s %-10s %-11s %-12s %-10s\n", "Índice", "Altura", "LeftMax", "RightMax", "Mínimo(L,R)", "Agua")
	fmt.Println("------------------------------------------------------------")

	var total int64 = 0
	for i := 0; i < n; i++ {
		minH := leftMax[i]
		if rightMax[i] < minH {
			minH = rightMax[i]
		}
		water := minH - height[i]
		if water < 0 {
			water = 0
		}
		total += int64(water)
		fmt.Printf("%-8d %-8d %-10d %-11d %-12d %-10d\n", i, height[i], leftMax[i], rightMax[i], minH, water)
	}
	fmt.Printf("Total acumulado de agua atrapada: %d\n", total)
}

func main() {
	// Ejercicio 1
	a, b := 8, -3
	swap(&a, &b)
	fmt.Printf("Ejercicio 1: a = %d, b = %d\n", a, b)

	// Ejercicio 2
	values := []int{4, -2, 0, 7, -5}
	fmt.Printf("Ejercicio 2: %v\n", values)
	fmt.Printf("Suma: %d Mínimo: %d Máximo: %d Negativos: %d\n",
		sum(values), minimum(values), maximum(values), countNegatives(values))
	fmt.Print("Recorrido inverso: ")
	printReverse(values)

	// Ejercicio 3
	original := []int{4, -2, 0, 7, -5}
	other := copyArray(original)
	other[0] = 99
	fmt.Printf("Ejercicio 3: original %v\n", original)
	fmt.Printf("             copia    %v\n", other)

	// Ejercicio 4
	sorted := copyArray(values)
	sortAscending(sorted)
	fmt.Printf("Ejercicio 4: ordenado %v\n", sorted)
	fmt.Printf("             entrada sin cambios %v\n", values)

	// Ejercicio 5: pruebas de búsqueda de pares
	fmt.Println("Ejercicio 5: búsqueda de pares")
	pairTests := []struct {
		values []int
		target int
	}{
		{[]int{-4, -1, 0, 2, 5}, 1},
		{[]int{-4, -1, 0, 2, 5}, 8},
		{[]int{0, 0}, 0},
		{[]int{3}, 6},
	}
	for i, test := range pairTests {
		fmt.Printf("Prueba %d: arreglo %v, objetivo %d -> ", i+1, test.values, test.target)
		if !findPair(test.values, test.target) {
			fmt.Println("No existe")
		}
	}

	// ==========================================
	// PRUEBAS DE RAÚL (Ejercicios 6, 7 y 8)
	// ==========================================
	fmt.Println("\n--- Ejercicio 6: 3Sum ---")
	threeSumCases := []struct {
		name   string
		values []int
	}{
		{"Caso 1", []int{-1, 0, 1, 2, -1, -4}},
		{"Caso 2", []int{0, 1, 1}},
		{"Caso 3", []int{0, 0, 0}},
		{"Caso 4", []int{0, 0, 0, 0}},
		{"Caso 5", []int{-2, 0, 0, 2, 2}},
		{"Insuficiente", []int{1, -1}},
	}
	for _, c := range threeSumCases {
		fmt.Printf("Entrada %s: %v\n", c.name, c.values)
		total := threeSum(c.values)
		fmt.Printf("Total de ternas únicas encontradas: %d\n\n", total)
	}

	fmt.Println("--- Ejercicio 7: Máximos Acumulados ---")
	sampleHeight := []int{3, 0, 2, 0, 4}
	leftMax := make([]int, len(sampleHeight))
	rightMax := make([]int, len(sampleHeight))
	buildLeftMax(sampleHeight, leftMax)
	buildRightMax(sampleHeight, rightMax)
	fmt.Printf("height:   %v\n", sampleHeight)
	fmt.Printf("leftMax:  %v\n", leftMax)
	fmt.Printf("rightMax: %v\n\n", rightMax)

	fmt.Println("--- Ejercicio 8: Trapping Rain Water ---")
	rainCases := []struct {
		name     string
		height   []int
		expected int64
	}{
		{"Un solo elemento", []int{5}, 0},
		{"Ascendente", []int{0, 1, 2, 3}, 0},
		{"Descendente", []int{3, 2, 1, 0}, 0},
		{"Plano", []int{2, 2, 2}, 0},
		{"Pozo simple", []int{3, 0, 3}, 3},
		{"Múltiples pozos 1", []int{3, 0, 2, 0, 4}, 7},
		{"Múltiples pozos 2", []int{4, 2, 0, 3, 2, 5}, 9},
		{"LeetCode clásico", []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{"Caso Propio", []int{5, 1, 2, 0, 3}, 8},
	}
	for _, c := range rainCases {
		res := trappedWater(c.height)
		fmt.Printf("%-20s -> Agua: %d (Esperado: %d)\n", c.name, res, c.expected)
	}

	printTrappedWaterDetail([]int{3, 0, 2, 0, 4})
}
