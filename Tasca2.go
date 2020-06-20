package main
import( 
    "fmt"
    "flag"
    "time"
    "net/http"
    "sync"    
)

//funció per mesurar la latencia (temps de resposta a una petició)
func RealicePetition(url string, totalLatency *float64, nErrors *int, k bool){
    //enregistrem un contat de temps a start
    start := time.Now()
    
    
    //realitza la petició
    resp, err := http.Get(url)
    
    //detecta les peticions fallides i enregistrem el nombre d'aquetes
    if err != nil {
        *nErrors++
    }
    
    //per activar KeepAlive s'ha de fer un close del body
    if(k==true){ 
        defer resp.Body.Close()
    }
    //actualitzem valors    
    latency := time.Since(start).Seconds()
    *totalLatency = *totalLatency + latency
    
}

func main(){
    //declarem totalLatency per asignarli el tipus float64. aqui tindrem la average latency i mes tard servira per calcular el tps
    var totalLatency float64 = 0
    
    //declarem nErrors per calcular mes tard el % d'errors de les peticions
    nErrors := 0
    
    //obtenim la url a fer peticions, el nombre de peticions dessitjades i el nombre de concurrencia mitjançant els flags
    url := flag.String("url", "url", "url a qui fer les peticions")
    n := flag.Int("n", 1, "nombre de peticions")
    c := flag.Int("c", 1, "nivell de concurrencia")
    k := flag.Bool("k", false, "KeepAlive enable o desable")
    
    flag.Parse()

    aux := 0;
    
    //funcio per aplicar concurrencia i limitarla en -c
    nivellConcurrencia := flag.Int("nivellConcurrencia", *c, "nombre de subrutines maxim (nivell de concurrencia)")
    nPeticions := flag.Int("nPeticions", *n, "nombre de peticions que farem")
    flag.Parse()
    
    concurrentGoroutines := make(chan struct{}, *nivellConcurrencia)
    //es crea un WaitGroup per utilitzarlo al final del bucle de subrutines
    var wg sync.WaitGroup
    
    //inici del for de subrutines per fer n peticions
    for i := 0; i < *nPeticions; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            concurrentGoroutines <- struct{}{}
            RealicePetition(*url, &totalLatency, &nErrors, *k)            
            //indica cada 100 requests completes l'estat del proces al client
            if(i%100==0){
                aux++;
                fmt.Printf("Completed %v00 requests \n", aux)
            }
            <-concurrentGoroutines
        }(i)
    }
    //esperem a que totes les subrutines acabin
    wg.Wait()
    
    //indica que hem acabat de fer les peticions
    fmt.Println("Finished ",*n, "requests")
    
    //Imprimir per pantalla els resultats del testing    
    //url on estem fent el test
    fmt.Println("Server Hostname:", *url, "\n")
    
    //Average latency
    fmt.Printf("Time taken for tests: %.3f seconds \n", totalLatency)
    
    //Errored responses (amount, percentage %)
    fmt.Println("Failed requests:", nErrors)
    fmt.Println("% Failed requests:", ((nErrors*100)/(*n)), "%")
    
    //Transactions Per Second (TPS)
    fmt.Printf("Requests per second: %.3f [#/sec] (mean)\n", ((float64(*n))/totalLatency))    
}
