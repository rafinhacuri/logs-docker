package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"strings"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

var clientDocker *client.Client

func initDocker() error {
	var err error
	clientDocker, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	return nil
}

func scanStream(r io.Reader, onLine func(ts time.Time, line string)) {
	s := bufio.NewScanner(r)
	buf := make([]byte, 0, 64*1024)
	s.Buffer(buf, 1024*1024)
	for s.Scan() {
		raw := s.Text()
		ts, msg := parseTimestampPrefix(raw)
		onLine(ts, msg)
	}
	if err := s.Err(); err != nil && err != io.EOF {
		log.Println("scanner erro:", err)
	}
}

func streamContainerLogs(ctx context.Context, containerID, containerName string) {
	reader, err := clientDocker.ContainerLogs(ctx, containerID, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
		Tail:       "1",
		Timestamps: true,
	})
	if err != nil {
		log.Printf("[%s] erro abrindo logs: %v\n", containerName, err)
		return
	}
	defer reader.Close()

	prOut, pwOut := io.Pipe()
	prErr, pwErr := io.Pipe()
	go func() {
		defer pwOut.Close()
		defer pwErr.Close()
		_, _ = stdcopy.StdCopy(pwOut, pwErr, reader)
	}()

	go scanStream(prOut, func(ts time.Time, line string) {
		if containerName != "dozzle" && containerName != "doku" {
			if err := enviarLog(containerName, line, ts, false); err != nil {
				log.Printf("[%s] erro ao enviar log: %v\n", containerName, err)
			}
		}
	})
	go scanStream(prErr, func(ts time.Time, line string) {
		if containerName != "dozzle" && containerName != "doku" {
			if err := enviarLog(containerName, line, ts, true); err != nil {
				log.Printf("[%s] erro ao enviar log: %v\n", containerName, err)
			}
		}
	})

	<-ctx.Done()
}

func followNewContainers(ctx context.Context) error {
	containers, err := clientDocker.ContainerList(ctx, container.ListOptions{All: false})
	if err != nil {
		return err
	}
	for _, c := range containers {
		name := "unknown"
		if len(c.Names) > 0 {
			name = strings.TrimPrefix(c.Names[0], "/")
		}
		go streamContainerLogs(ctx, c.ID, name)
	}

	f := filters.NewArgs()
	f.Add("type", "container")
	f.Add("event", "start")
	eventsCh, errsCh := clientDocker.Events(ctx, events.ListOptions{Filters: f})
	for {
		select {
		case e := <-eventsCh:
			name := e.Actor.Attributes["name"]
			if name == "" {
				name = e.ID[:12]
			}
			go streamContainerLogs(ctx, e.ID, name)
		case err := <-errsCh:
			if err == context.Canceled || err == io.EOF {
				return nil
			}
			log.Println("Erro nos eventos do Docker:", err)
			return err
		}
	}
}
