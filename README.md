# TechnicalTest_Owlint
---
Build the image with:

```sh
$> docker build -t technical_test_owlint .
```
Then run it with
```sh
$> docker run -p 4242:8080 technical_test_owlint
```
You can replace the `4242` by the port you prefer.

> For Mac user, add the environment variable *LOCALHOST="0.0.0.0"*:
> ```sh
> $> docker run -p 4242:8080 -env $LOCALHOST="0.0.0.0" technical_test_owlint
> ```
> It allows you to listen on $LOCALHOST IP instead of localhost (usually  127.0.0.1).


Note: You can use the interactive terminal to see the database:

```sh
#inside the container
root$ mongosh
```
The database is the `technical_test_owlint`, and the collection is `comment`.

---
Note: In a respectable team it should have more test 😅🫣😆