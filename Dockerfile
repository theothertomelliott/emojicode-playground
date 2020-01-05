FROM theothertomelliott/emojicode:0.8.4

COPY . /app
WORKDIR /app

EXPOSE 8080

ENTRYPOINT /app/emojicode-playground