import errno
import io
import os
import signal
import socket
import sys
import time


def grim_reaper(signum, frame):
    while True:
        try:
            pid, status = os.waitpid(-1, os.WNOHANG)
            print(
                'Child {pid} terminated with status {status}'
                '\n'.format(pid=pid, status=status)
            )
        except OSError:
            return
        if pid == 0:
            return


class WSGIServer(object):
    address_family = socket.AF_INET
    socket_type = socket.SOCK_STREAM
    request_queue_size = 1024

    def __init__(self, server_address):
        # 创建网络件监听

        self.listen_socket = listen_socket = socket.socket(
            self.address_family,
            self.socket_type
        )
        # 允许重复使用地址
        listen_socket.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
        # 绑定地址
        listen_socket.bind(server_address)
        # 监听
        listen_socket.listen(self.request_queue_size)
        # 获取  host 和 port
        host, port = self.listen_socket.getsockname()[:2]
        self.server_name = socket.getfqdn(host)
        self.server_port = port
        self.headers_set = []
        self.client_connection = None
        self.application = None
        self.request_data = None

    def set_app(self, application):
        self.application = application

    def server_forever(self):
        listen_socket = self.listen_socket
        while True:
            try:
                self.client_connection, client_address = listen_socket.accept()
            except IOError as e:
                code, msg = e.args
                if code == errno.EINTR:
                    continue
                else:
                    raise
                    # 创建新链接
            pid = os.fork()
            if pid == 0:  # child
                listen_socket.close()
                # time.sleep(10)
                self.handle_one_request()
                os._exit(0)
            else:
                self.client_connection.close()

    def parse_request(self, text):
        """解析出请求协议，path 和method"""
        # 获取第一行
        request_line = text.splitlines()[0]
        request_line = request_line.rstrip('\r\n')
        (self.request_method, self.path, self.request_version) = request_line.split()

    def get_environ(self):
        """WSGI必须要的参数"""
        env = {'wsgi.version': (1, 0), 'wsgi.url_scheme': 'http', 'wsgi.input': io.StringIO(self.request_data),
               'wsgi.errors': sys.stderr, 'wsgi.multithread': False, 'wsgi.multiprocess': False, 'wsgi.run_once': False,
               'REQUEST_METHOD': self.request_method, 'PATH_INFO': self.path, 'SERVER_NAME': self.server_name,
               'SERVER_PORT': str(self.server_port)}

        return env

    def start_response(self, status, response_headers, exc_info=None):
        server_headers = [
            ('Date', 'Mon, 15 Jul 2019 5:54:48 GMT'),
            ('Server', 'WSGIServer 0.2'),
        ]
        self.headers_set = [status, response_headers + server_headers]

    def handle_one_request(self):
        request_data = self.client_connection.recv(1024)
        self.request_data = request_data = request_data.decode('utf-8')
        print(''.join(f'<{line}\n' for line in request_data.splitlines()))
        # 解析请求协议第一行
        self.parse_request(request_data)

        env = self.get_environ()

        result = self.application(env, self.start_response)
        self.finish_response(result)

    def finish_response(self, result):
        """返回response"""
        try:
            status, response_headers = self.headers_set
            response = f'HTTP/1.1 {status}\r\n'
            # 返回响应头
            for header in response_headers:
                response += '{0}: {1}\r\n'.format(*header)
            response += '\r\n'
            for data in result:
                response += data.decode('utf-8')
            print(''.join(f'> {line}\n' for line in response.splitlines()))
            response_bytes = response.encode()
            self.client_connection.sendall(response_bytes)
        finally:
            self.client_connection.close()


SERVER_ADDRESS = (HOST, PORT) = '', 8888


def make_server(server_address, application):
    """创建服务"""
    signal.signal(signal.SIGCHLD,grim_reaper)
    server = WSGIServer(server_address)
    server.set_app(application)
    return server


if __name__ == '__main__':
    if len(sys.argv) < 2:
        sys.exit('Provide a WSGI application object as module:callable')
    # python main.py pyramidapp:app
    app_path = sys.argv[1]
    module, application = app_path.split(':')
    module = __import__(module)
    application = getattr(module, application)
    httpd = make_server(SERVER_ADDRESS, application)
    print(f'WSGIServer: Serving HTTP on port {PORT} ...\n')
    # 启动服务
    httpd.server_forever()
