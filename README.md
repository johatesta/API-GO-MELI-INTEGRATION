INTEGRACION DE MERCADO LIBRE PARA LABOROTARIO 3 DE LA UBP - Johana Testa
========================================================================


Esto es un proyectp que se lleva a cabo con el lenguaje de programacion go que consiste consumir los servicios que brinda mercado libre a traves de su api




Comenzando 🚀
==============
Para empezar a trabajar con la api a traves de la terminal de go podes usar el comando go get -u 
https://github.com/johatesta/API-GO-MELI-INTEGRATION con esto ya empezas a trabajar con la API
Pre-requisitos 📋
Que cosas necesitas para instalar el software y como instalarlas



Mercado libre y la autenticacion 🔧
====================================
Al iniciar el flujo de autorización, la aplicación que desarrolles deberá redireccionar a Mercado Libre para que los usuarios puedan autenticarse y posteriormente autorizar tu aplicación. En el navegador ingresa la siguiente dirección:
https://auth.mercadolibre.com.ar/authorization?response_type=code&client_id=$APP_ID&redirect_uri=$YOUR_URL

EN ESTE CASO NUESTRA URL QUEDARÍA DE LA SIGUIENTE MANERA: https://auth.mercadolibre.com.ar/authorization?response_type=code&client_id=5291933962243912&redirect_uri=http://localhost:8080/auth

Al poner esta url en nuestro navegador nos devolverá otra url con un codigo: http://localhost:8080/auth?code=TG-5faeb662a8096e0007167cd4-398763624
y un archivo en formato JSON que nos dará el token de acceso y el id de usuario con el que trabajaremos para acceder a los demas recursos que nos brinda la api de MercadoLibre

{"Access_token":"APP_USR-5291933962243912-111316-619d68ab8adfbbe03fcd70dc1bf16bc1-398763624","Token_type":"bearer","Expires_in":21600,"Scope":"read write","User_id":398763624,"Refresh_token":""}



Ejecutando las pruebas ⚙️
=============================
/items/all?token=$ACCESS_TOKEN&userid=$USER_ID Este endpoint devuelve todos los items con sus respectivas preguntas de un vendedor y las ventas concretadas. Trayendo un JSON como el siguiente:
[{"Id":"MLA896876339","Title":"No Ofertar Item De Prueba Para Api","Price":300,"Quantity":2,"SoldQuantity":0,"Picture":"http://http2.mlstatic.com/D_789705-MLA44060588210_112020-O.jpg","Question":[{"date_created":"2020-11-12T10:20:58.359-04:00","item_id":"MLA896876339","status":"UNANSWERED","text":"Tenes otros colores?","id":11598329310,"answer":""}]}]
