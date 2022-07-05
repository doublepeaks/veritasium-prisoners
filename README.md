# veritasium-prisoners
Testing out the riddle from a [veritasium video](https://www.youtube.com/watch?v=iSNsgj1OCLA)

Results check out!

## How To Run

### Go Version
````
cd go-prisoners
go build
./prisoners 10 10 100 random
2022-07-05T13:27:34Z 0.00% escapes

./prisoners 10 10 100 sequence
2022-07-05T13:28:09Z 33.00% escapes
````

### Python Version
````
cd py-prisoners
./prisoners.py 10 10 100 random
0.00% escapes

./prisoners.py 10 10 100 sequence
35.00% escapes
````
