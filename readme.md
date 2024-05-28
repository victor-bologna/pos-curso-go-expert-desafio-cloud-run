# Desafio Google Cloud Run

Objetivo: Desenvolver um sistema em Go que receba um CEP, identifica a cidade e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin). Esse sistema deverá ser publicado no Google Cloud Run.

# Subir a aplicação localmente/Rodar o Dockerfile

Executar no cmd o seguinte comando: `docker compose up -d`.

# Executar o projeto no Cloud Run

Acessar a seguinte url `https://weather-zipcode-imsmlbaiga-uc.a.run.app/temperature?cep=CEP` Onde CEP é o CEP desejado.