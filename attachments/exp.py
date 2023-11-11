from pwn import *
from Crypto.Util.number import *
from hashlib import sha256
from itertools import product
import string

table = string.ascii_letters+string.digits


rec = remote('127.0.0.1', 12345)
_ = rec.recvuntil(b'XXXX:')
'''
[+] sha256(XXXX+p9epXzl7VsnfAGl7) == 5cf00945d39011df5e8ff006b112c0637da80e7ac96bd8797e42c7253a1ca4ca
[+] Plz tell me XXXX:

'''
tail,h = _[16:32],_[37:101]
for head in product(table,repeat=4):
    m = "".join(head)+tail.decode()
    h_ = sha256(m.encode())
    if h_.hexdigest() == h.decode():
        print('find!')
        break
rec.sendline("".join(head).encode())
rec.interactive()