# !/usr/bin/env python
from Crypto.Util.number import *
from hashlib import sha256
import socketserver
import signal
import string
from os import popen, urandom
from Crypto.Util.number import *
from random import *


class Task(socketserver.BaseRequestHandler):
    def _recvall(self):
        BUFF_SIZE = 2048
        data = b''
        while True:
            part = self.request.recv(BUFF_SIZE)
            data += part
            if len(part) < BUFF_SIZE:
                break
        return data.strip()

    def send(self, msg, newline=True):
        try:
            if newline:
                msg += b'\n'
            self.request.sendall(msg)
        except:
            pass

    def recv(self, prompt=b'> '):
        self.send(prompt, newline=False)
        return self._recvall()

    def close(self):
        self.request.close()

    def proof_of_work(self):
        seed(urandom(8))
        proof = ''.join(
            [choice(string.ascii_letters+string.digits) for _ in range(20)])
        _hexdigest = sha256(proof.encode()).hexdigest()
        self.send(f"[+] sha256(XXXX+{proof[4:]}) == {_hexdigest}".encode())
        x = self.recv(prompt=b'[+] Plz tell me XXXX: ')
        if len(x) != 4 or sha256(x+proof[4:].encode()).hexdigest() != _hexdigest:
            return False
        return True

    def handle(self):
        try:
            if not self.proof_of_work():
                self.send(b"try again!")
                self.close()
                exit()

            key = hex(getrandbits(128))[2:].rjust(32, "0")

            pts = hex(getrandbits(64))[2:].rjust(16, "0")
            ptsh = pts.encode().hex()

            cts1 = popen("./cipher "+key+" "+ptsh).read()[:-1].encode()
            self.send(b"(DEBUG) key:" +key.encode())
            self.send(b"(DEBUG) ptsh:" +ptsh.encode()) 
            cts2 = popen("./cipher_broken "+key+" "+ptsh).read()[:-1].encode()
            self.send(b"T: "+cts1)
            self.send(b"F: "+cts2)
            self.send(b"(DEBUG) pts:" +pts.encode())
            signal.alarm(10)

            ans = self.recv()

            if ans == pts.encode():
                self.send(b"114514")
            else:
                self.send(b"sorry,plz try again")
        except:
            self.send(b"something wrong! plz try again!")
            self.close()


class ThreadedServer(socketserver.ThreadingMixIn, socketserver.TCPServer):
    pass


class ForkedServer(socketserver.ForkingMixIn, socketserver.TCPServer):
    pass


if __name__ == "__main__":
    HOST, PORT = '0.0.0.0', 12345
    server = ForkedServer((HOST, PORT), Task)
    server.allow_reuse_address = True
    server.serve_forever()
