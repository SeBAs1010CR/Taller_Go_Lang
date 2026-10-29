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
}
