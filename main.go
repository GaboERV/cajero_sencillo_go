package main

import (
	"cajero/cuenta"
	"fmt"
)

func main() {
	fmt.Println("bienvenido al cajero")
	for {
		var opcion int
		fmt.Println("1) crear cuenta \n 2) consultar saldo \n 3) retirar saldo \n 4) ingresar saldo \n 5)cambiar pin \n 6)salir")
		fmt.Print("por favor ingrese una opcion: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var Nombre_titular string
			var Pin uint
			fmt.Print("ingresa tu nombre: ")
			fmt.Scanln(&Nombre_titular)

			fmt.Println("ingresa tu PIN: ")
			fmt.Scanln(&Pin)
		    info := cuenta.CrearCuenta(Nombre_titular, Pin)

			fmt.Println(info)
		case 2:
			Numero_cuenta, Pin := ingresarCredendciales()
			cuenta.SaldoCuenta(Numero_cuenta, Pin)
		case 3:
			var monto uint
			Numero_cuenta, Pin := ingresarCredendciales()
			fmt.Print("Ingresa el monto a retirar: ")
			fmt.Scanln(&monto)
			cuenta.RetirarMonto(Numero_cuenta, Pin, monto)
		case 4:
			var monto uint
			Numero_cuenta, Pin := ingresarCredendciales()
			fmt.Print("Ingresa el monto a ingresar: ")
			fmt.Scanln(&monto)
			cuenta.IngresarMonto(Numero_cuenta, Pin, monto)
		case 5:
			var nuevo_pin uint
			Numero_cuenta, Pin := ingresarCredendciales()
			fmt.Print("Ingresa tu nuevo pin: ")
			fmt.Scanln(&nuevo_pin)
			cuenta.CambiarPIN(Numero_cuenta, Pin, nuevo_pin)
		case 6:
			fmt.Println("saliendo del progarama")
		default:
			fmt.Println("opcion no valida")
			continue
		}
		if opcion == 6 {
			break
		}
	}

}

func ingresarCredendciales() (string, uint) {
	var Numero_cuenta string
	var Pin uint
	fmt.Print("ingresa tu numero de cuenta: ")
	fmt.Scanln(&Numero_cuenta)

	fmt.Println("ingresa tu PIN: ")
	fmt.Scanln(&Pin)

	return Numero_cuenta, Pin
}
