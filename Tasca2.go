package main
import( 
"fmt"
"flag"
"time"
"net/http"

)

//funció per mesurar la latencia (temps de resposta a una petició)
func RealicePetition(url string, totalLatency *float64, nErrors *int){
    //enregistrem un contat de temps a start
    start := time.Now()
    
    //realitza la petició
    result, failReq := http.Get(url)
    
    //detecta les peticions fallides i enregistrem el nombre d'aquetes
    if failReq != nil {
        *nErrors++
    }
    defer result.Body.Close()
    
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
    //per poguer compilar (BORRAR quan estigui implementat c i k)
    fmt.Println("c: ",*c)
    fmt.Println("h: ",*k)
    
    //indica que comença el test al usuari
    fmt.Println("Benchmarking", *url, "(be patient)")
    aux := 0;
    
    //bucle per fer n peticions
    for i := 0; i < *n; i++ {
        //mostrem per pantalla un missatge cada 100 peticions per donar una referencia del estat del proces a l'usuari
        if(i%100==0){
            aux++;
            fmt.Printf("Completed %v00 requests \n", aux)
        }
        RealicePetition(*url, &totalLatency, &nErrors)
    }
    
    fmt.Println("Finished ",*n, "requests")
    
    //Imprimir per pantalla els resultats del testing    
    //url on estem fent el test
    fmt.Println("Server Hostname:", *url, "\n")
    
    //Average latency
    fmt.Println("Time taken for tests:", totalLatency, "seconds")
    
    //Errored responses (amount, percentage %)
    fmt.Println("Failed requests:", nErrors)
    fmt.Println("% Failed requests:", ((nErrors*100)/(*n)), "%")
    
    //Transactions Per Second (TPS)
    fmt.Println("Requests per second:", ((float64(*n))/totalLatency), "[#/sec] (mean)")
    
}
