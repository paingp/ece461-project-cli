package main

import (
    "net/http"
	"strings"
    "github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
    //router := gin.Default()
    //router.GET("/albums", getAlbums)
    //router.GET("/albums/:id", getAlbumByID)
    //router.POST("/albums", postAlbums)

    //router.Run("localhost:8080")
	
	filePath := os.Args[1]

	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var urls []string

	error := godotenv.Load(".env")
	if error != nil {
		panic("Error loading .env file")
	}

	// Create a token to and HTTP client to access the GitHub API
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	for scanner.Scan() {
		url := scanner.Text()
		module := ratom.Analyze(url, httpClient)
		fmt.Println(module)
		urls = append(urls, url)
	}

	Load .env file
	
	for i := 0; i < len(urls) {
		if strings.Contains(urls[i], "www.npmjs.com") {
			requestURL := urls[i]
		}
	}
	resp, error := httpClient.Get(requestURL)

	if error != nil {
		panic(error)
	}

	defer resp.Body.Close()

	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		panic(error)
	}

	fmt.Println(string(body))

	
}
// getAlbums responds with the list of all albums as JSON.
//func getAlbums(c *gin.Context) {
//   c.IndentedJSON(http.StatusOK, albums)
//}

// postAlbums adds an album from JSON received in the request body.
//func postAlbums(c *gin.Context) {
//    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
//    if err := c.BindJSON(&newAlbum); err != nil {
//        return
//    }

    // Add the new album to the slice.
//    albums = append(albums, newAlbum)
//    c.IndentedJSON(http.StatusCreated, newAlbum)
//}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
//func getLink(c *gin.Context) {
//    id := c.Param("id")

    // Loop through the list of albums, looking for
    // an album whose ID value matches the parameter.
//    for _, a := range albums {
//       if a.ID == id {
//            c.IndentedJSON(http.StatusOK, a)
//            return
//        }
//    }
//    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
//}
