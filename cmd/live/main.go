package main

import (
	"io"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/quells/LennardJonesGo/sim"
)

func main() {
	r := gin.Default()

	{
		s := DefaultSim()
		s.Step()
	}

	r.StaticFile("/", "./web/index.html")
	r.StaticFile("/app.js", "./web/app.js")

	sseGroup := r.Group("/sse").Use(sseHeaders())
	sseGroup.GET("/lj", sseLJ)

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

func sseHeaders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for h, v := range map[string]string{
			"Content-Type":      "text/event-stream",
			"Cache-Control":     "no-cache",
			"Connection":        "keep-alive",
			"Transfer-Encoding": "chunked",
		} {
			ctx.Writer.Header().Set(h, v)
		}
		ctx.Next()
	}
}

func sseLJ(ctx *gin.Context) {
	lj := DefaultSim()

	sendToClient := make(chan interface{})

	go func() {
		updateClient := time.NewTicker(100 * time.Millisecond)

		sendUpdate := func() {
			K, U, T := sim.Stats(lj.Rs, lj.Vs, lj.L, lj.M)
			pos := make(byZ, len(lj.Rs))
			copy(pos, lj.Rs)
			sort.Sort(pos)

			sendToClient <- map[string]interface{}{
				"t":   lj.Steps,
				"K":   K,
				"U":   U,
				"T":   T,
				"pos": pos,
			}
		}

		sendUpdate()
		for {
			select {
			case <-ctx.Request.Context().Done():
				log.Println("disconnected")
				return

			case <-updateClient.C:
				sendUpdate()

			default:
				lj.Step()
			}
		}
	}()

	ctx.SSEvent("reset", map[string]interface{}{
		"L": lj.L,
	})
	ctx.Stream(func(w io.Writer) bool {
		if msg, ok := <-sendToClient; ok {
			ctx.SSEvent("update", msg)
			return true
		}
		return false
	})
}
