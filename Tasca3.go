package main
import( 
    "net/http"
    "fmt"
)

//creem l'estructura Handler (Interfaç que controla les peticions de http entrants)
type HttpHandler struct{}

//implementació del HTTP server amb l'estructura HttpHandler
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request){
    
    //definim la resposta del server mitjançant la funció per imprimir per pantalla FMT.Fprint
    fmt.Fprint(res, "Hola món ! / Hello World !")
}


func main(){
    
    //creem el handler
    handler := HttpHandler{}
    
    //iniciem el proces ListenAndServe per escoltar peticions
    http.ListenAndServe(":9000", handler)
}
