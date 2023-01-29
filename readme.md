# **CRM**

## Final Project of Golang Course
### What's it ?

#### This project consists of creating apiRest on a fictitious database called datacustomers.With a previous Customers data structure with the attributes: 
#### - *ID:        string* 
#### - *Name      string*
#### - *Role      string*
#### - *Email     string*
#### - *Phone     int*    
#### - *Contacted bool* 

### This is our mock database : 
#### >>> var dataCustomers = map[int]customer{
####  	   1: {
####		    ID:        "1",
####        	Name:      "Miguel",
####		    Role:      "Administrateur",
#### 	    	Email:     "mig@gmail.com",
####		    Phone:     690505814,
####		    Contacted: true,
####	    },
####	   2: {
####		    ID:        "2",
####		    Name:      "Florian",
####		    Role:      "Developpeur web",
####		    Email:     "flor@gmail.com",
####		    Phone:     690507851,
####		    Contacted: true,
####	    },
####	   3: {
####		    ID:        "3",
####		    Name:      "Loic",
####		    Role:      " Graphic Designer",
####		    Email:     "loic@gmail.com",
####		    Phone:     670565842,
####		    >>>Contacted: true,
####	    },
####	   4: {
####		    ID:        "4",
####		    Name:      "kevin",
####		    Role:      "UI Designer",
####		    Email:     "kev@gmail.com",
####		    Phone:     677879747,
####		    Contacted: true,
####	},
#### }

## About ours ApiRest

### Here are our different api
### *NB : all these routes defined in the main, and you must first instantiate mux as newRouter() and their func defined out the main, in top.

#### * _Get All Customers_

#### router.HandleFunc("/customers", getCustomers).Methods("GET")

#### * _Get a specific Customer_

#### router.HandleFunc("/customer/{id}", getCustomer).Methods("GET")

#### * _Create a Customer_

#### router.HandleFunc("/customer", addCustomer).Methods("POST")

#### * _Update a Customer_

#### router.HandleFunc("/customer/{id}", updateCustomer).Methods("PUT")

#### * _Index page_

#### router.HandleFunc("/index", Index).Methods("GET")

### Router Port is **3000**

## The launch app setting

### To launch the app, run `go run main.go`

#### if you see the message : *Server is starting on port 3000...*, so everything is ok

### Open postman app or use the curl command and type :
### *http://localhost:3000/customers* to view all customers
### *http://localhost:3000/customer/{id}* to view a customers if you select the get method, to update if you select the put method, delete if you select delete
### *http://localhost:3000/customer* to add a customer inside our mock database 