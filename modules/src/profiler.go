package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.9" )

func aT9OK0zZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZbfRpDA6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func niPD2fhwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9LreW8i5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k6nFSdgaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ZEpAo6sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MlSwyJTcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7vtJQ7qrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2o4ncJTZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func viigzLN3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7fu6F0uDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func thZS6YAPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NT00vR1RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hIVtSIeEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OFcxjNsXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GU9RnXMsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hJ9vWqTpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6HeN2sqnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j9V6JNtDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dQu2UUnzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fow2h6oHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j1zjq7QqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jM1XhccbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Osp4aSCWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q9MBcOZvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KGEeR5pJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wkBVR0hfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 73bEWaQkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RXFt3B7FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p5ZyNstSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gsO7GIJ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1pPT45ERWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cb1bn2xMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func daURnKlqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 38BkdRigWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SxMeGPKVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UKJJOz8WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dlY8VzAfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xqqCn3nuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ve6vmf5UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V8jsZ6knWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cjsQhF5NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SgnclYLZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pP2vU1zRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q54K7YdZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PJJDKeXxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2herLmZoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fUYKsEMZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1oWePWstWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u5fxwW47Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WKaBWr3LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hbzWvnIUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZKmJGD4AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OtqXeE5pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zywpaBGCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 68flS31XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p2uv2UOFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WPhAUZO0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aYylz3peWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0mZTKPwJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FWZKwMG7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vDeaawDgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Btu4BcKBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rxDiXdarWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wCo1i0JoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZeOcAgaFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ALJcKtWFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WuXG7Bb1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VLOk4A5UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7XdUF0oKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eou0vPOyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CbiP0AfKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BNTCARTBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rl4UVdv4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1APudtPpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0xaXQamBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AOuzvP7rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gPhoYToRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KsE3r1d2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xfm6paGwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GimNDgcHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ndn2N4nZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2rxeKXDQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fMzc9w1cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hnSdvLGFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ttww56CqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G5mXoMNDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9UNrgRmGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2CuspcwVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2iYf9vUCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xb2FjmVgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JSRDAe1XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U6kUSpLOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3oqKjqVJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FzpcEafqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func knqNItkXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pcHJY598Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZIwt5SivWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G1sEgGRiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uAxqIoGKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UFnjAe6FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZXOMwOpMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eLpU5gb5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t7tRZ5xDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func STEOLpqRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1yGNjygvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4tQZe6RUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H79hPkn2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8FmmY51nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q7u7anHwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func inGsmTpYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A6SnPuPdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3GbRb9LLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 76WeoVVPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NrQm28q0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VvnYQY6sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j3l8Wn7pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hUI7hnILWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q1SU2dwPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ugPmzIpyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DBpk9NKNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1WK05x80Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H5vyBieQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bSB6NDJcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IcpbpX0kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JSiV3urSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zA74uQEoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2qJ5HzwwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P6tW0oPUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yCT2xtYKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jna2M4FSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uhPurfDpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A2zW06XcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vEn0XapMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iPI4skcjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j9L6D9p1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rNR75BNyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CaV9I7mnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LITadNWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 95FdBpYZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4LgBeIMjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mh0FNFxjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tr9Z7D2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5r0vxV9LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OwkjgRpHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CJLlzoFtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MzC9QftGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hfHtEuPcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wb0bMBZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rQyt4no7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qX0YYRqPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xTQmnaiMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vhZe82YFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0rrThizCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oJOFtUTGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4eLflahOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f6PWe57fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vrUBbT4GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V3tTbwEvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4oFvmOiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MMAnuppqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qgbCMcEEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GmqmnzbSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u6zEM3CSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MH2i5NZKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 80Q3wUyfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yoLibfvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yhBTBmM0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L0ZYqV2eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u2Jr7WfUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lNgXhYoyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jn8t1xhTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func geuckNNqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cUuzh5BwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J8UOI1GsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ucTKg7FfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ICpaXlGzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1UmBZA4WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vMCZsNBEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 74zUmXKqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SyOLODN7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AZjFCuwSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yArFsRcyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 228jy3XZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eSAVuReXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xrpKK1aqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C80bSzfmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c4bXGpVBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zGEWycWOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yDCKw512Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7mteFIsHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RHTw972eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S9ApAeLeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xr5jS4HAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yj7XYSvSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FBUTeItTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TzgrhqQBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z16sCWTeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UVzu5XwBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2eFOEsNXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VpUdrozhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KGijmKbqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z8MgaHZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4AnBB6xVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7LxVWQnoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0HD2fl59Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bzSZ7ZO6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LZR2wLdeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bI6E9XuQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jdsIKkyWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P36gLNxJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mZUuwWXQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func djzvUXFoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oHsBeDOdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BgcVbHpQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FGBlRDF1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6OoPZuT8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TMyHtRKUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UuKjZ0oeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9cIsXMKxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ppq1NL90Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M2TCnV2BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cdXxiyKCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VPOfBhIVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TU5QlQZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VdaNMQ8WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0E8uM7X5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t2FbnfPVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LhK9op2IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FZPEE7l6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uFICxHsxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VPunsKPOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ZO4TboSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ElFPJrT9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s1buYg6RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 30WebQTlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BINTB1PfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MeAEngd5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RnfG9cBmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VK3CcPA3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HThISmUNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ns1n4SbRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EiPVR3iJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pasUr0lQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0q8XDTbnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ycdiy5HrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eMRc3QeuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2szy59WhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3c6WVVrRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EAvmitD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 33AcZDulWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 77eflAvoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iX609MoMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xKhuXwIZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FyHS8XcEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NVOzM5W8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MJ9A5CWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tRDsJPjZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FxCVIM0jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UF4N1QuIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UCuQFQoUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nApyPfLvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rzv8lRb9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7v4LruEHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M8gfDsSmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UusY99fFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cnOa3jyqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qaS0SKD3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GTIRBYfFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GJqcrrzaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func syH76XUAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yfsKVxrkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C3OjyNKJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ID8dgzzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ezHdimN9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q4YVMTkEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LjyjkfW8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nmNepTYjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3QLnLW4uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hroENrduWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FG2303HzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iApvTO8MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TVkmS484Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DjmCC0zTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fCJsqj32Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WdlHUmcbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qTKVxm10Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eozCwybkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pSjeds81Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JNyunbCCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EF2safz8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gbIOKqafWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eTr4U9XFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gM1dvJD2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ezfSru4XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func REVzJRWMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BxSVajP1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RtPZUCtxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JdLO3YAlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ALrEbpsgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hqKvQtklWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bNyjPUdVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cZXIeGsLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2sCOmWHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qb0ruTa5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CEkt8Fp4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DNHHXIaeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Xpq5fUwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eXdygy8JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z2NMBBhWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AxbZxpKzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x1Wq4bnfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6o2e1n8lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HNCroSCrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uwVwT8GzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yhyvUYJ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HVmxX0FMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WYu9K86xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WsLot2LFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5592iUyGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3JzJ1mf7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A6Qd6A7eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jgmJsENMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JaJ2rPCtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zZF4c2LgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L636nYAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func COBF14WFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tx8XSJL8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F9LCPFIXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LFwepX7ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MsUvrWOCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eS4oSXSHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o7eyNPbGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZhBU8CY8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NwAj8cnCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vbqq2qb5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GfuqjGzoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mViuIBlsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xkj1DPmsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wAtUYzmAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ucx9z5ocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o997YJYCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yRinITroWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XnM4GVLFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kSwR6iaBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kvK13YU5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LFGmCbYAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nKEf0V9wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M8D5YsGXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ebde2rdfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1d4GaLgFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ijDUFmQnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pVYhe19CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7FTCkLQiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cHqlKyLBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 41lOSK2gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ioTJomkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bYaNnTDIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6acDz1J1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LzuYQO4JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bsohfNpQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IXAqXIVeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fwI34OVcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jpAqwGBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5OeOP6BwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nZ1OKoleWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iidujNKpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tvPIrbobWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hMcM4S8oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qskai6jbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1PzSiNS9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HmUHL5EkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0wH3aWy1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BhdlVdXeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2MtIoVdgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SttSjyEBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bKBjaH11Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AJGJyO38Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jgGNUII4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ipsm8Vo8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MNgsiXIGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LnD34pt9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mh0wrZG5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6J4SMcoLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NNciam9kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SwMTQLnyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YrVl79M3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EftuVXdHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eBBd0mrMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WKeyCB5TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1iotejJTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hk8We1KbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NapXBMIUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hNnKeK5QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ocMeaDmVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wlLRqUJBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DseC4lqUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 33ezTU02Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XZCjLiClWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fe1qWItaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ShMe9cKnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K6kcRSnfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h0I11Ge3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MdHYq0DKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6JMaQPdZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func upzgJOSUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wu0zkfAnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WhTXgWqIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Emz5zIxqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5t5mTpdmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KJI62yI4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mC1N9hpAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6McI26chWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7hOfygUiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BNuodjzfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jyeXWpjfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tkmge3vtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9g4IsVUnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pc4lDtKlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0qEyLCRYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RD5Qro39Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PKQB1VzXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9JkLIiX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KyQslbptWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gr5kSeyEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HbJjAZvwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VE4KTpJBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yuoITdL0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func um5J9PesWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SJ4K0jMNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VfjGBxW8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0uqFGQ7oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tvO9SttbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LMWcRCdbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9iDsxQJWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MiuqFtjjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eg4m69BFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c4fSnFF4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uMtfyKfXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lfQiMPTaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HEWhK4pVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bPle39fYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ul5u0WVOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YoT0vk2EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fXEp7aRkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XXyIgcIBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qbl2BjtMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UPpcwlCjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GfRU8ePAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h8JZKOTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0rSk74pjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uuLlzoQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 81TNVa1cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J71LveSDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SdbIuQsSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MajSUwMNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tQPgWERgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8498sfulWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G3vo2AVCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aak7bFXAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mRqxK75QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AhboDJ8QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LGAwWLH3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xWERSYmzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bwh7S6THWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mmDvR2atWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0pEMDVSfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u7ssJARWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7RzOp04vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TzMXFjDtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PBfyB6TMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5TzFriT1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KwYtEvLyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ox0SWW13Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tVBpyXi3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func toFYjjPyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U14rnNnoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9x29afGBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7FgezVQVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zJMH5bwdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZlhYjvRfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GgyEHFnwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mIoTrNRqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9nhtjpPdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DYhypbtnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jNbmpiHiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vaGPzcHEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2rEKanASWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QUd0JurNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H4mfrZ6mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ufl6zvwPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ct8qrzOrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R9XWvGtrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DkCTLXtnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5IMGWbdfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tfWJ8PAXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eeOLK4FkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ZhdTcb9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ZvHSz8fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JF9wpyGVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sDOQfajpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bzjG9HOoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eQpyHIpnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L5GVFQJ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SeKN1JhdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func peTmjSDmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QiiGinRXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iUsCp7TkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2CZsMHHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y9cViHRdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j1tSLtiuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8IU9PIiqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PJbbCbNNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BnPqJiWLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IRM15HmKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CExRU15SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XlB8PxUcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CdVXBUJ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9WLIsRoXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KlvSm9BEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IElMLoAoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func layUR1prWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VSRRbmifWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 512hqMHiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h273gDeBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6EFlDkimWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vMr8xJj0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Q1p2ZGVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zuPhPRaVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BNF1FmcIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RYEFzqENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JcXZdvo6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aBsepd8wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y6a5tnO2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WxL0vXgbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KCyWjPjmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hf8bVuecWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MPtPTbmFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OKx0TixkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u9GyHMG9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uVMNd1jwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gFtaLtiTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f9DaQdG6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AvweZKxRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func semlBhKEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T6MCYYUkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T5mtbS9lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ANOx8tjrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZzgxvGHqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kw19C0mZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yP6ASAi2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yNZeVXYZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 92lplJhlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QGztOt9LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ve9WZTqBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TI0NbTO5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l6b5zRLbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 99th3mjFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GUd9DVDbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ThgXTVXVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o3efXgNFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ODm4YX1zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WG45Z5BbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TTJasQprWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pyXcWhTYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SJkOUlwCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YirrVDIVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HEgQAVf4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2wP396JOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VdkNWQViWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OuAx0JXzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HtBP8AexWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q2yE73PnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gx6SBKr2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3XNKGJzoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QNmnJW8XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nxmmPYwkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aflYw6AJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lYwuZaNeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z54ceTQLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5OmnyEWHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F5L0ja63Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XYnnzCzmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8agkeuYHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L9KMtaLbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xYHxqXktWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PfHTPDaoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WRqz7A9JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JO69war9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xkqb7vNpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZoIoX96CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LeTByYVTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jeRp8JD7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XCoMD4ANWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tEjvpcITWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func elYbpZYKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hJ7DIt4BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HZwGf27TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func twcOAANeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iKQdXeWmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A8eSuv7UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WR98M4B2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3QLCUw3jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WVSira0GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8RGa458CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QMbRQpgRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kaoCI4zpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OsjWJ5akWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oXU3YRhSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DypCzU9pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4XCMt0byWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZSlfLO0VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VAvQP4ReWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 85Ha9AMiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i1YmQ06kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kBimTH86Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func chUdOMoKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MOxIBKuvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VT4C8cFdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Er0gbaojWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gf4dLMQfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pAwA333dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OLV9zrEfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wD8Jkhv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EdH7gyRbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4AyLF68GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ucT8cqeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sZxwBSDpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fp1EVCJVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UXBAdoCMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RLncwrytWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SIL1NgYKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S7wwxof5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eCmV6S4vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aKNNGdNVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I7kRNo1SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Otd8z2IGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IPL5ZrZRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d0ux6JJhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oE6ypsSeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ybj56vrUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xre58BZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jQKzmmHvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xNGn5rDgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NoczbIEDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func msIypfpAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 047wZhw6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CarwECcOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IkrY5chMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pq6EOsCsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UpzEdcM7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G9Y8spA7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PdiguBYDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JbqhDirzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a0IxEqPoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U4RescfWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eeeWYMFQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DVyWdjsGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zDlmJaniWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jalvhhAHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 07YpZrjYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qnuxO8h4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DsMcbwDOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hRXTpsxHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5p35ICZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func beU6I7QMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0zsdxpAUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lu4iEbjwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IcG4hHCcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gxrBeS48Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AQBheIXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 433o3BBpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fYq2Xn61Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ClgCtTd8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ByvoaE02Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aP5YAJXPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xzI7OstVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vtxTNjlqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IExcQOs4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YEbkTwaeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5eLtd3IjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 96ppsw0EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NbVJAlwGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KxaYYPWpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fLx6Z4fnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TIHBzvo4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pd2jW8fcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X1HqCIuHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lh76E0OUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xl0J4AyuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func srZLBEBqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rr9LOXOBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FapJJpA5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8XCrAObZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DKFeVujYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func isfDT8lqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EvCpoCE7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ows3sEs2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zk4WvKp3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lKMGqmE2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pzFm3Vj0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aA5aD76FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NTmKpywCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n3NRxSTeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func veblU1qKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yVolGEWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xEeJW5JTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6avsv7gGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V77WD0YLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ME58z65rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CEAtTQDAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zo0AEXXcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KpFXbKl0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gptPTKajWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hbz7OKznWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JxGW0nPWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kn3KspLRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5PAduMedWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ph0BaZ4dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TSUzsPDIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IkVraEEFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9g6JYsSOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fbBjEROEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NaFvQa0DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func foZow4kLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XQprXirUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eykzqY5TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QXiAFNMWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lyd7bxJlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tInvMMTHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F6dh0OHAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hHMcBrBpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1WLMKZAkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nc5aEUgzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QTAGjUp1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 324UpWJBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lXToGDFKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kzdGMFTDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ztzLRvnkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Q2rwo4zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JsePyQBRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func skKn4egnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BdelTvoaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rr3QRffAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UUQIcxxUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fxmlODiQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y6pnEW99Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kddUuSqqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bnHY47mzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hqFtYHpbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tW9CUDczWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0V7igvZmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BzO1P3iKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gvjFmbb0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WvQQnx2pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pYNjPYJJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GejdTdxdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fC4dMRBuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bkevN4LoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HXknOHb8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6mzmg4sRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

