FROM theothertomelliott/emojicode:0.8.4

COPY . /app
RUN mkdir -p /working
WORKDIR /app

EXPOSE 8080

ENTRYPOINT /app/emojicode-playground /working