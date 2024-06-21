

## Background
Users wanted more intelligent search capabilities where users could type in hiragana, katakana, Kanjis or various combinations of it and still get relevant results.
- Japanese token dictionary `ipadic` is powerful but limited in figuring out relevant kanjis.
Google-backed dictionary `mozc` is more powerful and does Kanji translations based on context.
- This project aims to return a combination of
  - User-typed search keywords as is.
  - Search keywords converted to Hiragana. (Only katakana keywords are converted to hiraganas)
  - Feed Hiragana search keywords combination to `mecab` and leverage `mozc` dictionary to get relevant results. Several combinations (N) needed can be specified and customized.

## Dependencies
- Docker
- docker-compose

## How to test and deploy

```sh
docker-compose up --build
```

## Example

- Once the application is up and running, you can check the results like this:
- http://localhost/results?q=さかいさんは元気ですか&&n=4