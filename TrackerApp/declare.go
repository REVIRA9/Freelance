package main

const NMAX int = 10

type tracking struct {
	proyek                                   string
	klien, status                            string
	deadlineDay, deadlineMonth, deadlineYear int
	bayaran                                  int
}

type tabTrack [NMAX]tracking