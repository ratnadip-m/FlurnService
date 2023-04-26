package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SeatPricing struct {
	ID          string `gorm:"primary_key"`
	SeatClass   string `gorm:"unique_index:idx_seatclass"`
	MinPrice    string
	MaxPrice    string
	NormalPrice string
}

type SeatBooking struct {
	ID        string `gorm:"primary_key"`
	SeatID    string `gorm:"unique_index:idx_seat"`
	SeatClass string `gorm:"unique_index:idx_seat"`
}

var db *gorm.DB
var err error

// var movies []SeatPricing

func main() {
	// var err error
	router := gin.Default()
	dsn := "root:cdac@tcp(localhost:3306)/moviedb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// db, err = sql.Open("mysql", "root:cdac@tcp(localhost:3306)/moviedb")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()
	db.AutoMigrate(&SeatPricing{}, &SeatBooking{})

	router.POST("/uploadfile", handleXLSXUpload)
	// router.POST("/updatemovies", UpdateMovie)
	// router.GET("/searchmovies", GetMovie)
	// router.GET("/searchmoviesbyidyeargenres", GetMovieByIDYearRatingGenres)

	router.Run(":8087")
}
func readAndStoreXLSXFile(filepath string) error {
	// Open the file for reading
	f, err := xlsx.OpenFile(filepath)
	if err != nil {
		return err
	}

	// Iterate over each sheet in the XLSX file
	for _, sheet := range f.Sheets {
		// Iterate over each row in the sheet
		for _, row := range sheet.Rows {
			// Create a new instance of YourModel
			model := SeatPricing{}

			// Map the values from the XLSX row to the model fields
			// ID, _ := strconv.Atoi(row.Cells[0].Value)
			// model.ID = uint(ID)
			// model.SeatClass = row.Cells[1].Value
			// min, _ := strconv.Atoi(row.Cells[2].Value)
			// model.MinPrice = min
			// nprice, _ := strconv.Atoi(row.Cells[3].Value)
			// model.NormalPrice = nprice
			// maxprice, _ := strconv.Atoi(row.Cells[4].Value)
			// model.MaxPrice = maxprice
			// ID, _ := strconv.Atoi()
			model.ID = row.Cells[0].Value
			model.SeatClass = row.Cells[1].Value
			// min, _ := strconv.Atoi()
			model.MinPrice = row.Cells[2].Value
			// nprice, _ := strconv.Atoi()
			model.NormalPrice = row.Cells[3].Value
			// maxprice, _ := strconv.Atoi()
			model.MaxPrice = row.Cells[4].Value

			// Save the model to the database
			if err := db.Create(&model).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

func handleXLSXUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error while uploading file: %s", err.Error()))
		return
	}

	// Save the file to a temporary location
	filepath := "/tmp/" + file.Filename
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error while saving file: %s", err.Error()))
		return
	}

	// Read the XLSX file and store the data into the database
	if err := readAndStoreXLSXFile(filepath); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error while reading and storing XLSX file: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded and stored successfully", file.Filename))
}

func addSeatPrice(c *gin.Context) {
	var seatPrice SeatPricing

	if err := c.ShouldBindJSON(&seatPrice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign a new ID to the movie
	// movie.ID = len(movies) + 1

	// Add the movie to the list of movies
	// movies = append(movies, movie)
	// result, err := db.Exce("INSERT INTO orders (id, title, year, rating, generes) VALUES (?, ?, ?, ?, ?)",
	// 	movie.ID, movie.Title, movie.Year, movie.Rating, movie.Genres)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	db.Save(&seatPrice)

	c.JSON(200, "Movie Added")
}

func addSeatBooking(c *gin.Context) {
	var seatBooking SeatBooking

	if err := c.ShouldBindJSON(&seatBooking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&seatBooking)

	c.JSON(200, "Movie Added")
}

// func GetSeatPrice(seatClass string, numBookings int) (int, error) {
// 	var seatPricing SeatPricing
// 	var minBookings, maxBookings int

// 	// Retrieve the pricing for the seat class
// 	err := db.Where("seat_class = ?", seatClass).First(&seatPricing).Error
// 	if err != nil {
// 		return 0, err
// 	}

// 	// Determine the minimum and maximum number of bookings
// 	// for the current seat class based on the percentage
// 	// fullness thresholds
// 	totalSeats := 100
// 	minBookings = 0
// 	maxBookings = int(0.6 * float64(totalSeats))

// 	if numBookings >= maxBookings {
// 		// Use the max price for all further seats booked
// 		return seatPricing.MaxPrice, nil
// 	} else if numBookings >= minBookings {
// 		// Use the normal price for the seat booking
// 		return seatPricing.NormalPrice, nil
// 	} else {
// 		// Use the min price if there is no price available for
// 		// the current range
// 		if seatPricing.MinPrice != 0 {
// 			return seatPricing.MinPrice, nil
// 		} else if seatPricing.NormalPrice != 0 {
// 			return seatPricing.NormalPrice, nil
// 		} else {
// 			return seatPricing.MaxPrice, nil
// 		}
// 	}
// }

// func BookSeat(seatID uint, seatClass string) (int, error) {
// 	var numBookings int

// 	// Retrieve the number of bookings for the current seat class
// 	err := db.Model(&SeatBooking{}).Where("seat_class = ?", seatClass).Count(&numBookings).Error
// 	if err != nil {
// 		return 0, err
// 	}

// 	// Get the seat price based on the number of bookings
// 	seatPrice, err := GetSeatPrice(seatClass, numBookings)
// 	if err != nil {
// 		return 0, nil
// 	}

// }
