# YT-DL

Este es un proyecto para descargar audio de videos de YouTube y enviarlos a un chat de Telegram.

## Instalación

### Prerrequisitos

- Docker
- Docker Compose

### Pasos

1. **Clona el repositorio:**

   ```bash
   git clone https://github.com/tu-usuario/tu-repositorio.git
   cd tu-repositorio
   ```

2. **Configura las variables de entorno:**

   Crea un archivo `.env` a partir del archivo `.env.template` y completa los valores:

   ```bash
   cp .env.template .env
   ```

   Edita el archivo `.env` con tus valores:

   ```
   TELEGRAM_BOT_TOKEN=tu-token-de-telegram
   TELEGRAM_CHAT_ID=tu-chat-id-de-telegram
   ```

3. **Construye la imagen de Docker:**

   Puedes construir la imagen usando el `Makefile`:

   ```bash
   make build-image
   ```

   O directamente con Docker:

   ```bash
   docker build -t yt-dl .
   ```

4. **Ejecuta el contenedor:**

   Usa Docker Compose para iniciar el servicio:

   ```bash
   docker-compose up -d
   ```

## Endpoints

- `GET /`: Muestra la página principal con la lista de archivos descargados.
- `POST /download`: Descarga el audio de un video de YouTube.
  - **Body (JSON):**
    ```json
    {
      "url": "la-url-del-video"
    }
    ```
- `GET /logs`: Muestra los logs en tiempo real.

