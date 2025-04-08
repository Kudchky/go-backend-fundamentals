
# 🧠 # The Secret to Solving Any Problem

1. Understand the problem
2. Make a plan
3. Carry out the plan
4. Look back and improve your solution 

Desarrollar un programa en Go que genere un número aleatorio entre 1 y 100, y 
permita al usuario adivinarlo en un máximo de 10 intentos. El programa debe 
proporcionar pistas ("mayor" o "menor") después de cada intento fallido y 
mostrar un mensaje de victoria o derrota al finalizar.

## 1. 📖 Entender el Enunciado

- [ ] Identificar requisitos clave:

    1. Generar un numero aleatorio entre 1 y 100.
    2. Permitir que al usuario adivinarlo maximo en 10 intentos.
    3. Dar pistas mayor o menor despues de c/intento fallido.
    4. Mostrar mensajes de victoria o derrota al finalizar juego.

## 2. 🧱 Diseñar la Estructura
- [ ] Pseudocódigo mental:
    - Entrada: Numeros ingresados por el usuario.
    
    - Salida: Mensajes de salida, victoria o derrota.

    - Proceso:
        1. Generar número aleatorio (1-10).
        2. Para cada intento (10 máx):
           a. Pedir input al usuario.
           b. Validar si es número.
           c. Comparar con el número aleatorio.
           d. Dar pista o terminar si acierta.
        3. Si no acierta, revelar el número.

## 3. 🧰 Elegir Herramientas
- [ ] Paquetes necesarios: "fmt", "math/rand"
- [ ] Estructuras que usaré (if, for):

## 4. 🛠️ Implementar en Etapas

### Etapa 1: Funcionalidad base
- [ ] Escribir código mínimo que funcione

### Etapa 2: Validaciones
- [ ] Manejo de errores
- [ ] Rango de valores, entradas inválidas, etc

### Etapa 3: UX Mejore la experiencia
- [ ] Mensajes claros
- [ ] Indicaciones al usuario

### Checklist
- [x] Entendí el problema y sus restricciones.  
- [] Escribí pseudocódigo/diagrama.  
- [] Elegí herramientas (paquetes, estructuras).  
- [] Implementé una versión mínima funcional.  
- [] Añadí validaciones y manejo de errores.  
- [] Optimicé y mejoré la UX.  
- [] Revisé casos límite y refactoricé.  
