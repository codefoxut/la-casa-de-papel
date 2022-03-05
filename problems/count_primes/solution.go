package main 



import (
	"fmt"
)

func isOdd(num int) bool {
	return num % 2 > 0
}

func isPrime(num int) bool {
	if num  <=  1 {
		return false
	}

	for i:= 3; i *  i  <= num; {
		if num % i == 0 {
			return false
		}
		i = i + 2
		

	}
	return true
}

func  updateMap(mapOfNums map[int]bool, num, divisor int) {
	for i := divisor; i * divisor < num; i++ {
		fmt.Println("divisor", i*divisor)
		mapOfNums[i*divisor] = false
	}
}


func countPrimes(num int) int {
        p := make(map[int]bool, num)
        for i:=2;i<num;i++ {
            p[i] = true
        }
        for i:=2;i*i<num;i++ {
            if value := p[i]; value {
                for j:=i*i;j<num;j=j+i {
                    p[j]=false
                }
            }
        }
        for  i:=2;i<num;i++ {
            if value := p[i]; value {
                count++
            }
        }
        return count
    }

func countPrimes2(num int) int {
        p := make(map[int]bool)
        p[0]=false 
        p[1]=false
        count:=0
        for  i:=2;i<num;i++ {
            if _, ok := p[i]; !ok {
                count++
                for j:=2*i;j<num;j=j+i {
                    p[j]=false
                }
            }
        }
        return count;
    }


func countPrimes1(num int) int {
	var numOfPrimes int
	mapOfNums := make(map[int]bool)


	if num > 2 {
		numOfPrimes++
	}
	for j:=3; j  < num; j++ {
		mapOfNums[j] = j % 2 > 0
	}

	for i := 3; i < num;  {
		// fmt.Println(i)
		if ok := mapOfNums[i]; ok {
			if isPrime(i) {
			fmt.Println("prime", i)
				numOfPrimes++
			} 
				updateMap(mapOfNums, num,  i)
			
		}
		
		i = i + 2
		
	}
	return numOfPrimes
}


 
func main() {
	fmt.Println("primes 5000000", countPrimes(50_00_000))
	fmt.Println("primes 10", countPrimes(10))
	fmt.Println("primes 5", countPrimes(5))
	fmt.Println("primes 2", countPrimes(2))

}