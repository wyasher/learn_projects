defmodule KV do
  @moduledoc """
  Documentation for `KV`.
  """
  use Application

  @impl true
  def start(_type,_args) do
    KV.Supervisor.start_link(name: KV.Supervisor)
  end

end
