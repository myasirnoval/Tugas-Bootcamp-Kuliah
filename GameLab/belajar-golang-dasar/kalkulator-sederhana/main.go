package main

import "fmt"

func addition(a int, b int) int {
	var result int

	result = a + b

	return result
}

func subtraction(a int, b int) int {
	var result int

	result = a - b

	return result
}
func multiplication(a int, b int) int {
	var result int

	result = a * b

	return result
}
func distribution(a int, b int) int {
	var result int

	if b != 0 {
		result = a / b
	} else {
		result = 0
	}

	return result
}
func modulus(a int, b int) int {
	var result int

	result = a % b

	return result
}

func main() {
	var numbers []int

	var operator string
	var answer string

	var a int
	var b int
	var x int

	fmt.Println("\t\tAPLIKASI KALKULATOR SEDERHANA")
MENU:
	fmt.Println("MENU APLIKASI:")
	fmt.Println("1. Menghitung nilai")
	fmt.Println("2. Menampilkan nilai")
	fmt.Println()

	fmt.Println("Pilih menu: ")
	fmt.Scanln(&answer)

	if answer == "1" {
		goto START
	} else if answer == "2" {
		goto VISUAL
	} else {
		fmt.Println("Nomor yang anda masukan salah")
		fmt.Println("Jalankan ulang program?[y/N] ")
		fmt.Scanln(&answer)

		if answer == "y" || answer == "Y" {
			goto MENU
		} else {
			goto FINISH
		}
	}

START:
	fmt.Println("Gunakan operator [+, -, *, /, %] untuk mulai menghitung")
	fmt.Print("Masukan Operator: ")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Print("Masukan nilai a: ")
		fmt.Scanln(&a)

		fmt.Print("Masukan nilai b: ")
		fmt.Scanln(&b)

		x = addition(a, b)
		numbers = append(numbers, x)

		fmt.Println("Hasil perhitungan : ", x)
		fmt.Println("lanjut menghitung?[y/N] ")
		fmt.Scanln(&answer)

		if answer == "y" || answer == "Y" {
			goto START
		} else {
			fmt.Println("Masuk ke menu?[y/N]")
			fmt.Scanln(&answer)
			if answer == "y" || answer == "Y" {
				goto MENU
			} else {
				goto FINISH
			}
		}
	case "-":
		fmt.Print("Masukan nilai a: ")
		fmt.Scanln(&a)

		fmt.Print("Masukan nilai b: ")
		fmt.Scanln(&b)

		x = subtraction(a, b)
		numbers = append(numbers, x)

		fmt.Println("Hasil perhitungan : ", x)
		fmt.Println("lanjut menghitung?[y/N] ")
		fmt.Scanln(&answer)

		if answer == "y" || answer == "Y" {
			goto START
		} else {
			fmt.Println("Masuk ke menu?[y/N]")
			fmt.Scanln(&answer)
			if answer == "y" || answer == "Y" {
				goto MENU
			} else {
				goto FINISH
			}
		}
	case "*":
		fmt.Print("Masukan nilai a: ")
		fmt.Scanln(&a)

		fmt.Print("Masukan nilai b: ")
		fmt.Scanln(&b)

		x = multiplication(a, b)
		numbers = append(numbers, x)

		fmt.Println("Hasil perhitungan : ", x)
		fmt.Println("lanjut menghitung?[y/N] ")
		fmt.Scanln(&answer)

		if answer == "y" || answer == "Y" {
			goto START
		} else {
			fmt.Println("Masuk ke menu?[y/N]")
			fmt.Scanln(&answer)
			if answer == "y" || answer == "Y" {
				goto MENU
			} else {
				goto FINISH
			}
		}
	case "/":
		fmt.Print("Masukan nilai a: ")
		fmt.Scanln(&a)

		fmt.Print("Masukan nilai b: ")
		fmt.Scanln(&b)

		x = distribution(a, b)
		numbers = append(numbers, x)

		fmt.Println("Hasil perhitungan : ", x)
		fmt.Println("lanjut menghitung?[y/N] ")
		fmt.Scanln(&answer)

		if answer == "y" || answer == "Y" {
			goto START
		} else {
			fmt.Println("Masuk ke menu?[y/N]")
			fmt.Scanln(&answer)
			if answer == "y" || answer == "Y" {
				goto MENU
			} else {
				goto FINISH
			}
		}
	case "%":
		fmt.Print("Masukan nilai a: ")
		fmt.Scanln(&a)

		fmt.Print("Masukan nilai b: ")
		fmt.Scanln(&b)

		x = modulus(a, b)
		numbers = append(numbers, x)

		fmt.Println("Hasil perhitungan : ", x)
		fmt.Println("lanjut menghitung?[y/N] ")
		fmt.Scanln(&answer)

		if answer == "y" || answer == "Y" {
			goto START
		} else {
			fmt.Println("Masuk ke menu?[y/N]")
			fmt.Scanln(&answer)
			if answer == "y" || answer == "Y" {
				goto MENU
			} else {
				goto FINISH
			}
		}
	default:
		fmt.Println("Operator yang anda masukan tidak sesuai")
		fmt.Println("Jalankan ulang program?[y/N] ")
		fmt.Scanln(&answer)

		if answer == "y" || answer == "Y" {
			goto MENU
		} else {
			goto FINISH
		}
	}

VISUAL:
	if numbers != nil {
		fmt.Println()
		fmt.Println("Data yang masuk")
		for i := 0; i < len(numbers); i++ {
			fmt.Printf("[%d] ", numbers[i])
		}

		fmt.Println()
		fmt.Println("Masuk ke menu?[y/N]")
		fmt.Scanln(&answer)

		if answer == "y" || answer == "Y" {
			goto MENU
		} else {
			goto FINISH
		}
	} else {
		fmt.Println()
		fmt.Println("data kosong, silahkan masukan data terlebih dahulu")
		fmt.Println()
		goto MENU
	}

FINISH:
	fmt.Println()
	fmt.Println("Terimakasih sudah menggunakan Aplikasi saya")
}
