A “Real” World Example
In one of my production applications, I was tasked with creating an API that interfaced with a tonne of other APIs and aggregated the results up into one response.

Each of these API calls took roughly 2-3 seconds to return a response and due to the sheer number of API calls I had to make, doing this synchronously was out of the question.

In order to make this endpoint usable, I would have to employ goroutines and perform these requests asynchronously.

package main

import (
    "fmt"
    "log"
    "net/http"
)

var urls = []string{
    "https://google.com",
    "https://tutorialedge.net",
    "https://twitter.com",
}

func fetch(url string) {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(resp.Status)
}

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("HomePage Endpoint Hit")
    for _, url := range urls {
        go fetch(url)
    }
    fmt.Println("Returning Response")
    fmt.Fprintf(w, "All Responses Received")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
    handleRequests()
}

However, when I first employed this tactic, I noticed that any calls I made to my API endpoint were returning before my goroutines had a chance to complete and populate the results.

This is where I learned of WaitGroups and dived deeper into the sync package.

By employing a WaitGroup I could effectively fix this unexpected behavior, and only return the results once all of my goroutines had finished.
package main

import (
    "fmt"
    "log"
    "net/http"
    "sync"
)

var urls = []string{
    "https://google.com",
    "https://tutorialedge.net",
    "https://twitter.com",
}

func fetch(url string, wg *sync.WaitGroup) (string, error) {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        return "", err
    }
    wg.Done()
    fmt.Println(resp.Status)
    return resp.Status, nil
}

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("HomePage Endpoint Hit")
    var wg sync.WaitGroup

    for _, url := range urls {
        wg.Add(1)
        go fetch(url, &wg)
    }

    wg.Wait()
    fmt.Println("Returning Response")
    fmt.Fprintf(w, "Responses")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
    handleRequests()
}

Now that I’ve added the WaitGroup to this endpoint, it will perform a HTTP GET request on all of the URLs listed and, only upon completion, will return a response to the client calling that particular endpoint.

$ go run wg.go
HomePage Endpoint Hit
200 OK
200 OK
200 OK
Returning Response