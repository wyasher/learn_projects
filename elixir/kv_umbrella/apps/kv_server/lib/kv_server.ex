defmodule KVServer do
  require Logger

  def accept(port) do
    {:ok, socket} =
      :gen_tcp.listen(port, [:binary, packet: :line, active: false, reuseaddr: true])

    Logger.info("Listening on port #{port}")
    loop_acceptor(socket)
  end

  def loop_acceptor(socket) do
    {:ok, client} = :gen_tcp.accept(socket)
    {:ok, pid} = Task.Supervisor.start_child(KVServer.TaskSupervisor, fn -> serve(client) end)

    #    使得子进程成为client套接字的“控制进程”。如果我们不这样做，如果它崩溃，接受器将关闭所有客户端，因为套接字将绑定到接受它们的进程
    :ok = :gen_tcp.controlling_process(client, pid)
    loop_acceptor(socket)
  end

  defp serve(socket) do
#    msg =
#      case read_line(socket) do
#        {:ok, data} ->
#          case KVServer.Command.parse(data) do
#            {:ok, command} -> KVServer.Command.run(command)
#            {:error, _} = err -> err
#          end
#
#        {:error, _} = err ->
#          err
#      end
    #with将检索<-右侧返回的值，并将其与左侧的模式匹配。如果值与模式匹配，则with移动到下一个表达式。如果没有匹配，则返回不匹配的值。
    msg = with {:ok,data} <- read_line(socket),
               {:ok,command} <- KVServer.Command.parse(data),
               do: KVServer.Command.run(command)

    write_line(socket, msg)
    serve(socket)
  end

  defp read_line(socket) do
    :gen_tcp.recv(socket, 0)
  end

  defp write_line(socket, {:ok, text}) do
    :gen_tcp.send(socket, text)
  end

  defp write_line(socket, {:error, :unknown_command}) do
    :gen_tcp.send(socket, "UNKNOWN COMMAND\r\n")
  end

  defp write_line(_socket,{:error,:closed}) do
    exit(:shutdown)
  end
  defp write_line(socket, {:error, :not_found}) do
    :gen_tcp.send(socket, "NOT FOUND\r\n")
  end
  defp write_line(socket,{:error,error}) do
    :gen_tcp.send(socket, "ERROR\r\n")
    exit(error)
  end
end
