package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// レコードに関する構造体データの宣言
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {
	// Defaultでginのルーターを初期化する
	router := gin.Default()
	// GET関数を使用して、GETHTTPメソッドと/albumsパスをハンドラー関数に関連付ける
	// getAlbums()ではないことに注意。関数の戻り値をいれてはいない。
	router.GET("/albums", getAlbums)
	// 下で作成したpostAlbumsがrouter.POST(/albums)で使用出来るように、関連付ける
	router.POST("/albums", postAlbums)

	// Run関数を使用して、ルーターをhttp.Serverに接続し、サーバーを起動します。
	router.Run("localhost:8083")
}

// アルバムのシードデータ
// 本来はデータはDBにあるもんだけど、今回はここに
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// json で全てのアルバム情報を返す
func getAlbums(c *gin.Context) { // ここの引数の(c *gin.Context)はどうも、ginのメソッドContext(gin.Context)をcに省略する、みたいな意味みたいね
	// なので、↓はContext.IndentedJSONが性格なメソッド名なんだろう。
	c.IndentedJSON(http.StatusOK, albums)
}

// json でアルバムデータを追加する
func postAlbums(c *gin.Context) {
	var newAlbum album
	// BindJSONを呼び出して、受信したJSONをnewAlbumにバインド。
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// 新しく取得したデータをくわえる
	albums = append(albums, newAlbum)
	// 新しくアルバムデータが追加されたJSONと、201ステータスが返される
	c.IndentedJSON(http.StatusCreated, newAlbum)
	// main次のように、関数が含まれるように関数を変更router.POSTします。から
}
