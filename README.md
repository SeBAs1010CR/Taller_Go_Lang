
# Taller de programación imperativa en Go

Curso: Paradigmas de Programación, Ingeniería en Computadores, ITCR.

## Integrantes

- Sebastian Azofeifa Vilallobos (ejercicios 1 al 5)
- Raúl Stuart Ramírez Villegas (ejercicios 6 al 8)

## Lenguaje

Go (versión probada: 1.27.3).

## Ejecución

```bash
go run taller_go_ejercicios_1_al_8.go

```

El programa no pide entrada: las pruebas están definidas en `main`.

## Resultados de las pruebas

```
Ejercicio 1: a = -3, b = 8
Ejercicio 2: [4 -2 0 7 -5]
Suma: 4 Mínimo: -5 Máximo: 7 Negativos: 2
Recorrido inverso: -5 7 0 -2 4
Ejercicio 3: original [4 -2 0 7 -5]
             copia    [99 -2 0 7 -5]
Ejercicio 4: ordenado [-5 -2 0 4 7]
             entrada sin cambios [4 -2 0 7 -5]
Ejercicio 5: búsqueda de pares
Prueba 1: arreglo [-4 -1 0 2 5], objetivo 1 -> Par válido: (-4, 5)
Prueba 2: arreglo [-4 -1 0 2 5], objetivo 8 -> No existe
Prueba 3: arreglo [0 0], objetivo 0 -> Par válido: (0, 0)
Prueba 4: arreglo [3], objetivo 6 -> No existe

Ejercicio 6: 3Sum
- Caso 1 [-1, 0, 1, 2, -1, -4] -> Terna 1: (-1, -1, 2), Terna 2: (-1, 0, 1) | Total: 2
- Caso 2 [0, 1, 1] -> Total: 0
- Caso 3 [0, 0, 0] -> Terna 1: (0, 0, 0) | Total: 1
- Caso 4 [0, 0, 0, 0] -> Terna 1: (0, 0, 0) | Total: 1
- Caso 5 [-2, 0, 0, 2, 2] -> Terna 1: (-2, 0, 2) | Total: 1

Ejercicio 7: Máximos desde ambos extremos
height:   [3, 0, 2, 0, 4]
leftMax:  [3, 3, 3, 3, 4]
rightMax: [4, 4, 4, 4, 4]

Ejercicio 8: Trapping Rain Water
- Un solo elemento [5] -> Agua: 0
- Ascendente [0, 1, 2, 3] -> Agua: 0
- Descendente [3, 2, 1, 0] -> Agua: 0
- Plano [2, 2, 2] -> Agua: 0
- Pozo simple [3, 0, 3] -> Agua: 3
- Múltiples pozos 1 [3, 0, 2, 0, 4] -> Agua: 7
- Múltiples pozos 2 [4, 2, 0, 3, 2, 5] -> Agua: 9
- LeetCode clásico [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1] -> Agua: 6
- Caso Propio [5, 1, 2, 0, 3] -> Agua: 8

```

## Alcance de variables

Ejemplo tomado de `swap`:

```go
func swap(a *int, b *int) {
	tmp := *a   // tmp existe solo dentro de swap
	*a = *b
	*b = tmp
}

```

`tmp` se crea al entrar a `swap` y deja de existir al salir. Nadie fuera de la función puede verla. Por eso el intercambio se hace con punteros: `a` y `b` apuntan a las variables `a` y `b` de `main`, y `*a`, `*b` permiten modificar esos valores.

## Cómo se ejecuta la búsqueda de pares sin modificar la entrada

Ver `findPair` en `taller_go_ejercicios_1_al_8.go`. Se usan dos índices (`left` y `right`) sobre un slice ordenado. Si la suma es menor que el objetivo se avanza `left`, y si es mayor se retrocede `right`. Cuando la suma es menor, `values[left]` no puede formar par con ningún valor restante, porque el mayor posible es `values[right]` y la suma seguiría siendo menor. Por eso se descarta el extremo izquierdo; el razonamiento inverso aplica al derecho.

## Manejo de memoria: Go frente a C

* En C, `malloc` reserva memoria y hay que liberarla con `free`. Olvidarlo produce fugas; liberar dos veces o usar memoria ya liberada produce errores difíciles de detectar.
* En Go, `make` reserva el slice y el recolector de basura libera la memoria cuando ya no es alcanzable. No hay `free`.
* En Go, `other := original` no copia los elementos: ambos slices comparten el mismo arreglo subyacente. Copiar exige reservar un slice nuevo con `make` y recorrer los elementos, como hace `copyArray`.

## Cómo se evitan ternas duplicadas en 3Sum

Para evitar generar ternas repetidas manteniendo una complejidad $O(n^2)$ tras el ordenamiento, implementamos los siguientes controles:

1. **Deduplicación del primer elemento ($i$):** Al iterar con el primer puntero `i`, si `sorted[i] == sorted[i-1]`, la iteración actual se omite con `continue`. Esto evita procesar el mismo valor inicial dos veces.
2. **Deduplicación de los extremos (`left` y `right`):** Cuando encontramos una terna cuya suma es `0`, registramos la solución y avanzamos ambos punteros (`left++` y `right--`). Inmediatamente aplicamos ciclos `for` para desplazar `left` a la derecha y `right` a la izquierda mientras apunten a valores repetidos respecto a los recién procesados:
```go
for left < right && sorted[left] == sorted[left-1] { left++ }
for left < right && sorted[right] == sorted[right+1] { right-- }

```



## Por qué el agua depende del menor de los dos máximos

En el problema de la acumulación de agua (*Trapping Rain Water*), la cantidad de líquido que puede sostenerse sobre una posición específica $i$ está limitada físicamente por la pared más baja que la encierra por la izquierda y por la derecha:

$$\text{Agua}[i] = \max(0, \min(\text{leftMax}[i], \text{rightMax}[i]) - \text{height}[i])$$

* **Vasos comunicantes:** Si la pared izquierda tiene una altura máxima de $3$ y la derecha una de $4$, el agua no puede superar el nivel $3$. Si sobrepasara esa altura, se desbordaría por el extremo más bajo.
* **Cálculo en $O(n)$:** Al precalcular dos arreglos `leftMax` y `rightMax` en una sola pasada cada uno, determinamos en tiempo $O(1)$ para cada columna cuál es la cota superior del nivel del agua, restando la altura propia de la barra.

## Caso propio de Trapping Rain Water y desglose

* **Entrada:** `[5, 1, 2, 0, 3]`
* **Agua total esperada:** `8`
* **Justificación y desglose por posición:**
* `i=0` (h=5): `leftMax=5`, `rightMax=5` $\rightarrow \min(5,5) - 5 = 0$
* `i=1` (h=1): `leftMax=5`, `rightMax=3` $\rightarrow \min(5,3) - 1 = 2$
* `i=2` (h=2): `leftMax=5`, `rightMax=3` $\rightarrow \min(5,3) - 2 = 1$
* `i=3` (h=0): `leftMax=5`, `rightMax=3` $\rightarrow \min(5,3) - 0 = 3$
* `i=4` (h=3): `leftMax=5`, `rightMax=3` $\rightarrow \min(5,3) - 3 = 0$
* **Suma acumulada:** $0 + 2 + 1 + 3 + 0 = 8$ unidades de agua.



