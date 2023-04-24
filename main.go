package main

type SeatPricing struct {
	ID          uint   `gorm:"primary_key"`
	SeatClass   string `gorm:"unique_index:idx_seatclass"`
	MinPrice    int
	MaxPrice    int
	NormalPrice int
}

type SeatBooking struct {
	ID        uint   `gorm:"primary_key"`
	SeatID    uint   `gorm:"unique_index:idx_seat"`
	SeatClass string `gorm:"unique_index:idx_seat"`
}

var db *gorm.DB
var err error

var movies []SeatPricing

func main() {
	var err error
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
	db.AutoMigrate(&SeatPricing{})

	// router.POST("/movies", addMovie)
	// router.POST("/updatemovies", UpdateMovie)
	// router.GET("/searchmovies", GetMovie)
	// router.GET("/searchmoviesbyidyeargenres", GetMovieByIDYearRatingGenres)

	router.Run(":8087")
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
