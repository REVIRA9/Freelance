package function

const NMAX int = 1024

type tracking struct {
	proyek        string
	klien, status string
	deadline      int
}

type tabTrack [NMAX]tracking