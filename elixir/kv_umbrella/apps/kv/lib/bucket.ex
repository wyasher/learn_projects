defmodule KV.Bucket do
  use Agent, restart: :temporary
  @doc """
      start a new bucket
  """
  def start_link(_opts) do
    Agent.start_link(fn -> %{} end)
  end
  @doc """
      get a value from a bucket
  """
  def get(bucket,key) do
    Agent.get(bucket,&Map.get(&1,key))
  end
  @doc """
      put a value in a bucket
  """
  def put(bucket,key,value) do
    Agent.update(bucket,&Map.put(&1,key,value))
  end

  @doc """
      delete a value from a bucket
  """
  def delete(bucket,key) do
    Agent.get_and_update(bucket,&Map.pop(&1,key))
  end



end
