package siamauth

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

	AddCourse struct {
		MasaKRS string

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

	Krs struct {
		MasaKRS string

		MataKuliah []MataKuliahKrs
	}

	MataKuliahKrs struct {
		Kode         string
		MataKuliah   string
		SKS          int
		Keterangan   string
		Kelas        string
		ProgramStudi string
	}
)