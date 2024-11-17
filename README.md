## Run in Docker

1. Build Docker Image
```bash
docker build -t <IMAGE_NAME>:<TAG> .
```

2. Run as Container
```bash
docker run -d --name <CONTAINER_NAME> -p 2701:2701 \
-e PORT=2701 \
-e DATABASE_URL=postgresql://orias:orias@127.0.0.1:5432/db_orias?sslmode=disable \
-e SALT_KEY=12 \
-e JWT_SECRET_KEY=Orimi_027@ \
-e AES_SECRET_KEY=INDONESIAORIAS27 \
-e GOOGLE_AUTH_CLIENT_ID=1028992494328-p8joiknfk0otg90ivl8s5s057kdse2gg.apps.googleusercontent.com \
-e GOOGLE_AUTH_CLIENT_SECRET=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 \
-e GOOGLE_AUTH_REDIRECT_URL=http://localhost:2701/api/v1/google_auth/callback \
-e GOOGLE_OAUTH_STATE_STRING=Orias_271@ \
-e GMAIL_EMAIL=oriastan999@gmail.com \
-e GMAIL_PASSWORD=qjcylukvwpqibepr \
-e EMAIL_VERIFICATION_LINK=http://localhost:2701/api/v1/auth/verify-email \
-e EMAIL_FORGET_PASSWORD_FRONTEND_LINK=http://localhost:3000/forget-password \