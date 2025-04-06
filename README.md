
# Proyecto Cajero Automático en Go

## Descripción

Este es un proyecto simple de simulación de un Cajero Automático (ATM) desarrollado en Go. Permite a los usuarios crear cuentas bancarias, consultar saldos, retirar e ingresar dinero, y cambiar su PIN. Los datos de las cuentas se persisten en un archivo `cuentas.json`.

El proyecto está dividido en dos paquetes principales:
1.  `main`: Contiene el punto de entrada de la aplicación y maneja la interfaz de usuario basada en consola (menú de opciones).
2.  `cuenta`: Encapsula toda la lógica relacionada con la gestión de las cuentas bancarias (creación, operaciones, persistencia).

## Características

*   Crear nuevas cuentas bancarias con un número de cuenta único generado aleatoriamente.
*   Consultar el saldo de una cuenta existente, validando el número de cuenta y PIN.
*   Retirar fondos de una cuenta, verificando el PIN y el saldo disponible.
*   Ingresar fondos a una cuenta, verificando el PIN.
*   Cambiar el PIN de una cuenta, validando el PIN actual.
*   Persistencia de datos de cuentas en formato JSON (`cuentas.json`).
*   Menú interactivo en la consola para una fácil navegación.

## Estructura del Proyecto

    


cajero/
├── main.go # Punto de entrada principal, maneja la interfaz de usuario
├── go.mod # Módulo de Go (si se inicializó)
├── cuenta/
│ └── cuenta.go # Lógica principal del manejo de cuentas (struct, funciones)
└── cuentas.json # Archivo para almacenar los datos de las cuentas (se crea/actualiza al ejecutar)

      
## Cómo Empezar

### Prerrequisitos

*   Tener instalado Go (versión 1.16 o superior recomendada). Puedes descargarlo desde [https://golang.org/dl/](https://golang.org/dl/).

### Instalación y Ejecución

1.  **Clonar o Descargar:** Obtén los archivos del proyecto. Si es un repositorio Git, clónalo:
    ```bash
    git clone <URL_DEL_REPOSITORIO>
    cd cajero
    ```
    Si solo tienes los archivos, asegúrate de que `main.go` y la carpeta `cuenta` (con `cuenta.go` dentro) estén en el mismo directorio raíz del proyecto (`cajero/`).

2.  **Ejecutar:** Abre una terminal en el directorio raíz del proyecto (`cajero/`) y ejecuta el siguiente comando:
    ```bash
    go run main.go
    ```

    Esto compilará y ejecutará el programa. La primera vez que se cree una cuenta, se generará el archivo `cuentas.json` en el mismo directorio.

## Detalles de Implementación

### Paquete `main` (`main.go`)

*   Presenta un menú de opciones al usuario en un bucle infinito hasta que se selecciona la opción de salir (6).
*   Utiliza `fmt.Scanln` para leer la opción del usuario y los datos necesarios (nombre, PIN, número de cuenta, monto).
*   Llama a las funciones correspondientes del paquete `cuenta` según la opción seleccionada.
*   Incluye una función auxiliar `ingresarCredendciales` para obtener el número de cuenta y el PIN del usuario, reutilizando este proceso en varias opciones.

### Paquete `cuenta` (`cuenta.go`)

*   **`cuenta_struct`**: Define la estructura de datos para una cuenta bancaria, conteniendo:
    *   `Nombre_titular` (string): Nombre del titular de la cuenta.
    *   `Numero_cuenta` (string): Número único de la cuenta (generado aleatoriamente).
    *   `Pin` (uint): Número de Identificación Personal.
    *   `Saldo` (uint): Saldo actual de la cuenta (inicializado en 0).

*   **Persistencia (`obtener_lista`, `sycnJson`)**:
    *   `obtener_lista()`: Lee el archivo `cuentas.json`. Si el archivo existe y es válido, decodifica el JSON en un slice de `cuenta_struct`. Si hay un error (ej. el archivo no existe la primera vez), devuelve un slice vacío.
    *   `sycnJson()`: Toma un slice de `cuenta_struct`, lo codifica a formato JSON y sobrescribe el archivo `cuentas.json`. Se llama después de cualquier modificación en la lista de cuentas (crear, retirar, ingresar, cambiar PIN).

*   **`CrearCuenta(Nombre_titular string, PIN uint) string`**:
    *   Genera un número de cuenta pseudoaleatorio concatenando llamadas a `rand.Intn` y eliminando espacios.
    *   Utiliza `validarNumerodecuenta` para asegurarse de que el número generado no exista ya en `cuentas.json`. Repite la generación si el número ya existe.
    *   Crea una nueva instancia de `cuenta_struct` con los datos proporcionados y el número generado.
    *   Añade la nueva cuenta al slice obtenido de `obtener_lista`.
    *   Guarda la lista actualizada usando `sycnJson`.
    *   Devuelve un mensaje informativo con el número de cuenta creado.

*   **`validarNumerodecuenta(numero_ramdom string) bool`**:
    *   Obtiene la lista actual de cuentas.
    *   Itera sobre las cuentas para verificar si alguna coincide con el `numero_ramdom` proporcionado.
    *   Devuelve `false` si el número ya existe, `true` en caso contrario.

*   **Operaciones de Cuenta (`SaldoCuenta`, `RetirarMonto`, `IngresarMonto`, `CambiarPIN`)**:
    *   Todas estas funciones toman el `Numero_cuenta` y el `Pin` (o `PIN`) como argumentos para identificar y autenticar la cuenta.
    *   Obtienen la lista completa de cuentas usando `obtener_lista`.
    *   Iteran sobre la lista buscando la cuenta por `Numero_cuenta`.
    *   Si encuentran la cuenta, verifican el `Pin` proporcionado.
    *   Si el PIN es correcto, realizan la operación correspondiente (mostrar saldo, restar del saldo, sumar al saldo, actualizar el PIN).
    *   Para operaciones que modifican la cuenta (`RetirarMonto`, `IngresarMonto`, `CambiarPIN`), actualizan el objeto `cuenta_struct` directamente en el slice y luego llaman a `sycnJson` para persistir los cambios.
    *   Muestran mensajes informativos o de error al usuario a través de `fmt.Println`.
    *   La función `RetirarMonto` incluye una verificación adicional para asegurar que el monto a retirar no exceda el saldo disponible.

## Uso

Al ejecutar el programa, se mostrará el siguiente menú:

    

bienvenido al cajero

    crear cuenta

    consultar saldo

    retirar saldo

    ingresar saldo

    cambiar pin

    salir
    por favor ingrese una opcion:

      
El usuario debe ingresar el número de la opción deseada y presionar Enter. Dependiendo de la opción, el programa solicitará información adicional como nombre, PIN, número de cuenta o monto.

*   **Crear Cuenta (1):** Pide nombre y PIN. Genera un número de cuenta y lo muestra.
*   **Consultar Saldo (2):** Pide número de cuenta y PIN. Muestra el saldo si las credenciales son correctas.
*   **Retirar Saldo (3):** Pide número de cuenta, PIN y monto. Realiza el retiro si es posible.
*   **Ingresar Saldo (4):** Pide número de cuenta, PIN y monto. Realiza el ingreso.
*   **Cambiar PIN (5):** Pide número de cuenta, PIN actual y nuevo PIN. Actualiza el PIN si el actual es correcto.
*   **Salir (6):** Termina la ejecución del programa.