package main
import ("fmt")

func main() {
    for a:=1; a<1000; a++ {
        for b:=1; b<1000; b++ {
             if (float64(a + b) -(float64(a*b)/1000)) == 500{
                fmt.Println(a," ",b)
            }
        }
    }
}
/* if a+b+c = 1000   so   c = 1000 -a-b
a^2 + b^2 = c^2 = (1000 -a -b)^2
a^2 +b^2 = 1e6 -2000a -2000b +a^2 +b^2 +2ab
1e6 -2000a -2000b +2ab = 0 --> divide 2000->
500 -a -b +ab/1000 = 0 -->
a + b - ab/1000 = 500
now it's so faster than normal
Finally a is 200 and b 375
So 200 + 375 +c = 1000 --> c = 425 */
