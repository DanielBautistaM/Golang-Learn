// Función que se ejecuta cuando se hace clic en la imagen
function saludar() {
    // Muestra una alerta con un mensaje
    alert("¡Hola desde Go!");
    
    // También podemos manipular el DOM
    console.log("Imagen clickeada");
}

// Ejemplo de función que se ejecuta cuando el documento está listo
document.addEventListener('DOMContentLoaded', function() {
    // Podemos agregar más funcionalidad aquí
    console.log("Documento cargado completamente");
});
