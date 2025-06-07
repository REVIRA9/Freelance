package main

const NMAX int = 10

type tracking struct {
	proyek                                   string
	klien, status                            string
	deadlineDay, deadlineMonth, deadlineYear int
	bayaran, id                              int
}

type tabTrack [NMAX]tracking