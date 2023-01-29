package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type customer struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Role      string `json:"role,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     int    `json:"phone,omitempty"`
	Contacted bool   `json:"contacted,omitempty"`
}

var dataCustomers = map[int]customer{
	1: {
		ID:        "1",
		Name:      "Miguel",
		Role:      "Administrateur",
		Email:     "mig@gmail.com",
		Phone:     690505814,
		Contacted: true,
	},
	2: {
		ID:        "2",
		Name:      "Florian",
		Role:      "Developpeur web",
		Email:     "flor@gmail.com",
		Phone:     690507851,
		Contacted: true,
	},
	3: {
		ID:        "3",
		Name:      "Loic",
		Role:      " Graphic Designer",
		Email:     "loic@gmail.com",
		Phone:     670565842,
		Contacted: true,
	},
	4: {
		ID:        "4",
		Name:      "kevin",
		Role:      "UI Designer",
		Email:     "kev@gmail.com",
		Phone:     677879747,
		Contacted: true,
	},
}

// Getting all customers through a /customers/{id} path
func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dataCustomers)
}

// Getting a single customer through a the /customers path
func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	for k, customer := range dataCustomers {
		if customer.ID == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(dataCustomers[k])
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Customer Not Found\n")
			break
		}
	}
}

// Creating a customer through a /customers path
func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCustomer customer
	json.NewDecoder(r.Body).Decode(&newCustomer)
	for _, v := range dataCustomers {
		if newCustomer.ID == v.ID {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusCreated)
			custom := customer{
				ID:        newCustomer.ID,
				Name:      newCustomer.Name,
				Role:      newCustomer.Role,
				Email:     newCustomer.Email,
				Phone:     newCustomer.Phone,
				Contacted: true,
			}
			dataCustomers[len(dataCustomers)+1] = custom
			json.NewEncoder(w).Encode(dataCustomers)
			break

		}
	}

	/*for k, v := range newCustomer {
		if _, ok := dataCustomers[k]; ok {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusCreated)
			dataCustomers[k+1] = v
			json.NewEncoder(w).Encode(dataCustomers)
		}
	}*/
}

// Updating a customer through a /customers/{id} path
func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	var updateCustomer customer
	json.NewDecoder(r.Body).Decode(&updateCustomer)

	for k, _ := range dataCustomers {
		if strconv.Itoa(k) == id {
			w.WriteHeader(http.StatusOK)
			custom := customer{
				ID:        updateCustomer.ID,
				Name:      updateCustomer.Name,
				Role:      updateCustomer.Role,
				Email:     updateCustomer.Email,
				Phone:     updateCustomer.Phone,
				Contacted: true,
			}
			dataCustomers[k] = custom
			/*custom.ID = updateCustomer.ID
			custom.Name = updateCustomer.Name
			custom.Role = updateCustomer.Role
			custom.Email = updateCustomer.Email
			custom.Phone = updateCustomer.Phone
			custom.Contacted = updateCustomer.Contacted*/

			json.NewEncoder(w).Encode(dataCustomers)
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Customer Not Found\n")
			break
		}
	}

}

// Deleting a customer through a /customers/{id} path
func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	for k, customer := range dataCustomers {
		if customer.ID == id {
			delete(dataCustomers, k)
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(dataCustomers)
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Customer Not Found\n")
			break
		}
	}
}

// Index route
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customer/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customer", addCustomer).Methods("POST")
	router.HandleFunc("/customer/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customer/{id}", deleteCustomer).Methods("DELETE")
	router.HandleFunc("/index", Index).Methods("GET")

	fmt.Println("Server is starting on port 3000...")
	// Pass the customer router into ListenAndServe
	http.ListenAndServe(":3000", router)

}
