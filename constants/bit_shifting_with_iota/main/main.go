package main

import "fmt"

const(
        _ = iota
        shift_bit = 1 << (iota * 10)
        shift_bit_again = 1 << (iota * 10)
        
        shift_bits = 1 << iota * 10

        shift_bits_right = 32 >> iota
        shift_bits_left = 32 << iota
        shift_bits_left1 = shift_bits_right << iota

        shift_bits_diff_value = 3 << iota
        
    )

func main() {
    fmt.Printf("%b\n",shift_bit)
    fmt.Printf("%d\n",shift_bit)
    fmt.Printf("%b\n",shift_bit_again)
    fmt.Printf("%d\n",shift_bit_again)
    fmt.Printf("%b\n",shift_bits)
    fmt.Printf("%d\n",shift_bits)
    fmt.Printf("%b\n",shift_bits_right)
    fmt.Printf("%d\n",shift_bits_right)
    fmt.Printf("%b\n",shift_bits_left)
    fmt.Printf("%d\n",shift_bits_left)
    fmt.Printf("%b\n",shift_bits_left1)
    fmt.Printf("%d\n",shift_bits_left1)
    fmt.Printf("%b\n",shift_bits_diff_value)
    fmt.Printf("%d\n",shift_bits_diff_value)
}


/* Notes:

1) The above program shows the bit shifting using IOTA.

2) We use "<<" or ">>" for bit shifting.

3) IOTA is defined in the const, it cannot be defined in the function.

4) "<<" is a precedence over arithmatic operators.

5) In the above program, we are moving 1, 10 places to the left. This will be the output of shift_bit.

6) shift_bit_again : the iota values becomes to 2. Hence, iota * 10 = 2 * 10 = 20.
So 1 << (iota * 10) => 1 << 20 - which means we move 1 by 20 places to the left.

7) shift_bits: the operation performed is 1 << iota * 10. IOTA now becomes 3. so we shift the value of 1, 3 places to the left. which means "1 << 3". So the binary value becomes 1000, which in decimals is 8. Then we multiply the value by 10, so in decimal, the value becomes 80. In binary, the value becomes 1010000.

8) Binary value is represented in %b, its corresponding integer value is represented by %d.

9) shift_bits_right : the IOTA value now becomes 4. so we move the bits 4 places to the right. The expression becomes 32 >> 4 => 100000 >> 4 => 000010. Hence the integer value becomes 2.

10) shift_bits_left : the IOTA value now becomes 5. so we move bits 5 places to the left of the value 32. The value of 32 in bits is 100000. The expression becomes 100000 << 5, which will be 10000000000, which in digits will be 1024.

11) shift_bits_left1: IOTA value becomes 6. The expression states, move 1 6 places to the left of the value shift_bits_right. From the above note, the value of shift_bits_right is 2, which in bits is 10. So it has to move 6 places to the left, which now becomes 1000000, which in decimal becomes 128.

12) shift_bits_diff_value : IOTA value now becomes 7. Expression : 3 << 7. Number 3 in bits is 011. Now we shift 7 places to the left, meaning, we add 7 0's to the right. It becomes 11000000. Now the decimal for 11000000 binary value is (2^7 + 2^8) = (128 + 256 = 384). Hence the output of 3 << 7 is 384.

*/
