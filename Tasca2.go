package main
import( 
    "fmt"
    "flag"   
)

/*funció per mesurar la latencia (temps de resposta a una petició)
func RealicePetition(url string, totalLatency *float64, nErrors *int){

}
*/

func main(){

    //declarem totalLatency per asignarli el tipus float64. aqui tindrem la average latency i mes tard servira per calcular el tps
    var totalLatency float64 = 3000.78
    
    //declarem nErrors per calcular mes tard el % d'errors de les peticions
    nErrors := 1
    
    //obtenim la url a fer peticions, el nombre de peticions dessitjades i el nombre de concurrencia mitjançant els flags
    url := flag.String("url", "url", "url a qui fer les peticions")
    n := flag.Int("n", 1, "nombre de peticions")
    c := flag.Int("c", 1, "nivell de concurrencia")
    k := flag.Bool("k", false, "KeepAlive enable o desable")
    
    flag.Parse()
    
    //imprimir per pantalla per testejar si he agafat els flags correctament
    fmt.Println("url:", *url)
    fmt.Println("n:", *n)
    fmt.Println("c:", *c)
    if (*k){
        fmt.Println("-k enable\n")
    }else{ 
        fmt.Println("-k desable\n")
    }
    
    for i := 0; i < *n; i++ {
      //  RealicePetition(*url, &totalLatency, &nErrors)
    }
    
    
    
    //Imprimir per pantalla els resultats del testing
    fmt.Println("Transactions Per Second (TPS):", (totalLatency/float64(*n)))
    fmt.Println("Average latency:", totalLatency)
    fmt.Println("Errored responses (amount, percentage %):", nErrors,",", ((nErrors*100)/(*n)), "%")
}
