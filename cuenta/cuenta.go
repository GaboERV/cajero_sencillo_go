package cuenta

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type cuenta_struct struct {
	Nombre_titular string
	Numero_cuenta  string
	Pin            uint
	Saldo          uint
}

func CrearCuenta(Nombre_titular string, PIN uint) string {
	var numero_ramdom string
	for {

		numero_r := fmt.Sprint(rand.Intn(10000), rand.Intn(10000), rand.Intn(10000), rand.Intn(10000))
		numero_r = strings.ReplaceAll(numero_r," ","")
		fmt.Print(numero_ramdom)

		if !validarNumerodecuenta(numero_r) {
			fmt.Println("numero creado por el sistema no valido intentando de nuevo")
			continue
		}
		numero_ramdom = numero_r
		break
	}

	var cuentas []cuenta_struct = obtener_lista()

	cuenta := cuenta_struct{Nombre_titular: Nombre_titular, Numero_cuenta: numero_ramdom, Pin: PIN}

	cuentas = append(cuentas, cuenta)
	sycnJson(cuentas)
	info_cuenta := fmt.Sprintln("tu numero de cuenta es",numero_ramdom)
	return info_cuenta
}

func validarNumerodecuenta(numero_ramdom string) bool {
	var cuentas []cuenta_struct = obtener_lista()
	for _, value := range cuentas {
		if value.Numero_cuenta == numero_ramdom {
			return false
		}
	}
	return true
}

func sycnJson(cuentas []cuenta_struct){
	file, err := os.Create("cuentas.json")
	if err != nil {
		fmt.Println("Error al guardar archivo:", err)
		return
	}
	defer file.Close()
	json.NewEncoder(file).Encode(cuentas)
} 

func obtener_lista() []cuenta_struct {
	var cuentas []cuenta_struct
	fileData, err := os.ReadFile("cuentas.json")
	if err == nil {
		_ = json.Unmarshal(fileData, &cuentas)
	}
	return cuentas
}

func SaldoCuenta(Numero_cuenta string, Pin uint) {
	var cuentas []cuenta_struct = obtener_lista()
	for _, value := range cuentas {
		if value.Numero_cuenta == Numero_cuenta {
			fmt.Println("Numero de cuenta encontrado procesando respuesta ...")
		} else {
			continue
		}
		if value.Pin == Pin {
			fmt.Println("tu saldo actual es de", value.Saldo)
			return
		} else {
			fmt.Println("Pin incorrecto por favor intente de nuevo")
			return
		}
	}
	fmt.Println("numero de cuenta no encotrado")
}

func RetirarMonto(Numero_cuenta string, PIN uint, monto uint) {
	var cuentas []cuenta_struct = obtener_lista()
	for index, value := range cuentas {
		if value.Numero_cuenta == Numero_cuenta {
			fmt.Println("Numero de cuenta encontrado procesando respuesta ...")
		} else {
			continue
		}
		if value.Pin == PIN {
			if monto > value.Saldo {
				fmt.Println("el monto excede la cantidad disponible quea hay")
				return
			}
			fmt.Println("procediendo con el retiro ...")
			cuentas[index] = cuenta_struct{
				Nombre_titular: value.Nombre_titular, 
				Numero_cuenta: value.Numero_cuenta, 
				Pin: value.Pin, 
				Saldo: value.Saldo - monto,
			}
			sycnJson(cuentas)
			return
		} else {
			fmt.Println("Pin incorrecto por favor intente de nuevo")
			return
		}
	}
	fmt.Println("numero de cuenta no encotrado")
}



func CambiarPIN(Numero_cuenta string, PIN uint, nuevo_pin uint) {
	var cuentas []cuenta_struct = obtener_lista()
	for index, value := range cuentas {
		if value.Numero_cuenta == Numero_cuenta {
			fmt.Println("Numero de cuenta encontrado procesando respuesta ...")
		} else {
			continue
		}
		if value.Pin == PIN {
			fmt.Println("actualizando PIN")
			cuentas[index] = cuenta_struct{
				Nombre_titular: value.Nombre_titular, 
				Numero_cuenta: value.Numero_cuenta, 
				Pin: nuevo_pin, 
				Saldo: value.Saldo,
			}
			sycnJson(cuentas)
			return
		} else {
			fmt.Println("Pin incorrecto por favor intente de nuevo")
			return
		}
	}
	fmt.Println("numero de cuenta no encotrado")
}

func IngresarMonto(Numero_cuenta string, PIN uint, monto uint) {
	var cuentas []cuenta_struct = obtener_lista()
	for index, value := range cuentas {
		if value.Numero_cuenta == Numero_cuenta {
			fmt.Println("Numero de cuenta encontrado procesando respuesta ...")
		} else {
			continue
		}
		if value.Pin == PIN {
			fmt.Println("procediendo con el ingreso de dinero ...")
			cuentas[index] = cuenta_struct{
				Nombre_titular: value.Nombre_titular, 
				Numero_cuenta: value.Numero_cuenta, 
				Pin: value.Pin, 
				Saldo: value.Saldo + monto,
			}
			sycnJson(cuentas)
			return
		} else {
			fmt.Println("Pin incorrecto por favor intente de nuevo")
			return
		}
	}
	fmt.Println("numero de cuenta no encotrado")
}

