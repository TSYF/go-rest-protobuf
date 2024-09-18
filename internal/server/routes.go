package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"maps"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"google.golang.org/protobuf/proto"
	__ "protobuf.go/internal/proto/employee"
	__e "protobuf.go/internal/proto/employee"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.HelloWorldHandler)

	mux.HandleFunc("/health", s.healthHandler)

	// employeeHandler := http.HandlerFunc(employeeGet)
	// mux.Handle("/employee/", enableCORS(employeeHandler))
	mux.HandleFunc("GET /employee/{id}", employeeGet)
	mux.HandleFunc("GET /employee", employeesGet)
	mux.HandleFunc("POST /employee", employeePost)

	return mux
}

var employees map[uint64]*__e.Employee = map[uint64]*__e.Employee{
	1: &__e.Employee{
		Id: 1,
		Username: "tomastuf",
		ProfileImage: "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/132.png",
		Salary: 1500.95,
	},
	2: &__e.Employee{
		Id: 2,
		Username: "tomasnix",
		ProfileImage: "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/123.png",
		Salary: 1200.50,
	},
}

func employeeGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=== FETCHING EMPLOYEE ===")
	employeeId, err := strconv.Atoi(r.PathValue("id"))

	id := uint64(employeeId)

	if err != nil {
		fmt.Fprintln(w, "ID format invalid")
	}

	employee := employees[id]

	employeePayload, err := proto.Marshal(employee)
	if err != nil {
		http.Error(w, "Failed to serialize employee", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Write(employeePayload)
	fmt.Println("=== END EMPLOYEE GET ===")
}

func employeesGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=== FETCHING EMPLOYEES ===")

	employeesPrototype := &__.Employees{
		Employees: make([]*__e.Employee, len(employees)),
	}

	i := 0
	for e := range maps.Values(employees) {
		fmt.Println("e")
		fmt.Println(e)
		employeesPrototype.Employees[i] = e
		i++
	}
	
	employeePayload, err := proto.Marshal(employeesPrototype)
	if err != nil {
		http.Error(w, "Failed to serialize employee", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Write(employeePayload)
	fmt.Println("=== END EMPLOYEE GET ===")
}

func employeePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=== EMPLOYEE GET ===")
	body, err := io.ReadAll(r.Body)
	
	if err != nil {
		http.Error(w, "Failed to obtain request body", http.StatusInternalServerError)
		return
	}

	employee := &__.Employee{}
	err = proto.Unmarshal(body, employee)

	if err != nil {
		http.Error(w, "Failed to deserialize employee", http.StatusInternalServerError)
		return
	}

	pw, err := bcrypt.GenerateFromPassword([]byte(*employee.Password), 14)

	if err != nil {
		http.Error(w, "Failed to encrypt pasword", http.StatusInternalServerError)
		return
	}

	pwd := string(pw)
	employee.Password = &pwd
	employees[employee.Id] = employee

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	employeePayload, err := proto.Marshal(employee)
	if err != nil {
		http.Error(w, "Failed to serialize employee", http.StatusInternalServerError)
		return
	}

	w.Write(employeePayload)
	fmt.Println("=== END EMPLOYEE POST ===")
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	fmt.Println("ACAH")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(http.StatusOK)
	fmt.Println("END ACAH")
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON maCross-Origin Request Blocked: The Same Origin Policy disallows reading the remote resource at http://127.0.0.1:8080/employee/1. (Reason: CORS header ‘Access-Control-Allow-Origin’ missing). Status code: 200.rshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
