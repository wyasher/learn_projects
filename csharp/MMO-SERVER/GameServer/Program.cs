using GameServer.Network;

namespace GameServer
{
    internal class Program
    {
        static void Main(string[] args)
        {
            TcpSocketListener listener = new TcpSocketListener("0.0.0.0", 32500);
            listener.Start();
        }
    }
}