using System;
using System.Collections.Generic;
using System.Linq;
using System.Net;
using System.Net.Sockets;
using System.Text;
using System.Threading.Tasks;

namespace GameServer.Network
{
    /// <summary>
    /// 负责监听TCP网络端口，异步接收Socket连接
    /// </summary>
    public class TcpSocketListener
    {
        private IPEndPoint endPoint;
        private Socket? serverSocket;

        public event EventHandler<Socket> SocketConnected; // 客户端接入事件

        public TcpSocketListener(string host, int port) {
            endPoint = new IPEndPoint(IPAddress.Parse(host), port);
        }

        public bool isRunning {
            get {
                return serverSocket != null;
            }
        }


        public void Start() { 
        
            if (!isRunning)
            {
                serverSocket = new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);
                serverSocket.Bind(endPoint);
                serverSocket!.Listen();
                Console.WriteLine("开始监听端口"+endPoint.ToString());

                SocketAsyncEventArgs args = new SocketAsyncEventArgs();
                args.Completed += onAccept;
                serverSocket!.AcceptAsync(args);

            }
        }
        private void onAccept(object? sender, SocketAsyncEventArgs args) { 
        
            if (args.SocketError == SocketError.Success) {
                Socket? client = args.AcceptSocket;
                if (client != null)
                {
                    SocketConnected?.Invoke(this, client);
                }
            }
            //继续接收下一位
            args.AcceptSocket = null;
            serverSocket?.AcceptAsync(args);
        }


        public void Stop()
        {

            serverSocket?.Close();
            serverSocket = null;
        }

    }
}
