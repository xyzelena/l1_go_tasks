package sleep

import (
	"fmt"
	"time"
)

func Sleep(duration time.Duration) {
	fmt.Println("Sleeping for", duration)

	done := make(chan struct{})

	go func() {
		defer close(done)

		start := time.Now()

		for time.Since(start) < duration {
			fmt.Println("Sleeping...")

		}
	}()

	<-done
}
