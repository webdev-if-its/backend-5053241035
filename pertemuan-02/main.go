package main

import "fmt"

// TODO(Level 1): lihat SOAL.md untuk kontrak lengkap tiap fungsi di bawah.
// Ganti setiap "panic" dengan implementasi yang benar.

func HitungSubtotal(qty int, hargaSatuan float64) float64 {
return float64(qty) * hargaSatuan
}

func HitungTotalPesanan(qty []int, hargaSatuan []float64) float64 {
	var total float64
	for i := 0; i < len(qty); i++ {
		total += HitungSubtotal(qty[i], hargaSatuan[i])
	}
	return total
}

func TerapkanPajak(total float64, tarifPajak float64) float64 {
	return total + (total * tarifPajak)
}

func HitungDiskon(total float64) float64 {
	if total >= 100000 {
		return total * 0.1 
	} else if total >= 50000 {
		return total * 0.05
	}
	return 0
}

func TotalSetelahDiskon(qty []int, hargaSatuan []float64, tarifPajak float64) float64 {
	total := HitungTotalPesanan(qty, hargaSatuan)
	total = TerapkanPajak(total, tarifPajak)
	diskon := HitungDiskon(total)
	return total - diskon
}

func ValidasiPesanan(qty []int, hargaSatuan []float64) (bool, string) {
if len(qty) != len(hargaSatuan) {
	return false, "Jumlah item dan harga satuan tidak sesuai"
}
for i := 0; i < len(qty); i++ {
	if qty[i] <= 0 {
		return false, "jumlah barang (qty) harus lebih dari 0"
	}
	if hargaSatuan[i] <= 0 {
		return false, "harga satuan harus lebih dari 0"
}
return true, ""
}

func TentukanStatus(total float64) string {
	if total >= 1000000 {
		return "Prioritas"
	} else if total >= 100000 { 
		return "Reguler"
	}
	return "Hemat"
}

func RingkasanPesanan(qty []int, hargaSatuan []float64, tarifPajak float64) string {
	panic("belum diimplementasikan")
}

// TODO(Level 9): signature ini SUDAH benar (cari tahu sendiri kenapa
// bentuknya begini - lihat SOAL.md) - tinggal implementasikan isinya.
func Total(harga ...float64) float64 {
	panic("belum diimplementasikan")
}

// TODO(Level 10, bonus): signature ini SUDAH benar (cari tahu sendiri
// kenapa ada dua nilai balik - lihat SOAL.md) - tinggal implementasikan isinya.
func HitungOngkosKirim(beratKg float64, jarakKm float64) (float64, error) {
	panic("belum diimplementasikan")
}

func main() {
	fmt.Println("Sales Order Processor - pertemuan 2")
}
