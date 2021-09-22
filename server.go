package main

import (
	    "fmt"
		"html/template"
		"log"
		"net/http"
		"path"
		"os"
		"encoding/json"
		"io"
		"io/ioutil"
		// "github.com/joho/godotenv"
)


type Airlines []struct {
	Airport struct {
		Code string `json:"Code"`
		Name string `json:"Name"`
	} `json:"Airport"`
	Time struct {
		Label string `json:"Label"`
		Month int `json:"Month"`
		MonthName string `json:"Month Name"`
		Year int `json:"Year"`
	} `json:"Time"`
	Statistics struct {
		OfDelays struct {
			Carrier int `json:"Carrier"`
			LateAircraft int `json:"Late Aircraft"`
			NationalAviationSystem int `json:"National Aviation System"`
			Security int `json:"Security"`
			Weather int `json:"Weather"`
		} `json:"# of Delays"`
		Carriers struct {
			Names string `json:"Names"`
			Total int `json:"Total"`
		} `json:"Carriers"`
		Flights struct {
			Cancelled int `json:"Cancelled"`
			Delayed int `json:"Delayed"`
			Diverted int `json:"Diverted"`
			OnTime int `json:"On Time"`
			Total int `json:"Total"`
		} `json:"Flights"`
		MinutesDelayed struct {
			Carrier int `json:"Carrier"`
			LateAircraft int `json:"Late Aircraft"`
			NationalAviationSystem int `json:"National Aviation System"`
			Security int `json:"Security"`
			Total int `json:"Total"`
			Weather int `json:"Weather"`
		}
	} `json:"Statistics"`
}

//array for storing filtered airports

type Airports_fitered []struct {
	Airport struct {
		Code string `json:"Code"`
		Name string `json:"Name"`
	} `json:"Airport"`
	Time struct {
		Label string `json:"Label"`
		Month int `json:"Month"`
		MonthName string `json:"Month Name"`
		Year int `json:"Year"`
	} `json:"Time"`
	Statistics struct {
		OfDelays struct {
			Carrier int `json:"Carrier"`
			LateAircraft int `json:"Late Aircraft"`
			NationalAviationSystem int `json:"National Aviation System"`
			Security int `json:"Security"`
			Weather int `json:"Weather"`
		} `json:"# of Delays"`
		Carriers struct {
			Names string `json:"Names"`
			Total int `json:"Total"`
		} `json:"Carriers"`
		Flights struct {
			Cancelled int `json:"Cancelled"`
			Delayed int `json:"Delayed"`
			Diverted int `json:"Diverted"`
			OnTime int `json:"On Time"`
			Total int `json:"Total"`
		} `json:"Flights"`
		MinutesDelayed struct {
			Carrier int `json:"Carrier"`
			LateAircraft int `json:"Late Aircraft"`
			NationalAviationSystem int `json:"National Aviation System"`
			Security int `json:"Security"`
			Total int `json:"Total"`
			Weather int `json:"Weather"`
		}
	} `json:"Statistics"`
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)

	// Get handler for filename, size and headers
	airport :=r.FormValue("code")
	file, handler, err := r.FormFile("filename")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	
	dst, err := os.Create(handler.Filename)
	defer dst.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonFile, err := os.Open("airlines.json")
	
    if err != nil {
		fmt.Println(err)
	}
    byteValue, _ := ioutil.ReadAll(jsonFile)
    var airlines Airlines
	 
	json.Unmarshal(byteValue, &airlines)
    count :=0
	
	var airports_fitered Airports_fitered
    for i := 0; i < len(airlines); i++ {
		if airlines[i].Airport.Code == airport{
			count=count+1
			
			airports_fitered=append(airports_fitered,airlines[i])
		}
	}
	for i := 0; i < len(airports_fitered); i++ {
		// fmt.Println(airports_fitered[i].Airport.Code)
		// fmt.Println(airports_fitered[i].Airport.Name)
		// fmt.Println(airports_fitered[i].Time.Label)	
	}
	
	defer jsonFile.Close()
	    tr:=airports_fitered
        fp := path.Join("static", "data.html")
		tmpl, err := template.ParseFiles(fp)
		// 
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w,tr); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}



}




func uploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		
	case "POST":

		uploadFile(w, r)

	}
}
	


func main() {
	type Book struct {
		Title  string
		Author string
	}
	// os.Setenv("PORT", "8081")
	port:=os.Getenv("PORT")
	fmt.Println(os.Getenv("PORT"))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		book := Book{"hii", "hello"}
        fp := path.Join("static", "index.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w,book); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
    })

    http.HandleFunc("/readAirport", uploadHandler)


    log.Fatal(http.ListenAndServe(":"+port, nil))

}