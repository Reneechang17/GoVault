package main

import "github.com/Reneechang17/Distributed-File-System/p2p"

type FileServerOpts struct {
	StorageRoot string
	PathTransformFunc PathTransformFunc
	Transport p2p.Transport
}

type FileServer struct {
	FileServerOpts

	store *Store
}

func NewFileServer(opts FileServerOpts) *FileServer {
	StoreOpts := StoreOpts {
		Root: opts.StorageRoot,
		PathTransformFunc: opts.PathTransformFunc,
	}

	return &FileServer{
		FileServerOpts: opts,
		store: NewStore(StoreOpts),
	}
}

func (s *FileServerOpts) Start() error {
	if err := s.Transport.ListenAndAccept(); err != nil {
		return err
	}
	return nil
}