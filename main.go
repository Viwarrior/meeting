package main

import (
	"fmt"
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"encoding/json"
)

type Meeting struct{
    Id  primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Title string `json:"Title,omitempty" bson:"Title,omitempty"`
    Participants bson.A `json:"Participants,omitempty" bson:"Participants,omitempty"`
    StartTime time.Time `json:"StartTime,omitempty" bson:"StartTime,omitempty"`
    EndTime time.Time `json:"EndTime,omitempty" bson:"EndTime,omitempty"`
    CreationTime time.Time `json:"CreationTime,omitempty" bson:"CreationTime,omitempty"`
    
}

var client *mongo.Client

func establishConnect(){
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://avi:mongoavi@cluster0.zdcb2.gcp.mongodb.net/test?w=majority"))
    if err != nil {
        log.Fatal(err)
    }
    ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }
	// meeting := Meeting{"HEllo",bson.A{"a@c.com","b@t.com"},time.Now(),time.Now(),time.Now()}
    // quickstartDatabase := client.Database("quickstart")
    // podcastsCollection := quickstartDatabase.Collection("meetings")
	// podcastResult, err := podcastsCollection.InsertOne(ctx, meeting)
	// defer client.Disconnect(ctx)
	// if err != nil { 
	// 	log.Fatal(err)
	// }
	// fmt.Println(podcastResult.InsertedID)
}


func getMeeting(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 		http.NotFound(w, r)
	// 		return
	// }
	switch r.Method {
	case "GET":
			for k, v := range r.URL.Query() {
				   fmt.Printf("%s: %s\n", k, v)
				   if k=="id"{
						quickstartDatabase := client.Database("quickstart")
						mCollection := quickstartDatabase.Collection("meetings")
						var meeting Meeting
						err := mCollection.FindOne(context.TODO(), bson.D{{"_id",v}}).Decode(&meeting)
						if err != nil {
							log.Fatal(err)
								}
						
						json.NewEncoder(w).Encode(meeting)
				   }
			}
			
			w.Write([]byte("Received a GET request\n"))
   default:
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
   }

}


func postMeeting(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 		http.NotFound(w, r)
	// 		return
	// }
	switch r.Method {
	case "POST":
		defer r.Body.Close()
		var meeting Meeting
		json.NewDecoder(r.Body).Decode(&meeting)
		quickstartDatabase := client.Database("quickstart")
		mCollection := quickstartDatabase.Collection("meetings")
		ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
		podcastResult, err := mCollection.InsertOne(ctx, meeting)
		defer client.Disconnect(ctx)
		if err != nil { 
			log.Fatal(err)
		}
		fmt.Println(podcastResult.InsertedID)


   default:
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
   }

}

func helloWorld(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Meeting App API")
}

 

func main() {
	establishConnect()
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/meeting", getMeeting)
	http.HandleFunc("/meetings", postMeeting)
    http.ListenAndServe(":8081", nil)
    


	// deleteResult, err := podcastsCollection.DeleteOne(context.TODO(), bson.D{{"_id",podcastResult.InsertedID}})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(deleteResult.DeletedCount)
}