import socket

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.connect(("localhost", 8888))
sock.sendall(b'test')
data = sock.recv(1024)
print(data.decode())
