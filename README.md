# Taller de programación imperativa en Go

Curso: Paradigmas de Programación, Ingeniería en Computadores, ITCR.

## Integrantes

- Sebastian Azofeifa Vilallobos (ejercicios 1 al 5)
- Raúl (ejercicios 6 al 8)

## Lenguaje

Go (versión probada: 1.27.3).

## Ejecución

```
go run taller_go_ejercicios_1_al_5.go
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

Ver `findPair` en `taller_go_ejercicios_1_al_5.go`. Se usan dos índices (`left` y `right`) sobre un slice ordenado. Si la suma es menor que el objetivo se avanza `left`, y si es mayor se retrocede `right`. Cuando la suma es menor, `values[left]` no puede formar par con ningún valor restante, porque el mayor posible es `values[right]` y la suma seguiría siendo menor. Por eso se descarta el extremo izquierdo; el razonamiento inverso aplica al derecho.

## Manejo de memoria: Go frente a C

- En C, `malloc` reserva memoria y hay que liberarla con `free`. Olvidarlo produce fugas; liberar dos veces o usar memoria ya liberada produce errores difíciles de detectar.
- En Go, `make` reserva el slice y el recolector de basura libera la memoria cuando ya no es alcanzable. No hay `free`.
- En Go, `other := original` no copia los elementos: ambos slices comparten el mismo arreglo subyacente. Copiar exige reservar un slice nuevo con `make` y recorrer los elementos, como hace `copyArray`.

## Pendiente: ejercicios 6 al 8 

### Evitar ternas duplicadas en 3Sum

_Pendiente._

### Por qué el agua depende del menor de los dos máximos

_Pendiente._
