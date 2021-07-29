INTEGRACION DE MERCADO LIBRE PARA LABOROTARIO 3 DE LA UBP - Johana Testa
========================================================================


Esto es un proyectp que se lleva a cabo con el lenguaje de programacion go que consiste consumir los servicios que brinda mercado libre a traves de su api




Comenzando 🚀
==============
Para clonar el proyecto https://github.com/johatesta/API-GO-MELI-INTEGRATION.git
git clone 
Descargar dependencias desde la terminal
 go get
con esto ya empezas a trabajar con la API
Cerrar sesión de mercadolibre
.Ejecutar el archivo main.go (go run main.go) para iniciar el servidor local en el puerto :8080, e iniciar el servicio de Apache y MySQL en XAMPP para usar la base de datos (Aclaracion: Se debe tener creada la Database netspace y la tabla items dentro de la misma para el correcto funcionamiento, ademas de la estructura misma de la tabla) .
. Una vez que se este ejecutando el servidor local y XAMPP, iremos a: http://localhost:8080/ingresar .
. Iniciaremos sesion en nuestra cuenta y daremos los permisos correspondientes a la aplicacion. (si usted ya esta logueado, se lo redijira directamente al dashboard)
. Una vez hecha la autenticacion, se lo redirijira al dashboard donde en la parte superior izquierda estara su "nickname" de la plataforma, verificando asi que se haya logueado correctamente.

Mercado libre y la autenticacion 🔧
====================================
Al iniciar el flujo de autorización, la aplicación que desarrolles deberá redireccionar a Mercado Libre para que los usuarios puedan autenticarse y posteriormente autorizar tu aplicación. En el navegador ingresa la siguiente dirección:
https://auth.mercadolibre.com.ar/authorization?response_type=code&client_id=$APP_ID&redirect_uri=$YOUR_URL

EN ESTE CASO NUESTRA URL QUEDARÍA DE LA SIGUIENTE MANERA: https://auth.mercadolibre.com.ar/authorization?response_type=code&client_id=5291933962243912&redirect_uri=http://localhost:8080/auth

Al poner esta url en nuestro navegador nos devolverá otra url con un codigo: http://localhost:8080/auth?code=TG-5faeb662a8096e0007167cd4-398763624




Ejecutando las pruebas ⚙️
=============================
2. El apartado de Preguntas no redijira a una seccion donde se mostraran las preguntas realizadas en los items del usuario actualmente logueado junto con un boton correspondiente a la pregunta para ser respondida dentro de la plataforma, seguido de un mensaje de confirmacion de envio de respuesta.
3. En el Apartado de Estadisticas, veremos valores correspondientes a los datos guardados por el usuario en la base de datos. Esto varia segun el usuario logueado y la informacion guardada por el mismo.
4. Y por ultimo tenemos el apartado Productos que consta de 3 secciones.
4. 1. La primera es de publicaciones donde el usuario puede ver los items publicados por el mismo, junto con un boton de guardado para almacenar en la base de datos.
4. 2. La segunda seccion es para Crear una Publicacion en MercadoLibre desde la plataforma NetSpace. Se despliega un formulario en la pantalla para completar los datos del itema a publicar, y una vez completado se publica mostranso posteriormente un mensaje de confirmacion.
4. 3. Por ultimo, esta la seccion de Ventas donde el usuario puede acceder a ver las Ventas recibidas en su cuenta, ademas de los datos correspondientes a cada venta en particular.

SE UTILIZARON LOS SIGUIENTES USUARIOS PARA EJECUTAR PRUEBAS
============================================================

USUARIO COMPRADOR 1
{   "id": 798890199,
    "nickname": "TT614680",
    "password": "qatest5397",
    "site_status": "active",
    "email": "test_user_91934992@testuser.com"
}
    
USUARIO COMPRADOR 2
{
    "id": 798922756,
    "nickname": "TESTNF5WF7ZU",
    "password": "qatest6083",
    "site_status": "active",
    "email": "test_user_31086430@testuser.com"

}


USUARIO VENDEDOR 
{
      
    "id": 798719013,
    "nickname": "TESTEJ4QQXY",
    "password": "qatest5154",
    "site_status": "active",
    "email": "test_user_56844852@testuser.com"

}

