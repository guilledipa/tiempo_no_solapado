// Este programa calcula el tiempo total (sin solapamiento)
// dada una lista de segmentos temporales.
// Por ejemplo, dados s1 = [1, 3] y s2= [2, 4]
// el tiempo total es Ttot = 3
package main

import (
	"fmt"
	"sort"
)

type periodo struct {
	timeIni, timeFin float32
}

type porTimeIni []periodo

// porTimeIni implementa sort.Interface basado en el campo timeIni del struct periodo.
func (p porTimeIni) Len() int           { return len(p) }
func (p porTimeIni) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p porTimeIni) Less(i, j int) bool { return p[i].timeIni < p[j].timeIni }

var (
	tSum     float32
	periodos = []periodo{
		periodo{1, 2},
		periodo{1.2, 12},
		periodo{0, 0.5},
		periodo{5, 6},
		periodo{1, 2.5},
	}
)

func main() {
	// Ordeno periodos
	sort.Sort(porTimeIni(periodos))
	fmt.Println(periodos)

	// periodos esta ordenado, por lo cual, puedo inicializar tSum con la
	// duracion del 1er segmento.
	tSum := periodos[0].timeFin - periodos[0].timeIni
	tSumFin := periodos[0].timeFin

	// Como inicialice tSum con el 1er segmento (periodos[0]), arranco el for loop
	// desde el segundo elemento periodos(1)
	for i := 1; i < len(periodos); i++ {
		if periodos[i].timeFin > tSumFin {
			if periodos[i].timeIni >= tSumFin {
				tSum = tSum + (periodos[i].timeFin - periodos[i].timeIni)
			} else {
				tSum = tSum + (periodos[i].timeFin - tSumFin)

			}
			tSumFin = periodos[i].timeFin
		} // else periodos[i].timeFin <= periodos[i-1].timeFin -> Ignorar. Contenido dentro del segmento anterior.
	}

	fmt.Printf("Tiempo total no solapado: %v\n", tSum)
}
