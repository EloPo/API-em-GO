package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

type Carro struct {
    ID        string   `json:"id,omitempty"`
    Modelo 	  string   `json:"firstname,omitempty"`
    Marca	  string   `json:"lastname,omitempty"`
	Ano		  int16    `json:"lastname,omitempty"`
	Cor		  string   `json:"lastname,omitempty"`
}

var veiculo []Carro

func GetVeiculo(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(veiculo)
}

func GetCarro(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range veiculo {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Carro{})
}

func CreateCarro(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var Carro Carro
    _ = json.NewDecoder(r.Body).Decode(&Carro)
    Carro.ID = params["id"]
    veiculo = append(veiculo, Carro)
    json.NewEncoder(w).Encode(veiculo)
}

func DeleteCarro(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range veiculo {
        if item.ID == params["id"] {
            veiculo = append(veiculo[:index], veiculo[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(veiculo)
    }
}

func main() {
    router := mux.NewRouter()
    veiculo = append(veiculo, Carro{ID: "1", Modelo: "Mobi", Marca: "Fiat", Ano:"2019", Cor "Vermelho"})
    veiculo = append(veiculo, Carro{ID: "2", Modelo: "March", Marca: "Nissan", Ano: "2019", Cor "Branco"})
    router.HandleFunc("/contato", GetVeiculo).Methods("GET")
    router.HandleFunc("/contato/{id}", GetCarro).Methods("GET")
    router.HandleFunc("/contato/{id}", CreateCarro).Methods("POST")
    router.HandleFunc("/contato/{id}", DeleteCarro).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000", router))
}