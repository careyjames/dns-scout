# DNS-Scout 🇨🇴 Carey James Balboa - Mac Help Nashville, Inc

DNS Scout para Linux/macOS extrae y muestra los registros DNS en una
salida de consola codificada por colores que es fácil de ver y copiar/pegar.

Registrador, NS, MX, SPF, DMARC, ASN y PTR para un fácil reconocimiento y
resolución de problemas de DNS.

![Example DNS records](example-domain.png)
![Example IP records](example-IP.png)

## Características

**Salida curada para mayor claridad:**
DNS Scout se destaca por filtrar la información no esencial, presentando a
los usuarios una vista más limpia y enfocada de los datos DNS, y optimizando
para la claridad y la relevancia.

**Interfaz CLI interactiva mejorada:**
DNS Scout aprovecha ```readline``` para ofrecer una interfaz de línea de
comandos avanzada que es **fácil de ver y copiar/pegar.**

**Ciclado de memoria basado en sesión:**
La interfaz interactiva de DNS Scout tiene una función de ciclo de memoria,
controlada por las teclas de flecha arriba y abajo. Ayuda a navegar rápidamente
por las búsquedas recientes de la sesión. Esta función es útil cuando se
realizan varias búsquedas y se necesita consultar una entrada anterior.

**Búsqueda de WHOIS simplificada:**
DNS Scout analiza de manera eficiente los datos de registro de dominios,
presentando al usuario detalles del registrador y servidores de nombres
concisos, eliminando el desorden que se ve típicamente en las salidas
WHOIS sin procesar.

**Visualización clara de registros TXT:**
DNS Scout enumera los registros TXT en un formato fácil de digerir,
lo que hace que tareas como la revisión de la verificación de SPF o DMARC
sean más sencillas.

**Registrador**
Servidores de nombres NS
**Registros MX**
**Muestra registros TXT, útiles para verificar la verificación de dominio,**
**la configuración de SPF, etc.**
**Registros DMARC**
**PTR**
**ASN**
**Datos DNS exactos, sin desplazamiento**

### Guía de instalación de DNS Scout
[![Instalar desde Snap Store](https://snapcraft.io/static/images/badges/es/snap-store-white.svg)](https://snapcraft.io/dns-scout)
#### Instalación manual de Nerd para MacOS/Linux

Requisitos previos: Go 1.21
Para aquellos que les gusta ensuciarse las manos:

1. Descargar el binario:
    Descargue el binario compilado para su sistema operativo desde la
    página Releases: [Release](https://github.com/careyjames/dns-scout/releases).

1. Hazlo ejecutable:
    Después de descargar, navegue al directorio de descarga y ejecute:

    ```chmod +x dns-scout```  
    
1. Mover a PATH:
    Mueva el ejecutable a un directorio en la PATH de su sistema. Por
    ejemplo, puede moverlo a ```/usr/local/bin/``` en un sistema basado en Unix/Mac:
    ```sudo mv dns-scout /usr/local/bin/```

1. Obtén un token gratuito o de pago de ```https://ipinfo.io```
    [Website](https://ipinfo.io)

1. Ejecutar DNS Scout:
    Abra una nueva ventana de terminal y escriba dns-scout para comenzar
    a usar la herramienta.

¡Eso es! Ha instalado manualmente DNS-Scout como un verdadero nerd.
ahora hay .deb también para los nerds de Debian.

**Aquí hay un desglose de cómo cada método de almacenamiento del token**
**de API podría ser útil:**

Variable de entorno: útil para usuarios que ejecutan el programa en un
entorno controlado como un servidor, donde establecer variables de
entorno es una práctica común.
La secuencia de comandos ```/setup-api-token.sh``` les sería útil.

Argumento de línea de comandos: útil para aquellos que desean especificar
diferentes tokens de API para diferentes ejecuciones sin cambiar las
variables de entorno. Podría ser útil para pruebas.

**Si estás en MacOS**, ve a Ajustes del sistema > Seguridad y privacidad
y dale a ```dns-scout``` permisos.

![Dev not Verified](dev-not-verified.png)
![Example IP records](mac-click-allow.png)
