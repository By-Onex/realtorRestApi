package parser

import (
	"container/list"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/By-Onex/realtorRestApi/models"

	"github.com/PuerkitoBio/goquery"
)

func GetApartments(count int) (*list.List, error) {
	apart := list.New()

	url := "https://www.avito.ru/novokuznetsk/kvartiry/prodam-ASgBAgICAUSSA8YQ"
	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	onlyNumber := regexp.MustCompile(`[^0-9]`)
	dotNumber := regexp.MustCompile(`[^0-9.]`)

	doc.Find(".item.item_table").Each(func(i int, s *goquery.Selection) {
		item := &models.Apartment{}

		price := s.Find(".snippet-price").First().Text()
		price = onlyNumber.ReplaceAllString(price, "")
		item.Price, err = strconv.Atoi(price)
		if err != nil {
			fmt.Println(err)
			return
		}

		info := strings.Split(s.Find(".snippet-link").First().Text(), ",")
		if len(info) != 3 {
			return
		}

		roomCount := onlyNumber.ReplaceAllString(info[0], "")
		item.RoomCount, err = strconv.Atoi(roomCount)
		if err != nil {
			fmt.Println(err)
			return
		}

		//area := strings.ReplaceAll(dotNumber.ReplaceAllString(info[1], ""), ".", ",")
		area := dotNumber.ReplaceAllString(info[1], "")

		item.Area, err = strconv.ParseFloat(area, 64)
		if err != nil {
			fmt.Println(err)
			return
		}

		floors := strings.Split(info[2], "/")

		floor := onlyNumber.ReplaceAllString(floors[0], "")
		item.Floor, err = strconv.Atoi(floor)
		if err != nil {
			fmt.Println(err)
			return
		}

		storeys := onlyNumber.ReplaceAllString(floors[1], "")
		item.Storeys, err = strconv.Atoi(storeys)
		if err != nil {
			fmt.Println(err)
			return
		}

		address := strings.Split(s.Find(".item-address__string").First().Text(), ",")

		if len(address) < 2 {
			return
		}

		streetInfo := strings.Split(address[len(address)-2], " ")
		if len(streetInfo) == 0 {
			return
		}
		street := strings.Join(streetInfo, " ")
		item.Street = street

		num := address[len(address)-1]
		item.Num = num

		apart.PushBack(item)
	})

	return apart, nil
}
