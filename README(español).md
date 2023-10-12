# DNS-Scout
DNS Scout para Linux/MacOS extrae y muestra los registros DNS en una salida de consola codificada por colores que es fácil de ver y copiar/pegar.

Registrador, NS, MX, SPF, DMARC, ASN y PTR para un fácil reconocimiento y resolución de problemas de DNS.

<img src="example-domain.png" alt="Example DNS records" width="800"> 

## Características:

**Salida curada para mayor claridad:**
DNS Scout se destaca por filtrar la información no esencial, presentando a los usuarios una vista más limpia y enfocada de los datos DNS, y optimizando para la claridad y la relevancia.

**Interfaz CLI interactiva mejorada:**
DNS Scout aprovecha ```readline``` para ofrecer una interfaz de línea de comandos avanzada que es **fácil de ver y copiar/pegar.**  

**Ciclado de memoria basado en sesión:**  
La interfaz interactiva de DNS Scout tiene una función de ciclo de memoria, controlada por las teclas de flecha arriba y abajo. Ayuda a navegar rápidamente por las búsquedas recientes de la sesión. Esta función es útil cuando se realizan varias búsquedas y se necesita consultar una entrada anterior.

**Búsqueda de WHOIS simplificada:**  
DNS Scout analiza de manera eficiente los datos de registro de dominios, presentando al usuario detalles del registrador y servidores de nombres concisos, eliminando el desorden que se ve típicamente en las salidas WHOIS sin procesar.

**Visualización clara de registros TXT:**  
DNS Scout enumera los registros TXT en un formato fácil de digerir, lo que hace que tareas como la revisión de la verificación de SPF o DMARC sean más sencillas.

**Registrador**  
Servidores de nombres NS
**Registros MX**  
**Muestra registros TXT, útiles para verificar la verificación de dominio, la configuración de SPF, etc.**  
**Registros DMARC**  
**PTR**  
**ASN**  
**Datos DNS exactos, sin desplazamiento**  

### Guía de instalación de DNS Scout  
#### Instalación manual de Nerd para MacOS/Linux  
Requisitos previos: Go 1.21
Para aquellos que les gusta ensuciarse las manos:

1. Descargar el binario:
Descargue el binario compilado para su sistema operativo desde la página Releases: https://github.com/careyjames/dns-scout/releases.

2. Hazlo ejecutable:
Después de descargar, navegue al directorio de descarga y ejecute:

```chmod +x dns-scout-<version>``` (macos-silicon linux-amd64)

3. Mover a PATH:
Mueva el ejecutable a un directorio en la PATH de su sistema. Por ejemplo, puede moverlo a ```/usr/local/bin/``` en un sistema basado en Unix/Mac:
```sudo mv dns-scout /usr/local/bin/```

4. Obtén un token gratuito o de pago de ```https://ipinfo.io```

5. Ejecutar DNS Scout:
Abra una nueva ventana de terminal y escriba dns-scout para comenzar a usar la herramienta.

¡Eso es! Ha instalado manualmente DNS-Scout como un verdadero nerd.

**Aquí hay un desglose de cómo cada método de almacenamiento del token de API podría ser útil:**  

Variable de entorno: útil para usuarios que ejecutan el programa en un entorno controlado como un servidor, donde establecer variables de entorno es una práctica común.  
La secuencia de comandos ```/share/setup-api-token.sh``` les sería útil.

Argumento de línea de comandos: útil para aquellos que desean especificar diferentes tokens de API para diferentes ejecuciones sin cambiar las variables de entorno. Podría ser útil para pruebas.

Almacenado en un archivo: ideal para usuarios habituales que desean establecer el token de API una vez y olvidarlo. El token se leerá de un archivo en el directorio de inicio del usuario, lo que lo hace conveniente para ellos.

Si estás en MacOS, ve a Ajustes del sistema > Seguridad y privacidad y dale a ```dns-scout-<version>``` permisos de disco completo.