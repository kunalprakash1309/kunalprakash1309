package main

import (
	"database/sql"
	"encoding/json"
	"github.com/emicklei/go-restful"
	"github.com/kunalprakash1309/RailApi/dbutils"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"time"
)

// DB Driver visible to whole program
var DB *sql.DB

// TrainResource is for holding rail information or train table
type TrainResource struct {
	ID              int
	DriverName      string
	OperatingStatus bool
}

// StationResource holds information about locations
type StationResource struct {
	ID          int
	Name        string
	OpeningTime time.Time
	ClosingTime time.Time
}

// ScheduleResource links both trains and stations
type ScheduleResource struct {
	ID          int
	TrainID     int
	StationID   int
	ArrivalTime time.Time
}

// Register add paths and routes to container
func (t *TrainResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/v1/trains").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/{train-id}").To(t.getTrain))
	ws.Route(ws.POST("").To(t.createTrain))
	ws.Route(ws.DELETE("/{train-id}").To(t.removeTrain))

	container.Add(ws)
	
}

// GET http://localhost:8000/v1/trains/1
func (t TrainResource) getTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")
	err := DB.QueryRow("select ID, DRIVER_NAME, OPERATING_STATUS from train where id=?", id).Scan(&t.ID, &t.DriverName, &t.OperatingStatus)
	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "Train could not be found")
	} else {
		response.WriteEntity(t)
	}
}

// POST http://localhost:8000/v1/trains
func (t TrainResource) createTrain(request *restful.Request, response *restful.Response) {
	log.Println(request.Request.Body)
	log.Println("---------------------1")
	decoder := json.NewDecoder(request.Request.Body)
	var b TrainResource
	err := decoder.Decode(&b)
	log.Println(b.DriverName, b.OperatingStatus)
	log.Println("---------------------2")
	// Error handling is obvious here. So omitting...
	statement, err := DB.Prepare("INSERT INTO train (DRIVER_NAME, OPERATING_STATUS) VALUES (?, ?)")
	log.Println("---------------------3")
	if err != nil {
		log.Println("error in creating sql post command", err)
	}
	log.Println("---------------------4")
	result, err := statement.Exec(b.DriverName, b.OperatingStatus)
	log.Println("---------------------5")
	if err == nil {
		newID, _ := result.LastInsertId()
		b.ID = int(newID)
		response.WriteHeaderAndEntity(http.StatusCreated, b)
		log.Println("-------------------6")
	} else {
		log.Println("error in creating post", err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		log.Println("------------------7")
	}
	log.Println("1")
}

// DELETE http://localhost:8000/v1/trains/1
func (t TrainResource) removeTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")
	statement, _ := DB.Prepare("delete from train where id=?")
	_, err := statement.Exec(id)
	if err == nil {
		response.WriteHeader(http.StatusOK)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}


func main() {
	var err error
	// Connect to Database
	DB, err = sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}

	dbutils.Initialize(DB)

	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})

	t := TrainResource{}
	t.Register(wsContainer)
	log.Printf("start listening on localhost:8000")


	log.Fatal(http.ListenAndServe(":8000", wsContainer))
}

