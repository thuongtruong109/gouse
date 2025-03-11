package gouse

import (
	"context"
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
)

func Starter() {
	println(Capitalize("hello world from Gouse!"))
}

/* Context */

type Ctx context.Context

var CtxBg = context.Background()

func SetCtx(ctx context.Context) {
	CtxBg = ctx
}

func GetCtx() Ctx {
	return CtxBg
}

func Profile(cpuprofile, memprofile string) {
	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer func() {
			err := f.Close()
			if err != nil {
				log.Printf("error closing CPU profile file: %v", err)
			}
		}()

		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	// ... rest of the program ...

	if memprofile != "" {
		f, err := os.Create(memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer func() {
			err := f.Close()
			if err != nil {
				log.Printf("error closing memory profile file: %v", err)
			}
		}()

		runtime.GC()

		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}

func DetectError(err interface{}) string {
	switch e := err.(type) {
	case error:
		return e.Error()
	case string:
		return e
	default:
		return fmt.Sprintf("%v", err)
	}
}
