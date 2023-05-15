/*
Copyright © 2023 Jinlong jlonmyway@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"flag"
	"github.com/hanwen/go-fuse/v2/fs"
	_ "github.com/hanwen/go-fuse/v2/fs"
	"github.com/jin06/NexusFS/nfs"
	"log"
	"os"
	"os/signal"
	"syscall"

	//"github.com/hanwen/go-fuse/v2/fuse/nodefs"
	//"github.com/jin06/NexusFS/cmd"
)

func main() {
	var err error
	debug := flag.Bool("debug", false, "print debug data")
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("Usage:\n  hello MOUNTPOINT")
	}
	opts := &fs.Options{}
	opts.AllowOther = true
	opts.Debug = *debug
	opts.Name = "testnexus5"
	clientFS := &nfs.Client{}
	clientFS.Local, err = fs.NewLoopbackRoot("/Users/jinlong/tmp/fs/data1")
	if err != nil {
		panic(err)
	}

	server, err := fs.Mount(flag.Arg(0),clientFS.Local, opts)
	if err != nil {
		log.Fatalf("Mount fail: %v\n", err)
	}
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		server.Unmount()
	}()

	server.Wait()
}
