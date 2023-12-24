from z3 import *

n1 = Int('n1')
n2 = Int('n2')
n3 = Int('n3')
n4 = Int('n4')

x1 = Int('x1')
y1 = Int('y1')
z1 = Int('z1')

x2 = Int('x2')
y2 = Int('y2')
z2 = Int('z2')

x3 = Int('x3')
y3 = Int('y3')
z3 = Int('z3')

x4 = Int('x4')
y4 = Int('y4')
z4 = Int('z4')

px = Int('px')
py = Int('py')
pz = Int('pz')

vx = Int('vx')
vy = Int('vy')
vz = Int('vz')

solve(
    248315803897794 - 89*n1==x1,
    386127890875011 - 119*n1==y1,
    326651351825022 + 32*n1==z1,
    px + vx * n1 == x1,
    py + vy * n1 == y1,
    pz + vz * n1 == z1,
    332497633176671 - 120*n2 == x2,
    319768494554521 - 49*n2 == y2,
    308514709214883 - 91*n2 == z2,
    px + vx * n2 == x2,
    py + vy * n2 == y2,
    pz + vz * n2 == z2,
    362310144548603 + 80*n3 == x3,
    372801373228571 - 599*n3 == y3,
    154999941640943 + 249*n3 == z3,
    px + vx * n3 == x3,
    py + vy * n3 == y3,
    pz + vz * n3 == z3,
    147600732651147 + 31*n4 == x4,
    257653897951867 + 20*n4 == y4,
    363204898979065 - 18*n4 == z4,
    px + vx * n4 == x4,
    py + vy * n4 == y4,
    pz + vz * n4 == z4,
    n1 >= 0,
    n2 >= 0,
    n3 >= 0,
    n4 >= 0)
