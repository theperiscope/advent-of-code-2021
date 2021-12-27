 aaaa
b    c
b    c
 dddd
e    f
e    f
 gggg

INPUT: acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf

0 = ?, 1 = ?, 2 = ?, 3 = ?, 4 = ?, 5 = ?, 6 = ?, 7 = ?, 8 = ?, 9 = ?

# easy: unique # of segments
0 = ?, 1 = ab, 2 = ?, 3 = ?, 4 = eafb, 5 = ?, 6 = ?, 7 = dab, 8 = acedgfb, 9 = ?

 ....
.    .
.    .
 ....
.    .
.    .
 ....


# ab = 1, dab = 7
> aaaa = d

 dddd
.    .
.    .
 ....
.    .
.    .
 ....

# 0,6,9
> share 4 (a, b, f, g) segments and have total 6 segments
> 0,6,9 = cefabd | cdfgeb | cagedb
> CALC 4 => cbed, REMAINING 3 => afg

# 6
> 6 = `cbed` + 2 of 3 `afg`
> because `a` is part of #1 cannot be part of #6
> 6 => cdfgeb (select one without an `a`)
> also because 6 has 'b' but not 'a' ... b is in the `f` segment and `a` therefore  in the c

0 = ?, 1 = ab, 2 = ?, 3 = ?, 4 = eafb, 5 = ?, 6 = cdfgeb, 7 = dab, 8 = acedgfb, 9 = ?


 dddd
.    a
.    a
 ....
.    b
.    b
 ....

# 2, 3, 5


 dddd
y    a
y    a
 xxxx
z    b
z    b
 wwww

> share dxw
> 2, 3, 5 (length 5) =  cdfbe | gcdfa | fbcad
> shared cdf
> ...but we know where `d` is
> ==> xw = cf
> ==> 3 = `7` + cf ... fbcad

0 = ?, 1 = ab, 2 = ?, 3 = fbcad, 4 = eafb, 5 = ?, 6 = cdfgeb, 7 = dab, 8 = acedgfb, 9 = ?

# 4
> 4 = eafb
> 3 = fbcad
> ==> exclude from 3 characters in 3 ... remains `e` ... therefore y = e

 dddd
e    a
e    a
 xxxx
z    b
z    b
 wwww

# 4

> 4 = eabx = eabf
> therefore xxxx = f

 dddd
e    a
e    a
 ffff
z    b
z    b
 wwww

# 3
> 3 = fbcad = dadbw
> therefore wwww = c

# 6 = cdfgeb

 dddd
e    a
e    a
 ffff
z    b
z    b
 cccc

# 8
> 8 = acedgfb = abcdefz
> therefore z = g

 dddd
e    a
e    a
 ffff
g    b
g    b
 cccc
