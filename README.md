### Nota a errores

EL siguiente error: ERROR_POLICY: attempt to perform an operation not allowed by the security policy `PDF' @ error/constitute.c/IsCoderAuthorized/408

## se debe a un problema de seguridad, màs detalles en este link: https://bugs.launchpad.net/ubuntu/+source/imagemagick/+bug/1796563

La soluciòn fue comentar la linea correspondiente al formato a convertir:

- cd /etc/ImageImagick-6
- sudo nano policy.xml

- y comente la linea donde se encuentre:
 <code> policy domain="coder" rights="none" pattern="PDF" </code>



## librería para convertir jpg to webp
https://github.com/kolesa-team/go-webp


#### Imagick golang - Instalación - Link: https://gopkg.in/gographics/imagick.v2/imagick
go get gopkg.in/gographics/imagick.v2/imagick

import "gopkg.in/gographics/imagick.v2/imagick"

