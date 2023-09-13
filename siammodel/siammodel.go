package siammodel

type (
	UserData struct {
		NIM          string
		Nama         string
		Jenjang      string
		Fakultas     string
		Jurusan      string
		ProgramStudi string
		Seleksi      string
		NomorUjian   string
		FotoProfil   string
	}

	AddCourseStruct struct {
		Hari           string // Nama hari
		Jam            string // HH:MM - HH:MM
		Kelas          string
		ProgramStudi   string
		Kode           string
		MataKuliah     string
		TahunKurikulum string // YEAR
		Peminat        int    // [PEMINAT]/[kuota]
		Kuota          int    // [peminat]/[KUOTA]
		SKS            int
		Ruang          string
		Jenis          string // Luring/Daring/Hybrid
	}
)