package publisher

import "io"

type JsonPublisher struct {
}

func (jp JsonPublisher) Publish(fileReader io.ReadCloser) (*string, error) {

}
